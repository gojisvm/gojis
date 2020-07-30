package agent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gojisvm/gojis/internal/runtime/agent/job"
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
	"github.com/oklog/ulid"
)

// QueueKind represents the description of 8.4, Table 25.
type QueueKind uint8

// Available QueueKinds. QueueUnknown indicates a severe programming error,
// since that value must never be used. It is the default value for QueueKind
// and indicates that something has not been initialized properly.
const (
	QueueUnknown QueueKind = iota
	QueueScript
	QueuePromise
)

// ID is a type alias for uuid.UUID, in order to be independent of
// uuid.UUID as a unique identifier.
type ID ulid.ULID

var (
	entropyLock sync.Mutex
	entropy     = ulid.Monotonic(rand.New(rand.NewSource(time.Now().Unix())), 0)
)

// NewID returns a new ID that is guaranteed to be globally unique.
func NewID() ID {
	entropyLock.Lock()
	defer entropyLock.Unlock()

	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), entropy))
}

// Agent comprises a set of ECMAScript execution contexts, an execution context
// stack, a running execution context, a set of named job queues, an Agent
// Record, and an executing thread. Except for the executing thread, the
// constituents of an agent belong exclusively to that agent. An agent's
// executing thread executes the jobs in the agent's job queues on the agent's
// execution contexts independently of other agents, except that an executing
// thread may be used as the executing thread by multiple agents, provided none
// of the agents sharing the thread have an Agent Record whose [[CanBlock]]
// property is true. While an agent's executing thread executes the jobs in the
// agent's job queues, the agent is the surrounding agent for the code in those
// jobs. The code uses the surrounding agent to access the specification level
// execution objects held within the agent: the running execution context, the
// execution context stack, the named job queues, and the Agent Record's fields.
// Agent is specified in 8.7.
type Agent struct {
	ExecutionContextStack *ExecutionContextStack

	LittleEndian       bool
	CanBlock           bool
	Signifier          ID
	IsLockFree1        bool
	IsLockFree2        bool
	CandidateExecution interface{} // #37: Implement Table 26, CandidateExecutionRecord

	ScriptJobs  *job.Queue
	PromiseJobs *job.Queue
}

// New creates a new agent that is ready to use.
func New() *Agent {
	a := new(Agent)
	a.ExecutionContextStack = NewExecutionContextStack()
	a.LittleEndian = false
	a.CanBlock = false
	a.Signifier = NewID()
	a.ScriptJobs = job.NewQueue()
	a.PromiseJobs = job.NewQueue()
	return a
}

// AgentSignifier returns the Signifier of the agent.
func (a *Agent) AgentSignifier() ID {
	return a.Signifier
}

// AgentCanSuspend returns the CanBlock value of the agent.
func (a *Agent) AgentCanSuspend() bool {
	return a.CanBlock
}

// RunningExecutionContext returns the currently executing ExecutionContext of
// this agent.
func (a *Agent) RunningExecutionContext() *ExecutionContext {
	return a.ExecutionContextStack.Peek()
}

// GetActiveScriptOrModule returns the active script or module.
// GetActiveScriptOrModule is specified in 8.3.1.
func (a *Agent) GetActiveScriptOrModule() lang.InternalValue {
	if a.ExecutionContextStack.IsEmpty() {
		return lang.Null
	}

	panic("#40: 8.3.1")
}

// ResolveBinding is used to determine the binding with the given name. The
// optional argument env can be used to explicitly provide the Lexical
// Environment that is to be searched for the binding. During execution of
// ECMAScript code, ResolveBinding is performed using the following algorithm.
// ResolveBinding is specified in 8.3.2.
func (a *Agent) ResolveBinding(name lang.String, env binding.Environment) *binding.Reference {
	if env == nil {
		env = a.RunningExecutionContext().LexicalEnvironment
		_ = env
	}

	strict := false // FIXME: 8.3.2, Step 3
	_ = strict
	panic("#38: 8.3.2, Step 3")

	// return binding.GetIdentifierReference(env, name, strict)
}

// GetThisEnvironment returns the current environment that currently supplies
// the binding for the keyword 'this'.
// GetThisEnvironment is specified in 8.3.3.
func (a *Agent) GetThisEnvironment() binding.Environment {
	lex := a.RunningExecutionContext().LexicalEnvironment

	/*
		Step out until an environment has a this binding.
		The Global Environment has a this binding, and is the
		only environment which has no outer environment, so this
		will always terminate.
	*/
	for !lex.HasThisBinding() {
		lex = lex.Outer()

		if lex == nil {
			panic("Outer environment cannot be nil, this means that the global object does not have a this binding, which must not happen, or that an environment which is not the global environment has a nil reference as an outer environment.")
		}
	}

	return lex
}

// ResolveThisBinding determines the binding of the keyword 'this' using the
// LexicalEnvironment of the running execution context. ResolveThisBinding is
// specified in 8.3.4.
func (a *Agent) ResolveThisBinding() (lang.Value, errors.Error) {
	return a.GetThisEnvironment().GetThisBinding()
}

// GetNewTarget determines the NewTarget using the lexical environment of the
// running execution context.
// GetNewTarget is specified in 8.3.5.
func (a *Agent) GetNewTarget() lang.Value {
	panic("#8: 8.3.5")
}

// GetGlobalObject returns the global object used by the running execution context.
// GetGlobalObject is specified in 8.3.6.
func (a *Agent) GetGlobalObject() lang.Value {
	return a.RunningExecutionContext().Realm.GlobalObj
}

// EnqueueJob enqueues a new job into a given kind of queue, script or promise.
// EnqueueJob is specified in 8.4.1.
func (a *Agent) EnqueueJob(q QueueKind, jobName string, arguments []lang.Value) {
	callerCtx := a.RunningExecutionContext()
	callerRealm := callerCtx.Realm
	callerScriptOrModule := callerCtx.ScriptOrModule
	pending := job.PendingJob{
		Job:            jobName,
		Arguments:      arguments,
		Realm:          callerRealm,
		ScriptOrModule: callerScriptOrModule,
		HostDefined:    lang.Undefined,
	}
	// do we need to modify pending in any way?

	switch q {
	case QueueScript:
		a.ScriptJobs.Enqueue(pending)
	case QueuePromise:
		a.PromiseJobs.Enqueue(pending)
	default:
		panic(fmt.Sprintf("Unknown queue kind: %v", q))
	}
}

// InitializeHostDefinedRealm is specified in 8.5.
func (a *Agent) InitializeHostDefinedRealm() {
	r := realm.CreateRealm()
	newCtx := &ExecutionContext{
		Function:       lang.Null,
		Realm:          r,
		ScriptOrModule: lang.Null,
	}
	a.ExecutionContextStack.Push(newCtx)

	/*
		If the host requires use of an exotic object to serve as realm's global object, let global be such an object created in
		an implementation-defined manner. Otherwise, let global be undefined, indicating that an ordinary object
		should be created as the global object.
	*/
	// Let global be undefined
	global := lang.Undefined

	/*
		If the host requires that the thisthis binding in realm's global scope return an object other than the global object,
		let thisValue be such an object created in an implementation-defined manner. Otherwise, let thisValue be
		undefined, indicating that realm's global thisthis binding should be the global object.
	*/
	// Let thisValue be undefined
	thisValue := lang.Undefined

	r.SetRealmGlobalObject(global, thisValue)

	globalObj := r.SetDefaultGlobalBindings()
	// #39: Create any implementation-defined global object properties on globalObj.
	_ = globalObj
}

// RunJobs is specified in 8.6.
func (a *Agent) RunJobs() {
	a.InitializeHostDefinedRealm()

	panic("#9: 8.6")
}

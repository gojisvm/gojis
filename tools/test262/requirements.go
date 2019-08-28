package test262

// Phase represents the phase, that the test is expected to produce an error.
type Phase uint8

const (
	// PhaseUnknown is the default value and can be interpreted as "not set".
	PhaseUnknown Phase = iota
	// PhaseParse indicates, that an error occurs while parsing the source text.
	PhaseParse
	// PhaseEarly indicates, that an error occurs prior to evaluation.
	PhaseEarly
	// PhaseResolution indicates, that an error occurs during module resolution.
	PhaseResolution
	// PhaseRuntime indicates, that an error occurs during evaluation.
	PhaseRuntime
)

// Requirements contains all information that were included in the metadata of a
// Test262-testcase.
type Requirements struct {
	// If a test configured with the negative attribute completes without
	// throwing an exception, or if the name of the thrown exception's
	// constructor does not match the specified constructor name, or if the
	// error occurs at a phase that differs from the indicated phase, the test
	// must be interpreted as "failing."
	Negative Negative
	// One or more files whose content must be evaluated in the test realm's
	// global scope prior to test execution. These files are located within the
	// harness/ directory of the Test262 project.
	Includes []string
	// Flags are flags that specify the execution environment further.
	Flags Flag
	// Locale allows tests to declare explicit information regarding locale
	// specificity. Its value is an array of one or more valid language tags or
	// subtags.
	Locale []string
}

// Negative indicates, that these tests are expected to generate an uncaught
// exception. The Phase and Type specify, when and what error is expected.
type Negative struct {
	// Phase is the stage of the test interpretation process that the error is
	// expected to be produced.
	Phase Phase
	// Type is the name of the constructor of the expected error.
	Type string
}

// Flag effectively is a bitmask with 16 bits.
type Flag uint16

const (
	// FlagUnknown is the default value and can be interpreted as "not set".
	FlagUnknown Flag = 1 << iota
	// FlagOnlyStrict indicates that the test must be executed just once, in
	// strict mode, only.
	FlagOnlyStrict
	// FlagNoStrict indicates, that the test must be executed just once, in
	// non-strict mode, only.
	FlagNoStrict
	// FlagModule indicates, that the test source code must be interpreted as
	// module code. In addition, this flag negates the default requirement to
	// execute the test both in strict mode and in non-strict mode. In other
	// words, the transformation described by the section titled "Strict Mode"
	// must not be applied to these tests.
	FlagModule
	// FlagRaw indicates, that the test source code must not be modified in any
	// way, and the test must be executed just once (in non-strict mode, only).
	FlagRaw
	// FlagAsync indicates, that the file harness/doneprintHandle.js must be
	// evaluated in the test realm's global scope prior to test execution. The
	// test must not be considered complete until the implementation-defined
	// print function has been invoked or some length of time has passed without
	// any such invocation. In the event of a passing test run, this function
	// will be invoked with the string 'Test262:AsyncTestComplete'. If invoked
	// with a string that is prefixed with the character sequence
	// Test262:AsyncTestFailure:, the test must be interpreted as failed. The
	// implementation is free to select an appropriate length of time to wait
	// before considering the test "timed out" and failing.
	FlagAsync
	// FlagGenerated indicates, that the test file was created procedurally
	// using the project's test generation tool. This flag is specified for
	// informational purposes only and has no bearing on how the test should be
	// interpreted.
	FlagGenerated
	// FlagCanBlockIsFalse indicates, that the test file should only be run when
	// the [[CanBlock]] property of the Agent Record executing the file is
	// false.
	FlagCanBlockIsFalse
	// FlagCanBlockIsTrue indicates, that the test file should only be run when
	// the [[CanBlock]] property of the Agent Record executing the file is true.
	FlagCanBlockIsTrue
)

// Set sets all flags that are set in the given flag, in this flag.
//
//  f := new(Flag)
//  f.Set(FlagModule)
//  f.Set(FlagAsync)
//  f == FlagModule | FlagAsync // true
func (f *Flag) Set(flag Flag) {
	*f |= flag
}

// IsSet determines if this flag contains all flags set in the given flag.
//
//	f := new(Flag)
//	f.Set(FlagModule)
//	f.Set(FlagAsync)
//	f.IsSet(FlagModule) // true
//	f.IsSet(FlagRaw)    // false
func (f Flag) IsSet(flag Flag) bool {
	return f&flag == flag
}

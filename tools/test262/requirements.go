package test262

type Phase uint8

const (
	PhaseUnknown Phase = iota
	PhaseParse
	PhaseEarly
	PhaseResolution
	PhaseRuntime
)

// Requirements contains all information that were
// included in the metadata of a Test262-testcase.
type Requirements struct {
	Negative Negative
	Includes []string
	Flags    Flag
	Locale   []string
}

type Negative struct {
	Phase Phase
	Type  string
}

type Flag uint16

const (
	FlagUnknown Flag = 1 << iota
	FlagOnlyStrict
	FlagNoStrict
	FlagModule
	FlagRaw
	FlagAsync
	FlagGenerated
	FlagCanBlockIsFalse
	FlagCanBlockIsTrue
)

func (f *Flag) Set(flag Flag) {
	*f |= flag
}

func (f Flag) IsSet(flag Flag) bool {
	return f&flag == flag
}

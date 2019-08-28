package test262

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseHeader(t *testing.T) {
	tests := []struct {
		file    string
		want    Expectation
		wantErr bool
	}{
		{
			"file000.js",
			Expectation{},
			true,
		},
		{
			"file001.js",
			Expectation{
				Negative: Negative{
					Phase: PhaseRuntime,
					Type:  "ReferenceError",
				},
			},
			false,
		},
		{
			"file002.js",
			Expectation{
				Negative: Negative{
					Phase: PhaseParse,
					Type:  "SyntaxError",
				},
			},
			false,
		},
		{
			"file003.js",
			Expectation{
				Negative: Negative{
					Phase: PhaseResolution,
					Type:  "ReferenceError",
				},
				Flags: FlagModule,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			require := require.New(t)

			file, err := os.Open(filepath.Join("testdata", tt.file))
			require.NoError(err)

			got, err := ParseHeader(file)
			if tt.wantErr {
				require.Error(err)
			}
			require.Equal(tt.want, got)
		})
	}
}

func Test_toPhase(t *testing.T) {
	tests := []struct {
		phase string
		want  Phase
	}{
		{"parse", PhaseParse},
		{"early", PhaseEarly},
		{"resolution", PhaseResolution},
		{"runtime", PhaseRuntime},
		// negative tests
		{"Runtime", PhaseUnknown},
		{"parser", PhaseUnknown},
		{"RUNTIME", PhaseUnknown},
		{"PARSE", PhaseUnknown},
		{"unknown", PhaseUnknown},
		{"Unknown", PhaseUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.phase, func(t *testing.T) {
			if got := toPhase(tt.phase); got != tt.want {
				t.Errorf("toPhase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toFlag(t *testing.T) {
	tests := []struct {
		flag string
		want Flag
	}{
		{"onlyStrict", FlagOnlyStrict},
		{"noStrict", FlagNoStrict},
		{"module", FlagModule},
		{"raw", FlagRaw},
		{"async", FlagAsync},
		{"generated", FlagGenerated},
		{"CanBlockIsFalse", FlagCanBlockIsFalse},
		{"CanBlockIsTrue", FlagCanBlockIsTrue},
		// negative tests
		{"canBlockIsFalse", FlagUnknown},
		{"canBlockIsTrue", FlagUnknown},
		{"Generated", FlagUnknown},
		{"GENERATED", FlagUnknown},
		{"OnlyStrict", FlagUnknown},
		{"nostrict", FlagUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.flag, func(t *testing.T) {
			if got := toFlag(tt.flag); got != tt.want {
				t.Errorf("toFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

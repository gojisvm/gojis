package test262

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type sentinel string

func (s sentinel) Error() string { return string(s) }

const (
	// ErrNoTest262Metadata indicates, that a file that was attempted to parse,
	// did not contain a Test262-Metadata header.
	//
	// To see the structure of a test metadata header, please see
	// https://github.com/tc39/test262/blob/master/INTERPRETING.md
	ErrNoTest262Metadata = sentinel("File does not contain a Test262-Metadata header")
)

var (
	magicHeader = []byte("/*---")
	magicFooter = []byte("---*/")
)

type metadata struct {
	Negative struct {
		Phase   string `yaml:"phase"`
		ErrType string `yaml:"type"`
	} `yaml:"negative"`
	Includes []string `yaml:"includes"`
	Flags    []string `yaml:"flags"`
	Locale   []string `yaml:"locale"`
}

func ParseHeader(r io.Reader) (Expectation, error) {
	content, err := ioutil.ReadAll(r) // files are small, so this should be fine
	if err != nil {
		return Expectation{}, fmt.Errorf("read data source: %v", err)
	}

	meta, err := parseMetadata(content)
	if err != nil {
		return Expectation{}, fmt.Errorf("parse metadata: %v", err)
	}

	flags := new(Flag)
	for _, f := range meta.Flags {
		flags.Set(toFlag(f))
	}

	exp := Expectation{
		Negative: Negative{
			Phase: toPhase(meta.Negative.Phase),
			Type:  meta.Negative.ErrType,
		},
		Includes: meta.Includes,
		Flags:    *flags,
		Locale:   meta.Locale,
	}

	return exp, nil
}

func parseMetadata(data []byte) (md metadata, err error) {
	if !(bytes.Contains(data, magicHeader) &&
		bytes.Contains(data, magicFooter)) {
		return metadata{}, ErrNoTest262Metadata
	}

	yamlStart := bytes.Index(data, magicHeader) + len(magicHeader)
	yamlEnd := bytes.Index(data, magicFooter)
	yamlData := bytes.TrimSpace(data[yamlStart:yamlEnd])

	err = yaml.Unmarshal(yamlData, &md)
	if err != nil {
		return metadata{}, fmt.Errorf("unmarshal yaml: %v", err)
	}

	return
}

func toPhase(phase string) Phase {
	switch phase {
	case "parse":
		return PhaseParse
	case "early":
		return PhaseEarly
	case "resolution":
		return PhaseResolution
	case "runtime":
		return PhaseRuntime
	default:
		return PhaseUnknown
	}
}

func toFlag(flag string) Flag {
	switch flag {
	case "onlyStrict":
		return FlagOnlyStrict
	case "noStrict":
		return FlagNoStrict
	case "module":
		return FlagModule
	case "raw":
		return FlagRaw
	case "async":
		return FlagAsync
	case "generated":
		return FlagGenerated
	case "CanBlockIsFalse":
		return FlagCanBlockIsFalse
	case "CanBlockIsTrue":
		return FlagCanBlockIsTrue
	default:
		return FlagUnknown
	}
}

package options

import "fmt"

//noinspection GoSnakeCaseUsage
const MISSING_VALUE = "???-Missing-Entry-???"

func For(sliceLength int, ifEmpty string) *Options {
	return &Options{sliceLength: sliceLength, ifEmpty: ifEmpty}
}

type Options struct {
	sliceLength     int
	ifEmpty         string
	entriesConsumed int
	built           string
}

func (in *Options) Add(value string) {
	in.entriesConsumed++
	switch {
	case in.entriesConsumed > in.sliceLength:
		panic(fmt.Sprintf("Options (%d) full, attempted to add: %s", in.sliceLength, value))
	case in.entriesConsumed < in.sliceLength:
		if in.entriesConsumed != 1 {
			in.built += ", "
		}
		in.built += value

		// in.entriesConsumed == in.sliceLength
	case in.sliceLength == 1:
		in.built += value
	case in.sliceLength == 2:
		in.built += " or " + value
	default: // in.sliceLength > 2
		in.built += ", or " + value
	}
}

func (in *Options) Done() string {
	for in.entriesConsumed < in.sliceLength {
		in.Add(MISSING_VALUE)
	}
	if in.sliceLength == 0 {
		return in.ifEmpty
	}
	return in.built
}

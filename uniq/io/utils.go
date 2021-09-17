package io

type Flags struct {
	CntF          bool
	RepeatF       bool
	NotRepeatF    bool
	CntSkipWordsF int
	CntSkipCharsF int
	RegisterSkipF bool
	FNameIn       string
	FNameOut      string
}

type UniqRes struct {
	Str string
	Cnt int
}

func boolToInt(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

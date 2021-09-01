package uniq

import (
	"math"
	"strings"
	"uniq/read_write"
)

func SkipWords(prev string, cur string, cnt int) (string, string, error) {
	if cnt < 0 {
		return prev, cur, IncorrectArgs
	}
	if cnt == 0 {
		return prev, cur, nil
	}

	tmp := strings.Split(prev, " ")
	cntSkipWords := math.Min(float64(len(tmp)), float64(cnt))
	prev = strings.Join(tmp[int(cntSkipWords):], " ")

	tmp = strings.Split(cur, " ")
	cntSkipWords = math.Min(float64(len(tmp)), float64(cnt))
	cur = strings.Join(tmp[int(cntSkipWords):], " ")

	return prev, cur, nil
}
func SkipChars(prev string, cur string, cnt int) (string, string, error) {
	if cnt < 0 {
		return prev, cur, IncorrectArgs
	}
	if cnt == 0 {
		return prev, cur, nil
	}
	cntSkipChars := math.Min(float64(len(prev)), float64(cnt))
	prev = prev[int(cntSkipChars):]

	cntSkipChars = math.Min(float64(len(cur)), float64(cnt))
	cur = cur[int(cntSkipChars):]

	return prev, cur, nil
}

func GetUniqStrings(src []string, flags read_write.Flags) ([]read_write.UniqRes, error) {
	if len(src) == 0 {
		return []read_write.UniqRes{}, nil
	}

	cnts := make([]read_write.UniqRes, 0)
	var cur string
	var prev string
	var err error

	for idx, val := range src {
		if idx != 0 {
			prev = src[idx-1]
			cur = val
			if flags.RegisterSkipF {
				prev = strings.ToLower(src[idx-1])
				cur = strings.ToLower(cur)
			}
		}
		if flags.CntSkipWordsF != 0 {
			prev, cur, err = SkipWords(prev, cur, flags.CntSkipWordsF)
			if err != nil {
				return []read_write.UniqRes{}, err
			}
		}

		if flags.CntSkipCharsF != 0 {
			prev, cur, err = SkipChars(prev, cur, flags.CntSkipCharsF)
			if err != nil {
				return []read_write.UniqRes{}, err
			}
		}

		if prev != cur || idx == 0 {
			cnts = append(cnts, read_write.UniqRes{Str: val, Cnt: 1})
		} else {
			cnts[len(cnts)-1].Cnt += 1
		}
	}
	return cnts, nil
}
func Uniq(data []string, flags read_write.Flags) ([]read_write.UniqRes, error) {
	var err error
	if len(data) == 0 {
		return []read_write.UniqRes{}, err
	}

	res := make([]read_write.UniqRes, 0)

	res, err = GetUniqStrings(data, flags)

	return res, err

}

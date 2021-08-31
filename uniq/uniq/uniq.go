package uniq

import (
	"strings"
	"uniq/read_write"
)

func getCountStrings(src []string, flags read_write.Flags) ([]read_write.UniqRes, error) {
	cnts := make([]read_write.UniqRes, 0)
	for idx, val := range src {
		var cur string
		var prev string
		if idx != 0 {
			prev = src[idx-1]
			cur = val
			if flags.RegisterSkipF {
				prev = strings.ToLower(src[idx-1])
				cur = strings.ToLower(cur)
			}
		}
		if prev != cur || idx == 0 {
			cnts = append(cnts, read_write.UniqRes{Str: val, Cnt: 1})
		} else {
			cnts[len(cnts)-1].Cnt += 1
		}
	}
	//if idx == 0 || src[idx - 1] != val {
	//	cnts = append(cnts, read_write.UniqRes{Str: val, Cnt: 1})
	//} else {
	//}
	//}
	return cnts, nil
}
func Uniq(data []string, flags read_write.Flags) ([]read_write.UniqRes, error) {
	var err error
	if len(data) == 0 {
		return []read_write.UniqRes{}, err
	}

	//uniqStr := make([]string, 0, len(data))
	//for _, val := range data {
	//	cur := val
	//	idx := len(uniqStr) - 1
	//
	//	if idx == -1 || strings.Compare(uniqStr[idx], cur) != 0 {
	//		uniqStr = append(uniqStr, cur)
	//	}
	//}
	res := make([]read_write.UniqRes, 0)

	res, err = getCountStrings(data, flags)

	return res, err

}

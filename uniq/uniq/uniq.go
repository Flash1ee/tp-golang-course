package uniq

import (
	"strings"
	"uniq/read_write"
)

func Uniq(data []string, flags read_write.Flags) ([]string, error) {
	var err error
	if len(data) == 0 {
		return []string{}, err
	}

	res := make([]string, 0, len(data))
	for _, val := range data {
		cur := val
		idx := len(res) - 1

		if idx == -1 || strings.Compare(res[idx], cur) != 0 {
			res = append(res, cur)
		}
	}
	return res, err

}

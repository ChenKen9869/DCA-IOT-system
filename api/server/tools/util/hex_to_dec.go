package util

import "strconv"

func Hex2Dec(val string) (int, error) {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}
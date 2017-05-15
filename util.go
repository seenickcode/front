package front

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/jmcvetta/randutil"
)

func RandomString() string {
	return RandomStringWithLength(12)
}

func RandomStringWithLength(len int) string {
	rand, _ := randutil.String(12, randutil.Alphabet)
	ts := time.Now().Unix()
	s := TruncateString(fmt.Sprintf("%v%v", rand, ts), len, "")
	return s
}

func TruncateString(s string, i int, filler string) string {
	if len(s) < i {
		return s
	}
	if utf8.ValidString(s[:i]) {
		to := i - len(filler)
		return s[:to] + filler
	}
	return s[:i+1] + filler
}

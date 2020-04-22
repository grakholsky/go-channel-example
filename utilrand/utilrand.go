package utilrand

import (
	"math/rand"
	"strings"
)

const dictionary = "abcdefghijklmnopqrstuvwxyz"

func Str(length int) string {
	b := strings.Builder{}
	b.Grow(length)
	for i := 0; i < length; i++ {
		b.WriteByte(dictionary[Num(len(dictionary))])
	}
	return b.String()
}

func Num(length int) int64 {
	return rand.Int63() % int64(length)
}

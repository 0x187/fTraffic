package randomINT

import (
	"math/rand"
	"time"
)

func RandomINT(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

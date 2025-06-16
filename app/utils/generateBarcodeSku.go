package utils

import (
	"fmt"
	"math/rand"
)

func GenerateBarcodeSKU() string {
	return fmt.Sprintf("8%012d", rand.Int63n(1e12))
}
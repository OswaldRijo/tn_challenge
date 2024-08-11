package utils_test

import (
	"fmt"
	"testing"

	"truenorth/packages/utils"
)

func TestRandStringBytesMask(t *testing.T) {
	s := utils.RandStringBytesMask(64)

	fmt.Printf("%s\n", s)
}

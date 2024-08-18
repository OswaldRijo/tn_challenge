package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"truenorth/packages/utils"
	"truenorth/services/users_service/test"
)

func TestHashString(t *testing.T) {
	readableStr := "SOME STRING"
	encodedStr, err := utils.HashString(readableStr, test.RandSalt)

	assert.NoError(t, err)
	assert.NotEqual(t, readableStr, encodedStr)
}

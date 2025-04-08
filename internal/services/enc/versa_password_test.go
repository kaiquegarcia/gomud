package enc_test

import (
	"gomud/internal/services/enc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VersaPassword(t *testing.T) {
	type TestCase struct {
		Title          string
		Password       string
		ExpectedResult string
		ExpectedErr    error
	}

	testCases := []TestCase{
		{
			Title:          "should return versa encoded password '123456'",
			Password:       "123456",
			ExpectedResult: "9f4b48560204479916e1cb1b49d0fc92f69235ef43908debec601f267fc9223eacef7a8a4b9d7c693a9ee609e9665fb7:7c4a8",
			ExpectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			result, err := enc.VersaPassword(tc.Password)
			assert.Equal(t, tc.ExpectedResult, result)
			assert.Equal(t, tc.ExpectedErr, err)
		})
	}
}

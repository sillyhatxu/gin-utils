package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	code := "TEST_CODE"
	message := "This is a message"
	date := map[string]string{
		"id":   "1",
		"name": "test-name",
	}
	result := New(code, Msg(message), Data(date))
	assert.Equal(t, code, result.Code)
	assert.Equal(t, message, result.Msg)
	assert.Equal(t, date, result.Data)
	assert.Nil(t, result.Extra)
}

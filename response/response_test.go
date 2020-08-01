package response

import (
	"errors"
	"github.com/sillyhatxu/gin-utils/v2/entity"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	type Test struct {
		Id   int
		Name string
	}
	code := "TEST_CODE"
	msg := "have some error"
	data := &Test{Id: 1, Name: "test-name"}
	err := entity.New(code, entity.Msg(msg), entity.Data(data)).Err()
	assert.NotNil(t, err)
	assert.Equal(t, "go-error: code = TEST_CODE ,message = have some error ,data = &{%!s(int=1) test-name}", err.Error())
	e, ok := err.(*entity.Entity)
	assert.Equal(t, true, ok)
	assert.Equal(t, code, e.Code)
	assert.Equal(t, msg, e.Msg)
	assert.Equal(t, data, e.Data)
	assert.Nil(t, e.Extra)
}

func TestError(t *testing.T) {
	code1 := "TEST_CODE"
	msg1 := "unknown error"
	err := NewError(code1, msg1)
	e, ok := FromError(err)
	assert.EqualValues(t, ok, true)
	assert.EqualValues(t, code1, e.Code)
	assert.EqualValues(t, msg1, e.Msg)

	err = errors.New(msg1)
	e, ok = FromError(err)
	assert.EqualValues(t, ok, false)
	assert.EqualValues(t, gincodes.ServerError, e.Code)
	assert.EqualValues(t, msg1, e.Msg)
}

func TestErrorf(t *testing.T) {
	type Test struct {
		Id   int
		Name string
	}
	type ExtraTest struct {
		TotalCount int
	}
	code1 := "TEST_CODE"
	msg1 := "unknown error"
	data1 := &Test{Id: 1, Name: "test-name"}
	extra1 := &ExtraTest{TotalCount: 500}
	err := NewError(code1, msg1, entity.Data(data1), entity.Extra(extra1))
	e, ok := FromError(err)
	assert.EqualValues(t, true, ok)
	assert.EqualValues(t, code1, e.Code)
	assert.EqualValues(t, msg1, e.Msg)
	assert.EqualValues(t, data1, e.Data)
	assert.EqualValues(t, extra1, e.Extra)

	msg2 := "test error"
	err = errors.New(msg2)
	e, ok = FromError(err)
	assert.EqualValues(t, false, ok)
	assert.EqualValues(t, gincodes.ServerError, e.Code)
	assert.EqualValues(t, msg2, e.Msg)
}

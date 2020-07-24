package response

import (
	"errors"
	"fmt"
	"github.com/sillyhatxu/gin-utils/entity"
	"github.com/sillyhatxu/gin-utils/gincodes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	type Test struct {
		Id   int
		Name string
	}
	err := entity.New("TEST_CODE", entity.Msg("have some error"), entity.Data(&Test{Id: 1, Name: "test-name"})).Err()
	fmt.Println(err)

	if e, ok := err.(*entity.Entity); ok {
		fmt.Println("Code:", e.Code)
		fmt.Println(fmt.Sprintf("Data:%#v", e.Data))
		fmt.Println("Message:", e.Msg)
	}
}

func TestError(t *testing.T) {
	err := NewError("TEST_CODE", "unknown error")
	e, ok := FromError(err)
	assert.EqualValues(t, ok, true)
	assert.EqualValues(t, e.Code, "TEST_CODE")
	assert.EqualValues(t, e.Msg, "unknown error")

	err = errors.New("test error")
	e, ok = FromError(err)
	assert.EqualValues(t, ok, false)
	assert.EqualValues(t, e.Code, gincodes.ServerError)
	fmt.Println(e.Msg)
}

func TestErrorf(t *testing.T) {
	type Test struct {
		Id   int
		Name string
	}
	type ExtraTest struct {
		TotalCount int
	}
	err := NewError("TEST_CODE", "unknown error", entity.Data(&Test{Id: 1, Name: "test-name"}), entity.Extra(&ExtraTest{TotalCount: 500}))
	e, ok := FromError(err)
	assert.EqualValues(t, ok, true)
	assert.EqualValues(t, e.Code, "TEST_CODE")
	assert.EqualValues(t, e.Msg, "unknown error")
	fmt.Println(fmt.Sprintf("e : %#v", e))
	fmt.Println(fmt.Sprintf("Data : %#v", e.Data))
	fmt.Println(fmt.Sprintf("Extra : %#v", e.Extra))

	err = errors.New("test error")
	e, ok = FromError(err)
	assert.EqualValues(t, ok, false)
	assert.EqualValues(t, e.Code, gincodes.ServerError)
	fmt.Println(e.Msg)
}

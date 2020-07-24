package entity

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var data interface{}
	test := New("code", Msg("msg"), Data(data))
	fmt.Println(fmt.Sprintf("%#v", test))
}

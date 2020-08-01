package response

import (
	"github.com/sillyhatxu/gin-utils/v2/entity"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
)

func NewError(code, msg string, opts ...entity.Option) error {
	opts = append(opts, func(e *entity.Entity) {
		e.Msg = msg
	})
	return entity.New(code, opts...).Err()
}

func Success(opts ...entity.Option) *entity.Entity {
	return entity.New(gincodes.OK, opts...)
}

func Error(err error) *entity.Entity {
	return Convert(err)
}

func Errorf(code string, err error) *entity.Entity {
	e := Convert(err)
	e.Code = code
	return e
}

func Convert(err error) *entity.Entity {
	s, _ := FromError(err)
	return s
}

func FromError(err error) (e *entity.Entity, ok bool) {
	if err == nil {
		return nil, false
	}
	if e, ok := err.(*entity.Entity); ok {
		return e, true
	}
	return entity.New(gincodes.ServerError, entity.Msg(err.Error())), false
}

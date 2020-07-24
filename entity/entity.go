package entity

import "fmt"

type Entity struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"message"`
	Extra interface{} `json:"extra"`
}

func (e Entity) Error() string {
	if e.Data == nil && e.Extra == nil {
		return fmt.Sprintf("go-error: code = %s ,message = %s", e.Code, e.Msg)
	} else if e.Extra == nil {
		return fmt.Sprintf("go-error: code = %s ,message = %s ,data = %s", e.Code, e.Msg, e.Data)
	} else if e.Data == nil {
		return fmt.Sprintf("go-error: code = %s ,message = %s ,extra = %s", e.Code, e.Msg, e.Extra)
	} else {
		return fmt.Sprintf("go-error: code = %s ,message = %s ,data = %s ,extra = %s", e.Code, e.Msg, e.Data, e.Extra)
	}
}

func (e *Entity) Err() error {
	return e
}

type Option func(*Entity)

func Msg(msg string) Option {
	return func(e *Entity) {
		e.Msg = msg
	}
}

func Data(data interface{}) Option {
	return func(e *Entity) {
		e.Data = data
	}
}

func Extra(extra interface{}) Option {
	return func(e *Entity) {
		e.Extra = extra
	}
}

func New(code string, opts ...Option) *Entity {
	//default
	entity := &Entity{
		Code: code,
		Msg:  "",
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(entity)
	}
	return entity
}

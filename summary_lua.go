package service

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

func (su *Summary) pipeL(L *lua.LState) int {
	pip := pipe.NewByLua(L, pipe.Env(xEnv))

	for _, s := range su.ssv {
		pip.Do(s, L, func(err error) {
			L.RaiseError("pipe call fail %v", err)
		})
	}
	return 0
}

func (su *Summary) Meta(L *lua.LState, key lua.LValue) lua.LValue {
	switch key.Type() {
	case lua.LTString:
		return su.Index(L, key.String())
	case lua.LTInt:
		n := int(key.(lua.LInt))
		if n-1 < 0 || n-1 >= len(su.ssv) {
			return lua.LNil
		}
		return su.ssv[n-1]
	}

	return lua.LNil
}

func (su *Summary) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "size":
		return lua.LInt(len(su.ssv))

	case "pipe":
		return lua.NewFunction(su.pipeL)

	}

	return lua.LNil
}

package service

import (
	"github.com/vela-ssoc/vela-kit/vela"
	cond "github.com/vela-ssoc/vela-cond"
	"github.com/vela-ssoc/vela-kit/lua"
)

var xEnv vela.Environment

func lookupL(L *lua.LState) int {
	cnd := cond.CheckMany(L)
	su := New()
	su.collect(cnd)
	L.Push(su)
	return 1
}

func indexL(L *lua.LState, key string) lua.LValue {
	cnd := cond.New("name = " + key)
	su := New()
	su.collect(cnd)
	return su
}

func WithEnv(env vela.Environment) {
	xEnv = env
	xEnv.Set("service", lua.NewExport("vela.service.export", lua.WithFunc(lookupL), lua.WithIndex(indexL)))
}

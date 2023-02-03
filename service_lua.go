package service

import "github.com/vela-ssoc/vela-kit/lua"

func (s *Service) String() string                         { return lua.B2S(s.Byte()) }
func (s *Service) Type() lua.LValueType                   { return lua.LTObject }
func (s *Service) AssertFloat64() (float64, bool)         { return 0, false }
func (s *Service) AssertString() (string, bool)           { return "", false }
func (s *Service) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (s *Service) Peek() lua.LValue                       { return s }

func (s *Service) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "name":
		return lua.S2L(s.Name)
	case "start_type":
		return lua.S2L(s.StartType)
	case "exec_path":
		return lua.S2L(s.ExecPath)
	case "display":
		return lua.S2L(s.DisplayName)
	case "description":
		return lua.S2L(s.Description)
	case "state":
		return lua.S2L(s.State)
	case "pid":
		return lua.LInt(s.Pid)
	case "exit_code":
		return lua.LInt(s.ExitCode)
	}

	return lua.LNil
}

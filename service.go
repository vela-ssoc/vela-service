package service

import "github.com/bytedance/sonic"

type Service struct {
	Name        string `json:"name"`
	StartType   string `json:"start_type"`
	ExecPath    string `json:"exec_path"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	State       string `json:"state"`
	Pid         uint32 `json:"pid"`
	ExitCode    uint32 `json:"exit_code"`
}

func (s *Service) Byte() []byte {
	chunk, _ := sonic.Marshal(s)
	return chunk
}

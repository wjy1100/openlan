package ctrlc

import (
	"github.com/danieldin95/openlan-go/src/libol"
	"github.com/danieldin95/openlan-go/src/olctl/libctrl"
)

type Hello struct {
	libctrl.Listen
	cc *CtrlC
}

func (h *Hello) GetCtl(id string, m libctrl.Message) error {
	libol.Cmd("Hello.GetCtl %s %s", id, m.Data)
	return nil
}

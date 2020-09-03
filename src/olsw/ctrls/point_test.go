package ctrls

import (
	"github.com/danieldin95/openlan-go/src/libol"
	"github.com/danieldin95/openlan-go/src/models"
	"github.com/danieldin95/openlan-go/src/olsw/storage"
	"testing"
	"time"
)

func TestCtl_Point(t *testing.T) {
	libol.SetLog(libol.STACK)
	cc := &CtrlC{
		Url:      "http://localhost:10088/ctrl",
		Name:     "admin",
		Password: "123",
	}
	err := cc.Open()
	if err != nil {
		t.Error(err)
		return
	}
	cc.Start()
	point := models.Point{
		UUID:   "xxxx",
		Alias:  "alias-test",
		Client: libol.NewTcpClient("xx", nil),
	}
	storage.Point.Add(&point)
	time.Sleep(5 * time.Second)
	storage.Point.Del(point.Client.Address())
	time.Sleep(5 * time.Second)
	cc.Stop()
	time.Sleep(5 * time.Second)
}
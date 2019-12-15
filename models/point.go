package models

import (
	"github.com/lightstar-dev/openlan-go/libol"
)

type Point struct {
	Alias  string
	Client *libol.TcpClient
	Device *TapDevice
}

func NewPoint(c *libol.TcpClient, d *TapDevice) (w *Point) {
	w = &Point{
		Client: c,
		Device: d,
	}

	return
}
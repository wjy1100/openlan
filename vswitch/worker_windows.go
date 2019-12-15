package vswitch

import (
	"github.com/lightstar-dev/openlan-go/config"
	"github.com/lightstar-dev/openlan-go/libol"
	"github.com/lightstar-dev/openlan-go/models"
)

type Worker struct {
	*WorkerBase
	Br *Bridger
}

func NewWorker(server *libol.TcpServer, c *config.VSwitch) *Worker {
	w := &Worker{
		WorkerBase: NewWorkerBase(server, c),
		Br:         NewBridger(c.BrName, c.IfMtu),
	}
	if w.Br.Name == "" {
		w.Br.Name = w.BrName()
	}

	w.Init(w)

	return w
}

func (w *Worker) NewBr() {
	w.Br.Open(w.Conf.IfAddr)
}

func (w *Worker) FreeBr() {
	w.Br.Close()
}

func (w *Worker) NewTap() (*models.TapDevice, error) {
	//TODO
	libol.Warn("Worker.NewTap: TODO")
	return nil, nil
}

func (w *Worker) Start() {
	w.NewBr()
	w.WorkerBase.Start()
}

func (w *Worker) Stop() {
	w.WorkerBase.Stop()
	w.FreeBr()
}
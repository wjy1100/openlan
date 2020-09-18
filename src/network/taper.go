package network

import (
	"fmt"
	"sync"
)

type Taper interface {
	Type() string
	IsTun() bool
	Name() string
	Read([]byte) (int, error)  // read data from kernel to user space
	Write([]byte) (int, error) // write data from user space to kernel
	Send([]byte) (int, error)  // send data from virtual bridge to kernel
	Recv([]byte) (int, error)  // recv data from kernel to virtual bridge
	Close() error
	Slave(br Bridger)
	Up()
	Down()
	Tenant() string
	Mtu() int
	SetMtu(mtu int)
}

func NewTaper(tenant string, c TapConfig) (Taper, error) {
	if c.Provider == ProviderVir {
		return NewVirtualTap(tenant, c)
	}
	return NewKernelTap(tenant, c)
}

type tapers struct {
	lock    sync.RWMutex
	index   int
	devices map[string]Taper
}

func (t *tapers) GenName() string {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.index++
	return fmt.Sprintf("vir%d", t.index)
}

func (t *tapers) Add(tap Taper) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.devices == nil {
		t.devices = make(map[string]Taper, 1024)
	}
	t.devices[tap.Name()] = tap
}

func (t *tapers) Get(name string) Taper {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t.devices == nil {
		return nil
	}
	if t, ok := t.devices[name]; ok {
		return t
	}
	return nil
}

func (t *tapers) Del(name string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.devices == nil {
		return
	}
	if _, ok := t.devices[name]; ok {
		delete(t.devices, name)
	}
}

func (t *tapers) List() <-chan Taper {
	data := make(chan Taper, 32)
	go func() {
		t.lock.RLock()
		defer t.lock.RUnlock()
		for _, obj := range t.devices {
			data <- obj
		}
		data <- nil
	}()
	return data
}

var Taps = &tapers{}

const (
	_ = iota
	TUN
	TAP
)

type TapConfig struct {
	Provider string
	Type     int
	Network  string
	Name     string
	WriteBuf int
	SendBuf  int
}

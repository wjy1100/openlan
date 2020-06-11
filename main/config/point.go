package config

import (
	"flag"
	"github.com/danieldin95/openlan-go/libol"
	"runtime"
)

type Interface struct {
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Mtu      int    `json:"mtu" yaml:"mtu"`
	Address  string `json:"address,omitempty" yaml:"address,omitempty"`
	Bridge   string `json:"bridge,omitempty" yaml:"bridge,omitempty"`
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`
}

type Point struct {
	Alias       string    `json:"name,omitempty" yaml:"name,omitempty"`
	Network     string    `json:"network,omitempty" yaml:"network,omitempty"`
	Addr        string    `json:"connection" yaml:"connection"`
	Timeout     int       `json:"timeout"`
	Username    string    `json:"username,omitempty" yaml:"username,omitempty"`
	Password    string    `json:"password,omitempty" yaml:"password,omitempty"`
	Protocol    string    `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Intf        Interface `json:"interface" yaml:"interface"`
	Log         Log       `json:"log" yaml:"log"`
	Http        *Http     `json:"http,omitempty" yaml:"http,omitempty"`
	RequestAddr bool      `json:"-" yaml:"-"`
	SaveFile    string    `json:"-" yaml:"-"`
}

var pd = Point{
	Alias:    "",
	Addr:     "openlan.net",
	Protocol: "tls", // udp, kcp, tcp, tls, ws and wss etc.
	Timeout:  30,
	Log: Log{
		File:    "./point.log",
		Verbose: libol.INFO,
	},
	Intf: Interface{
		Mtu:      1518,
		Provider: "tap",
		Name:     "",
	},
	Http: &Http{
		Listen: "0.0.0.0:10001",
	},
	SaveFile:    "./point.json",
	Network:     "default",
	RequestAddr: true,
}

func NewPoint() (c *Point) {
	c = &Point{
		Http:        &Http{},
		RequestAddr: true,
	}
	flag.StringVar(&c.Alias, "alias", pd.Alias, "Alias for this point")
	flag.StringVar(&c.Network, "net", pd.Network, "Network name")
	flag.StringVar(&c.Addr, "conn", pd.Addr, "Virtual switch connect to")
	flag.StringVar(&c.Username, "user", pd.Username, "Accessed username")
	flag.StringVar(&c.Password, "pass", pd.Password, "Accessed password")
	flag.StringVar(&c.Protocol, "proto", pd.Protocol, "Connection protocol")
	flag.IntVar(&c.Timeout, "timeout", pd.Timeout, "Time in secs socket dead")
	flag.IntVar(&c.Log.Verbose, "log:level", pd.Log.Verbose, "Log level")
	flag.StringVar(&c.Log.File, "log:file", pd.Log.File, "Log saved to file")
	flag.StringVar(&c.Intf.Name, "if:name", pd.Intf.Name, "Configure interface name")
	flag.StringVar(&c.Intf.Address, "if:addr", pd.Intf.Address, "Configure interface address")
	flag.StringVar(&c.Intf.Bridge, "if:br", pd.Intf.Bridge, "Configure bridge name")
	flag.StringVar(&c.Intf.Provider, "if:provider", pd.Intf.Provider, "Interface provider")
	flag.StringVar(&c.Http.Listen, "http:listen", pd.Http.Listen, "Http listen on")
	flag.StringVar(&c.SaveFile, "conf", pd.SaveFile, "the configuration file")
	flag.Parse()

	if err := c.Load(); err != nil {
		libol.Warn("NewPoint.load %s", err)
	}
	c.Default()
	libol.Init(c.Log.File, c.Log.Verbose)
	return c
}

func (c *Point) Right() {
	if c.Alias == "" {
		c.Alias = GetAlias()
	}
	RightAddr(&c.Addr, 10002)
	if runtime.GOOS == "darwin" {
		c.Intf.Provider = "tun"
	}
}

func (c *Point) Default() {
	c.Right()

	//reset zero value to default
	if c.Addr == "" {
		c.Addr = pd.Addr
	}
	if c.Intf.Mtu == 0 {
		c.Intf.Mtu = pd.Intf.Mtu
	}
	if c.Timeout == 0 {
		c.Timeout = pd.Timeout
	}
}

func (c *Point) Load() error {
	return libol.UnmarshalLoad(c, c.SaveFile)
}

func init() {
	pd.Right()
}

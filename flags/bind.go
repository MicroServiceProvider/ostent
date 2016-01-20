package flags

import (
	"fmt"
	"net"
	"strings"
)

type Bind struct {
	string
	defport string // const
	Host    string // available after flag.Parse()
	Port    string // available after flag.Parse()
}

func NewBind(defport int) Bind {
	b := Bind{defport: fmt.Sprintf("%d", defport)}
	b.Set("")
	return b
}

// satisfying flag.Value interface
func (b Bind) String() string { return string(b.string) }
func (b *Bind) Set(input string) error {
	if input == "" {
		input = ":" + b.defport
	}
	var err error
	b.Host, b.Port, err = net.SplitHostPort(input)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "missing port in address") {
			return err
		}
		b.Host, b.Port = input, b.defport
	}
	if _, err = net.LookupPort("tcp", b.Port); err != nil {
		return err
	}
	b.string = b.Host + ":" + b.Port
	return nil
}

// Type of pflag.Value interface
func (b Bind) Type() string { return "bind" }
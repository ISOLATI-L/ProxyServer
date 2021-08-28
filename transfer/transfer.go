package transfer

import (
	"io"
)

type OneWayTransferor struct {
	Destination io.WriteCloser
	Source      io.ReadCloser
}

func (t OneWayTransferor) Start() {
	go transfer(t.Destination, t.Source)
}

type TwoWayTransferor struct {
	Stream1 io.ReadWriteCloser
	Stream2 io.ReadWriteCloser
}

func (t TwoWayTransferor) Start() {
	go transfer(t.Stream1, t.Stream2)
	go transfer(t.Stream2, t.Stream1)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

//go:build !windows
// +build !windows

package main

import (
	"fmt"

	_ "github.com/abmpio/webserver/starter/pprof"
)

func init() {
	fmt.Printf("plugin healthcheck init function called\r\n")
}

type Bootstrap struct {
}

func newBootstrap() Bootstrap {
	b := Bootstrap{}
	return b
}

func (b Bootstrap) BootstrapPlugin() (err error) {
	return nil
}

var PluginBootstrap = newBootstrap()

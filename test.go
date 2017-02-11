package main

import (
	"honnef.co/go/augeas"

	"fmt"
)

func main() {
	ag, err := augeas.New("/", "", augeas.None)
	if err != nil {
		panic(err)
	}

	err = ag.Set("/augeas/load/Xmgmt/lens", "Sshd.lns")
	if err != nil {
		panic("err on set lens")
	}
	err = ag.Set("/augeas/load/Xmgmt/incl", "/tmp/foo/bar")
	if err != nil {
		panic("err on set incl")
	}
	err = ag.Load()
	if err != nil {
		panic("err on load")
	}
	err = ag.Set("/files/tmp/foo/bar/X11Forwarding", "no")
	if err != nil {
		panic("err on set x11forward")
	}
	err = ag.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

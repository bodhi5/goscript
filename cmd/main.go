package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bodhi5/goscript/goscript"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: goscript [code]")
	os.Exit(2)
}

func Exec() {
	if len(os.Args) < 2 {
		usage()
	}
	code := os.Args[1]
	if code == "" {
		usage()
	}

	gs, _ := goscript.NewFromString(code, os.Args[1:]...)
	gs.Stderr = os.Stderr
	gs.Stdin = os.Stdin
	gs.Stdout = os.Stdout

	defer gs.Clean()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		gs.Clean()
		os.Exit(1)
	}()

	if err := gs.Run(); err != nil {
		log.Fatal(err)
	}
}

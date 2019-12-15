package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

const (
	version = "0.0.1"
)

var (
	// ErrShowVersion returns when set version flag.
	ErrShowVersion = errors.New("show version")
)

type lgen struct {
	params               Params
	template, dist       string
	outStream, errStream io.Writer
}

// Params includes template parameters.
type Params struct {
	Action string
	Model  string
}

func fill(args []string, outStream, errStream io.Writer) (*lgen, error) {
	var v bool
	var a, m, t, d string
	vdesc := "print version information and quit."
	adesc := "action name"
	mdesc := "model name"
	cn := args[0]
	flags := flag.NewFlagSet(cn, flag.ContinueOnError)
	flags.SetOutput(errStream)

	flags.BoolVar(&v, "version", false, vdesc)
	flags.BoolVar(&v, "v", false, vdesc)
	flags.StringVar(&a, "action", "", adesc)
	flags.StringVar(&a, "a", "", adesc)
	flags.StringVar(&m, "model", "", mdesc)
	flags.StringVar(&m, "m", "", mdesc)
	tdesc := "templates directory"
	flags.StringVar(&t, "template", "", tdesc)
	flags.StringVar(&t, "t", "./templates", tdesc)
	ddesc := "output directory"
	flags.StringVar(&d, "dist", "", ddesc)
	flags.StringVar(&d, "d", "./", ddesc)

	if err := flags.Parse(args[1:]); err != nil {
		return nil, err
	}

	if v {
		fmt.Fprintf(errStream, "%s version %s\n", cn, version)
		return nil, ErrShowVersion
	}
	if len(a) == 0 || len(m) == 0 {
		msg := "need to set action name and model name"
		return nil, fmt.Errorf(msg)
	}

	nargs := flags.Args()
	if len(nargs) > 0 {
		msg := "non-flag option must be zero."
		return nil, fmt.Errorf(msg)
	}
	return &lgen{
		params: Params{
			Action: a,
			Model:  "m",
		},
		template:  "",
		dist:      "",
		outStream: outStream,
		errStream: errStream,
	}, nil
}

func (l *lgen) run() error {
	// TODO: load templates in directory.
	// TODO: compile template.
	// TODO: build saved file path.
	// TODO: create saved directory.
	// TODO: write directory.
	return nil
}

// Run is entry point.
func Run(args []string, outStream, errStream io.Writer) error {
	lgen, err := fill(args, outStream, errStream)
	if err != nil {
		return err
	}
	fmt.Printf("%v", lgen)
	return lgen.run()
}

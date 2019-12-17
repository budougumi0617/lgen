package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	var err error
	t, err = filepath.Abs(t)
	if err != nil {
		panic(err)
	}
	d, err = filepath.Abs(d)
	if err != nil {
		panic(err)
	}
	return &lgen{
		params: Params{
			Action: a,
			Model:  "m",
		},
		template:  t,
		dist:      d,
		outStream: outStream,
		errStream: errStream,
	}, nil
}

func (l *lgen) run() error {
	return filepath.Walk(l.template, l.walk)
}

var fmap = template.FuncMap{
	"title": strings.Title,
}

// make file name with action and model.
func (l *lgen) buildFileName(base string) string {
	return strings.ToLower(strings.Join([]string{l.params.Action, l.params.Model, base}, "_"))
}

func (l *lgen) walk(path string, info os.FileInfo, err error) error {
	p, err := filepath.Rel(l.template, path)
	if err != nil {
		panic(err)
	}
	fp := filepath.Join(l.dist, p)

	if info.IsDir() {
		// make same directory structure in distribution.
		if err := os.MkdirAll(fp, 0777); err != nil {
			panic(err)
		}
		return nil
	}
	dn, fn := filepath.Split(fp)
	sp := filepath.Join(dn, l.buildFileName(fn))

	buf := bytes.Buffer{}
	// TODO: load template from file.
	tmpl := `action = "{{- .Action }}"
model: "{{- .Model }}"
`
	dtmpl = string(b)
	if err := template.Must(template.New(sp).Funcs(fmap).Parse(dtmpl)).Execute(&buf, l.params); err != nil {
		panic(err)
	}

	// execute gofmt
	codes, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Printf("l.params.Action = %q\n", l.params.Action)
	fmt.Printf("codes = %q\n", codes)
	fmt.Printf("codes = %q\n", codes)
	// TODO: Need warning if overwrite file?
	f, err := os.Create(sp)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	if _, err = f.Write(codes); err != nil {
		panic(err)
	}
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

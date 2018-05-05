package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/luc10/zykgen"
	"os"
)

const usage = `Zyxel VMG8823-B50B WPA Keygen

Usage:
  zykgen (-m|-n|-c) [-l <length>] <serial>
  zykgen -h | --help

Options:
  -l <length>     Output key length [default: 10].
  -h --help       Show this screen.`

func main() {
	var cocktail zykgen.Cocktail
	var args struct {
		Serial       string `docopt:"<serial>"`
		Length       int    `docopt:"-l"`
		Mojito       bool   `docopt:"-m"`
		Negroni      bool   `docopt:"-n"`
		Cosmopolitan bool   `docopt:"-c"`
	}

	opts, err := docopt.DefaultParser.ParseArgs(usage, os.Args[1:], "")
	if err != nil {
		return
	}

	opts.Bind(&args)
	if args.Mojito {
		cocktail = zykgen.Mojito
	}
	if args.Negroni {
		cocktail = zykgen.Negroni
	}
	if args.Cosmopolitan {
		cocktail = zykgen.Cosmopolitan
	}

	fmt.Println(zykgen.Wpa(args.Serial, args.Length, cocktail))
}

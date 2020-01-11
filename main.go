package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/srfirouzi/rsrc/rsrc"
)

var usage = `USAGE:

rsrc [-i main.rc] [-o rsrc.syso] [-arch amd64]
  Generates a .syso file with specified resources embedded in .rsrc section,
  aimed for consumption by Go linker when building Win32 excecutables.

The generated *.syso files should get automatically recognized by 'go build'
command and linked into an executable/library, as long as there are any *.go
files in the same directory.

OPTIONS:
`

func main() {
	//FIXME: verify that data file size doesn't exceed uint32 max value
	var fnamein, fnameout, arch string
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.StringVar(&fnamein, "i", "main.rc", "path of input rc file")
	flags.StringVar(&fnameout, "o", "rsrc.syso", "name of output COFF (.res or .syso) file")
	flags.StringVar(&arch, "arch", "amd64", "architecture of output file - one of: 386,amd64 ")
	_ = flags.Parse(os.Args[1:])
	data, err := ioutil.ReadFile(fnamein)
	if (arch != "386" && arch != "amd64") || err != nil {
		if err != nil {
			fmt.Print("Error: ")
			fmt.Println(err.Error())

		}
		fmt.Fprintln(os.Stderr, usage)
		flags.PrintDefaults()
		os.Exit(1)
	}
	rcfile := rsrc.NewRCFile(arch, fnameout)
	rcfile.AddLines(strings.Split(string(data), "\n"))
	err = rsrc.EmbedRCFile(rcfile)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

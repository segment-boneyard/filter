package main

import "github.com/robertkrimen/otto"
import "github.com/tj/docopt"
import "encoding/json"
import "io"
import "os"

// Version.
var Version = "0.0.1"

// Usage.
var Usage = `
  Usage: 
    filter <js>

    filter -h | --help
    filter -v | --version

  Options:
    -h, --help      show help information
    -v, --version   show version information

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	check(err)

	js := args["<js>"].(string)
	fn := "(function(){ " + js + " })()"
	vm := otto.New()

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v interface{}
		err := dec.Decode(&v)

		if err == io.EOF {
			return
		}

		check(err)

		vm.Set("j", v)
		val, err := vm.Run(fn)
		check(err)

		if val.IsDefined() {
			v, err := val.Export()
			check(err)
			enc.Encode(v)
		}
	}
}

func check(err error) {
	if err != nil {
		println("filter: " + err.Error())
		os.Exit(1)
	}
}

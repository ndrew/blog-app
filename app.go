package main

/**
             _
  _ __   __| |_ ____      __
 | '_ \ / _` | '__\ \ /\ / /
 | | | | (_| | |   \ V  V /
 |_| |_|\__,_|_|    \_/\_/

------------------------------

CLI for blog generator

*/

import (
	"./lib"
	"encoding/json"
	"flag"
	"fmt"
	engine "github.com/ndrew/stagosaurus"
	"io/ioutil"
	"path/filepath"
	"sort"
)

//
//
func printHeader() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("╟ ░░░░░░░░░░░ BLOG GENERATOR ░░░░░░░░░░░ ╢")
	fmt.Println("╚════════════════════════════════════════╝")
}

//
//
func printUsage() {
	fmt.Println("Usage: blog [--config <cfg-file>] [--help]")
	fmt.Println("             <command> [<args>]")
	fmt.Println("──────────────────────────────────────────")
	fmt.Println("Available commands:")

}

//
//
func listCommands(prettyprint bool) {
	cmds := blog.AvailableCommands()

	var keys []string
	for k := range cmds {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, command := range keys {
		if prettyprint {
			fmt.Println("  ■", command, cmds[command])
		} else {
			fmt.Println(command)
		}
	}
}

//
//
func readConfig(path string, defaults engine.Config) *engine.Config {
	cfg := engine.NewConfig(defaults)

	var realpath, err = filepath.Abs(path)
	if err != nil {
		return &cfg
	}

	source, err := ioutil.ReadFile(realpath)
	if err != nil {
		fmt.Printf("ERR: Config file '%v' is not found.\n", realpath)
		return &cfg
	}

	//
	// JSON stuff
	//
	// TODO: move config reading to stago lib
	var data map[string]*json.RawMessage
	err = json.Unmarshal(source, &data)

	if err != nil || len(data) == 0 {
		fmt.Printf("ERR: can't parse JSON from '%v'\n", realpath)
		return &cfg
	}

	for k, v := range data {
		var value interface{}
		err = json.Unmarshal(*v, &value)

		if err == nil {
			cfg.Set(k, value)
		} else {
			// does this really occurs?
			fmt.Printf("ERR: couldn't interpret json, '%v':%v \n", k, *v)
		}
	}

	return &cfg
}

//
// stagosaurus cli
//
func main() {
	var configFile = ""
	var help = false

	flag.StringVar(&configFile, "config", "default.cfg", "")
	flag.BoolVar(&help, "help", false, "") // to override the annoying behaviour on --help
	flag.Parse()

	var args = flag.Args()   // shift params for action handlers
	var action = flag.Arg(0) // internal bash stuff in order to get autocomplete for terminal

	if action == "autocomplete" {
		listCommands(false)
		return
	}

	printHeader()

	if help || len(args) == 0 {
		// TODO: --help <command>
		printUsage()
		listCommands(true)
		return
	}

	defaults := engine.EmptyConfig()
	// TODO: add hardcoded defaults

	config := readConfig(configFile, defaults)
	params := args[1:len(args)]

	result, err := blog.Workflow(*config, action, params)

	if err != nil {
		fmt.Println("ERR:\n")
		fmt.Println(err)
	}

	fmt.Println(result)

}

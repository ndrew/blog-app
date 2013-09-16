/**
             _
  _ __   __| |_ ____      __
 | '_ \ / _` | '__\ \ /\ / /
 | | | | (_| | |   \ V  V /
 |_| |_|\__,_|_|    \_/\_/

------------------------------

Static blog generator for http://sernyak.com

*/
package main

import (
	"flag"
	"fmt"
	//"github.com/ndrew/stagosaurus"
	"path/filepath"
	"sort"
)

var COMMANDS = map[string]string{
	"new":     "<post-name> [<params>] - creates new post and opens editor",
	"edit":    "<post-name>           - opens post in editor",
	"publish": "[<post-name>]      - renders markdown posts to html"}

func printHeader() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("╟ ░░░░░░░░░░░ BLOG GENERATOR ░░░░░░░░░░░ ╢")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println("Usage: blog [--config <cfg-file>] [--help]")
	fmt.Println("             <command> [<args>]")
	fmt.Println("──────────────────────────────────────────")
}

func listCommands(full_description bool) {
	if full_description {
		fmt.Println("Available commands:")
	}

	var keys []string
	for k := range COMMANDS {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, command := range keys {
		if full_description {
			fmt.Println(" ■", command, COMMANDS[command])
		} else {
			// TODO: params
			fmt.Println(command)
		}
	}
}

/*func newPost(args []string, engine *stagosaurus.Engine) {
	// do something here
	fmt.Println("new post!")
	editPost(args, engine)
}

func editPost(args []string, engine *stagosaurus.Engine) {
	// do something here
	fmt.Println("edit post!")
}

func publishPosts(args []string, engine *stagosaurus.Engine) {
	// do something here
	fmt.Println("publish!")
}*/

func main() {

	// TODO: loading from filesystem or --config
	/*
		    cfg := new(stagosaurus.Config)
			cfg.BaseUrl = "http://localhost:666/blog/"

			renderingStrategy := new(stagosaurus.RenderingStrategy)

			postsFactory := new(stagosaurus.FolderPostFactory)
			postsFactory.PostsDir = "/Users/ndrw/Desktop/dev/site/blog/posts"

			posts := postsFactory.GetPosts()

			engine := stagosaurus.New(cfg, renderingStrategy, posts)
			// engine.RunServer(".", "bla-bla")
	*/

	var configFile = ""
	flag.StringVar(&configFile, "config", "default.cfg", "help message for flagname")
	flag.Parse()

	var configFileAbs, err = filepath.Abs(configFile)

	if err != nil {
		println(configFileAbs)
	}
	/*var metadata *stagosaurus.AppCfg = &stagosaurus.AppCfg{Foo: "tttt"}
	if err != nil {
		metadata = &stagosaurus.AppCfg{Foo: "xyz"}
	} else {
		metadata.ReadConfig(configFileAbs)
	}*/

	//fmt.Println(metadata.Foo)

	var action = flag.Arg(0)
	if action == "autocomplete" {
		listCommands(false)
		return
	}

	var args = flag.Args()
	var commandParams = []string{}
	if len(args) > 1 {
		commandParams = args[1:len(args)]

		println(commandParams)
	}

	switch {
	/*case action == "new":
		newPost(commandParams, engine)
	case action == "edit":
		editPost(commandParams, engine)
	case action == "publish":
		publishPosts(commandParams, engine)
	*/
	default:
		{
			printHeader()
			listCommands(true)
		}
	}
}

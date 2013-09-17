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
	blog "github.com/ndrew/stagosaurus"
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

func getConfig(configFile string) *blog.Config {
	cfg := new(blog.Config)
	// set defaults
	cfg.BaseUrl = "http://localhost:666/blog/"

	// read config if needed
	var configFileAbs, err = filepath.Abs(configFile)
	if err == nil {
		err = cfg.ReadConfig(configFileAbs)
		if err != nil {
			println(err)
		}
	} else {
		println("Can't load config " + configFileAbs)
	}
	return cfg
}

func main() {

	var configFile = ""
	flag.StringVar(&configFile, "config", "default.cfg", "help message for flagname")
	flag.Parse()

	config := getConfig(configFile)

	renderingStrategy := new(blog.RenderingStrategy)
	postsFactory := new(blog.FileSystem)
	// todo: do this via config
	postsFactory.PostsDir = "/Users/ndrw/Desktop/dev/site/blog/posts"

	engine := blog.New(config, postsFactory, renderingStrategy, nil) // todo: add depoloyer

	// internal cmd app stuff
	var action = flag.Arg(0)
	if action == "autocomplete" {
		listCommands(false)
		return
	}

	// shift params for action handlers
	var args = flag.Args()
	var actionParams = []string{}
	if len(args) > 1 {
		actionParams = args[1:len(args)]
	}

	// debug output
	fmt.Printf("action %v, params %v \n", action, actionParams)
	fmt.Printf("engine %v \n", engine)

	switch {
	/*case action == "new":
	      newPost(actionParams, engine)
	  case action == "edit":
	      editPost(actionParams, engine)
	  case action == "publish":
	      publishPosts(actionParams, engine)
	*/
	default:
		{
			printHeader()
			listCommands(true)
		}
	}
}

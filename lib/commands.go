/**
             _
  _ __   __| |_ ____      __
 | '_ \ / _` | '__\ \ /\ / /
 | | | | (_| | |   \ V  V /
 |_| |_|\__,_|_|    \_/\_/

------------------------------

Static blog generator for http://sernyak.com

*/
package blog

import (
	stago "./../../stagosaurus" //"github.com/ndrew/stagosaurus"
)

/// commands go here

// fit this to some kind of workflow
//
func build(cfg stago.Config, args []string) (string, error) {
	generator, err := New()
	if err != nil {
		return "can't init stagosaurus", err
	}
	err = generator.BuildAll(cfg)
	if err != nil {
		return "Build error", err
	}
	return "Build completed.", nil
}

///////////////////////
// command api stuff
//	 <?> extend this to an rpc api
//

var COMMANDS = map[string]func(stago.Config, []string) (string, error){
	"publish": build,
}

var DESCRIPTIONS = map[string]string{
	//"new":     "<post-name> [<params>] - creates new post and opens editor",
	//"edit":    "<post-name>           - opens post in editor",
	//"publish": "[<post-name>]      - renders markdown posts to html",
	"publish": "- generates html from templates",
}

//
//
func AvailableCommands() map[string]string {
	return DESCRIPTIONS
}

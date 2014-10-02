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
	engine "./../../stagosaurus" //"github.com/ndrew/stagosaurus"
	"fmt"
)

///////////////////////
// command api stuff
//	 <?> extend this to an rpc api
//

var COMMANDS = map[string]func(engine.Config, []string) (string, error){
	"build": build,
}

var DESCRIPTIONS = map[string]string{
	//"new":     "<post-name> [<params>] - creates new post and opens editor",
	//"edit":    "<post-name>           - opens post in editor",
	//"publish": "[<post-name>]      - renders markdown posts to html",
	"build": "- generates html from templates",
}

//
//
func AvailableCommands() map[string]string {
	return DESCRIPTIONS
}

//
//
func Workflow(config engine.Config, action string, params []string) (string, error) {
	// TODO: create stagosaurus engine

	var command = COMMANDS[action]
	if nil != command {
		return command(config, params)
	}

	return fmt.Sprintf("Can't do action '%v' with params: %v \n", action, params), nil
}

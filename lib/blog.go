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
	"fmt"
	engine "github.com/ndrew/stagosaurus"
)

func newPost(args []string) (string, error) {
	fmt.Println("NEW!\n")

	return "You've did it", nil
}

///////////////////////
// command api stuff
//	 <?> extend this to an rpc api
//

var COMMANDS = map[string]func([]string) (string, error){
	"new": newPost,
}

var DESCRIPTIONS = map[string]string{
	"new":     "<post-name> [<params>] - creates new post and opens editor",
	"edit":    "<post-name>           - opens post in editor",
	"publish": "[<post-name>]      - renders markdown posts to html",
}

func AvailableCommands() map[string]string {
	return DESCRIPTIONS
}

func Workflow(config *engine.Config, action string, params []string) (string, error) {
	// TODO: create stagosaurus engine

	var command = COMMANDS[action]
	if nil != command {
		return command(params)
	}

	return fmt.Sprintf("Can't do action '%v' with params: %v \n", action, params), nil
}

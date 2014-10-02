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
	"errors"
)

/// commands go here

//
//
func build(config stago.Config, args []string) (string, error) {

	generator, err := New()
	if err != nil {
		return "can't init stagosaurus", err
	}
	//fmt.Println(generator)

	// validate the config
	validator := map[interface{}](func(interface{}) bool){
		"index-template": func(v interface{}) bool {
			return v != nil
		},
	}

	if original, _ := config.Validate(validator); !original {
		return "incorrect configuration", errors.New("You've provided too trivial value! Try again, be original!")
	}

	posts, err := generator.GetPosts(config)
	if err != nil {
		return "", err
	}

	renderedPosts, err := generator.Render(config, posts)
	if err != nil {
		return "", err
	}

	_, err = generator.Deploy(config, renderedPosts)
	if err != nil {
		return "", err
	}

	return "You've did it", nil
}

///////////////////////
// command api stuff
//	 <?> extend this to an rpc api
//

var COMMANDS = map[string]func(stago.Config, []string) (string, error){
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

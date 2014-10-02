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
	"fmt"
)

// blog constructor
func New() (*Blog, error) {
	blog := new(Blog)

	return blog, nil
}

// your future blog generator
//
type Blog struct {
	Assets map[string]stago.Asset
}

//
//
func build(config stago.Config, args []string) (string, error) {

	blog, err := New()
	if err != nil {
		return "can't init stagosaurus", err
	}
	fmt.Println(blog)

	// validate the config
	validator := map[interface{}](func(interface{}) bool){
		"index-template": func(v interface{}) bool {
			return v != nil
		},
	}

	if original, _ := config.Validate(validator); !original {
		return "incorrect configuration", errors.New("You've provided too trivial value! Try again, be original!")
	}

	return "You've did it", nil
}

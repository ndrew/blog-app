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
	engine "github.com/ndrew/stagosaurus"
)

/*func newPost(args []string, engine *stagosaurus.Site) {
    // do something here
    fmt.Println("new post!")
    editPost(args, engine)
}

func editPost(args []string, engine *stagosaurus.Site) {
    // do something here
    fmt.Println("edit post!")
}

func publishPosts(args []string, engine *stagosaurus.Site) {
    // do something here
    fmt.Println("publish!")
}*/

/*func getConfig(configFile string) *engine.Config {
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
}*/

func Workflow(config *engine.Config, action string, params []string) (bool, error) {

	/*renderingStrategy := new(engine.RenderingStrategy)
	postsFactory := new(engine.FileSystem)
	// todo: do this via config
	postsFactory.PostsDir = "/Users/ndrw/Desktop/dev/site/blog/posts"

	engine := engine.New(config, postsFactory, renderingStrategy, nil) // todo: add depoloyer
	*/

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
		}
	}

	return false, nil
}

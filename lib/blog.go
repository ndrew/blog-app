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
	"strings"
)

// blog constructor
func New() (*BlogGenerator, error) {
	generator := new(BlogGenerator)

	return generator, nil
}

// your future blog generator
//
type BlogGenerator struct {
	Assets map[string]stago.Asset
}

// PostsSource
//
func (this *BlogGenerator) GetPosts(meta stago.Config) ([]stago.Post, error) {
	conf := stago.HumanConfig(meta)

	source, _ := conf.String("source-dir")

	fmt.Print("* retrieving posts from ")
	fmt.Println(source)

	fs, _ := stago.NewFileSystem(conf)
	fmt.Println(fs)

	//assets := []stago.Asset{}
	posts := []stago.Post{}

	post, err := stago.NewPost("INDEX", "foo!", meta, []stago.Asset{})
	if err != nil {
		return []stago.Post{}, err
	}

	posts = append(posts, post)

	/*post, err := stago.NewPost("INDEX", "{HEADER}\nHere are my posts: \n{POSTS}", meta, assets)
	if err != nil {
		return []stago.Post{}, err
	}

	posts = append(posts, post) */
	return posts, nil
}

// rendering text in a fancy border
//
func (this *BlogGenerator) Render(cfg stago.Config, posts []stago.Post) ([]stago.Post, error) {
	var results = []stago.Post{}

	var p stago.Post = nil
	var err error = nil

	for _, post := range posts {
		if "INDEX" == post.GetName() {
			p, err = this.renderPost(post) //this.renderIndex(post)
		} else {
			p, err = this.renderPost(post)
		}

		if err != nil {
			return results, err
		}
		results = append(results, p)

	}
	return results, nil
}

// renders index post
//
/*func (this *BlogGenerator) renderIndex(post stago.Post) (stago.Post, error) {
	hello := ""
	world := ""

	// cast to type manually
	helloProperty := post.GetConfig().Get("greeting")
	if helloProperty != nil {
		var ok bool = true
		if hello, ok = helloProperty.(string); !ok {
			return nil, errors.New("hello is not a string!")
		}
	}

	// or use shorthand for common types: string/bool/int
	world, err := stago.ToString(post.GetConfig().Get("blogName"))
	if err != nil {
		return nil, err
	}

	header := hello + " " + world + "!"

	indexContent, err := post.GetContents()
	if err != nil {
		return nil, err
	}

	postsListing := ""

	// usually here you have to sort posts on some criteria (i.e. post date from meta-data), but I'll ommit it here
	for _, asset := range post.GetAssets() {
		if p, ok := asset.(stago.Post); ok {
			postsListing += "\t - " + p.GetName() + "\n"
		}
	}

	content := strings.Replace(string(*indexContent), "{HEADER}", BlogHeader+"\n"+header, 1)
	content = strings.Replace(content, "{POSTS}", postsListing, 1)

	return stago.NewPost("index.html", content, stago.EmptyConfig(), []stago.Asset{})
} */

// renders post
//
func (this *BlogGenerator) renderPost(post stago.Post) (stago.Post, error) {
	data, err := post.GetContents()
	if err != nil {
		return nil, err
	}

	content := strings.Replace(string(*data), "{NO SUCH SUBSTITUION}", "NO!", 1)
	return stago.NewPost(strings.Replace(post.GetName(), " ", "_", 10)+".htm", content, stago.EmptyConfig(), []stago.Asset{})
}

//
//
func (this *BlogGenerator) Deploy(config stago.Config, posts []stago.Post) ([]stago.Post, error) {
	// here usually posts are being saved to filesystems, but for simplicity we will 'deploy' posts to screen
	for _, post := range posts {
		fmt.Println(post.GetName())
		contents, err := post.GetContents()
		if err != nil {
			return []stago.Post{}, err
		}
		fmt.Println(string(*contents))
		fmt.Println(strings.Repeat("=", 80))
	}

	return []stago.Post{}, nil
}

//
//
func (generator *BlogGenerator) BuildAll(cfg stago.Config) error {
	// 0) setup
	config := stago.HumanConfig(cfg)

	// validate the config
	validator := map[interface{}](func(interface{}) bool){
		"templates": func(v interface{}) bool {
			dict := v.(map[string]interface{})

			return dict["index"] != nil && dict["post"] != nil
		},
	}
	if original, _ := config.Validate(validator); !original {
		return errors.New("Error with configuration!")
	}

	// 1) posts
	postsCfg, _ := config.SubConfig("posts")
	posts, err := generator.GetPosts(postsCfg)
	if err != nil {
		return err
	}
	// println(posts)
	// return nil

	// 2) render
	renderCfg, _ := config.SubConfig("templates")
	renderedPosts, err := generator.Render(renderCfg, posts)
	if err != nil {
		return err
	}

	// 3) deploy
	_, err = generator.Deploy(config, renderedPosts)
	if err != nil {
		return err
	}

	// 4) promote
	return nil
}

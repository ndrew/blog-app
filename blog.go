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

import(
"flag"
"fmt"
"sort"
"github.com/ndrew/blog"
)


var COMMANDS = map[string]string{
    "new": "<post-name> [<params>] - creates new post and opens editor", 
    "edit": "<post-name>           - opens post in editor", 
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
    if (full_description) {
        fmt.Println("Available commands:")
    }
    
    var keys []string
    for k := range COMMANDS {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, command := range keys {
        if (full_description) {
            fmt.Println(" ■", command, COMMANDS[command])
        } else {
            // TODO: params
            fmt.Println(command)
        }
    }
}


func NewPost(args []string) {
    // do something here
    fmt.Println("new post!")
}

func main() {
    var configFile = ""
    flag.StringVar(&configFile, "config", "default.xml", "help message for flagname")
    flag.Parse()

    blog.HelloFromLib() //ReadConfig(configFile)
    //fmt.Println( cfg.Foo )
    
        
    var action = flag.Arg(0)
    if (action == "autocomplete") {
        listCommands(false)
        return 
    } 


    var args = flag.Args()
    var commandParams = []string{} 
    if (len(args) > 1) {
        commandParams = args[1:len(args)]    
    }    

    switch {
        case action == "new": NewPost(commandParams)
        case action == "edit": fmt.Println("edit")
        case action == "publish": fmt.Println("publish")

        default: {
            printHeader()
            listCommands(true)
        } 
    }
}

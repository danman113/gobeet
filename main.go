package main

import (
	"fmt"
	"github.com/danman113/gobeet/config"
	"github.com/danman113/gobeet/ping"
	// 	. "github.com/danman113/gobeet/re"
	// 	"runtime"
	// 	"time"
)

func main() {
	conf, _ := config.ParseConfigFile("sampleconfig.json")
	fmt.Println(*conf)
	for _, site := range conf.Sites {
		fmt.Println(site)
		for _, page := range site.Pages {
			fmt.Println(page)
			ret := make(chan error)
			ping.PingStatus(page, ret)
			for val := range ret {
				fmt.Println("Error: ")
				fomt.Println(val)
			}
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"github.com/danman113/gobeet/config"
	"github.com/danman113/gobeet/email"
	"github.com/danman113/gobeet/ping"
)

func main() {
	configPath := flag.String("config", "config.json", "Location for a config file")
	emailPassword := flag.String("password", "", "Password for your email account")
	flag.Parse()
	conf, err := config.ParseConfigFile(*configPath)
	if err != nil {
		fmt.Println("Could Not Parse Config File")
		fmt.Println("Error: ")
		fmt.Println(err)
		return
	}
	fmt.Println(conf.Email)
	email.Config = conf.Email
	email.EmailPassword = emailPassword
	// email.SendAlert(err, conf.Sites[0])
	if err != nil {
		fmt.Println("Could not parse config file: " + *configPath)
		return
	}
	fmt.Println(*conf)
	for i, _ := range conf.Sites {
		go ping.PingWebsite(&conf.Sites[i])
	}
	select {}
}

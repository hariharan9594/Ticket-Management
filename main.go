package main

import (
	"log"
	"time"

	"github.com/jessevdk/go-flags"
	"gitlab.com/vipindasvg/ticket-management/common"
	"gitlab.com/vipindasvg/ticket-management/routers"
)

var opts struct {
	Port string `short:"p" long:"port" description:"set TCP port to listen to"`
}


func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalln("error parsing flags", err)
	}

	defer common.Db.Close()
	//wait for the database service
	time.Sleep(time.Second * 3)
	common.StartUp()
	e := routers.InitRoutes()
	common.Log.Info("STARTING THE WEB SERVICE...")
	if opts.Port != "" {
		log.Panic(e.Start(":" + opts.Port))
	} else {
		log.Panic(e.Start(":8080"))
	}
}
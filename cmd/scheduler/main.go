package main

import (
	"goabtest/internal/cache"
	"goabtest/internal/utils"
	"log"
	"time"
)

func main() {
	log.Println("starting scheduler")
	conn := utils.GetDBDriver()
	t := time.NewTicker(1 * time.Hour)
	for {
		time := <-t.C
		log.Println("running scheduler at ", time)
		cache.RefreshFeatures(conn)
	}
}

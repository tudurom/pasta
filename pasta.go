package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/tudurom/pasta/util"
)

func main() {
	var configfile = flag.String("config", util.DefaultConfigFile, "path to the config file")
	flag.Parse()

	engine := gin.Default()

	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})

	util.GlobalConfig.ReadConfig(*configfile)
	util.GlobalStorage.NewStorage()
	util.ConnectToDB()
	util.MigrateSchema()
	defer util.GlobalDB.Close()

	util.MakeRoutes(engine)
	engine.Run(":8080")
}

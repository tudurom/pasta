package main

import (
	"fmt"

	"github.com/tudurom/pasta/util"
)

func main() {
	util.GlobalConfig.ReadConfig(util.DefaultConfigFile)
	util.ConnectToDB()
	util.MigrateSchema()

	fmt.Print(util.GlobalConfig, util.GlobalDB)
}

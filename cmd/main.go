package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jekabolt/imgproc"
	"github.com/kardianos/osext"

	log "github.com/InVisionApp/go-logger"
	"github.com/spf13/viper"
)

func main() {
	folderPath, _ := osext.ExecutableFolder()
	logger := log.NewSimple().WithFields(log.Fields{"(" + folderPath: ")"})

	configPath := flag.String("conf", "./imgproc.config", "set a config path")
	viper.SetConfigType("json")

	confFile, err := os.Open(*configPath)
	if err != nil {
		fmt.Println("err:os.Open: %v", err.Error())
	}

	viper.ReadConfig(confFile)

	globalConfig := imgproc.Config{}
	viper.GetViper().UnmarshalExact(&globalConfig)

	globalConfig.Router.InitRouter()
	logger.Infof("Router inited on port %v", globalConfig.Router.HTTPPort)

	block := make(chan struct{})
	<-block

}

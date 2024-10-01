package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	datasource "github.com/myklst/packer-plugin-tencentcloud/datasource/image"
	"github.com/myklst/packer-plugin-tencentcloud/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("images", new(datasource.Datasource))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

package main

import (
	"os"
	"log"
	"fmt"
	"plugin"

	"github.com/kardianos/osext"

	. "github.com/Rojuinex/go-plugin-sql/dataprovider"
)

func main() {
	folderPath, err := osext.ExecutableFolder()

	if err != nil {
		log.Fatal(err)
	}

	pluginPath := folderPath + "/plugins"

	stats, err := os.Stat(pluginPath)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("plugins folder does not exist")
		} else {
			fmt.Printf("Cannot stat '%s'\n", pluginPath)
			log.Fatal(err)
		}
	}

	if !stats.IsDir() {
		log.Fatal("File 'plugins' exists but is not a directory")
	}

	mod := pluginPath + "/sql-provider.so"

	providerPlugin, err := plugin.Open(mod)

	if err != nil {
		log.Fatal(err)
	}

	symProvider, err := providerPlugin.Lookup("ProviderImplementation")
	if err != nil {
		fmt.Println("Invalid or corrupted plugin.")
		fmt.Println("Cannot lookup symbol ProviderImplementation.")
		log.Fatal(err)
	}

	var dataProvider DataProvider

	dataProvider, ok := symProvider.(DataProvider)
	if !ok {
		log.Fatal("Exported symbol 'ProviderImplementation' is not compatable with DataProvider interface.")
	}

	records, err := dataProvider.GetData()

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Printf("ID: %d, Name: %s", record.ID, record.Name)
	}
}


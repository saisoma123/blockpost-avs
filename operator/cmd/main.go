package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli"

	"operator"

	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/types"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{config.ConfigFileFlag}
	app.Name = "blockpost-avs"
	app.Usage = "BlockPost Message Storage/Retrieval"
	app.Description = "Service that stores and retrieve messages with off chain validation."

	app.Action = operatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func operatorMain(ctx *cli.Context) error {

	// Reads in config.yaml to initialize Operator
	log.Println("Initializing Operator")
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	log.Println("initializing operator")
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}
	log.Println("initialized operator")

	log.Println("starting operator")
	// Starts the subscriber and AVS process
	err = operator.StartMessageProcessing(context.Background())
	if err != nil {
		return err
	}
	log.Println("started operator")

	return nil

}

package main

import (
	"flag"
	"fmt"
	"os"

	utils "github.com/devops-unicorn/buoy/pkg/utils"
)

func main() {
	// https://github.com/dhirajgidde/Go-Basics/blob/9b144b463e0ab30166393af2e0eb265929a4a22d/Fibonacii/cli/CLI.go
	// config
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	componentFlag := configCommand.String("component", "", "component {storage|network|image}. (Required)")
	durationFlag := configCommand.String("duration", "", "Set scan duration (cron expression)")

	//help
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Command option is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "config":
		configCommand.Parse(os.Args[2:])
		fmt.Println("sub-command 'config'")
		fmt.Println("  component:", *componentFlag)
		fmt.Println("  duration:", *durationFlag)
	case "help":
		fmt.Println("")
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if configCommand.Parsed() {
		// Required Flags
		if *componentFlag == "" {
			fmt.Println("Please specify the scanning component.")
			os.Exit(1)
		}
		if *durationFlag == "" {
			fmt.Println("Please specify cron duration.")
			os.Exit(1)
		}
		componentChoices := map[string]bool{"storage": true, "network": true, "image": true}
		if _, validChoice := componentChoices[*componentFlag]; !validChoice {
			configCommand.PrintDefaults()
			os.Exit(1)
		}
		utils.ToTime(*durationFlag)
		fmt.Printf("You set cron job %q to run\n", *durationFlag)

	}

	if helpCommand.Parsed() {
		fmt.Println("")
	}

}

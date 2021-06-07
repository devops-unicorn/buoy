package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!!")
	// Subcommands
	configCommand := flag.NewFlagSet("config", flag.ExitOnError)
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Command option is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "config":
		configCommand.Parse(os.Args[2:])
	case "help":
		break
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if configCommand.Parsed() {
		// Required Flags
		// if *listTextPtr == "" {
		//     configCommand.PrintDefaults()
		//     os.Exit(1)
		// }
		// //Choice flag
		// metricChoices := map[string]bool{"chars": true, "words": true, "lines": true}
		// if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
		//     listCommand.PrintDefaults()
		//     os.Exit(1)
		// }
		// // Print
		// fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n",iu
		//     *listTextPtr,
		//     *listMetricPtr,
		//     *listUniquePtr,
		// )
	}

}

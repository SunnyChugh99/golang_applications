package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/klog"
)

func main() {
	// Create a new FlagSet for logging flags
	loggingFlags := flag.NewFlagSet("logging", flag.ExitOnError)

	// Initialize Kubernetes logging flags
	klog.InitFlags(loggingFlags)

	// Parse command-line arguments
	loggingFlags.Parse(os.Args[1:])

	// Print the command-line arguments
	fmt.Println("Logging flags initialized successfully.")
	fmt.Println("Command-line arguments:", os.Args[1:])

	// ... rest of your code ...
}

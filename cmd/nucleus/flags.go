package main

import (
	"github.com/spf13/cobra"
)

//AttachCLIFlags attaches command line flags to command
func AttachCLIFlags(rootCmd *cobra.Command) error {

	rootCmd.PersistentFlags().StringP("config", "c", "", "the config file to use")
	rootCmd.PersistentFlags().StringP("port", "p", "", "Port for api server to run")
	rootCmd.PersistentFlags().StringP("payloadAddress", "l", "", "Payload address")
	rootCmd.PersistentFlags().BoolP("verbose", "", false, "Run in verbose mode")
	rootCmd.PersistentFlags().BoolP("coverage", "", false, "Run coverage only mode")
	rootCmd.PersistentFlags().BoolP("parser", "", false, "Run YML parsing only mode")
	rootCmd.PersistentFlags().BoolP("discover", "", false, "Run nucleus in test discovery mode")
	rootCmd.PersistentFlags().BoolP("execute", "", false, "Run nucleus in test execution mode")
	rootCmd.PersistentFlags().StringP("env", "e", "prod", "Environment.")
	rootCmd.PersistentFlags().String("taskID", "", "The unique ID for a task")
	rootCmd.PersistentFlags().String("locators", "", "The test locators for a task")
	rootCmd.PersistentFlags().String("locatorAddress", "", "The test locators address for a task")
	rootCmd.PersistentFlags().String("buildID", "", "The unique ID for a build")
	rootCmd.PersistentFlags().String("targetCommit", "", "The target commit for nucleus")
	rootCmd.PersistentFlags().String("baseCommit", "", "The base commit for nucleus")
	rootCmd.PersistentFlags().StringP("synapsehost", "", "", "Local Ip of proxy server.")
	rootCmd.PersistentFlags().BoolP("local", "", false, "local mode")

	return nil
}

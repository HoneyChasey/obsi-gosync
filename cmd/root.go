/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


// thx to https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43
// and to https://cobra.dev/


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ogs",
	Short: "CLI written in Golang. Backup your Obisdian vault with encryption in the cloud or in your homelab ",
	Long: `Obsidian-gosync is a CLI written in go. You can create automated backup, sync to your favorite drive and add some encryption. Under developement, you may have some crash or problems, so please take time to make me an ussue ;)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.obsi-gosync.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/HoneyChasey/obsi-gosync/internal"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup <scr> <dst>",
	Short: "Create and archive of your obsidian folder, encrypting if needed and send it to a server",
	Long: `Please when running this command, <scr> and <dest> are only usable with absolute path
<src>   absolute path to the obsidian vault
<dest> absolute path where you wanna save your back in zip format"
`,
	Args:  cobra.ExactArgs(2), //Checking with cobra 
	Run: func(cmd *cobra.Command, args []string) {
		if !filepath.IsAbs(args[0]) || !filepath.IsAbs(args[1]) {fmt.Println("The path must be absolute"); cmd.Help(); return} 
		err := internal.CreateZip(args[0], args[1])
		if err != nil {fmt.Println("Error:", err)}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

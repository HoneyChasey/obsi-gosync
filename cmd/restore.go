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

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore <scr> <dst>",
	Short: "Restore backup from your server",
	Long: `Please when running this command, <scr> <dest> are only usable with absolute path. <scr> is where your archive is stored. <dst> is where your obsidian vault is setupt`,
	Run: func(cmd *cobra.Command, args []string) {
		if !filepath.IsAbs(args[0]) || !filepath.IsAbs(args[1]) {fmt.Println("The path must be absolute"); cmd.Help(); return} 
		err := internal.Unzip_archvie(args[0], args[1])
		if err != nil {fmt.Println("Error",err)}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

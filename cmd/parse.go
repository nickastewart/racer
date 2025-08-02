/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"racer/parser"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Given a Daytona HTML email, prints out the parsed data",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		event := parser.Parse(file)
		fmt.Println(event.Location)
		fmt.Println(event.RaceType)
		for _, row := range event.DriverTimes {
			fmt.Println(row)
		} 
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

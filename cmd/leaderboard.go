/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"racer/model"
	"racer/parser"
	"racer/sort"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var driver string
var competitors []string
var isMarkdownOutput bool

var leaderboardCmd = &cobra.Command{
	Use:   "leaderboard",
	Short: "Displays the leaderboard for a given set of events",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		createLeaderboard()
	},
}

func init() {
	rootCmd.AddCommand(leaderboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	leaderboardCmd.PersistentFlags().StringVarP(&driver, "driver", "d", "", "Name of the Driver")
	leaderboardCmd.PersistentFlags().StringSliceVarP(&competitors, "competitors", "c", []string{}, "List of competitor names")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaderboardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	leaderboardCmd.Flags().BoolVarP(&isMarkdownOutput, "markdownOuput", "m", false, "Display Output In Markdown")
}

func createLeaderboard() {
	var baseDir string = "./"
	entries, err := os.ReadDir(baseDir)

	if err != nil {
		log.Panic(err)
	}

	var rows []model.Row = []model.Row{}

	for _, entry := range entries {
		event := parser.Parse(baseDir + entry.Name())
		for _, time := range event.DriverTimes {
			if time.Racer == driver || slices.Contains(competitors, time.Racer) {
				rows = append(rows, model.Row{
					DriverTime: &time,
					Event:      &event,
				})
			}
		}
	}

	sort.Sort(&rows)

	var sb strings.Builder
	sb.WriteString("| Pos | " + rightPad("**Driver**", 24) + " | **Best**   | **Avg**    | **Race Type** \n")
	for i, row := range rows {
		sb.WriteString(markDownRow(&row, i))
	}
	fmt.Print(sb.String())
}

func markDownRow(row *model.Row, pos int) string {
	return fmt.Sprintf("| %d   | %s | %.3f | %.3f | %s \n",
		pos, rightPad(row.DriverTime.Racer, 20), float64(row.DriverTime.Best)/1000, float64(row.DriverTime.Avg)/1000, row.Event.RaceType)
}

func rightPad(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
}

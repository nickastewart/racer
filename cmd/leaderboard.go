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
	"slices"

	"github.com/spf13/cobra"
)

var driver string
var competitors []string

type Leaderboard struct {
	Drivertimes []model.DriverTime
}

// leaderboardCmd represents the leaderboard command
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
}

func createLeaderboard() {
	var baseDir string = "./"
	entries, err := os.ReadDir(baseDir)

	if err != nil {
		log.Panic(err)
	}

	var filteredTimes []model.DriverTime = []model.DriverTime{}

	for _, entry := range entries {
		event := parser.Parse(baseDir + entry.Name())
		for _, time := range event.DriverTimes {
			if time.Racer == driver || slices.Contains(competitors, time.Racer) {
				filteredTimes = append(filteredTimes, time)
			}
		}
	}

	sort(&filteredTimes)

	for _, time := range filteredTimes {
		fmt.Println(time)
	}

}

func sort(unsorted *[]model.DriverTime) {
	if len(*unsorted) < 2 {
		return
	}
	quicksort(unsorted, 0, len(*unsorted)-1)
}

func quicksort(unsorted *[]model.DriverTime, p int, r int) {
	if p < r {
		q := partition(unsorted, p, r)
		quicksort(unsorted, p, q-1)
		quicksort(unsorted, q+1, r)
	}
}

func partition(arr *[]model.DriverTime, p int, r int) int {
	x := &(*arr)[r]
	i := p - 1
	j := p
	
	for j <= r-1 {
		if (*arr)[j].Best <= (*x).Best {
			i++
			temp := (*arr)[j]
			(*arr)[j] = (*arr)[i]
			(*arr)[i] = temp 
		}
		j++
	}

	temp := (*arr)[i+1]
	(*arr)[i+1] = (*arr)[j]
	(*arr)[j] = temp
	
	return i+1
}

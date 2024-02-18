/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/NayronFerreira/stress_test_challenge/loadtester"
	"github.com/NayronFerreira/stress_test_challenge/report"
	"github.com/spf13/cobra"
)

// stressTestCmd represents the stressTest command
var stressTestCmd = &cobra.Command{
	Use:   "stressTest",
	Short: "Teste de carga para uma URL específica",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		header, _ := cmd.Flags().GetStringSlice("header")

		result := loadtester.RunLoadTest(url, requests, concurrency, header)

		fmt.Printf("Teste de carga concluído para %s\n", url)

		report.GenerateReport(result)
	},
}

func init() {

	var url string
	var header []string
	var requests, concurrency int

	rootCmd.AddCommand(stressTestCmd)

	stressTestCmd.Flags().StringVarP(&url, "url", "u", "", "URL to be tested")
	stressTestCmd.Flags().IntVarP(&requests, "requests", "r", 100, "Total number of requests")
	stressTestCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Number of concurrent requests")
	stressTestCmd.Flags().StringSliceVarP(&header, "header", "H", []string{}, "Header to be included in the request")

	stressTestCmd.MarkFlagRequired("url")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

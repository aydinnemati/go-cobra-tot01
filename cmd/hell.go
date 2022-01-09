/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// hellCmd represents the hell command
var hellCmd = &cobra.Command{
	Use:   "hell",
	Short: "hehe lucifer's in the house babe",
	Long: `
	
	just want to be long so what ??? (.\/.)
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("hell called")
		aydin, _ := cmd.Flags().GetString("aydin")
		if aydin != "" {
			getRandomJokeWithTerm(aydin)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(hellCmd)
	// println("heeloow there")
	hellCmd.PersistentFlags().String("aydin", "", "lucifer's name")

}
func getRandomJokeWithTerm(aydin string) {
	log.Println(aydin)
}
func getRandomJoke() {
	log.Println("its noting there ...")
}

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new reminder",
	Run: func(cmd *cobra.Command, args []string) {
		addReminder(args)
	},
}

func addReminder(args []string) {
	fileList, err := os.ReadDir("./")
	CheckErr(err)
	findDir(fileList)

	os.Create("./data/" + strings.Join(args, " "))
}

func findDir(dir []fs.DirEntry) bool {
	for _, f := range dir {
		if f.Name() == "data" {
			return true
		}
	}
	os.Mkdir("data", 0755)
	return false
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}

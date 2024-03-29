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
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var RipGrep string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List GCloud projects to which you have access",
	Long: `Equivalent to calling gcloud projects list.

This was mostly used as a test for learning Cobra, but does
speed up listing projects. Also allows for easy rg`,
	Run: func(cmd *cobra.Command, args []string) {
		gcloudCmd := exec.Command("gcloud", "projects", "list")
		rgCmd := exec.Command("rg", RipGrep)

		rgCmd.Stdin, _ = gcloudCmd.StdoutPipe()
		rgCmd.Stdout = os.Stdout

		_ = rgCmd.Start()
		err := gcloudCmd.Run()

		if err != nil {
			log.Fatalf("Unable to call gcloud, make sure it is installed: %v", err)
		}
		err = rgCmd.Wait()
		if err != nil {
			log.Fatalf("Unable to call rg, make sure it is installed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().StringVarP(&RipGrep, "rg", "r", "", "Pattern to rg for")
}

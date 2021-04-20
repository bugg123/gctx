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
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// setProjectCmd represents the setProject command
var setProjectCmd = &cobra.Command{
	Use:   "set-project",
	Short: "Sets your gcloud project in your config",
	Long: `Sets the project being used in your gcloud config.

The primary reason for the creation of this project is to 
allow for the autocompletion of gcloud projects`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single project argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gcloudCmd := exec.Command("gcloud", "config", "set", "project", args[0])
		gcloudCmd.Stdout = os.Stdout
		err := gcloudCmd.Run()
		if err != nil {
			log.Fatalf("unable to call gcloud config set project %s, make sure it is installed: %v", args[0], err)
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return getGcloudProjects(toComplete), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(setProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Project struct {
	Name string `json:"name"`
}

func getGcloudProjects(toComplete string) []string {
	cmd := exec.Command("gcloud", "projects", "list", `--format="json"`)

	var out bytes.Buffer
	var projects []Project
	cmd.Stdout = &out
	cmd.Run()
	json.Unmarshal(out.Bytes(), &projects)
	var names []string
	for _, project := range projects {
		if strings.HasPrefix(project.Name, toComplete) {
			names = append(names, project.Name)
		}
	}
	return names
}

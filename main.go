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
package main

import "github.com/bugg123/gctx/cmd"

func main() {
	cmd.Execute()
}

// type Project struct {
// 	Name string `json:"name"`
// }

// func main() {
// 	cmd := exec.Command("gcloud", "projects", "list", `--format="json"`)

// 	var out bytes.Buffer
// 	var projects []Project
// 	cmd.Stdout = &out
// 	cmd.Run()
// 	err := json.Unmarshal(out.Bytes(), projects)
// 	if err != nil {
// 		log.Printf("Json Error: %v", err)
// 	}
// 	var names []string
// 	log.Println("printing stuff")
// 	for _, project := range projects {
// 		log.Println(project.Name)
// 		names = append(names, project.Name)
// 	}

// }
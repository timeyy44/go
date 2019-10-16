/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use:   "agenda",
  Short: "agenda command program",
  Long: `be used for a test`,

  Run: func(cmd *cobra.Command, args []string) {
    _, _ = fmt.Fprintln(os.Stderr, "you should use the subCommands like `agenda register`.")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {

}


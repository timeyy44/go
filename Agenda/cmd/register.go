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
	"Agenda/agenda"
	"github.com/spf13/cobra"
)

var name string
var password string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register for your account",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if agenda.Check(name) {
			agenda.ShowMessage(name, password)
			agenda.AddUser(name, password)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&name, "name", "n", "", "name for used")
	registerCmd.Flags().StringVarP(&password, "password", "p", "000000", "password for used")
}

// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package foo

import (
	"os"

	"github.com/spf13/cobra"
)

// MainCmd represents the foo command
var MainCmd = &cobra.Command{
	Use:   "foo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("foo called")
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := MainCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.

	// TODO: disable completion command?
	MainCmd.Root().CompletionOptions.DisableDefaultCmd = true

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	MainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

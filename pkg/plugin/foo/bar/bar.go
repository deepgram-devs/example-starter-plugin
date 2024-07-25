// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package analyze

import (
	"fmt"

	"github.com/spf13/cobra"

	foo "github.com/deepgram-starters/example-starter-plugin/pkg/plugin/foo"
)

// BarCmd represents the foo command
var BarCmd = &cobra.Command{
	Use:   "bar",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bar called")
	},
}

func init() {
	foo.MainCmd.AddCommand(BarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

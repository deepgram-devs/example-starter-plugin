// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"plugin"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var SubCmd *cobra.Command

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	f1, err := p.Lookup("InitMain")
	if err == nil {
		f1.(func())()
	}

	b, err := p.Lookup("MainCmd")
	if err != nil {
		panic(err)
	}

	SubCmd = *(b.(**cobra.Command))
	RootCmd.AddCommand(SubCmd)
	err = RootCmd.Execute()
	if err != nil {
		fmt.Printf("RootCmd.Execute() Error: %s\n", err)
	}
}

func init() {
	// TODO: disable completion command?
	RootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	// Cobra also supports local flags, which will only run when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

import (
	"github.com/deepgram-starters/example-starter-plugin/pkg/plugin/foo"

	_ "github.com/deepgram-starters/example-starter-plugin/pkg/plugin/foo/bar"
)

// MainCmd exports the main command
var MainCmd = foo.MainCmd //nolint:deadcode // this is used when the plugin is in standalone mode

// Execute runs the main command
func Execute() { //nolint:deadcode // this is used when the plugin is in standalone mode
	foo.Execute()
}

// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package main

import (
	"fmt"
	//"errors"
	plugin "github.com/BFLB/monitoringplugin"
)

func main() {
	check := plugin.New()
	check.Status.Ok()
	check.Message("It works")
	fmt.Printf("%s\n", check.String())

	checkWarn := plugin.NewCheck()
	checkWarn.Status.Warning()
	checkWarn.Message("Warning")
	fmt.Printf("%s\n", checkWarn.String())
}

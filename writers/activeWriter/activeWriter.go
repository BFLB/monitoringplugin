// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package activeWriter

import (
	"fmt"
	"os"
	check "github.com/BFLB/monitoringplugin"
)

type Writer struct {}

func New() (*Writer){
	w := Writer{}
	return &w
}

// Write function for active check. Writes output to stdout and
// exits afterwards
func (w *Writer) Write(check check.Check) (error) {
	// Print result
	fmt.Printf("%s", check.String())

	// Done. Exit with exitcode
	os.Exit(check.Status.ReturnCode())

	return nil
}

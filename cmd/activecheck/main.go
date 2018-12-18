// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package main

import (
	"flag"
	plugin   "github.com/BFLB/monitoringplugin"
	writer   "github.com/BFLB/monitoringplugin/writers/activeWriter"
	r        "github.com/BFLB/monitoringplugin/Range"
	perfdata "github.com/BFLB/monitoringplugin/performancedata"
)

// Comman-line Arguments
var (
	value = flag.Float64("value", 25.0, "Value")
	warn  = flag.String("w", "20:30", "Warning Threshold")
	crit  = flag.String("c", "10:40", "Critical Threshold")
	min   = flag.Int64("min", 0.0, "Min")
	max  = flag.Int64("max", 100.0, "max")
)
	
func main() {

	// Parse command-line args
	flag.Parse()
	
	// Create check
	check := plugin.New()

	// Set status
	check.Status.Ok(true)

	// Set message
	check.Message("Go plugin Go")

	// Set ranges
	rangeWarn := r.New()
	rangeWarn.Parse(*warn)

	rangeCrit := r.New()
	rangeCrit.Parse(*crit)

	// Update Status
	check.Status.Threshold(*value, rangeWarn, rangeCrit, false)

	// Performancedata
	perf, err := perfdata.New("TESTPERF", *value, "", rangeWarn, rangeCrit, min, max)
	if err == nil {
		check.Perfdata(perf)
	}

	// Write result
	w := writer.New()
	w.Write(check)

}

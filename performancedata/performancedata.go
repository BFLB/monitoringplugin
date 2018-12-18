// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//
// Basic active-check plugin example
//

package performancedata

import (
	"fmt"
	"strconv"
	"math"
	"strings"
	"errors"
	r "github.com/BFLB/monitoringplugin/Range"
)

// FIXME rename to datapoint
type PerformanceData struct {
	label string
	value float64
	uom   string
	warn  *r.Range
	crit  *r.Range
	min   *int64
	max   *int64
}

func New(label string, value float64, uom string, warn *r.Range, crit *r.Range, min *int64, max *int64 ) (*PerformanceData, error) {

	// Validate label
	if strings.ContainsAny(label, "='\n") {
		return nil, errors.New(fmt.Sprintf("Invalid label format: %s", label))
	}

	// Validate UOM
	switch uom {
	case "":
		// Empty, ok
	case "s", "ms", "us": 	
		// Ok
	case "%":
		// Ok
	case "B", "KB", "MB", "TB": 
		// OK
	case "c":
		// OK
	default:
		return nil, errors.New(fmt.Sprintf("Invalid UOM: %s", uom))		
	}

	p := PerformanceData{label, value, uom, warn, crit, min, max}

	return &p, nil
}

func (p *PerformanceData) String() (string) {

	//Label
	label := p.label
	if strings.Contains(p.label, " ") {
		label = fmt.Sprintf("'%s'", label)
	}
	
	//Warn
	warn := ""
	if p.warn != nil {
		warn = p.warn.String()
	}
	
	//Warn
	crit := ""
	if p.crit != nil {
		crit = p.crit.String()
	}
	
	//Min
	min := ""
	if p.min != nil {
		min = strconv.FormatInt(int64(*p.min), 10)
	}
	
	//Max
	max := ""
	if p.max != nil {
		max = strconv.FormatInt(int64(*p.max), 10)
	}

	// Format output string
	str := fmt. Sprintf("%s=%s%s;%s;%s;%s;%s", label, floatToString(p.value), p.uom, warn, crit, min, max)

	// Strip unneeded semicolons
	str = strings.TrimRight(str, ";")

	return str
}

// TODO Move to utilities
// Return a string representation of float64. Convert to int representation if possible 
func floatToString(f float64) (string) {
	i, frac := math.Modf(f)
	if frac == 0 {
		return strconv.FormatInt(int64(i), 10)
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float64(x int64) *int64 {
    return &x
}
// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package Range

import (
	"math"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

type Range struct {
	start  float64
	end    float64
	inside bool
}

// Returns a pointer to a range object
func New() (*Range) {
	r := Range{0,0,false}
	return &r
}

// Parses string value of range
// Returns error if str of invalid format
func (r *Range) Parse(str string) (error) {
	var err error

	//Trim
	strings.TrimSpace(str)

	// Parse inside (@)
	i := strings.Index(str, "@")
	switch i {
	// Inside not set
	case -1:
		r.inside = false
	// Inside set
	case 0:
		r.inside = true
		str = str[1 : len(str)-1]
	// Inside at wrong position
	default:
		return errors.New("Invalid input detected: @ at wrong place")
	}

	// Parse Start
	i = strings.Index(str, ":")
	switch i {
	// No (:) Set. Min = 0
	case -1:
		r.start = 0
	// (:) at index 0, invalid
	case 0:
		return errors.New("Invalid input detected: : at wrong place")
	// (: set)
	default:
		startStr := str[0 : i]
		str = str[i+1 : len(str)]
		// If (~) set to infinite
		if startStr == "~" {
			r.start = math.Inf(-1)
		} else {
			var start float64
			start, err := parseNumber(startStr)
			if err != nil {
				return err
			}
			r.start = start
		}
	}

	// Parse End
	if str == "~" {
		r.end = math.Inf(1)
	} else {
		var end float64
		end, err = parseNumber(str)
		if err != nil {
			return err
		}
		r.end = end
	}

	return nil
}

func (r *Range) String() string {

	str := ""

	// Inside
	if r.inside == true {
		str = str + "@"
	}

	// Start
	switch r.start {
	case math.Inf(-1):
		str = str + "~:"

	case float64(0):
		// Do nothing

	default:
		str = str + floatToString(r.start) + ":"
	}

	// End
	switch r.end {
	case math.Inf(1):
		str = str + "~:"
	default:
		str = str + floatToString(r.end)
	}

	return str
}


// Returns true if value matches the ranche.
// False otherwise
func (r *Range) Match(value float64) (match bool) {
	match = false

	if value >= r.start && value <= r.end {
		match = true
		
	}

	if r.inside == false {
		return match
	} else {
		return !match
	}
}


func parseNumber(str string) (float64, error) {
	var err error

	// Float 64
	var f float64
	f, err = strconv.ParseFloat(str, 64)
	if err == nil {
		return f, nil
	}
	// Int
	var i int64
	i, err = strconv.ParseInt(str, 10, 64)
	if err == nil {
		return float64(i), nil
	}
	// Error
	msg := fmt.Sprintf("Invalid input. %s not a number", str)
	return 0, errors.New(msg)
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

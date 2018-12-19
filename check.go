// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package check

import (
	"fmt"
	"strings"
	//"errors"
	s "github.com/BFLB/monitoringplugin/status"
	p"github.com/BFLB/monitoringplugin/performancedata"
)

const VERSION = "v0.1.1"

type Check struct {
	Status	s.Status
	message	[]string
	perfdata []p.PerformanceData
}

func New() (Check) {
	// TODO properly init
	return Check{}
}

// Adds a message line to the result
func (c *Check) Message(messageLine string) {
	c.message = append(c.message, messageLine)
}

// Adds a performance datapoint
func (c *Check) Perfdata(datapoint *p.PerformanceData) {
	if datapoint != nil {
		c.perfdata = append(c.perfdata, *datapoint)
	}
}

// String representation of Result
func (c *Check) String() (string) {
	var s string

	// Service Status
	s = fmt.Sprintf("%s:", c.Status.String())

	// First line of message
	if len(c.message) == 0 {
		s += "Message must not be empty!!!"
		return s
	}
	s += fmt.Sprintf(" %s", c.message[0])

	// Performancedata. Note: Currently all perfdata will be sent in first line
	if len(c.perfdata) > 0 {
		s += "|"

		for _, p := range c.perfdata {
			s += fmt.Sprintf("%s ", p.String()) 
		}
		s = strings.TrimRight(s, " ")
	}

	// Additional message lines
	if len(c.message) > 1 {
		for i := 1; i < len(c.message); i++ {
			s += fmt.Sprintf("\n%s", c.message[i])
		}
	}

	return s
}

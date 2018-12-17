// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package status

import (
	r "github.com/BFLB/monitoringplugin/Range"
)

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

type Status struct {
	// return code
	rc int
}

// Creates a new status with default value ok
func New() Status {
	s := Status{OK}
	return s
}

// Merges a status into an existing one.
func (s *Status) Merge(status Status) {
	switch status.rc {
	case OK:
		s.Ok(false)
	case WARNING:
		s.Warning(false)
	case CRITICAL:
		s.Critical(false)
	default:
		s.Unknown()
	}
}

// Sets the status based on a threshold.
func (s *Status) Threshold(value float64, warn *r.Range, crit *r.Range, force bool) {
	if crit != nil {
		if crit.Match(value) == false {
			s.Critical(force)
			return
		}
	}
	if warn != nil {
		if warn.Match(value) == false {
			s.Warning(force)
			return
		}
	}
	s.Ok(force)
}


// If force == true, set to OK
// Else only takes effect if current state OK
// Noop function but still useful in some circumstances
func (s *Status) Ok(force bool) {
	if force == true {
		s.rc = OK
	} else {
		return
	}
}

// If force == true, set to WARNING. 
// Else set to warning if current state OK
func (s *Status) Warning(force bool) {
	if force == true {
		s.rc = WARNING
	} else {
		if s.rc < WARNING {
			s.rc = WARNING
		}
	}
}

// If force == true, set to CRITICAL. 
// Else set to CRITICAL if current state OK, WARNING
func (s *Status) Critical(force bool) {
	if force == true {
		s.rc = CRITICAL
	} else {
		if s.rc < CRITICAL {
			s.rc = CRITICAL
		}
	}
}

// Set to UNKNOWN. Since this is the highest status, force is not needed. 
func (s *Status) Unknown() {
	s.rc = UNKNOWN
}

// Returns String representations of status
func (s *Status) String() string {
	switch s.rc {
	case OK:
		return "OK"
	case WARNING:
		return "Warning"
	case CRITICAL:
		return "Critical"
	default:
		return "Unknown"
	}
}

// Returns the Returncode of the status
func (s *Status) ReturnCode() int {
	return s.rc
}

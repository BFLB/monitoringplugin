// Copyright (c) 2018 Bernhard Fluehmann. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.
//
// Library for Nagios/Icinga/etc. plugins
//

package status

import (
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
func NewStatus() Status {
	s := Status{OK}
	return s
}

// Forces a status to a given value
func (s *Status) Force(statusCode int) {
		switch statusCode {
		case OK:
			s.rc = OK
		case WARNING:
			s.rc = WARNING
		case CRITICAL:
			s.rc = CRITICAL
		default:
			s.rc = UNKNOWN
		}
	}

// Status update OK. 
// Only takes effect if current state OK
// Noop function but still useful in some circumstances
func (s *Status) Ok() {
	return
}

// Status update to WARNING. 
// Ignored it current state WARNING, CRITICAL, UNKNOWN
func (s *Status) Warning() {
	if s.rc < WARNING {
		s.rc = WARNING
	}
}

// Status update to CRITICAL. 
// Ignored it current state CRITICAL, UNKNOWN
func (s *Status) Critical() {
	if s.rc < CRITICAL {
		s.rc = CRITICAL
	}
}

// Status update to UNKNOWN. 
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
// Copyright 2021-2022 Peter Bigot Consulting, LLC
// SPDX-License-Identifier: Apache-2.0

// Package edcode provides basic encode/decode support for various types that
// aren't natively supported in golang.
package edcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrDurationInvalid = errors.New("invalid duration text")
)

// Duration is an extension of time.Duration that decodes non-negative integer
// values without units as durations in milliseconds, and strings as with
// time.ParseDuration.
type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	s := string(text)
	td, err := time.ParseDuration(s)
	if err != nil && strings.Contains(err.Error(), "time: missing unit") {
		var i int
		i, err = strconv.Atoi(s)
		if err == nil {
			td = time.Duration(i) * time.Millisecond
		}
	}
	if err == nil {
		if td < 0*time.Millisecond {
			err = fmt.Errorf("%w: %s", ErrDurationInvalid, s)
		} else {
			*d = Duration(td)
		}
	}
	return err
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

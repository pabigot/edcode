// Copyright 2021-2022 Peter Bigot Consulting, LLC
// SPDX-License-Identifier: Apache-2.0

package edcode

import (
	"encoding"
	"errors"
	"strings"
	"testing"
	"time"
)

// Run standard verification of expected errors, i.e. that err is an
// error and its text contains errstr.
func confirmError(t *testing.T, err error, base error, errstr string) {
	t.Helper()
	if err == nil {
		t.Fatalf("succeed, expected error %s", errstr)
	}
	if base != nil && !errors.Is(err, base) {
		t.Fatalf("err not from %s: %T: %s", base, err, err)
	}
	if testing.Verbose() {
		t.Logf("Error=`%v`", err.Error())
	}
	if !strings.Contains(err.Error(), errstr) {
		t.Fatalf("failed, missing %s: %v", errstr, err)
	}
}

func TestDuration(t *testing.T) {
	type ts struct {
		text   string
		dur    time.Duration
		exp    string
		err    error
		errstr string
	}
	tests := []ts{
		{
			text: "5",
			dur:  5 * time.Millisecond,
			exp:  "5ms",
		},
		{
			text: "123us",
			dur:  123 * time.Microsecond,
			exp:  "123µs",
		},
		{
			text:   "-1s",
			err:    ErrDurationInvalid,
			errstr: ": -1s",
		},
		{
			text: "0",
			dur:  time.Duration(0),
			exp:  "0s",
		},
	}
	var v Duration

	// First make sure we spelled everything so that the conversion
	// interfaces will be found.  This panics if we screwed up.
	_ = interface{}(&v).(encoding.TextUnmarshaler)
	_ = interface{}(&v).(encoding.TextMarshaler)

	for _, x := range tests {
		s := x.text
		err := v.UnmarshalText([]byte(s))
		if x.err != nil {
			confirmError(t, err, x.err, x.errstr)
		} else if err != nil {
			t.Error(err)
		} else {
			if d := time.Duration(v); d != x.dur {
				t.Errorf("%s: wrong FlushInterval: %v", s, d)
			}
			if y := v.String(); y != x.exp {
				t.Errorf("%s: text FlushInterval wrong: %s", x.exp, y)
			} else {
				r, err := v.MarshalText()
				if err != nil {
					t.Error(err)
				} else if s := string(r); s != x.exp {
					t.Errorf("%s: wrong round-trip: %s", x.exp, s)
				}
			}

		}
	}

}

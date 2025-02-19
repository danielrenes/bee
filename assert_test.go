package bee_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/danielrenes/bee"
)

type mockT struct {
	*testing.T
	wasErr bool
	wasLog bool
	errMsg string
	logMsg string
}

func (mt *mockT) Errorf(format string, args ...any) {
	mt.wasErr = true
	mt.errMsg = fmt.Sprintf(format, args...)
}

func (mt *mockT) Logf(format string, args ...any) {
	mt.wasLog = true
	mt.logMsg = fmt.Sprintf(format, args...)
}

type newBee func(*mockT) *bee.Bee

func newBeeWithoutColor() newBee {
	return func(mt *mockT) *bee.Bee {
		return bee.New(mt, bee.NoColor())
	}
}

func newBeeWithDefaultColor() newBee {
	return func(mt *mockT) *bee.Bee {
		return bee.New(mt)
	}
}

func newBeeWithCustomColor() newBee {
	return func(mt *mockT) *bee.Bee {
		return bee.New(
			mt,
			bee.WhatColor(3, 3, 3),
			bee.ExpectedColor(2, 2, 2),
			bee.ActualColor(1, 1, 1),
		)
	}
}

type matchText func(string) bool

func exact(text string) matchText {
	return func(s string) bool {
		return s == text
	}
}

func regex(pattern string) matchText {
	re := regexp.MustCompile(pattern)
	return func(s string) bool {
		return re.MatchString(s)
	}
}

func TestNil(t *testing.T) {
	tests := []struct {
		actual     any
		wantErr    bool
		errMsg     string
		newMatcher func(string) matchText
		newBee     newBee
	}{
		{
			actual:     nil,
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     (*int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     (error)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     ([]int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     (map[string]int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     func(v int) *int { return &v }(1),
			wantErr:    true,
			errMsg:     "0x[0-9a-f]+ != <nil>",
			newMatcher: regex,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     errors.New("oops"),
			wantErr:    true,
			errMsg:     "oops != <nil>",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     []int{},
			wantErr:    true,
			errMsg:     "[] != <nil>",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     map[string]int{},
			wantErr:    true,
			errMsg:     "map[] != <nil>",
			newMatcher: exact,
			newBee:     newBeeWithoutColor(),
		},
		{
			actual:     nil,
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     (*int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     (error)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     ([]int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     (map[string]int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     func(v int) *int { return &v }(1),
			wantErr:    true,
			errMsg:     "\x1b\\[38;2;250;40;25m0x[0-9a-f]+\x1b\\[0m != \x1b\\[38;2;18;181;32m<nil>\x1b\\[0m",
			newMatcher: regex,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     errors.New("oops"),
			wantErr:    true,
			errMsg:     "\x1b[38;2;250;40;25moops\x1b[0m != \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     []int{},
			wantErr:    true,
			errMsg:     "\x1b[38;2;250;40;25m[]\x1b[0m != \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     map[string]int{},
			wantErr:    true,
			errMsg:     "\x1b[38;2;250;40;25mmap[]\x1b[0m != \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newMatcher: exact,
			newBee:     newBeeWithDefaultColor(),
		},
		{
			actual:     nil,
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     (*int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     (error)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     ([]int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     (map[string]int)(nil),
			wantErr:    false,
			errMsg:     "",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     func(v int) *int { return &v }(1),
			wantErr:    true,
			errMsg:     "\x1b\\[38;2;1;1;1m0x[0-9a-f]+\x1b\\[0m != \x1b\\[38;2;2;2;2m<nil>\x1b\\[0m",
			newMatcher: regex,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     errors.New("oops"),
			wantErr:    true,
			errMsg:     "\x1b[38;2;1;1;1moops\x1b[0m != \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     []int{},
			wantErr:    true,
			errMsg:     "\x1b[38;2;1;1;1m[]\x1b[0m != \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
		{
			actual:     map[string]int{},
			wantErr:    true,
			errMsg:     "\x1b[38;2;1;1;1mmap[]\x1b[0m != \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newMatcher: exact,
			newBee:     newBeeWithCustomColor(),
		},
	}

	for _, test := range tests {
		mockT := &mockT{T: t}
		bee := test.newBee(mockT)
		bee.Nil(test.actual)
		if mockT.wasErr != test.wantErr {
			if test.wantErr {
				t.Error("expected error")
			} else {
				t.Error("expected no error")
			}
		}
		if !test.newMatcher(test.errMsg)(mockT.errMsg) {
			t.Errorf("%q != %q", mockT.errMsg, test.errMsg)
		}
	}
}

func TestNotNil(t *testing.T) {
	tests := []struct {
		actual  any
		wantErr bool
		errMsg  string
		newBee  newBee
	}{
		{
			actual:  func(v int) *int { return &v }(1),
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  errors.New(""),
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  []int{},
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  map[string]int{},
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  nil,
			wantErr: true,
			errMsg:  "<nil> == <nil>",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  (*int)(nil),
			wantErr: true,
			errMsg:  "<nil> == <nil>",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  (error)(nil),
			wantErr: true,
			errMsg:  "<nil> == <nil>",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  ([]int)(nil),
			wantErr: true,
			errMsg:  "[] == <nil>",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  (map[string]int)(nil),
			wantErr: true,
			errMsg:  "map[] == <nil>",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  func(v int) *int { return &v }(1),
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  errors.New(""),
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  []int{},
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  map[string]int{},
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  nil,
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25m<nil>\x1b[0m == \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  (*int)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25m<nil>\x1b[0m == \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  (error)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25m<nil>\x1b[0m == \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  ([]int)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25m[]\x1b[0m == \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  (map[string]int)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25mmap[]\x1b[0m == \x1b[38;2;18;181;32m<nil>\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  func(v int) *int { return &v }(1),
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  errors.New(""),
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  []int{},
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  map[string]int{},
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  nil,
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1m<nil>\x1b[0m == \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  (*int)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1m<nil>\x1b[0m == \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  (error)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1m<nil>\x1b[0m == \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  ([]int)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1m[]\x1b[0m == \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  (map[string]int)(nil),
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1mmap[]\x1b[0m == \x1b[38;2;2;2;2m<nil>\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
	}

	for _, test := range tests {
		mockT := &mockT{T: t}
		bee := test.newBee(mockT)
		bee.NotNil(test.actual)
		if mockT.wasErr != test.wantErr {
			if test.wantErr {
				t.Error("expected error")
			} else {
				t.Error("expected no error")
			}
		}
		if mockT.errMsg != test.errMsg {
			t.Errorf("%q != %q", mockT.errMsg, test.errMsg)
		}
	}
}

func TestTrue(t *testing.T) {
	tests := []struct {
		actual  bool
		wantErr bool
		errMsg  string
		newBee  newBee
	}{
		{
			actual:  true,
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  false,
			wantErr: true,
			errMsg:  "false != true",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  true,
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  false,
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25mfalse\x1b[0m != \x1b[38;2;18;181;32mtrue\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  true,
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  false,
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1mfalse\x1b[0m != \x1b[38;2;2;2;2mtrue\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
	}

	for _, test := range tests {
		mockT := &mockT{T: t}
		bee := test.newBee(mockT)
		bee.True(test.actual)
		if mockT.wasErr != test.wantErr {
			if test.wantErr {
				t.Error("expected error")
			} else {
				t.Error("expected no error")
			}
		}
		if mockT.errMsg != test.errMsg {
			t.Errorf("%q != %q", mockT.errMsg, test.errMsg)
		}
	}
}

func TestFalse(t *testing.T) {
	tests := []struct {
		actual  bool
		wantErr bool
		errMsg  string
		newBee  newBee
	}{
		{
			actual:  false,
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  true,
			wantErr: true,
			errMsg:  "true != false",
			newBee:  newBeeWithoutColor(),
		},
		{
			actual:  false,
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  true,
			wantErr: true,
			errMsg:  "\x1b[38;2;250;40;25mtrue\x1b[0m != \x1b[38;2;18;181;32mfalse\x1b[0m",
			newBee:  newBeeWithDefaultColor(),
		},
		{
			actual:  false,
			wantErr: false,
			errMsg:  "",
			newBee:  newBeeWithCustomColor(),
		},
		{
			actual:  true,
			wantErr: true,
			errMsg:  "\x1b[38;2;1;1;1mtrue\x1b[0m != \x1b[38;2;2;2;2mfalse\x1b[0m",
			newBee:  newBeeWithCustomColor(),
		},
	}

	for _, test := range tests {
		mockT := &mockT{T: t}
		bee := test.newBee(mockT)
		bee.False(test.actual)
		if mockT.wasErr != test.wantErr {
			if test.wantErr {
				t.Error("expected error")
			} else {
				t.Error("expected no error")
			}
		}
		if mockT.errMsg != test.errMsg {
			t.Errorf("%q != %q", mockT.errMsg, test.errMsg)
		}
	}
}

func TestEqual(t *testing.T) {
	type testStruct struct {
		a int
		b []int
		c map[string]int
		p *[]string
	}

	tests := []struct {
		actual   any
		expected any
		wantErr  bool
		errMsg   string
		newBee   newBee
	}{
		{
			actual:   true,
			expected: 1,
			wantErr:  true,
			errMsg:   "bool != int",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   false,
			expected: false,
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   true,
			expected: false,
			wantErr:  true,
			errMsg:   "true != false",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int(1),
			expected: int(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int(1),
			expected: int(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int8(1),
			expected: int8(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int8(1),
			expected: int8(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int16(1),
			expected: int16(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int16(1),
			expected: int16(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int32(1),
			expected: int32(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int32(1),
			expected: int32(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int64(1),
			expected: int64(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   int64(1),
			expected: int64(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint(1),
			expected: uint(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint(1),
			expected: uint(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint8(1),
			expected: uint8(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint8(1),
			expected: uint8(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint16(1),
			expected: uint16(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint16(1),
			expected: uint16(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint32(1),
			expected: uint32(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint32(1),
			expected: uint32(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint64(1),
			expected: uint64(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   uint64(1),
			expected: uint64(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   float32(1e9),
			expected: float32(1e9),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   float32(1e9),
			expected: float32(2e9),
			wantErr:  true,
			errMsg:   "1e+09 != 2e+09",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   float64(1.1),
			expected: float64(1.1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   float64(1.1),
			expected: float64(2.2),
			wantErr:  true,
			errMsg:   "1.1 != 2.2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   complex64(1 + 1i),
			expected: complex64(1 + 1i),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   complex64(1 + 1i),
			expected: complex64(2 + 2i),
			wantErr:  true,
			errMsg:   "(1+1i) != (2+2i)",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   complex128(1 + 1i),
			expected: complex128(1 + 1i),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   complex128(1 + 1i),
			expected: complex128(2 + 2i),
			wantErr:  true,
			errMsg:   "(1+1i) != (2+2i)",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   byte(1),
			expected: byte(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   byte(1),
			expected: byte(2),
			wantErr:  true,
			errMsg:   "1 != 2",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   "a",
			expected: "a",
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   "a",
			expected: "b",
			wantErr:  true,
			errMsg:   "a != b",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   any(testStruct{}),
			expected: any(testStruct{}),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   any(testStruct{a: 1}),
			expected: any(testStruct{a: 2}),
			wantErr:  true,
			errMsg:   "1 != 2 (.a)",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   [1]int{},
			expected: [1]int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   [1]int{1},
			expected: [1]int{2},
			wantErr:  true,
			errMsg:   "1 != 2 ([0])",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   [1]int{},
			expected: [2]int{},
			wantErr:  true,
			errMsg:   "[1]int != [2]int",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   []int{},
			expected: []int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   []int{1},
			expected: []int{2},
			wantErr:  true,
			errMsg:   "1 != 2 ([0])",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   []int{0},
			expected: []int{0, 0},
			wantErr:  true,
			errMsg:   "1 != 2 (len())",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   map[string]int{},
			expected: map[string]int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   map[string]int{"a": 1},
			expected: map[string]int{"a": 2},
			wantErr:  true,
			errMsg:   "1 != 2 ([a])",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   map[string]int{"a": 1},
			expected: map[string]int{"a": 1, "b": 2},
			wantErr:  true,
			errMsg:   "1 != 2 (len())",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{},
			expected: testStruct{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{a: 1},
			expected: testStruct{a: 2},
			wantErr:  true,
			errMsg:   "1 != 2 (.a)",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{b: []int{1}},
			expected: testStruct{b: []int{2}},
			wantErr:  true,
			errMsg:   "1 != 2 (.b[0])",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{b: []int{0}},
			expected: testStruct{b: []int{0, 0}},
			wantErr:  true,
			errMsg:   "1 != 2 (len(.b))",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{c: map[string]int{"a": 1}},
			expected: testStruct{c: map[string]int{"a": 2}},
			wantErr:  true,
			errMsg:   "1 != 2 (.c[a])",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{c: map[string]int{"a": 1}},
			expected: testStruct{c: map[string]int{"a": 1, "b": 2}},
			wantErr:  true,
			errMsg:   "1 != 2 (len(.c))",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{p: &[]string{"a"}},
			expected: testStruct{p: &[]string{"b"}},
			wantErr:  true,
			errMsg:   "a != b (*.p[0])",
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   testStruct{p: &[]string{"a"}},
			expected: testStruct{p: &[]string{"a"}},
			wantErr:  false,
			newBee:   newBeeWithoutColor(),
		},
		{
			actual:   true,
			expected: 1,
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25mbool\x1b[0m != \x1b[38;2;18;181;32mint\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   false,
			expected: false,
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   true,
			expected: false,
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25mtrue\x1b[0m != \x1b[38;2;18;181;32mfalse\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int(1),
			expected: int(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int(1),
			expected: int(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int8(1),
			expected: int8(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int8(1),
			expected: int8(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int16(1),
			expected: int16(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int16(1),
			expected: int16(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int32(1),
			expected: int32(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int32(1),
			expected: int32(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int64(1),
			expected: int64(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   int64(1),
			expected: int64(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint(1),
			expected: uint(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint(1),
			expected: uint(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint8(1),
			expected: uint8(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint8(1),
			expected: uint8(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint16(1),
			expected: uint16(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint16(1),
			expected: uint16(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint32(1),
			expected: uint32(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint32(1),
			expected: uint32(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint64(1),
			expected: uint64(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   uint64(1),
			expected: uint64(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   float32(1e9),
			expected: float32(1e9),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   float32(1e9),
			expected: float32(2e9),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1e+09\x1b[0m != \x1b[38;2;18;181;32m2e+09\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   float64(1.1),
			expected: float64(1.1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   float64(1.1),
			expected: float64(2.2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1.1\x1b[0m != \x1b[38;2;18;181;32m2.2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   complex64(1 + 1i),
			expected: complex64(1 + 1i),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   complex64(1 + 1i),
			expected: complex64(2 + 2i),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m(1+1i)\x1b[0m != \x1b[38;2;18;181;32m(2+2i)\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   complex128(1 + 1i),
			expected: complex128(1 + 1i),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   complex128(1 + 1i),
			expected: complex128(2 + 2i),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m(1+1i)\x1b[0m != \x1b[38;2;18;181;32m(2+2i)\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   byte(1),
			expected: byte(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   byte(1),
			expected: byte(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   "a",
			expected: "a",
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   "a",
			expected: "b",
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25ma\x1b[0m != \x1b[38;2;18;181;32mb\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   any(testStruct{}),
			expected: any(testStruct{}),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   any(testStruct{a: 1}),
			expected: any(testStruct{a: 2}),
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m.a\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   [1]int{},
			expected: [1]int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   [1]int{1},
			expected: [1]int{2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m[0]\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   [1]int{},
			expected: [2]int{},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m[1]int\x1b[0m != \x1b[38;2;18;181;32m[2]int\x1b[0m",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   []int{},
			expected: []int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   []int{1},
			expected: []int{2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m[0]\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   []int{0},
			expected: []int{0, 0},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250mlen()\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   map[string]int{},
			expected: map[string]int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   map[string]int{"a": 1},
			expected: map[string]int{"a": 2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m[a]\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   map[string]int{"a": 1},
			expected: map[string]int{"a": 1, "b": 2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250mlen()\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{},
			expected: testStruct{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{a: 1},
			expected: testStruct{a: 2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m.a\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{b: []int{1}},
			expected: testStruct{b: []int{2}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m.b[0]\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{b: []int{0}},
			expected: testStruct{b: []int{0, 0}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250mlen(.b)\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{c: map[string]int{"a": 1}},
			expected: testStruct{c: map[string]int{"a": 2}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250m.c[a]\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{c: map[string]int{"a": 1}},
			expected: testStruct{c: map[string]int{"a": 1, "b": 2}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25m1\x1b[0m != \x1b[38;2;18;181;32m2\x1b[0m (\x1b[38;2;2;118;250mlen(.c)\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   testStruct{p: &[]string{"a"}},
			expected: testStruct{p: &[]string{"b"}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;250;40;25ma\x1b[0m != \x1b[38;2;18;181;32mb\x1b[0m (\x1b[38;2;2;118;250m*.p[0]\x1b[0m)",
			newBee:   newBeeWithDefaultColor(),
		},
		{
			actual:   true,
			expected: 1,
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1mbool\x1b[0m != \x1b[38;2;2;2;2mint\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   false,
			expected: false,
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   true,
			expected: false,
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1mtrue\x1b[0m != \x1b[38;2;2;2;2mfalse\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int(1),
			expected: int(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int(1),
			expected: int(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int8(1),
			expected: int8(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int8(1),
			expected: int8(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int16(1),
			expected: int16(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int16(1),
			expected: int16(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int32(1),
			expected: int32(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int32(1),
			expected: int32(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int64(1),
			expected: int64(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   int64(1),
			expected: int64(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint(1),
			expected: uint(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint(1),
			expected: uint(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint8(1),
			expected: uint8(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint8(1),
			expected: uint8(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint16(1),
			expected: uint16(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint16(1),
			expected: uint16(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint32(1),
			expected: uint32(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint32(1),
			expected: uint32(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint64(1),
			expected: uint64(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   uint64(1),
			expected: uint64(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   float32(1e9),
			expected: float32(1e9),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   float32(1e9),
			expected: float32(2e9),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1e+09\x1b[0m != \x1b[38;2;2;2;2m2e+09\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   float64(1.1),
			expected: float64(1.1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   float64(1.1),
			expected: float64(2.2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1.1\x1b[0m != \x1b[38;2;2;2;2m2.2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   complex64(1 + 1i),
			expected: complex64(1 + 1i),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   complex64(1 + 1i),
			expected: complex64(2 + 2i),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m(1+1i)\x1b[0m != \x1b[38;2;2;2;2m(2+2i)\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   complex128(1 + 1i),
			expected: complex128(1 + 1i),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   complex128(1 + 1i),
			expected: complex128(2 + 2i),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m(1+1i)\x1b[0m != \x1b[38;2;2;2;2m(2+2i)\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   byte(1),
			expected: byte(1),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   byte(1),
			expected: byte(2),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   "a",
			expected: "a",
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   "a",
			expected: "b",
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1ma\x1b[0m != \x1b[38;2;2;2;2mb\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   any(testStruct{}),
			expected: any(testStruct{}),
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   any(testStruct{a: 1}),
			expected: any(testStruct{a: 2}),
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m.a\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   [1]int{},
			expected: [1]int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   [1]int{1},
			expected: [1]int{2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m[0]\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   [1]int{},
			expected: [2]int{},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m[1]int\x1b[0m != \x1b[38;2;2;2;2m[2]int\x1b[0m",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   []int{},
			expected: []int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   []int{1},
			expected: []int{2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m[0]\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   []int{0},
			expected: []int{0, 0},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3mlen()\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   map[string]int{},
			expected: map[string]int{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   map[string]int{"a": 1},
			expected: map[string]int{"a": 2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m[a]\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   map[string]int{"a": 1},
			expected: map[string]int{"a": 1, "b": 2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3mlen()\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{},
			expected: testStruct{},
			wantErr:  false,
			errMsg:   "",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{a: 1},
			expected: testStruct{a: 2},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m.a\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{b: []int{1}},
			expected: testStruct{b: []int{2}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m.b[0]\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{b: []int{0}},
			expected: testStruct{b: []int{0, 0}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3mlen(.b)\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{c: map[string]int{"a": 1}},
			expected: testStruct{c: map[string]int{"a": 2}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3m.c[a]\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{c: map[string]int{"a": 1}},
			expected: testStruct{c: map[string]int{"a": 1, "b": 2}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1m1\x1b[0m != \x1b[38;2;2;2;2m2\x1b[0m (\x1b[38;2;3;3;3mlen(.c)\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
		{
			actual:   testStruct{p: &[]string{"a"}},
			expected: testStruct{p: &[]string{"b"}},
			wantErr:  true,
			errMsg:   "\x1b[38;2;1;1;1ma\x1b[0m != \x1b[38;2;2;2;2mb\x1b[0m (\x1b[38;2;3;3;3m*.p[0]\x1b[0m)",
			newBee:   newBeeWithCustomColor(),
		},
	}

	for _, test := range tests {
		mockT := &mockT{T: t}
		bee := test.newBee(mockT)
		bee.Equal(test.actual, test.expected)
		if mockT.wasErr != test.wantErr {
			if test.wantErr {
				t.Error("expected error")
			} else {
				t.Error("expected no error")
			}
		}
		if mockT.errMsg != test.errMsg {
			t.Errorf("%q != %q", mockT.errMsg, test.errMsg)
		}
	}
}

func TestExpand(t *testing.T) {
	actual := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in tortor non enim vehicula malesuada at ut lorem. Curabitur placerat sem id pellentesque posuere. Proin porttitor ante a laoreet fringilla. Nam bibendum velit posuere est feugiat, condimentum ornare est malesuada. In sed consectetur mi. Vestibulum ante lectus, vulputate dapibus maximus ac, tincidunt sed nisl. Morbi eu felis faucibus, imperdiet quam in, accumsan libero. Phasellus lacinia mauris arcu, nec aliquam mi laoreet pharetra. Pellentesque non lorem magna."
	expected := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce cursus eros neque, in varius tellus bibendum ut. Nunc blandit eu lorem non mollis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus condimentum eros euismod mattis placerat. Donec id faucibus dolor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Donec in velit vitae lorem venenatis consectetur non et odio. Maecenas odio leo, tristique nec feugiat ac, blandit et nisl. Phasellus faucibus enim eu elit rhoncus mollis. Etiam nec rutrum orci. Ut eget pretium nisi."
	errMsg := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. ... != Lorem ipsum dolor sit amet, consectetur adipiscing elit. ..."
	logMsg := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit.     Lorem ipsum dolor sit amet, consectetur adipiscing elit.   
Nullam in tortor non enim vehicula malesuada at ut lorem.    Fusce cursus eros neque, in varius tellus bibendum ut. Nunc
Curabitur placerat sem id pellentesque posuere. Proin        blandit eu lorem non mollis. Lorem ipsum dolor sit amet,   
porttitor ante a laoreet fringilla. Nam bibendum velit       consectetur adipiscing elit. Phasellus condimentum eros    
posuere est feugiat, condimentum ornare est malesuada. In    euismod mattis placerat. Donec id faucibus dolor. Interdum 
sed consectetur mi. Vestibulum ante lectus, vulputate        et malesuada fames ac ante ipsum primis in faucibus. Donec 
dapibus maximus ac, tincidunt sed nisl. Morbi eu felis       in velit vitae lorem venenatis consectetur non et odio.    
faucibus, imperdiet quam in, accumsan libero. Phasellus      Maecenas odio leo, tristique nec feugiat ac, blandit et    
lacinia mauris arcu, nec aliquam mi laoreet pharetra.        nisl. Phasellus faucibus enim eu elit rhoncus mollis. Etiam
Pellentesque non lorem magna.                                nec rutrum orci. Ut eget pretium nisi.                     `
	mockT := &mockT{T: t}
	bee := bee.New(mockT, bee.NoColor(), bee.ColumnWidth(60))
	bee.Equal(actual, expected)
	if !mockT.wasErr {
		t.Error("expected error")
	}
	if mockT.errMsg != errMsg {
		t.Errorf("%q != %q", mockT.errMsg, errMsg)
	}
	if !mockT.wasLog {
		t.Error("expected log")
	}
	if mockT.logMsg != logMsg {
		t.Errorf("%q != %q", mockT.logMsg, logMsg)
	}
}

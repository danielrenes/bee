package bee

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func (b *Bee) Nil(actual any) {
	b.tb.Helper()
	if !isNil(b.tb, actual) {
		b.errorNotEquals(actual, nil, "")
	}
}

func (b *Bee) NotNil(actual any) {
	b.tb.Helper()
	if isNil(b.tb, actual) {
		b.errorEquals(actual, nil, "")
	}
}

func (b *Bee) True(actual bool) {
	b.tb.Helper()
	b.Equal(actual, true)
}

func (b *Bee) False(actual bool) {
	b.tb.Helper()
	b.Equal(actual, false)
}

func (b *Bee) Equal(actual, expected any) {
	b.tb.Helper()
	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)
	b.equals(actualValue, expectedValue, "")
}

func (b *Bee) equals(actual, expected reflect.Value, what string) {
	b.tb.Helper()
	if !actual.IsValid() || !expected.IsValid() {
		if actual.IsValid() != expected.IsValid() {
			b.errorNotEquals(actual, expected, what)
		}
		return
	}
	if actual.Type() != expected.Type() {
		b.errorNotEquals(actual.Type(), expected.Type(), what)
		return
	}
	switch actual.Kind() {
	case reflect.Bool:
		if actual.Bool() != expected.Bool() {
			b.errorNotEquals(actual.Bool(), expected.Bool(), what)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if actual.Int() != expected.Int() {
			b.errorNotEquals(actual.Int(), expected.Int(), what)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if actual.Uint() != expected.Uint() {
			b.errorNotEquals(actual.Uint(), expected.Uint(), what)
		}
	case reflect.Float32, reflect.Float64:
		if math.Abs(actual.Float()-expected.Float()) > 1e-9 {
			b.errorNotEquals(actual.Float(), expected.Float(), what)
		}
	case reflect.Complex64, reflect.Complex128:
		if actual.Complex() != expected.Complex() {
			b.errorNotEquals(actual.Complex(), expected.Complex(), what)
		}
	case reflect.String:
		if actual.String() != expected.String() {
			b.errorNotEquals(actual.String(), expected.String(), what)
		}
	case reflect.Interface:
		b.equals(actual.Elem(), expected.Elem(), what)
	case reflect.Array, reflect.Slice:
		if actual.Len() != expected.Len() {
			b.errorNotEquals(actual.Len(), expected.Len(), fmt.Sprintf("len(%s)", what))
			return
		}
		for i := 0; i < actual.Len(); i++ {
			b.equals(actual.Index(i), expected.Index(i), fmt.Sprintf("%s[%d]", what, i))
		}
	case reflect.Map:
		if actual.Len() != expected.Len() {
			b.errorNotEquals(actual.Len(), expected.Len(), fmt.Sprintf("len(%s)", what))
			return
		}
		for _, k := range actual.MapKeys() {
			b.equals(actual.MapIndex(k), expected.MapIndex(k), fmt.Sprintf("%s[%s]", what, k))
		}
	case reflect.Struct:
		for i := 0; i < actual.NumField(); i++ {
			b.equals(
				actual.Field(i),
				expected.Field(i),
				fmt.Sprintf("%s.%s", what, actual.Type().Field(i).Name),
			)
		}
	case reflect.Chan:
		// TODO: TryRecv until OK on both and compare the two read values, then TryRecv one more time to check whether both are finished
	case reflect.Func:
		// TODO: Pointer()?: func: not necessarily unique, runtime.FuncForPC to get func name from pointer
	case reflect.Pointer:
		// TODO: Pointer()
	case reflect.UnsafePointer:
		// TODO: Pointer()?
	}
}

func (b *Bee) errorEquals(actual, expected any, what string) {
	b.tb.Helper()
	b.error(actual, expected, what, "==")
}

func (b *Bee) errorNotEquals(actual, expected any, what string) {
	b.tb.Helper()
	b.error(actual, expected, what, "!=")
}

func (b *Bee) error(actual, expected any, what, relation string) {
	b.tb.Helper()
	sActual := fmt.Sprintf("%v", actual)
	sExpected := fmt.Sprintf("%v", expected)
	format := "%s %s %s"
	args := []any{
		b.cfg.actualTextStyle.Render(wrap(b.tb, sActual, b.cfg.actualTextStyle.GetMaxWidth())),
		relation,
		b.cfg.expectedTextStyle.Render(wrap(b.tb, sExpected, b.cfg.expectedTextStyle.GetMaxWidth())),
	}
	if what != "" {
		format += " (%s)"
		args = append(args, b.cfg.whatTextStyle.Render(what))
	}
	b.tb.Errorf(format, args...)
	if (len(sActual) + len(sExpected)) > (b.cfg.expectedColumnStyle.GetWidth() + b.cfg.actualColumnStyle.GetWidth()) {
		b.tb.Logf(
			"\n%s",
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				b.cfg.actualColumnStyle.Render(sActual),
				b.cfg.expectedColumnStyle.Render(sExpected),
			),
		)
	}
}

func isNil(tb testing.TB, value any) bool {
	tb.Helper()
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		return v.IsNil()
	}
	return false
}

func wrap(tb testing.TB, s string, w int) string {
	tb.Helper()
	s = strings.ReplaceAll(s, "\n", "")
	if len(s) > w {
		s = fmt.Sprintf("%s...", s[:w-3])
	}
	return s
}

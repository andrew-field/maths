package maths

import (
	"errors"
	"testing"
)

func checkResult[T comparable](t testing.TB, want T, got T, gotError error) {
	t.Helper()
	if gotError != nil {
		t.Errorf("Got error but didn't want one. Error: %v", gotError)
	}

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}

func checkError(t testing.TB, gotError, wantError error) {
	t.Helper()
	if !errors.Is(gotError, wantError) {
		t.Errorf("Got: %v, want: %v", gotError, wantError)
	}
}

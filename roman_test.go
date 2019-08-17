package main

import "testing"

func TestRoman(t *testing.T) {
	assertEqual := func(t *testing.T, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("get empty string when 0 is given", func(t *testing.T) {
		assertEqual(t, "", ArabicToRoman(0))
	})

	t.Run("get empty string when -1 is given", func(t *testing.T) {
		assertEqual(t, "", ArabicToRoman(-1))
	})

	t.Run("get I when 1 is given", func(t *testing.T) {
		assertEqual(t, "I", ArabicToRoman(1))
	})

	t.Run("get II when 2 is given", func(t *testing.T) {
		assertEqual(t, "II", ArabicToRoman(2))
	})

	t.Run("get III when 3 is given", func(t *testing.T) {
		assertEqual(t, "III", ArabicToRoman(3))
	})

	t.Run("get IV when 4 is given", func(t *testing.T) {
		assertEqual(t, "IV", ArabicToRoman(4))
	})

	t.Run("get V when 5 is given", func(t *testing.T) {
		assertEqual(t, "V", ArabicToRoman(5))
	})

	t.Run("get VI when 6 is given", func(t *testing.T) {
		assertEqual(t, "VI", ArabicToRoman(6))
	})

	t.Run("get IX when 9 is given", func(t *testing.T) {
		assertEqual(t, "IX", ArabicToRoman(9))
	})

}

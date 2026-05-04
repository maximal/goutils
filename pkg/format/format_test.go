package format

import (
	"testing"
)

func TestSiFloat(t *testing.T) {
	cases := []struct {
		input float64
		want  string
	}{
		{input: 0.0, want: "0"},
		{input: 0.9, want: "0.9"},
		{input: 0.99, want: "0.99"},
		{input: 0.992, want: "0.99"},
		{input: 0.995, want: "0.99"},
		{input: 0.997, want: "1"},
		{input: 0.999, want: "1"},

		{input: 9.0, want: "9"},
		{input: 9.9, want: "9.9"},
		{input: 9.92, want: "9.92"},
		{input: 9.95, want: "9.95"},
		{input: 9.99, want: "9.99"},
		{input: 9.992, want: "9.99"},
		{input: 9.995, want: "9.99"},
		{input: 9.997, want: "10"},
		{input: 9.999, want: "10"},

		{input: 99.0, want: "99"},
		{input: 99.9, want: "99.9"},
		{input: 99.92, want: "99.9"},
		{input: 99.95, want: "100"},
		{input: 99.99, want: "100"},
		{input: 99.992, want: "100"},

		{input: 990_000_000, want: "990 M"},
		{input: 992_000_000, want: "992 M"},
		{input: 998_000_000, want: "998 M"},
		{input: 998_100_000, want: "1 G"},
		{input: 999_000_000, want: "1 G"},

		// Quetta (Q) is the largest SI prefix, so we should not go beyond it
		{input: 990_000_000_000_000_000_000_000_000_000_000, want: "990 Q"},
		{input: 999_000_000_000_000_000_000_000_000_000_000, want: "999 Q"},
		{input: 999_500_000_000_000_000_000_000_000_000_000, want: "1000 Q"},
		{input: 999_900_000_000_000_000_000_000_000_000_000, want: "1000 Q"},
		{input: 1_234_500_000_000_000_000_000_000_000_000_000, want: "1234 Q"},
		{input: 1_234_600_000_000_000_000_000_000_000_000_000, want: "1235 Q"},
		{input: 999_999_000_000_000_000_000_000_000_000_000_000, want: "999999 Q"},
	}

	for _, tc := range cases {
		t.Run("SiFloat", func(t *testing.T) {
			got := SiFloat(tc.input, "")

			if got != tc.want {
				t.Errorf("wrong result for %#v\ngot:  %#v\nwant: %#v", tc.input, got, tc.want)
			}
		})
	}
}

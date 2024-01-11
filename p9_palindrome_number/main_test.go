package main

import "testing"

type IsPalindromeTestCase struct {
	number int
	want   bool
}

func TestIsPalindrome(t *testing.T) {
	test_cases := []IsPalindromeTestCase{
		{
			number: 0123,
			want:   false,
		},
		{
			number: 121,
			want:   true,
		},
		{
			number: 987,
			want:   false,
		},
		{
			number: -121,
			want:   false,
		},
	}

	for _, test_case := range test_cases {
		x := test_case.number
		want := test_case.want

		got := IsPalindrome(test_case.number)

		if want != got {
			t.Fatalf(`IsPalindrome(%d) expected to return %v, got = %v`, x, want, got)
		}
	}
}

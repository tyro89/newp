package main

import "testing"

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"Foo", "foo"},
		{"foo", "foo"},
		{"FooBar", "foo_bar"},
		{"foo_bar", "foo_bar"},
		{"ABC", "a_b_c"},
		{"a_b_c", "a_b_c"},
		{"___a___b___c___", "a_b_c"},
	}

	for _, test := range tests {
		if actual := snakeCase(test.in); actual != test.expected {
			t.Errorf("snakeCase(%q) => %q, want %q", test.in, actual, test.expected)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"Foo", "Foo"},
		{"foo", "Foo"},
		{"FooBar", "FooBar"},
		{"foo_bar", "FooBar"},
		{"ABC", "ABC"},
		{"a_b_c", "ABC"},
		{"___a___b___c___", "ABC"},
	}

	for _, test := range tests {
		if actual := camelCase(test.in); actual != test.expected {
			t.Errorf("camelCase(%q) => %q, want %q", test.in, actual, test.expected)
		}
	}
}

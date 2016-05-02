package gostr

import (
	"testing"
	"reflect"
)

func TestAppend(t *testing.T) {
	for _, c := range []struct {
		in []string
		want string
	}{
		{
			[]string{"f", "o", "o", "b", "a", "r"},
			"foobar",
		},
		{
			[]string{"foobar"},
			"foobar",
		},
		{
			[]string{"", "foobar"},
			"foobar",
		},
	} {
		got := Append(c.in...)
		if got != c.want {
			t.Errorf("Append(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAppendArray(t *testing.T) {
	for _, c := range []struct {
		in []string
		want string
	}{
		{
			[]string{"f", "o", "o", "b", "a", "r"},
			"foobar",
		},
		{
			[]string{"foobar"},
			"foobar",
		},
		{
			[]string{"", "foobar"},
			"foobar",
		},
	} {
		got := AppendArray("", c.in)
		if got != c.want {
			t.Errorf("AppendArray(\"\", %q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAt(t *testing.T) {
	for _, c := range []struct {
		str string
		index int
		want string
	}{
		{"foobar", 0, "f"},
		{"ofobar", 1, "f"},
		{"oobarf", -1, "f"},
		{"oobafr", -2, "f"},
		{"Hello, 世界", -2, "世"},
	} {
		got := At(c.str, c.index)
		if got != c.want {
			t.Errorf("At(%q, %q) == %q, want %q", c.str, c.index, got, c.want)
		}
	}
}

func TestAtRunes(t *testing.T) {
	for _, c := range []struct {
		r []rune
		index int
		want string
	}{
		{[]rune("foobar"), 0, "f"},
		{[]rune("ofobar"), 1, "f"},
		{[]rune("oobarf"), -1, "f"},
		{[]rune("oobafr"), -2, "f"},
		{[]rune("Hello, 世界"), -2, "世"},
	} {
		got := AtRunes(c.r, c.index)
		if got != c.want {
			t.Errorf("AtRunes(%q, %q) == %q, want %q", c.r, c.index, got, c.want)
		}
	}
}

func TestBetween(t *testing.T) {
	for _, c := range []struct {
		str, start, end string
		want []string
	}{
		{"[foo]", "[", "]", []string{"foo"}},
		{"<span>foo</span>", "<span>", "</span>", []string{"foo"}},
		{"<span>bar</span><span>foo</span>", "<span>", "</span>", []string{"bar", "foo"}},
	} {
		got := Between(c.str, c.start, c.end)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Between(%q, %q, %q) == %q, want %q", c.str, c.start, c.end, got, c.want)
		}
	}
}

func TestRunes(t *testing.T) {
	for _, c := range []struct {
		str string
		want []rune
	}{
		{"foobar", []rune{'f', 'o', 'o', 'b', 'a', 'r'}},
		{"barfoo", []rune{'b', 'a', 'r', 'f', 'o', 'o'}},
		{"Hello, 世界", []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界'}},
	} {
		got := Runes(c.str)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Runes(%q) == %q, want %q", c.str, got, c.want)
		}
	}
}

func TestCollapseWhitespace(t *testing.T) {
	for _, c := range []struct {
		str, want string
	}{
		{"foo    bar", "foo bar"},
		{"Fòô     Bàř", "Fòô Bàř"},
		// {"   世  界   ", " 世 界 "}, //TODO
	} {
		got := CollapseWhitespace(c.str)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CollapseWhitespace(%q) == %q, want %q", c.str, got, c.want)
		}
	}
}

func TestContains(t *testing.T) {
		for _, c := range []struct {
		str, sub string
		caseSensitive bool
		want bool
	}{
		{"foo bar", "foo", false, true},
		{"bar FOO", "foo", true, false},
		{"foobar", "foo", false, true},
		{"FOO", "foo", false, true},
		{"Hello, 世界", "世", false, true},
	} {
		got := Contains(c.str, c.sub, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Contains(%q, %q, %t) == %t, want %t", c.str, c.sub, c.caseSensitive, got, c.want)
		}
	}
}

func TestContainsAll(t *testing.T) {
		for _, c := range []struct {
		str, chars string
		caseSensitive bool
		want bool
	}{
		{"foo bar", "f", false, true},
		{"bar FOO", "f", true, false},
		{"foobar", "afa", false, true},
		{"FOO", "foo", false, true},
		{"Hello, 世界", "世", false, true},
	} {
		got := ContainsAll(c.str, c.chars, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ContainsAll(%q, %q, %t) == %t, want %t", c.str, c.chars, c.caseSensitive, got, c.want)
		}
	}
}

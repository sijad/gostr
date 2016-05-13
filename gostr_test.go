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
		{[]rune("foobar"), -99, ""},
		{[]rune("foobar"), 99, ""},
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
		{"foo bar", "", false, false},
		{"bar foo", "", false, false},
		{"foobar", "", false, false},

		{"foo bar", "foo", false, true},
		{"bar foo", "bar", false, true},
		{"foobar", "oba", false, true},

		{"foo bar", "dleitee", false, false},
		{"bar foo", "dleitee", false, false},
		{"foobar", "dleitee", false, false},

		{"foo bar", "FOO", true, false},
		{"bar foo", "BAR", true, false},
		{"foobar", "OBA", true, false},

		{"foo bar", "FOO", false, true},
		{"bar foo", "BAR", false, true},
		{"foobar", "OBA", false, true},

		{"Hello, 世界", "世", false, true},
	} {
		got := ContainsAll(c.str, c.chars, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ContainsAll(%q, %q, %t) == %t, want %t", c.str, c.chars, c.caseSensitive, got, c.want)
		}
	}
}

func TestContainsAny(t *testing.T) {
		for _, c := range []struct {
		str, chars string
		caseSensitive bool
		want bool
	}{
		{"foo bar", "", false, false},
		{"bar foo", "", false, false},
		{"foobar", "", false, false},

		{"foo bar", "f", false, true},
		{"bar foo", "b", false, true},
		{"foobar", "o", false, true},

		{"foo bar", "dleitee", false, false},
		{"bar foo", "dleitee", false, false},
		{"foobar", "dleitee", false, false},

		{"foo bar", "F", true, false},
		{"bar foo", "B", true, false},
		{"foobar", "O", true, false},

		{"foo bar", "F", false, true},
		{"bar foo", "B", false, true},
		{"foobar", "O", false, true},

		{"Hello, 世界", "世", false, true},
	} {
		got := ContainsAny(c.str, c.chars, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ContainsAny(%q, %q, %t) == %t, want %t", c.str, c.chars, c.caseSensitive, got, c.want)
		}
	}
}

func TestCountSubstr(t *testing.T) {
		for _, c := range []struct {
		str, sep string
		caseSensitive bool
		want int
	}{
		{"aaaaaAaaAA", "a", true, 7},
		{"faaaAAaaaaAA", "a", true, 7},
		{"aaAAaaaaafA", "a", true, 7},
		{"AAaaafaaaaAAAA", "a", true, 7},

		{"aaaaaaa", "a", false, 7},
		{"faaaaaaa", "a", false, 7},
		{"aaaaaaaf", "a", false, 7},
		{"aaafaaaa", "a", false, 7},

		{"Hello, 世界", "世", false, 1},
	} {
		got := CountSubstr(c.str, c.sep, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CountSubstr(%q, %q, %t) == %d, want %d", c.str, c.sep, c.caseSensitive, got, c.want)
		}
	}
}

func TestEndsWith(t *testing.T) {
		for _, c := range []struct {
		str, search string
		caseSensitive bool
		want bool
	}{
		{"foo bar", "bar", false, true},
		{"bar", "bar", false, true},

		{"foo bar", "dleitee", false, false},
		{"bar", "dleitee", false, false},

		{"foo bar", "BAR", false, true},
		{"bar", "BAR", false, true},

		{"foo bar", "BAR", true, false},
		{"bar", "BAR", true, false},

		{"Hello, 世界", "界", false, true},
	} {
		got := EndsWith(c.str, c.search, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("EndsWith(%q, %q, %t) == %t, want %t", c.str, c.search, c.caseSensitive, got, c.want)
		}
	}
}

func TestStartsWith(t *testing.T) {
		for _, c := range []struct {
		str, search string
		caseSensitive bool
		want bool
	}{
		{"bar foo", "bar", false, true},
		{"bar", "bar", false, true},

		{"bar foo", "dleitee", false, false},
		{"bar", "dleitee", false, false},

		{"bar foo", "BAR", false, true},
		{"bar", "BAR", false, true},

		{"bar foo", "BAR", true, false},
		{"bar", "BAR", true, false},

		{"世界, Hello", "世", false, true},
	} {
		got := StartsWith(c.str, c.search, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("StartsWith(%q, %q, %t) == %t, want %t", c.str, c.search, c.caseSensitive, got, c.want)
		}
	}
}

func TestEnsureLeft(t *testing.T) {
	for _, c := range []struct {
		str, substr string
		caseSensitive bool
		want string
	}{
		{"bar", "foo", false, "foobar"},
		{"foobar", "foo", false, "foobar"},

		{"bar", "FOO", false, "FOObar"},
		{"foobar", "FOO", false, "foobar"},

		{"bar", "FOO", true, "FOObar"},
		{"foobar", "FOO", true, "FOOfoobar"},

		{"世界, Hello", "世", false, "世界, Hello"},
	} {
		got := EnsureLeft(c.str, c.substr, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("EnsureLeft(%q, %q) == %q, want %q", c.str, c.substr, got, c.want)
		}
	}
}

func TestEnsureRight(t *testing.T) {
	for _, c := range []struct {
		str, substr string
		caseSensitive bool
		want string
	}{
		{"foo", "bar", false, "foobar"},
		{"foobar", "bar", false, "foobar"},

		{"foo", "BAR", false, "fooBAR"},
		{"foobar", "BAR", false, "foobar"},

		{"foo", "BAR", true, "fooBAR"},
		{"foobar", "BAR", true, "foobarBAR"},

		{"Hello, 世界", "界", false, "Hello, 世界"},
	} {
		got := EnsureRight(c.str, c.substr, c.caseSensitive)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("EnsureRight(%q, %q) == %q, want %q", c.str, c.substr, got, c.want)
		}
	}
}

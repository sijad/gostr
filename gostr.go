package gostr

import (
	"math/rand"
	"strings"
	"unicode/utf8"
)

// Append appends given strings strs.
func Append(strs ...string) string {
	return strings.Join(strs, "")
}

// AppendArray appends an array of strings to the string s.
func AppendArray(s string, strs []string) string {
	return s + strings.Join(strs, "")
}

// At returns the character at the specified index in the string s.
func At(s string, index int) string {
	return AtRunes(Runes(s), index)
}

// AtRunes returns the character at the specified index in given array of rune.
func AtRunes(r []rune, index int) string {
	if index < 0 {
		index += len(r)
	}
	if index < 0 || index > len(r)-1 {
		return ""
	}
	return string(r[index])
}

// Between returns an array of strings between the start and end string.
func Between(s, start, end string) (result []string) {
	splited := strings.Split(s, end)
	startLen := Length(start)
	for i := 0; i < len(splited); i++ {
		r := splited[i]
		str := Substr(r, strings.Index(r, start)+startLen)
		// if Length(strings.TrimSpace(str)) > 0 {
		result = append(result, str)
		// }
	}
	return result[:len(result)-1]
}

// Runes converts the string s to an array of rune and returns it.
func Runes(s string) []rune {
	return []rune(s)
}

// CollapseWhitespace replaces consecutive whitespace characters with a single space in string s.
func CollapseWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Contains reports whether substr is within s.
func Contains(s, substr string, caseSensitive bool) bool {
	if caseSensitive {
		return strings.Contains(s, substr)
	}
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// ContainsAll reports whether all of Unicode code points in chars contained in s.
func ContainsAll(s, chars string, caseSensitive bool) bool {
	if chars == "" {
		return false
	}

	if !caseSensitive {
		s = strings.ToLower(s)
		chars = strings.ToLower(chars)
	}

	for _, c := range chars {
		if !Contains(s, string(c), false) {
			return false
		}
	}

	return true
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string, caseSensitive bool) bool {
	if caseSensitive {
		return strings.ContainsAny(s, chars)
	}
	return strings.ContainsAny(strings.ToLower(s), strings.ToLower(chars))
}

// CountSubstr counts the number of non-overlapping instances of sep in s.
func CountSubstr(s, sep string, caseSensitive bool) int {
	if caseSensitive {
		return strings.Count(s, sep)
	}
	return strings.Count(strings.ToLower(s), strings.ToLower(sep))
}

// EndsWith tests whether the string s ends with string search.
func EndsWith(s, search string, caseSensitive bool) bool {
	if caseSensitive {
		return strings.HasSuffix(s, search)
	}
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(search))
}

// StartsWith tests whether the string s starts with string search.
func StartsWith(s, search string, caseSensitive bool) bool {
	if caseSensitive {
		return strings.HasPrefix(s, search)
	}
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(search))
}

// EnsureLeft prepends substr to s if s doesn't starts with substr.
func EnsureLeft(s, substr string, caseSensitive bool) string {
	if !StartsWith(s, substr, caseSensitive) {
		return substr + s
	}
	return s
}

// EnsureRight appends substr to s if s doesn't ends with substr.
func EnsureRight(s, substr string, caseSensitive bool) string {
	if !EndsWith(s, substr, caseSensitive) {
		return s + substr
	}
	return s
}

// Insert insert substr into string s at the provided index.
func Insert(s, substr string, index int) string {
	if index > Length(s) {
		return s
	}

	start := SubstrLen(s, 0, index)
	end := SubstrLen(s, index, Length(s))

	return start + substr + end
}

// SurroundPad repeats characters of pad around given string s.
// number of repeats in left and right side specifies by leftPad and rightPad.
func SurroundPad(s, pad string, leftPad, rightPad int) string {
	r := Runes(pad)
	padLen := len(r)

	if padLen == 0 {
		return s
	}

	leftStr, RightStr := "", ""

	for i := 0; i < leftPad; i++ {
		leftStr += AtRunes(r, i%padLen)
	}

	for i := 0; i < rightPad; i++ {
		RightStr += AtRunes(r, i%padLen)
	}

	return leftStr + s + RightStr
}

// PaddingLeft pad a string to a certain length with another string in left side.
func PaddingLeft(s, pad string, length int) string {
	strLen := Length(s)
	if length < 0 || length <= strLen {
		return s
	}
	return SurroundPad(s, pad, length-strLen, 0)
}

// PaddingRight pad a string to a certain length with another string in right side.
func PaddingRight(s, pad string, length int) string {
	strLen := Length(s)
	if length < 0 || length <= strLen {
		return s
	}
	return SurroundPad(s, pad, 0, length-strLen)
}

// PaddingBoth pad a string to a certain length with another string in both right and left side.
func PaddingBoth(s, pad string, length int) string {
	strLen := Length(s)
	if length < 0 || length <= strLen {
		return s
	}
	padLen := length - strLen
	leftPad := padLen / 2
	return SurroundPad(s, pad, leftPad, padLen-leftPad)
}

// Prepend returns string s starting with strigns strs.
func Prepend(s string, strs ...string) string {
	return strings.Join(strs, "") + s
}

// PrependArray returns string s starting with array strigns strs.
func PrependArray(s string, strs []string) string {
	return strings.Join(strs, "") + s
}

// RemoveLeft remove substr from start of string s if present.
func RemoveLeft(s, substr string, caseSensitive bool) string {
	if StartsWith(s, substr, caseSensitive) {
		return Substr(s, Length(substr))
	}
	return s
}

// RemoveRight remove substr from end of string s if present.
func RemoveRight(s, substr string, caseSensitive bool) string {
	if EndsWith(s, substr, caseSensitive) {
		return SubstrLen(s, 0, Length(s)-Length(substr))
	}
	return s
}

// Shuffle returns string s with its characters in random order.
func Shuffle(s string) string {
	r := Runes(s)
	for i := range r {
		j := rand.Intn(i + 1)
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Surround surrounds strings with given string substr.
func Surround(s, substr string) string {
	return substr + s + substr
}

// Length returns number of characters stirng s.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// Substr returns part of string s.
func Substr(s string, start int) string {
	return SubstrLen(s, start, Length(s))
}

// SubstrLen returns part of string s.
func SubstrLen(s string, start, length int) (str string) {
	str = ""
	strLen := Length(s)

	if length < 0 && -length > strLen {
		return
	} else if length > strLen {
		length = strLen
	}

	if start > strLen {
		return
	} else if start < 0 && -start > strLen {
		start = 0
	}

	if length < 0 && (length+strLen-start) < 0 {
		return
	}

	if start < 0 {
		start = strLen + start
		if start < 0 {
			start = 0
		}
	}

	if length < 0 {
		length = strLen - start + length
		if length < 0 {
			length = 0
		}
	}

	if start+length > strLen {
		length = strLen - start
	}

	r := Runes(s)

	for i := start; i < start+length; i++ {
		str += AtRunes(r, i)
	}

	return
}

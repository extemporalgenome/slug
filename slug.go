// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slug transforms strings into a normalized
// form well suited for use in URLs.
package slug

import (
	"code.google.com/p/go.text/unicode/norm"
	"unicode"
)

var nop = []*unicode.RangeTable{unicode.Mark, unicode.Sk, unicode.Lm}

// Slug replaces each run of characters which are not unicode letters or
// numbers with a single hyphen, except for leading or trailing runs. Letters
// will be stripped of diacritical marks and lowercased.
func Slug(s string) string {
	buf := make([]rune, 0, len(s))
	dash := false
	for _, r := range norm.NFKD.String(s) {
		switch {
		// unicode 'letters' like mandarin characters pass through
		case unicode.IsLetter(r):
			r = unicode.ToLower(r)
			fallthrough
		case unicode.IsNumber(r):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(nop, r):
			// skip
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}
	return string(buf)
}

func IsSlugAscii(s string) bool {
	dash := true
	for _, r := range s {
		switch {
		case r == '-':
			if dash {
				return false
			}
			dash = true
		case 'a' <= r && r <= 'z', '0' <= r && r <= '9':
			dash = false
		}
	}
	return !dash
}

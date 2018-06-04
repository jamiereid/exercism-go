// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	// it took me **far** too long to think of stripping white space from the end of remark instead of trying to solve every problem with regex. It only occured to me when I started looking how other people had solved this problem and I came across Thomas's far more elegant solution (which honestly is close to how I first started to tackle this challenge: https://github.com/ThomasZumsteg/exercism-go/blob/master/bob/bob.go
	remark = strings.TrimSpace(remark)

	// tried using [:punct:] here instead of the brute force `\:\)` but couldn't work out character class subtractions, and golang doesn't support negative lookahead
	shout := regexp.MustCompile(`^[^a-z\:\)\d]*$|^[^a-z]*\!$`)
	question := regexp.MustCompile(`^.*\?$`)
	silence := regexp.MustCompile(`^\s*$|^$`)

	if silence.MatchString(remark) == true {
		return "Fine. Be that way!"
	} else if shout.MatchString(remark) == true {
		if question.MatchString(remark) == true {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else if question.MatchString(remark) == true {
		return "Sure."
	} else {
		return "Whatever."
	}
}

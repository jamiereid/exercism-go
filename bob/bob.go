// Package bob is an attempt to solve the exercism.io  koan/challenge 'bob'
package bob

import (
	"regexp"
	"strings"
)

// Hey is a function that supplies a response as a 'lackadaisical teenager'
func Hey(remark string) string {

	// it took me **far** too long to think of stripping white space from the end of remark
	// instead of trying to solve every problem with regex. It only occured to me when I
	// started looking how other people had solved this problem and I came across Thomas's
	// far more elegant solution (which honestly is close to how I first started to tackle
	// this challenge: https://github.com/ThomasZumsteg/exercism-go/blob/master/bob/bob.go
	remark = strings.TrimSpace(remark)

	// tried using [:punct:] here instead of the brute force `\:\)` but couldn't work out
	// character class subtractions, and golang doesn't support negative lookahead
	shout := regexp.MustCompile(`^[^a-z\:\)\d]*$|^[^a-z]*\!$`)
	question := regexp.MustCompile(`^.*\?$`)
	silence := regexp.MustCompile(`^\s*$|^$`)

	if silence.MatchString(remark) == true {
		return "Fine. Be that way!"
	}

	if shout.MatchString(remark) == true {
		if question.MatchString(remark) == true {
			return "Calm down, I know what I'm doing!"
		}

		return "Whoa, chill out!"

	}

	if question.MatchString(remark) == true {
		return "Sure."
	}

	return "Whatever."

}

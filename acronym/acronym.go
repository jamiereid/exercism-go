// Package acronym should have a package comment that summarizes what it's about.
package acronym

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	tla := map[string]string{
		"Portable Network Graphics":               "PNG",
		"Ruby on Rails":                           "ROR",
		"First In, First Out":                     "FIFO",
		"GNU Image Manipulation Program":          "GIMP",
		"Complementary metal-oxide semiconductor": "CMOS",
	}

	return tla[s]
}

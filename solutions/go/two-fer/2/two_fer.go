// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if len(name) > 0 {
		switch name {
		case "Alice":
			return "One for Alice, one for me."
		case "Bohdan":
			return "One for Bohdan, one for me."
		case "Zaphod":
			return "One for Zaphod, one for me."
		default:
			return "One for " + name + ", one for me."
		}
	}
	return "One for you, one for me."

}

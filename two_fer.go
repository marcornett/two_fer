/*
Package twofer creates a phrase containing the name of a person you are sharing with.

	Sharewith:
		Returns a string of a person you share with.
		Defaults to "you" if no name given.

*/
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if len(name) > 0 {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}

// Package sourcing contains state and ways to change it
package sourcing

// Parse an input onto a valir operation or an error.
func Parse(input string) (StateChanger, error) {
	return Configure{NewConfig: Trust}, nil
}

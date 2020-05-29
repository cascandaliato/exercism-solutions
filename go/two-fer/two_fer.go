// Package twofer provides functions to manage two-for-one bargains.
package twofer

import "fmt"

// ShareWith generates instructions to distribute goods between two counterparties fairly.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}

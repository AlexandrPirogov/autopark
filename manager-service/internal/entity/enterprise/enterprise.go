// package enteprise holds structs of Enterprise
// This can be used as a message for communication between services
package enterprise

import "strings"

type Enterprise struct {
	Title string `json:"title"`
}

// Compare compares given enterprise with the compared one by title
//
// Pre-cond: given enterprise to compare
//
// Post-cond: returns 0 if they are equal
// -1 if with greater than our
// 1 if out greater than given
func (e Enterprise) Compare(with Enterprise) int {
	return strings.Compare(e.Title, with.Title)
}

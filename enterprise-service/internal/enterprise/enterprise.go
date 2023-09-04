// enterprise package works with enterprise entity
package enterprise

import "strings"

type Enterprise struct {
	Title string `json:"title"`
}

func (e Enterprise) Compare(with Enterprise) int {
	return strings.Compare(e.Title, with.Title)
}

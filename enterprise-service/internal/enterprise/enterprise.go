// enterprise package works with enterprise entity
package enterprise

import (
	"enterprise-service/internal/client"
	"strings"
)

type Enterprise struct {
	ID       int              `json:"id"`
	Title    string           `json:"title"`
	Managers []client.Manager `json:"managers"`
}

func (e Enterprise) Compare(with Enterprise) int {
	return strings.Compare(e.Title, with.Title)
}

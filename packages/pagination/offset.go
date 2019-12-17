package pagination

import "strconv"

type offset struct {
	page         string
	itemsPerPage int
}

func (o offset) calculateOffset() int {
	cast, err := strconv.Atoi(o.page)
	if err == nil {
		return ((cast - 1) * o.itemsPerPage)
	}

	return 0
}

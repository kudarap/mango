package util

import (
	"strconv"
	"strings"

	"github.com/gorethink/gorethink"
)

// ParseOptOrder parse sort option
func ParseOptOrder(order string) interface{} {
	orders := strings.Split(strings.ToLower(strings.TrimSpace(order)), ",")
	if len(orders) == 2 && orders[1] == "desc" {
		// its desc
		return gorethink.Desc(orders[0])
	}

	// its asc
	return orders[0]
}

// ParseOptSlice parse slice option
func ParseOptSlice(slice string) (start int, end int) {
	slices := strings.Split(strings.TrimSpace(slice), ",")
	// check slices
	if len(slices) < 2 {
		return
	}

	start, _ = strconv.Atoi(slices[0])
	end, _ = strconv.Atoi(slices[1])

	return
}

// ParseOptField prase field option
func ParseOptField(field string) (s []string) {
	for _, f := range strings.Split(strings.TrimSpace(field), ",") {
		s = append(s, f)
	}

	return
}

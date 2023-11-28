package service

import (
	"noda/data/types"
	"strings"
)

func doTrim(args ...*string) (trimmed int) {
	if 0 == len(args) {
		return 0
	}
	for _, str := range args {
		if nil != str {
			*str = strings.Trim(*str, " \a\b\f\n\r\t\v")
			trimmed++
		}
	}
	return trimmed
}

func doDefaultPagination(pagination *types.Pagination) {
	if nil == pagination {
		return
	}
	if 0 >= pagination.Page {
		pagination.Page = 1
	}
	if 0 >= pagination.RPP {
		pagination.RPP = 10
	}
}

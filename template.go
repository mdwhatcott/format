package format

import (
	"fmt"
	"strings"
)

type Values map[string]interface{}

func Template(layout string, values Values) string {
	for key, value := range values {
		layout = strings.Replace(layout, "{"+key+"}", fmt.Sprint(value), -1)
	}
	return layout
}

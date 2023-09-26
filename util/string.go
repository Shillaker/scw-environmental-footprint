package util

import "strings"

func CompressJsonString(jsonIn string) string {
	jsonOut := strings.Replace(jsonIn, "\n", "", -1)
	jsonOut = strings.Replace(jsonOut, " ", "", -1)
	jsonOut = strings.Replace(jsonOut, "\t", "", -1)

	return jsonOut
}

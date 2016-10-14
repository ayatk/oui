package main

import (
	"fmt"
	"os"
	"strings"
)

func parseAddress(mac string) []string {
	ret := strings.Split(mac, ":")
	if len(ret) == 1 {
		ret[0] = ""
		return ret
	}
	for m := range ret {
		if len(ret[m]) == 1 {
			ret[m] = "0" + ret[m]
		} else if len(ret[m]) > 2 {
			fmt.Println("Parse error: Unauthorized address.")
			os.Exit(1)
		}
	}
	return ret
}

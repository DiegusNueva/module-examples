package main

import (
	"strings"

	"github.com/DiegusNueva/module-examples/slices"
	"rsc.io/quote/v3"
)

func main() {

	list := []string{"EDteam", "gophers", "golang", quote.HelloV3()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

}

package main

import (
	"fmt"

	cloudflareips "github.com/sch8ill/cloudflare-ips"
)

func main() {
	for _, cidr := range cloudflareips.MustFetch() {
		fmt.Println(cidr)
	}
}

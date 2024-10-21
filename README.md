# cloudflare-ips

[![Release](https://img.shields.io/github/release/sch8ill/cloudflare-ips.svg?style=flat-square)](https://github.com/sch8ill/cloudflare-ips/releases)
[![doc](https://img.shields.io/badge/go.dev-doc-007d9c?style=flat-square&logo=read-the-docs)](https://pkg.go.dev/github.com/sch8ill/cloudflare-ips)
[![Go Report Card](https://goreportcard.com/badge/github.com/sch8ill/cloudflare-ips)](https://goreportcard.com/report/github.com/sch8ill/cloudflare-ips)
![MIT license](https://img.shields.io/badge/license-MIT-green)

---

Automatically fetch and update Cloudflare proxy IPs, ensuring your trusted proxies are always up to date.

---


## Installation

```bash
go get github.com/sch8ill/cloudflare-ips
```

---

## Usage

```go
package main

import (
        "time"

        "github.com/gin-gonic/gin"
        cloudflareips "github.com/sch8ill/cloudflare-ips"
)

func main() {
        r := gin.Default()

        // Set the current clouflare reverse proxies as trusted.
        r.SetTrustedProxies(cloudflareips.MustFetch())

        // OR

        // Start a go routine that updates to trusted proxies list every 24 hours.
        cloudflareips.Updater(time.Hour*24, r.SetTrustedProxies)

        r.GET("/", func(c *gin.Context) {
                c.String(200, "Hello World!")
        })

        r.Run()
}
```

## License

This package is licensed under the [MIT License](LICENSE).

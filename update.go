package cloudflareips

import (
	"fmt"
	"log"
	"time"
)

const (
	updateRetries int = 3

	retryDelay time.Duration = time.Second * 30
)

// Updater fetches the current Cloudflare proxy ips once and starts a
// goroutine that updates the trusted proxies list at specified intervals.
func Updater(interval time.Duration, callback func([]string) error) {
	if err := retryUpdate(callback, updateRetries); err != nil {
		log.Println(err)
	}

	go func() {
		for {
			if err := retryUpdate(callback, updateRetries); err != nil {
				log.Println(err)
			}
			time.Sleep(interval)
		}
	}()
}

func retryUpdate(callback func([]string) error, retries int) error {
	if err := update(callback); err != nil {
		if retries > 0 {
			time.Sleep(retryDelay)
			return retryUpdate(callback, retries-1)
		}
		return err
	}

	return nil
}

func update(callback func([]string) error) error {
	cidrs, err := Fetch()
	if err != nil {
		return fmt.Errorf("failed to fetch cloudflare proxy ips: %w", err)
	}

	if err := callback(cidrs); err != nil {
		return fmt.Errorf("failed to set trusted proxies: %w", err)
	}

	return nil
}

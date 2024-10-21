package cloudflareips

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://api.cloudflare.com/client/v4/ips?networks=jdcloud"

type response struct {
	Errors   []error   `json:"errors"`
	Messages []message `json:"messages"`
	Success  bool      `json:"success"`
	Result   result    `json:"result"`
}

type apiError struct {
	Code    int    `json:"error"`
	Message string `json:"message"`
}

type message apiError

type result struct {
	Etag         string   `json:"etag"`
	IPv4Cidrs    []string `json:"ipv4_cidrs"`
	IPv6Cidrs    []string `json:"ipv6_cidrs"`
	JDCloudCidrs []string `json:"jdcloud_cidrs"`
}

// Fetch retrieves the Cloudflare proxy IPs (IPv4, IPv6, and JDCloud CIDRs).
func Fetch() ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp := new(response)
	if err := json.NewDecoder(resp.Body).Decode(apiResp); err != nil {
		return nil, err
	}

	if apiResp.Success == false || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("api error: %d: %+v", resp.StatusCode, apiResp.Errors)
	}

	cidrs := append(
		apiResp.Result.IPv4Cidrs,
		append(
			apiResp.Result.IPv6Cidrs,
			apiResp.Result.JDCloudCidrs...,
		)...,
	)

	return cidrs, nil
}

// MustFetch retrieves the Cloudflare proxy IPs and panics if an error occurs.
func MustFetch() []string {
	cidrs, err := Fetch()
	if err != nil {
		panic(err)
	}
	return cidrs
}

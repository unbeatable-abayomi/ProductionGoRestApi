//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/health")
	if err != nil {
		t.Fail()
	}
	//fmt.Println(resp.StatusCode())
	assert.Equal(t, 200, resp.StatusCode())
}

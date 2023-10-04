//go:build e2e
// +build e2e

package test


import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)


func TestGetComment(t *testing.T){
	fmt.Println("Running E2E test for get comments  endpoint")
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment");
	if err != nil {
		t.Fail()
	}
		//fmt.Println(resp.StatusCode())
		assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T){
	fmt.Println("Running E2E test for post comment endpoint")
	client := resty.New()
	resp, err := client.R().SetBody(`{"slug":"/", "author":"123456","body":"hello world"}`).Post(BASE_URL + "/api/comment");
	assert.Equal(t, 200, resp.StatusCode())

	assert.NoError(t, err)
}
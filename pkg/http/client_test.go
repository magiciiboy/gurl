package http

import (
	"testing"

	"github.com/magiciiboy/gurl/pkg/url"
)

type ClientTestCase struct {
	req    Request
	client *Client
}

var clientTestSet = []ClientTestCase{
	{
		req: Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "http",
				Host:   "linktree.magicii.workers.dev",
				Port:   80,
				Path:   "/links",
				URL:    "https://linktree.magicii.workers.dev/links",
			},
			Headers:      Header{"User-Agent": []string{defaultUserAgent}},
			ProtoVersion: &HTTPVersion1_1,
		},
		client: DefaultClient,
	},
	{
		req: Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "https",
				Host:   "linktree.magicii.workers.dev",
				Port:   80,
				Path:   "/links",
				URL:    "https://linktree.magicii.workers.dev/",
			},
			Headers:      Header{"User-Agent": []string{defaultUserAgent}},
			ProtoVersion: &HTTPVersion1_1,
		},
		client: DefaultClient,
	},
}

func TestClientSendRequest(t *testing.T) {
	for _, tc := range clientTestSet {
		_, err := tc.client.SendRequest(&tc.req)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
}

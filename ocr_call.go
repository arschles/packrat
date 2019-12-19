package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type cogSvcsEndpoint struct {
	url   string
	subID string
}

func makeOCRCall(
	ctx context.Context,
	cl *http.Client,
	endpoint *cogSvcsEndpoint,
	rdr io.Reader,
) (*http.Response, error) {
	fullURL := fmt.Sprintf("%s?language=unk&detectOrientation=true", endpoint.url)
	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, rdr)
	req.Header.Add("Ocp-Apim-Subscription-Key", endpoint.subID)
	req.Header.Add("Content-Type", "application/octet-stream")
	if err != nil {
		return nil, err
	}
	resp, err := cl.Do(req)
	return resp, err
}

package ocr

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// CognitiveServicesEndpoint represents an endpoint to Azure Cognitive Services
type CognitiveServicesEndpoint struct {
	URL   string
	SubID string
}

// Run sends a request to Azure Cognitive Services for OCR
func Run(
	ctx context.Context,
	cl *http.Client,
	endpoint *CognitiveServicesEndpoint,
	rdr io.Reader,
) (*http.Response, error) {
	fullURL := fmt.Sprintf("%s?language=unk&detectOrientation=true", endpoint.URL)
	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, rdr)
	req.Header.Add("Ocp-Apim-Subscription-Key", endpoint.SubID)
	req.Header.Add("Content-Type", "application/octet-stream")
	if err != nil {
		return nil, err
	}
	resp, err := cl.Do(req)
	return resp, err
}

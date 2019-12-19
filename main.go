package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gobuffalo/envy"
)

// ported from:
// https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/quickstarts/csharp-print-text
func main() {
	filename, err := envy.MustGet("FILENAME")
	if err != nil {
		log.Fatal("You forgot to set FILENAME!")
	}
	subID, err := envy.MustGet("AZURE_SUBSCRIPTION_ID")
	if err != nil {
		log.Fatal("You forgot to set AZURE_SUBSCRIPTION_ID!")
	}
	endpoint, err := envy.MustGet("AZURE_COG_SVCS_ENDPOINT")
	if err != nil {
		log.Fatal("You forgot to set AZURE_COG_SVCS_ENDPOINT!")
	}

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Couldn't read file %s (%s)", filename, err)
	}
	ctx := context.Background()
	resp, err := makeOCRCall(
		ctx,
		http.DefaultClient,
		&cogSvcsEndpoint{
			url:   fmt.Sprintf("%s/vision/v2.1/ocr", endpoint),
			subID: subID,
		},
		bytes.NewReader(fileBytes),
	)
	if err != nil {
		log.Fatalf("Error with Cognitive Services (%s)", err)
	}
	defer resp.Body.Close()
	dc := decoded{}
	if err := json.NewDecoder(resp.Body).Decode(&dc); err != nil {
		log.Fatalf("Error decoding json (%s)", err)
	}
	fmt.Print(strings.Join(dc.allText(), "\n"))
}

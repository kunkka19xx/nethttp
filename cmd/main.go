package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello noobs")
	requestURL := "https://www.netflix.com/jp-en/"
	reqBody:= []byte(`{"appName":"android","events":[{"name":"advdevtag","data":{"advdevtag_type":"web","ad_vendor_sync_type":"netflix","event_type":"nmLanding","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:134.0) Gecko/20100101 Firefox/134.0","tags":["adwords_Simplicity_NMLanding","fb_simplicity_testPixel","fb_simplicity_nmLanding","tiktok_nmLanding"],"params":{"membership_status":"ANONYMOUS","country":"JP","region_code":"11","is_member":"ANONYMOUS","wasFormerMember":false,"referrer":"nmLanding","deniedConsentCookieGroups":"C0005"},"country":"JP","tagLoaderVersion":"20240924-1446"}}]}`)
	bodyReader := bytes.NewReader(reqBody)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}

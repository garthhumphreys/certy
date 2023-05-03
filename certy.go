package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CRTShCertificate struct {
	NameValue string `json:"common_name"`
}

func main() {
	domainPtr := flag.String("domain", "garthhumphreys.com", "a domain")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected 'domain'")
		os.Exit(1)
	}

	if *domainPtr != "" {
		fmt.Println("site:", *domainPtr)
		certSearch(*domainPtr)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func certSearch(domain string) {
	// domain := "example.com"
	url := fmt.Sprintf("https://crt.sh/?q=%s.%s&output=json", "%", domain)
	fmt.Println("URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("RESP:", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var certificates []CRTShCertificate
	err = json.Unmarshal(body, &certificates)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	unique := make(map[string]bool)
	for _, certificate := range certificates {
		if _, ok := unique[certificate.NameValue]; !ok {
			unique[certificate.NameValue] = true
			fmt.Println(certificate.NameValue)
		}
	}

	// prettyJSON, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	//
	// fmt.Println(string(prettyJSON))
}

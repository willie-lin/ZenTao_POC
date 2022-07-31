package ZenTao_POC

import (
	"flag"
	"fmt"
)

var (
	url        string
	Zentaopath = "/zentao/user-login.html"
)

func init() {
	flag.StringVar(&url,
		"u",
		"null",
		"url:http://127.0.0.1/")
}

func main() {
	flag.Parse()
	fmt.Println("author:willie-lin")
	if url != "null" {
		ZenTao()
		return
	}
	fmt.Println("ZenTao_POC: usage_poc:./ZentaoSqli.exe -u url")
}

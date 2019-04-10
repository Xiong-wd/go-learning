package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err := http.Get("https://www.zhipin.com/job_detail/?query=golang&city=101040100&industry=&position=")

	if err != nil {

	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error ",resp.StatusCode)
		return
	}
	//utf8R := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	all,err := ioutil.ReadAll(resp.Body)
	if err !=  nil {

	}
	fmt.Printf("%s\n",all)


}

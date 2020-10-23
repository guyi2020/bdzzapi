package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// 抓取博客文章url
func spider() ([]string, error) {

	url := "http://www.ancientone.cn/wp-json/wp/v2/posts"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res []map[string]interface{}

	if err := json.Unmarshal(body, &res); err == nil {
		var links []string
		for _, v := range res {
			links = append(links, v["link"].(string))
		}
		return links, nil
	}
	return nil, err
}

func main() {

	// 获取资源链接
	links, err := spider()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	url := "http://data.zz.baidu.com/urls?site=www.ancientone.cn&token=YHNkWOWa5e1k1ZTi"

	// 拼接格式
	urls := strings.Join(links, "\n")

	reader := bytes.NewReader([]byte(urls))
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println("post request err", err.Error())
		return
	}
	request.Header.Set("Content-Type", "Content-Type:text/plain;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("resp err", err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logFile, err := os.OpenFile("/home/soft/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err.Error())
	}
	// 将文件设置为log输出的文件
	log.SetOutput(logFile)
	// 输出前缀
	log.SetPrefix("[log]")
	// log格式
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log输出到文件
	log.Printf("%s", respBytes)
	defer logFile.Close()
}

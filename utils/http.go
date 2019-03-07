package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//HTTPGet 提交 get 请求
func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//HttpPost 提交post请求
func HttpPost(uri string, params map[string]string) ([]byte, error) {

	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}

	response, err := http.PostForm(uri, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
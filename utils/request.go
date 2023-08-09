package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	defaultDialTimeout = 15 * time.Second
)

// GetRequest GetRequest with Authorization.
func GetRequestWithAuth(method, url, version string, params map[string]interface{}, basicAuth string) ([]byte, error) {
	var apiURL string
	if params == nil {
		apiURL = url
	} else {
		strParams := Map2UrlParams(params)
		apiURL = url + "?" + strParams
	}
	request, err := http.NewRequest(method, apiURL, nil)
	if err != nil {
		return nil, err
	}
	if version != "" {
		request.Header.Set("Accept", "application/json;v="+version)
	} else {
		request.Header.Set("Accept", "application/json")
	}
	if basicAuth != "" {
		request.Header.Set("Authorization", basicAuth)
	}
	fmt.Printf("%v", request.Header)
	client := http.Client{Timeout: defaultDialTimeout}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetRequest GetRequest.
func GetRequest(method, url, version string, params map[string]interface{}) ([]byte, error) {
	var apiURL string
	if params == nil {
		apiURL = url
	} else {
		strParams := Map2UrlParams(params)
		apiURL = url + "?" + strParams
	}
	request, err := http.NewRequest(method, apiURL, nil)
	if err != nil {
		return nil, err
	}
	if version != "" {
		request.Header.Set("Accept", "application/json;v="+version)
	} else {
		request.Header.Set("Accept", "application/json")
	}

	client := http.Client{Timeout: defaultDialTimeout}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// PostRequest Post Request.
func PostRequest(url, version string, params map[string]interface{}) ([]byte, error) {
	jsonParams := ""
	if params != nil {
		bytesData, _ := json.Marshal(params)
		jsonParams = string(bytesData)
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(jsonParams))
	if err != nil {
		return nil, err
	}
	if version != "" {
		request.Header.Set("Content-Type", "application/json;v="+version)
	} else {
		request.Header.Set("Content-Type", "application/json")

	}
	client := &http.Client{Timeout: defaultDialTimeout}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// PostRequest Post Request with Authorization.
func PostRequestWithAuth(url, version string, params map[string]interface{}, basicAuth string) ([]byte, error) {
	jsonParams := ""
	if params != nil {
		bytesData, _ := json.Marshal(params)
		jsonParams = string(bytesData)
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(jsonParams))
	if err != nil {
		return nil, err
	}
	if version != "" {
		request.Header.Set("Content-Type", "application/json;v="+version)
	} else {
		request.Header.Set("Content-Type", "application/json")
	}
	if basicAuth != "" {
		request.Header.Set("Authorization", basicAuth)
	}
	fmt.Printf("%v", request.Header)
	client := &http.Client{Timeout: defaultDialTimeout}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Map2UrlParams Map 2Url Params.
func Map2UrlParams(params map[string]interface{}) string {
	var strParams string
	for k, v := range params {
		value := ToString(v)
		strParams += (k + "=" + value + "&")
	}
	if len(strParams) > 0 {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}
	return strParams
}

// ToString To tring.
func ToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case bool:
		return strconv.FormatBool(v.(bool))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	default:
		return fmt.Sprintf("%", v)
	}
}

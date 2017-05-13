package front

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func httpCall(method string, url string, body io.Reader) (data []byte, status int, err error) {
	return httpCallWithHeaders(method, url, body, make(map[string]string))
}

func httpCallWithAuthToken(method string, url string, body io.Reader, token string) (data []byte, status int, err error) {
	return httpCallWithHeaders(method, url, body, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	})
}

func httpCallWithHeaders(method string, url string, body io.Reader, headers map[string]string) (
	data []byte, status int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("couldn't instantiate new HTTP %s request to %v: %v", method, url, err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	// add additional headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	fmt.Printf(">>>>>>>>>>>> %+v", req)
	resp, err := client.Do(req)
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		err = fmt.Errorf("HTTP request failed for request %+v: %+v", req, err)
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	status = resp.StatusCode
	return
}

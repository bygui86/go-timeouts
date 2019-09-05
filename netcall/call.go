package netcall

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// real
	// url = "https://jsonplaceholder.typicode.com/todos/1"
	// mock
	url = "http://localhost:3000/todos/1"
)

type CallResponse struct {
	Resp *Response
	Err  error
}

func GetHttpResponse(ctx context.Context, enableTimeout bool) (*Response, error) {

	if enableTimeout {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context timeout, ran out of time")
		case respChan := <-withTimeout(ctx):
			return respChan.Resp, respChan.Err
		}
	} else {
		return noTimeout()
	}
}

func noTimeout() (*Response, error) {

	resp, respErr := http.Get(url)
	if respErr != nil {
		return nil, fmt.Errorf("error in http call")
	}
	defer resp.Body.Close()

	byteResp, byteErr := ioutil.ReadAll(resp.Body)
	if byteErr != nil {
		return nil, fmt.Errorf("error in reading response")
	}

	structResp := &Response{}
	unmErr := json.Unmarshal(byteResp, structResp)
	if unmErr != nil {
		return nil, fmt.Errorf("error in unmarshalling response")
	}

	return structResp, nil
}

func withTimeout(ctx context.Context) <-chan *CallResponse {

	respChan := make(chan *CallResponse, 1)

	go func() {
		resp, respErr := http.Get(url)
		if respErr != nil {
			respChan <- &CallResponse{nil, fmt.Errorf("error in http call")}
			return
		}
		defer resp.Body.Close()

		byteResp, byteErr := ioutil.ReadAll(resp.Body)
		if byteErr != nil {
			respChan <- &CallResponse{nil, fmt.Errorf("error in reading response")}
			return
		}

		structResp := &Response{}
		unmErr := json.Unmarshal(byteResp, structResp)
		if unmErr != nil {
			respChan <- &CallResponse{nil, fmt.Errorf("error in unmarshalling response")}
		}

		respChan <- &CallResponse{structResp, nil}
	}()

	return respChan
}

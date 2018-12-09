package rest

import (
	"bytes"
	"net/http"

	"github.com/senfix/logger"
)

type Caller interface {
	Get(action string, params map[string]interface{}) (buffer []byte, err error)
}

func NewCaller(log logger.Log) Caller {
	return &caller{log.Enable("REST")}
}

type caller struct {
	log logger.Log
}

func (c *caller) Get(action string, params map[string]interface{}) (buffer []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, action, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	buffer = []byte{}
	response, err := client.Do(req)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}
	buffer = buf.Bytes()
	defer response.Body.Close()
	return
}

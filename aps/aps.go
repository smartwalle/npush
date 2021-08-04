package aps

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

type Client struct {
	domain           string
	bundleIdentifier string
	Client           *http.Client
}

func New(bundleIdentifier, p12, password string, isProduction bool) (*Client, error) {
	client, err := NewClient(p12, password)
	if err != nil {
		return nil, err
	}

	var p = &Client{}
	p.bundleIdentifier = bundleIdentifier
	p.Client = client
	if isProduction {
		p.domain = kProductionURL
	} else {
		p.domain = kDevelopmentURL
	}
	return p, nil
}

func NewClient(p12, password string) (*http.Client, error) {
	cert, err := Load(p12, password)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	config.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: config}

	if err = http2.ConfigureTransport(transport); err != nil {
		return nil, err
	}

	return &http.Client{Transport: transport}, nil
}

func (this *Client) Push(deviceToken string, header *Header, payload Payload) (result string, err error) {
	pBytes, err := json.Marshal(payload.toMap())
	if err != nil {
		return "", err
	}

	if len(pBytes) > kMaxPayload {
		return "", errors.New("payload too large")
	}

	var urlStr = fmt.Sprintf(this.domain, deviceToken)
	//fmt.Println(urlStr)

	req, err := http.NewRequest("POST", urlStr, bytes.NewReader(pBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if header == nil {
		header = &Header{}
	}
	if header.Topic == "" {
		header.Topic = this.bundleIdentifier
	}
	header.set(req.Header)

	rsp, err := this.Client.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	if rsp.StatusCode == http.StatusOK {
		return rsp.Header.Get("apns-id"), nil
	}

	data, err := ioutil.ReadAll(rsp.Body)
	var pushRsp *PushResponse
	if err = json.Unmarshal(data, &pushRsp); err != nil {
		return "", err
	}
	return "", errors.New(pushRsp.Reason)
}

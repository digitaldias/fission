/*
Copyright 2016 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fission/fission"
	builder "github.com/fission/fission/builder"
)

type (
	Client struct {
		url string
	}
)

func MakeClient(serverUrl string) *Client {
	return &Client{
		url: strings.TrimSuffix(serverUrl, "/"),
	}
}

func (c *Client) Build(req *builder.PackageBuildRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fission.MakeErrorFromHTTP(resp)
	}
	return nil
}

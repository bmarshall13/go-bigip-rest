/*
 * BigIP iControl REST
 *
 * REST API for F5 BigIP. List of operations is not complete, nor known to be accurate.
 *
 * OpenAPI spec version: 12.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Client library for BigIP iControl REST
//
// REST API for F5 BigIP. List of operations is not complete, nor known to be accurate.
//
// API version: 12.0
// Package version: v0.0.1-3-g67a31cb
//
// Use NewClient to create a new client handle
//
// For more information, please visit [https://devcentral.f5.com/](https://devcentral.f5.com/)
//
// This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.
package f5api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/url"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-resty/resty"
)

// Client is the REST handle that should be used for all REST interactions
// It maintains connections, and login operations. All methods will generally
// make a query to the REST API server, unless otherwise noted
type Client struct {
	configuration *configuration

	Ltm    LtmApi
	Net    NetApi
	Shared SharedApi
	Sys    SysApi
}

// NewClient creates a new Clinet object.
// No queries are performed.
func NewClient(hostName string, username string, password string, skipTlsVerify bool) *Client {
	cfg := newConfiguration()
	cfg.BasePath = fmt.Sprintf("https://%s/mgmt", hostName)
	cfg.Password = password
	cfg.UserName = username

	r := resty.New()
	r.SetDebug(cfg.debug)
	r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: skipTlsVerify})
	cfg.restClient.Resty = r

	return &Client{
		configuration: cfg,

		Ltm: LtmApi{
			configuration: cfg,
		},
		Net: NetApi{
			configuration: cfg,
		},
		Shared: SharedApi{
			configuration: cfg,
		},
		Sys: SysApi{
			configuration: cfg,
		},
	}
}

// DoLogin performs the remote operations to obtain an auth token from the API server
func (self *Client) DoLogin() error {
	res, err := self.Shared.Login(LoginBody{
		LoginProviderName: "tmos",
		Password:          self.configuration.Password,
		Username:          self.configuration.UserName,
	})
	if err != nil {
		return err
	}

	self.configuration.APIKey["X-F5-Auth-Token"] = res.Token.Token
	return nil
}

// Set or unset debug mode. Debug mode will print debugging information,
// including all REST HTTP messages.
func (self *Client) SetDebug(debug bool) {
	self.configuration.SetDebug(debug)
}

type restClient struct {
	Resty *resty.Client
}

func (c *restClient) selectHeaderContentType(contentTypes []string) string {

	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *restClient) SelectHeaderAccept(accepts []string) string {

	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(source []string, containvalue string) bool {
	for _, a := range source {
		if strings.ToLower(a) == strings.ToLower(containvalue) {
			return true
		}
	}
	return false
}

func (c *restClient) callAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) (*resty.Response, error) {

	request := prepareRequest(c.Resty.R(), postBody, headerParams, queryParams, formParams, fileName, fileBytes)

	switch strings.ToUpper(method) {
	case "GET":
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err
	}

	return nil, fmt.Errorf("invalid method %v", method)
}

func (c *restClient) parameterToString(obj interface{}, collectionFormat string) string {
	if reflect.TypeOf(obj).String() == "[]string" {
		switch collectionFormat {
		case "pipes":
			return strings.Join(obj.([]string), "|")
		case "ssv":
			return strings.Join(obj.([]string), " ")
		case "tsv":
			return strings.Join(obj.([]string), "\t")
		case "csv":
			return strings.Join(obj.([]string), ",")
		}
	}

	return fmt.Sprintf("%v", obj)
}

func prepareRequest(request *resty.Request,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) *resty.Request {

	request.SetBody(postBody)

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetMultiValueQueryParams(queryParams)
	}

	// add form parameter, if any
	if len(formParams) > 0 {
		request.SetFormData(formParams)
	}

	if len(fileBytes) > 0 && fileName != "" {
		_, fileNm := filepath.Split(fileName)
		request.SetFileReader("file", fileNm, bytes.NewReader(fileBytes))
	}
	return request
}

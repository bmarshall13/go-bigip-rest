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

package f5api

import (
	"fmt"
	"net/url"
	"strings"

	"encoding/json"
)

// This class holds all the API methods for the Shared API sub tree
type SharedApi struct {
	configuration *configuration
}

func unused_import_hack_SharedApi() {
	strings.Replace("a", "a", fmt.Sprintf("%v", nil), -1)
}

// Login
//
//
// Login to generate an API token.
//
// loginBody is for User to authenticate as.
func (a SharedApi) Login(loginBody LoginBody) (*LoginResp, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.configuration.BasePath + "/shared/authn/login"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(basicAuth)' required
	// http basic authentication required
	if a.configuration.UserName != "" || a.configuration.Password != "" {
		headerParams["Authorization"] = "Basic " + a.configuration.getBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.configuration.DefaultHeader {
		headerParams[key] = a.configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{
		"application/json",
	}

	// set Content-Type header
	localVarHttpContentType := a.configuration.restClient.selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.configuration.restClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &loginBody

	httpResponse, err := a.configuration.restClient.callAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	var successPayload = new(LoginResp)
	if err == nil && httpResponse.StatusCode() == 200 {
		err = json.Unmarshal(httpResponse.Body(), &successPayload)
	}
	err = NewAPIResponse(httpResponse, err)
	if err != nil {
		return nil, err
	} else {
		return successPayload, err
	}

}

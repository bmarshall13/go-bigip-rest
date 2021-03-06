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

// This describes a message sent to or received from some operations
type LtmProfileAnalyticsTrafficCapture struct {
	AppService string `json:"appService,omitempty"`

	// Adds, deletes, or replaces a set of HTTP methods from/to which captured traffic is sent.
	Methods string `json:"methods,omitempty"`

	// Specifies the string by which the response data is filtered or none.
	ResponseContentFilterSearchString string `json:"responseContentFilterSearchString,omitempty"`

	// Specifies which part of the request is filtered by a specific string.
	RequestContentFilterSearchPart string `json:"requestContentFilterSearchPart,omitempty"`

	// Specifies the string by which the request data is filtered, or none.
	RequestContentFilterSearchString string `json:"requestContentFilterSearchString,omitempty"`

	// Adds, deletes, or replaces a set of HTTP response codes from which traffic is captured.
	ResponseCodes int64 `json:"responseCodes,omitempty"`

	// Specifies whether the system captures traffic data mitigated by DoS Layer 7 Enforcer or regardless of DoS activity.
	DosActivity string `json:"dosActivity,omitempty"`

	// Adds, deletes, or replaces a set of URL path prefixes of which to capture traffic.
	UrlPathPrefixes string `json:"urlPathPrefixes,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies what parts of the response data the system captures.
	ResponseCapturedParts string `json:"responseCapturedParts,omitempty"`

	// Adds, deletes, or replaces a set of client IPs from/to which captured traffic is sent.
	ClientIps string `json:"clientIps,omitempty"`

	// Specifies which part of the response is filtered by a specific string.
	ResponseContentFilterSearchPart string `json:"responseContentFilterSearchPart,omitempty"`

	// Specifies whether the system captures traffic data sent using all protocols, or only one type of protocol.
	CapturedProtocols string `json:"capturedProtocols,omitempty"`

	// Specifies what parts of the request data the system captures.
	RequestCapturedParts string `json:"requestCapturedParts,omitempty"`

	// Adds, deletes, or replaces a set of user agent substrings of which to capture traffic.
	UserAgentSubstrings string `json:"userAgentSubstrings,omitempty"`
}

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
type LtmProfileResponseAdapt struct {
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the maximum size of the preview buffer. The preview buffer is used to hold a copy of the HTTP response header and data sent to the internal virtual server in case the adaptation server reports that it does not need to adapt the HTTP response. Setting the preview-size to zero, disables buffering the response and should only be done if the adaptation server will always return with a modified HTTP response or the original HTTP response. The default value is 1024.
	PreviewSize int64 `json:"previewSize,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the action to take if the internal virtual server does not exist or returns an error. The default value is ignore.
	ServiceDownAction string `json:"serviceDownAction,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Enables adaptation of HTTP responses. If set to yes, HTTP responses will be forwarded to the specified internal virtual server for adaptation. The default value is yes.
	Enabled string `json:"enabled,omitempty"`

	// Specifies whether to forward HTTP version 1.0 responses for adaptation. By default only HTTP version 1.1 responses are forwarded. Version 1.0 is not supported. While it should work in most cases, it might be necessary to restrict adaptation on a site-specific basis. The default value is no.
	AllowHttp10 string `json:"allowHttp_10,omitempty"`

	// Specifies a timeout in milliseconds. If the internal virtual server does not return a result within the specified time, a timeout error will occur. A 0 value disables the timeout. The default value is 0.
	Timeout int64 `json:"timeout,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the name of the internal virtual server to use for adapting the HTTP response.
	InternalVirtual string `json:"internalVirtual,omitempty"`
}

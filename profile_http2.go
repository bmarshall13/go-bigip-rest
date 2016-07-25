/* 
 * BigIP iControl REST
 *
 * REST API for F5 BigIP. Only LTM is included, and most endpoints are untested.
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

type ProfileHttp2 struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
	ActivationModes string `json:"activationModes,omitempty"`

	// Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
	ConcurrentStreamsPerConnection int64 `json:"concurrentStreamsPerConnection,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	EnforceTlsRequirements string `json:"enforceTlsRequirements,omitempty"`

	// Specifies what table size will be used for the compression of headers (unused).
	HeaderTableSize int64 `json:"headerTableSize,omitempty"`

	// The size in bytes of the data frames that will be produced by HTTP/2.
	FrameSize int64 `json:"frameSize,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies in KB the size of the receive window for HTTP/2 flow-control.
	ReceiveWindow int64 `json:"receiveWindow,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// Specifies whether an HTTP header should be added to the HTTP request to show the request was received via HTTP/2.
	InsertHeader string `json:"insertHeader,omitempty"`

	// Specifies the name of the header that is added to the HTTP request when insert-header is enabled.
	InsertHeaderName string `json:"insertHeaderName,omitempty"`

	// The size in bytes of the SSL records that will be produced by HTTP/2.handled.
	WriteSize int64 `json:"writeSize,omitempty"`

	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion.
	ConnectionIdleTimeout int64 `json:"connectionIdleTimeout,omitempty"`
}

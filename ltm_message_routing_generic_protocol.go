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
type LtmMessageRoutingGenericProtocol struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// When set, the generic message profile parser will be disabled. It will ignore all incoming packets and not directly send message data. This mode supports iRule script protocol implementations that will generate messages from the incoming transport stream and send outgoing messages on the outgoing transport stream.
	DisableParser string `json:"disableParser,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// The string of characters used to terminate a message. If the message-terminator is empty, the generic message parser will not separate the input stream into messages.
	MessageTerminator string `json:"messageTerminator,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Specifies the maximum size of a received message. If a message exceeds this size, the connection will be reset.
	MaxMessageSize int64 `json:"maxMessageSize,omitempty"`

	// When set, matching of responses to requests is disabled.
	NoResponse string `json:"noResponse,omitempty"`

	// Specifies the maximum size of the send buffer in bytes. If the number of bytes in the send buffer for a connection exceeds this value, the generic message protocol will stop receiving outgoing messages from the router until the size of the size of the buffer drops below this setting.
	MaxEgressBuffer int64 `json:"maxEgressBuffer,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all of the settings and values from the specified parent profile.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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

type MessageRoutingDiameterProfileRouter struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Specifies descriptive text that identifies the profile.
	Description string `json:"description,omitempty"`

	// Specifies the maximum number of bytes contained within pending messages that will be held while waiting for a connection to a peer to be created. If the specified value is reached, any additional messages to the peer will be undeliverable, and held messages are delivered to the peer. The default value is 0, which disables this setting.
	MaxPendingBytes int64 `json:"maxPendingBytes,omitempty"`

	// Specifies the maximum number of pending messages held while waiting for a connection to a peer to be created. If the specified value is reached, any additional messages to the peer will be undeliverable, and held messages are delivered to the peer. The default value is 0, which disables this setting.
	MaxPendingMessages int64 `json:"maxPendingMessages,omitempty"`

	// Specifies the maximum period in seconds between a request and response. A provisional response restarts the timer. This setting might not affect all transactions. The default value is 10 seconds.
	TransactionTimeout int64 `json:"transactionTimeout,omitempty"`

	// When selected (enabled), controls whether connections that are established by the ingress TMM are preferred to connections that are established by another TMM when selecting an egress connection to a destination peer. The default value is enabled.
	UseLocalConnection string `json:"useLocalConnection,omitempty"`

	// Specifies the profile from which this profile inherits settings. The default is the system-supplied \"diameterrouter\" profile.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

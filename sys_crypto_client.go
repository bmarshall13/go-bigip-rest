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
type SysCryptoClient struct {

	// Specifies the interval in seconds between attempts to connect to the remote crypto server. The default value is 10 seconds.
	RetryInterval int64 `json:"retryInterval,omitempty"`

	// Resets the connection to the remote crypto server.
	ConnectionReset bool `json:"connectionReset,omitempty"`

	// Specifies the IP address of the remote crypto server.
	Addr string `json:"addr,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the timeout in milliseconds for crypto requests to complete. The default value is 5000 milliseconds.
	ReqTimeout int64 `json:"reqTimeout,omitempty"`

	// Specifies the number of seconds to wait before sending a heartbeat request. A value of 0 disables the sending of heartbeat requests. The default value is 30 seconds.
	Heartbeat int64 `json:"heartbeat,omitempty"`

	// Specifies the maximum number of times to retry connecting to the remote crypto server. If the maximum retries value is infinite, the crypto client retries connecting until a connection is made. The default value is infinite.
	MaxRetries string `json:"maxRetries,omitempty"`

	// Specifies the port used by the remote crypto server.
	Port int64 `json:"port,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}
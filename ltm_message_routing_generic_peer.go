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
type LtmMessageRoutingGenericPeer struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the ratio to be used for selection of a peer within a list of peers in a ltm genericmsg-route.
	Ratio int64 `json:"ratio,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies how the number of connections per host are to be limited. Note a host (specified in the referred pool) may also exist in other peer objects and those other peer objects may have different settings for connection-mode and number_connections. Thus these settings specify how messages routed through this peer are distributed between a set of connections, not the maximum number of connections to a specified host.
	ConnectionMode string `json:"connectionMode,omitempty"`

	Partition string `json:"partition,omitempty"`

	// The name of the ltm virtual or ltm transport-config to use for creating an outgoing connection.
	TransportConfig string `json:"transportConfig,omitempty"`

	// Specifies the distribution of connections between the BIG-IP and a remote host.
	NumberConnections int64 `json:"numberConnections,omitempty"`

	// Specifies the name of the pool that messages will be routed towards.
	Pool string `json:"pool,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

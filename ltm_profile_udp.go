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
type LtmProfileUdp struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Describe the outgoing UDP packet TTL mode.
	IpTtlMode string `json:"ipTtlMode,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 60 seconds.
	IdleTimeout string `json:"idleTimeout,omitempty"`

	// Enables or disables checksum processing. Note that if the datagram is IPv6, the system always performs checksum processing. The default value is disabled.
	NoChecksum string `json:"noChecksum,omitempty"`

	// Specifies the Type of Service level that the traffic management system assigns to UDP packets when sending them to clients.
	IpTosToClient string `json:"ipTosToClient,omitempty"`

	// Provides the ability to allow the passage of datagrams that contain header information, but no essential data. The default value is disabled.
	AllowNoPayload string `json:"allowNoPayload,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Make the system use the same max segment size on both ends.
	ProxyMss string `json:"proxyMss,omitempty"`

	// Provides the ability to load balance UDP datagram by datagram. The default value is disabled.
	DatagramLoadBalancing string `json:"datagramLoadBalancing,omitempty"`

	// Specifies the Quality of Service level that the system assigns to UDP packets when sending them to clients.
	LinkQosToClient string `json:"linkQosToClient,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

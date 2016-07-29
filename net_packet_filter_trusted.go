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
type NetPacketFilterTrusted struct {

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies a list of MAC addresses. Packets received from these addresses are allowed to bypass the packet filter.
	MacAddresses string `json:"macAddresses,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies a list of source IP addresses. Packets received from these addresses are allowed to bypass the packet filter.
	IpAddresses string `json:"ipAddresses,omitempty"`

	// Specifies a list of ingress VLANs. Packets received on these VLANs are allowed to bypass the packet filter.
	Vlans string `json:"vlans,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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
type NetRoute struct {

	// Specifies a gateway address for the route.
	Gw string `json:"gw,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Displays the administrative partition within which the route resides.
	Partition string `json:"partition,omitempty"`

	// Drops traffic addressed to the subnet.
	Blackhole bool `json:"blackhole,omitempty"`

	// Sets a specific maximum transmission unit (MTU).
	Mtu int64 `json:"mtu,omitempty"`

	// Specifies a VLAN for the route. This can be a VLAN or VLAN group.
	TmInterface string `json:"tmInterface,omitempty"`

	// Specifies a gateway pool, which allows multiple, load-balanced gateways to be used for the route.
	Pool string `json:"pool,omitempty"`

	// The destination subnet and netmask for the route.
	Network string `json:"network,omitempty"`
}

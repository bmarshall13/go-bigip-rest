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
type NetCmetrics struct {

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Displays the cached slow-start threshold.
	Ssthresh int64 `json:"ssthresh,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Displays the maximum transmit unit size on the route.
	Mtu int64 `json:"mtu,omitempty"`

	// Displays the MAC address for the route.
	Hwaddress string `json:"hwaddress,omitempty"`

	// Displays the round-trip time on the route.
	Rtt int64 `json:"rtt,omitempty"`

	// Displays the size of the channel.
	Bandwidth int64 `json:"bandwidth,omitempty"`

	// Specifies the destination IP address of the entry to display or delete.
	DestAddr string `json:"destAddr,omitempty"`

	// Displays the variance in the round-trip time.
	Rttvar int64 `json:"rttvar,omitempty"`

	// Displays the identifying number of the tmm (Traffic Management Microkernel).
	Tmm int64 `json:"tmm,omitempty"`
}

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
type NetVlanGroup struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// When enabled, specifies that the VLAN group forwards packets, even when the system is the standby unit in a redundant system. Note that this setting is designed for deployments in which the VLAN group exists on only one of the units. If that does not match your configuration, using this setting may cause adverse effects. The default value is disabled.
	BridgeInStandby string `json:"bridgeInStandby,omitempty"`

	// Displays the administrative partition within which the vlan-group resides.
	Partition string `json:"partition,omitempty"`

	// When enabled, specifies that the VLAN group forwards all frames, including non-IP traffic. The default value is disable.
	BridgeTraffic string `json:"bridgeTraffic,omitempty"`

	// Specifies whether the system will send keepalive frames (TCP keepalives and empty UDP packets depending on the connection type) when a node is moved from one vlan-group member to another vlan-group member for all existing connections that the system has to that node.
	MigrationKeepalive string `json:"migrationKeepalive,omitempty"`

	// When enabled, allows bridging of non-Internet Protocol (IP) Address Resolution Protocol (ARP) multicast frames across a VLAN group. An example of when you might want to use this option is when you are implementing the Spanning Tree Protocol (STP).
	BridgeMulticast string `json:"bridgeMulticast,omitempty"`

	// Specifies the level of exposure of remote MAC addresses within VLAN groups. The default value is translucent.
	Mode string `json:"mode,omitempty"`

	// Displays the index assigned to this VLAN group. It is a unique identifier assigned for all objects displayed in the SNMP IF-MIB.
	IfIndex int64 `json:"ifIndex,omitempty"`

	// Specifies whether auto lasthop is enabled or not. The default is to use next levels default.
	AutoLasthop string `json:"autoLasthop,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

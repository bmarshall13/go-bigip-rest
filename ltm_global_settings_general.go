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
type LtmGlobalSettingsGeneral struct {

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies how fast gratuitous ARPs can be sent. If it is 0, then gratuitous ARPs are sent without pause. Otherwise, it specifies how many gratuitous ARPs can be sent every second. The default value is 0.
	GratuitousArpRate int64 `json:"gratuitousArpRate,omitempty"`

	// Specifies, when enabled, that the unit is in maintenance mode. In maintenance mode, the system stops accepting new connections and slowly finishes off existing connections.
	MaintenanceMode string `json:"maintenanceMode,omitempty"`

	// Specifies the Media Access Control address (MAC address) that the system assigns to a VLAN. The value of unique indicates that a VLAN uses a mac address from a global mac address pool assigned to each hardware platform. The global value indicates that all of the VLANs on the system use the same MAC address.  The vmw-compat value indicates that the MAC address of a vlan is allocated in a manner compatible with VMware(tm) vSwitch and restricts VLANs to a single interface, with no trunks allowed.  Changing the value of this feature requires a manual restart of all daemons
	ShareSingleMac string `json:"shareSingleMac,omitempty"`

	// Enables or disables SNAT packet forwarding. The default is disable.
	SnatPacketForward string `json:"snatPacketForward,omitempty"`

	// Specifies, in seconds, the amount of time that records remain in the Layer 2 forwarding table, when the MAC address of the record is no longer detected on the network. The default is 300 seconds.
	L2CacheTimeout int64 `json:"l2CacheTimeout,omitempty"`
}

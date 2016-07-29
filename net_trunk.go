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
type NetTrunk struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Specifies the media access control (MAC) address, which is associated with the trunk, in case-insensitive hexadecimal colon notation, for example, 00:0b:09:88:00:9a.
	MacAddress string `json:"macAddress,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Resets STP, which forces a migration check.
	StpReset bool `json:"stpReset,omitempty"`

	// Specifies the number of configured members.
	CfgMbrCount int64 `json:"cfgMbrCount,omitempty"`

	// Specifies the number of working members associated with this trunk.
	WorkingMbrCount int64 `json:"workingMbrCount,omitempty"`

	// Specifies the operational bandwidth in Mobs.
	Bandwidth int64 `json:"bandwidth,omitempty"`

	// Specifies the rate at which the system sends the LACP control packets.
	LacpTimeout string `json:"lacpTimeout,omitempty"`

	// Specifies the operation mode for link aggregation control protocol (LACP), if LACP is enabled for the trunk.
	LacpMode string `json:"lacpMode,omitempty"`

	// Specifies, when enabled, that the system supports the link aggregation control protocol (LACP), which monitors the trunk by exchanging control packets over the member links to determine the health of the links. If LACP detects a failure in a member link, it removes the link from the link aggregation. LACP is disabled by default, for backward compatibility.
	Lacp string `json:"lacp,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the protocol identifier associated with the tagged mode of the trunk.
	QinqEthertype string `json:"qinqEthertype,omitempty"`

	// Displays the media settings for the trunk.
	Media string `json:"media,omitempty"`

	// Displays the ID of the trunk.
	Id int64 `json:"id,omitempty"`

	// Specifies the basis for the hash that the system uses as the frame distribution algorithm. The system uses the resulting hash to determine which interface to use for forwarding traffic.
	DistributionHash string `json:"distributionHash,omitempty"`

	// Enables or disables STP. If you disable STP, the system does not transmit or receive STP, RSTP, or MSTP packets on the trunk, and STP has no control over forwarding or learning on the trunk. The default value is enabled.
	Stp string `json:"stp,omitempty"`

	// Sets the LACP policy that the trunk uses to determine which member link (interface) can handle new traffic. Note that link aggregation is allowed only when all the interfaces are operating at the same media speed and connected to the same partner aggregation system. When there is a mismatch among configured members due to configuration errors or topology changes (auto-negotiation), link selection policy determines which links become working members and form the aggregation.
	LinkSelectPolicy string `json:"linkSelectPolicy,omitempty"`

	// Displays whether this trunk is managed by a VCMP hypervisor or not.
	Type_ string `json:"type,omitempty"`
}
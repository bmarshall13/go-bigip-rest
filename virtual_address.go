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

type VirtualAddress struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Enables or disables ARP for the specified virtual address. The default value is enabled.
	Arp string `json:"arp,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies whether a virtual server IP address is enabled. The default value is yes.
	Enabled string `json:"enabled,omitempty"`

	// Specifies the traffic group for the virtual address.  The default traffic group is inherited from the containing folder.
	TrafficGroup string `json:"trafficGroup,omitempty"`

	// Sets a concurrent connection limit in seconds for one or more virtual servers. The default value is 0 seconds.
	ConnectionLimit int64 `json:"connectionLimit,omitempty"`

	Address string `json:"address,omitempty"`

	Floating string `json:"floating,omitempty"`

	// Enables or disables ICMP echo replies for the specified virtual address. The default value is enabled.
	IcmpEcho string `json:"icmpEcho,omitempty"`

	Unit int64 `json:"unit,omitempty"`

	// Enables or disables route advertisement for the specified virtual address. The default value is disabled.
	RouteAdvertisement string `json:"routeAdvertisement,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Indicates if the virtual address will be deleted automatically on deletion of the last associated virtual server or not. The default value is true.
	AutoDelete string `json:"autoDelete,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// Sets the netmask or one or more network virtual servers only. This setting is required for network virtual servers.
	Mask string `json:"mask,omitempty"`

	// Specifies the server that uses the specified virtual address.
	ServerScope string `json:"serverScope,omitempty"`

	InheritedTrafficGroup string `json:"inheritedTrafficGroup,omitempty"`
}

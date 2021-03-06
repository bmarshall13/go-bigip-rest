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
type NetWccpServices struct {
	AppService string `json:"appService,omitempty"`

	// Specifies the method the router uses to redirect traffic: GRE or L2. The default value is gre.
	RedirectionMethod string `json:"redirectionMethod,omitempty"`

	// Specifies the protocol of the traffic to be redirected: TCP or UDP. The default value is tcp.
	Protocol string `json:"protocol,omitempty"`

	// Specifies to the router which traffic attributes to use to determine which BIG-IP system it should forward traffic to for load balancing: destination IP address (dest-ip), destination port (dest-port), source IP address (src-ip), and/or source port (src-port).
	HashFields string `json:"hashFields,omitempty"`

	// Specifies the relative importance of this traffic in a load balancing environment. The range is from 1 to 100. The default value is 50.
	Weight int64 `json:"weight,omitempty"`

	// Specifies the IP addresses of the WCCP-enabled routers that redirect traffic.
	Routers string `json:"routers,omitempty"`

	// Specifies the password for MD5 authentication or none.
	Password string `json:"password,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the method used to return pass-through traffic to the router: GRE or L2. The default value is gre.
	ReturnMethod string `json:"returnMethod,omitempty"`

	// Specifies the Router Identifier IP address of the router that redirects traffic.
	TunnelRemoteAddresses string `json:"tunnelRemoteAddresses,omitempty"`

	// Specifies one or more ports (up to 8) on which traffic is redirected.
	Ports int64 `json:"ports,omitempty"`

	// Specifies the precedence of the service group relative to the other service groups. The range is from 1 to 255. The default value is 100.
	Priority int64 `json:"priority,omitempty"`

	// Specifies an IP address on the BIG-IP system to which the WCCP-enabled routers should redirect traffic. Specify a self IP address of an external VLAN on the BIG-IP system.
	TunnelLocalAddress string `json:"tunnelLocalAddress,omitempty"`

	// Specifies whether the WCCP interception of traffic is based on the destination port (dest) or source port (source), or is not specified (none). The default value is none.
	PortType string `json:"portType,omitempty"`

	// Specifies whether load balancing is achieved by a hash algorithm or a mask. If you specify hash, specify one or more attributes using the option hash-fields.
	TrafficAssign string `json:"trafficAssign,omitempty"`
}

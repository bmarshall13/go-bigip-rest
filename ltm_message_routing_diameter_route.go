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
type LtmMessageRoutingDiameterRoute struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies how the system selects a peer to handle the Diameter traffic for this static route. The default is Sequential. Options are  1) Ratio: Peer selection is based on the ratio that is set for each peer that is in the Selected list.  2) Sequential: Peer selection is based on the order of the peers in the Selected list.
	PeerSelectionMode string `json:"peerSelectionMode,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Specifies the destination realm matching the Destination-Realm AVP value in the message. The blank default value routes all destination-realms.
	DestinationRealm string `json:"destinationRealm,omitempty"`

	// Specifies the origin realm matching the Origin-Realm AVP value in the message. The blank default value routes all origin-realms.
	OriginRealm string `json:"originRealm,omitempty"`

	// Specifies the virtual server from which the system receives client requests for this static route. If you do not select a virtual server, the system uses this static route to route Diameter messages originating from any client.
	VirtualServer string `json:"virtualServer,omitempty"`

	// Specifies the identifier matching the application-id in the Diameter message. A value of 0 matches every application-id.
	ApplicationId int64 `json:"applicationId,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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
type LtmMessageRoutingGenericRoute struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the destination address of the route. If this attribute is not present, it will be treated as a wildcard matching all message destination-addresses.
	DestinationAddress string `json:"destinationAddress,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the method to use when selecting a peer from the provided list of peers. Possible values are: sequential and ratio.
	PeerSelectionMode string `json:"peerSelectionMode,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the source address of the route. If this attribute is not present, it will be treated as a wildcard matching all message source-addresses.
	SourceAddress string `json:"sourceAddress,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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

type MessageRoutingSipRoute struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the pattern x@y where x=Username and y=host(domain or IP) to be matched against the From URI (sip:x@y:Port) field of a SIP message.
	FromUri string `json:"fromUri,omitempty"`

	// Specifies the mode of selecting a peer from a list of peers.
	PeerSelectionMode string `json:"peerSelectionMode,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Specifies the pattern x@y where x=Username and y=host(domain or IP) to be matched against the To URI (sip:x@y:Port) field of a SIP message.
	ToUri string `json:"toUri,omitempty"`

	// Specifies the virtual server on which connections will be routed to this route.
	VirtualServer string `json:"virtualServer,omitempty"`

	// Specifies the pattern x@y where x=Username and y=host(domain or IP) to be matched against the Request URI (sip:x@y:Port) field of a SIP message.
	RequestUri string `json:"requestUri,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

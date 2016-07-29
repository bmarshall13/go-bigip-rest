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
type NetTunnelsGre struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies whether the system includes a checksum on transmitted packets. The default value is disabled.
	TxCsum string `json:"txCsum,omitempty"`

	// Displays the admin-partition within which this component resides.
	Partition string `json:"partition,omitempty"`

	// Specifies whether the system verifies the checksum on received packets. The default value is disabled.
	RxCsum string `json:"rxCsum,omitempty"`

	// Specifies the flooding type to use to transmit broadcast and unknown destination frames. The default is none.
	FloodingType string `json:"floodingType,omitempty"`

	// The profile from which to inherit settings. The default value is gre.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the flavor of GRE header to use for encapsulation. The default value is standard.
	Encapsulation string `json:"encapsulation,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}
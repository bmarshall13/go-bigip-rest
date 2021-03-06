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
type LtmAuthKerberosDelegation struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies whether debug logging is enabled. The default value is disabled.
	DebugLogging string `json:"debugLogging,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the principal that the client sees. This is usually a value such as HTTP/ fqdn . The client principal may be in a different domain from the server principal. This setting is required. There is no default value.
	ClientPrincipal string `json:"clientPrincipal,omitempty"`

	// Specifies whether associated virtual should transition client certificate authentication into Kerberos credentials.
	ProtocolTransition string `json:"protocolTransition,omitempty"`

	// Specifies the principal of the back-end web server. This is usually a value such as HTTP/ fqdn of server . The server principal may be in a different domain from the client principal. This setting is required. There is no default value.
	ServerPrincipal string `json:"serverPrincipal,omitempty"`

	// Displays the administrative partition within which the configuration resides.
	Partition string `json:"partition,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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
type LtmAuthTacacs struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the protocol associated with the value specified in the service option, which is a subset of the associated service being used for client authorization or system accounting.
	Protocol string `json:"protocol,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the name of the service that the user is requesting to be authenticated to use. Identifying the service enables the TACACS+ server to behave differently for different types of authentication requests. This option is required.
	Service string `json:"service,omitempty"`

	// Enables or disables encryption of TACACS+ packets. Recommended for normal use. The default value is enabled.
	Encryption string `json:"encryption,omitempty"`

	// Displays the partition within which the server resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the process the system employs when sending authentication requests. The default is use-first-server. use-first-server specifies that the system sends authentication requests to only the first server in the list. use-all-servers specifies that the system sends an authentication request to each server until authentication succeeds, or until the system has sent a request to all servers in the list.
	Authentication string `json:"authentication,omitempty"`

	// Specifies a host name or IP address for the TACACS+ server. This option is required. Possible values are a user-specified string, and none. You must specify a server when you create a TACACS+ configuration object.
	Servers string `json:"servers,omitempty"`

	// Sets the secret key used to encrypt and decrypt packets sent or received from the server. This option is required.
	Secret string `json:"secret,omitempty"`

	// Enables syslog-ng debugging information at LOG DEBUG level. Not recommended for normal use. The default value is disabled.
	Debug string `json:"debug,omitempty"`

	// If multiple TACACS+ servers are defined and pluggable authentication module (PAM) session accounting is enabled, sends accounting start and stop packets to the first available server or to all servers. The default value is send-to-first-server.
	Accounting string `json:"accounting,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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
type LtmAuthRadius struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Specifies the number of authentication retries that the BIG-IP local traffic management system allows before authentication fails. The default value is 3.
	Retries int64 `json:"retries,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Enables or disables validation of the accounting response vector. This option should be necessary only on older servers. The default value is disabled.
	AccountingBug string `json:"accountingBug,omitempty"`

	// Displays the partition within which the component resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the type of service requested from the RADIUS server.  The default value is authenticate-only
	ServiceType string `json:"serviceType,omitempty"`

	// Sends a NAS-Identifier RADIUS attribute with string bar. If you do not specify a value for the client-id option, the system uses the pluggable authentication module (PAM) service type. You can disable this feature by specifying a blank client ID.
	ClientId string `json:"clientId,omitempty"`

	// Enables or disables syslog-ng debugging information at LOG DEBUG level. Not recommended for normal use. The default value is disabled.
	Debug string `json:"debug,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}
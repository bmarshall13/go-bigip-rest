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
type LtmAuthSslCrldp struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Specifies whether the system extracts the CRL distribution point from the client certificate. The default value is disabled.
	UseIssuer string `json:"useIssuer,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Displays the partition within which the server resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the number of seconds that CRLs will be cached. The default value is 86400 seconds.
	CacheTimeout int64 `json:"cacheTimeout,omitempty"`

	// Specifies the number of seconds before the connection times out. The default value is 15 seconds.
	ConnectionTimeout int64 `json:"connectionTimeout,omitempty"`

	// Specifies an update interval for CRL distribution points. The update interval for distribution points ensures that CRL status is checked at regular intervals, regardless of the CRL timeout value. This helps to prevent CRL information from becoming outdated before the BIG-IP system checks the status of a certificate. The default value is 0 (zero), which indicates an internal default value is active.
	UpdateInterval int64 `json:"updateInterval,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

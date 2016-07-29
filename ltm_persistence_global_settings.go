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
type LtmPersistenceGlobalSettings struct {

	// Specifies that the persistence session is limited by either the number of seconds before the persistence entry times out, or by a maximum number of requests to the destination address.
	DestAddrLimitMode string `json:"destAddrLimitMode,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies a group of servers that are configured to process all of the requests from a single source address during a persistence session.
	ProxyGroup string `json:"proxyGroup,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the maximum number of entries that can be in the persistence table at any one time when using the destination address affinity mode and when the option dest addr limit is set to max-count. The default value is 2048 entries.
	DestAddrMax int64 `json:"destAddrMax,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`
}
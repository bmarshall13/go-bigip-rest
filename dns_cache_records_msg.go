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

type DnsCacheRecordsMsg struct {

	// Chassis slot to query.
	Slot int64 `json:"slot,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Name of DNS Cache configuration object to inspect.
	Cache string `json:"cache,omitempty"`

	// DNS return code for message cache filtering.
	Rcode string `json:"rcode,omitempty"`

	// DNS query name for message cache filtering.
	Qname string `json:"qname,omitempty"`

	// TMM to query (typically unset and used for debugging).
	Tmm int64 `json:"tmm,omitempty"`
}

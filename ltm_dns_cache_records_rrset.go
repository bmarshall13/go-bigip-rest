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
type LtmDnsCacheRecordsRrset struct {

	// Chassis slot to query.
	Slot int64 `json:"slot,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Name of DNS Cache configuration object to inspect.
	Cache string `json:"cache,omitempty"`

	// Record class to filter on.
	TmClass string `json:"tmClass,omitempty"`

	// Range of TTLs to filter on, low:high.
	TtlRange string `json:"ttlRange,omitempty"`

	// DNS owner name for filtering RRs.
	Owner string `json:"owner,omitempty"`

	// Record type to filter on.
	Type_ string `json:"type,omitempty"`

	// TMM to query (typically unset and used for debugging).
	Tmm int64 `json:"tmm,omitempty"`
}

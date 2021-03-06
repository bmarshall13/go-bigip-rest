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
type SysHaGroupPools struct {
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Maximum value this attribute contributes to the ha score.
	Weight int64 `json:"weight,omitempty"`

	// Below this value the attribute contributes nothing to the ha score.
	Threshold int64 `json:"threshold,omitempty"`

	// Choose the pool attribute that contributes to the ha score.
	Attribute string `json:"attribute,omitempty"`

	// Displays the calculated percent of pool members that are up.
	PercentUp int64 `json:"percentUp,omitempty"`
}

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
type LtmPersistenceUniversal struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies, when enabled, that the system can use any pool that contains this persistence record. The default value is disabled.
	MatchAcrossPools string `json:"matchAcrossPools,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies, when enabled, that all persistent connections from a client IP address, which go to the same virtual IP address, also go to the same node. The default value is disabled.
	MatchAcrossServices string `json:"matchAcrossServices,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies, when enabled, that the pool member connection limits are not enforced for persisted clients. Per-virtual connection limits remain hard limits and are not disabled. The default value is disabled.
	OverrideConnectionLimit string `json:"overrideConnectionLimit,omitempty"`

	// Displays the administrative partition within which the profile resides.
	Partition string `json:"partition,omitempty"`

	// Specifies a rule name if you are using a rule for universal persistence.
	Rule string `json:"rule,omitempty"`

	// Specifies, when enabled, that all persistent connections from the same client IP address go to the same node. The default value is disabled.
	MatchAcrossVirtuals string `json:"matchAcrossVirtuals,omitempty"`

	Mode string `json:"mode,omitempty"`

	// Specifies the duration of the persistence entries. The default value is 180 seconds.
	Timeout string `json:"timeout,omitempty"`

	// Specifies whether the system mirrors persistence records to the high-availability peer. The default value is disabled.
	Mirror string `json:"mirror,omitempty"`

	// Specifies the existing profile from which the system imports settings for the new profile.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

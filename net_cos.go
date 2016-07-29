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
type NetCos struct {

	// Configures a traffic priority object.
	TrafficPriority string `json:"trafficPriority,omitempty"`

	// Configures the global configuration for class of service.
	GlobalSettings string `json:"globalSettings,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Configures IP DSCP field to traffic priority mapping.
	MapDscp string `json:"mapDscp,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Configures vlan 8021.p tag to traffic priority mapping.
	Map8021p string `json:"map_8021p,omitempty"`
}

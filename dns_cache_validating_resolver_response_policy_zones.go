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

type DnsCacheValidatingResolverResponsePolicyZones struct {

	AppService string `json:"appService,omitempty"`

	// If action is configured to walled-garden, the name of the local zone containing the records to use in the DNS response.
	WalledGarden string `json:"walledGarden,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// The action to take upon finding a match to the response policy data.
	Action string `json:"action,omitempty"`

	// If enabled, the action is not enforced, but the logs and statistics are updated as if it were.
	LogsAndStatsOnly string `json:"logsAndStatsOnly,omitempty"`
}

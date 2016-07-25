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

type PoolMembers struct {

	AppService string `json:"appService,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Enables or disables the pool member for new sessions. The default value is user-enabled.
	Session string `json:"session,omitempty"`

	// Specifies a range of numbers that you want the system to use in conjunction with the ratio load balancing method. The default value is 1.
	DynamicRatio int64 `json:"dynamicRatio,omitempty"`

	// Specifies the maximum number of concurrent connections allowed for a pool member. The default value is 0 (zero).
	ConnectionLimit int64 `json:"connectionLimit,omitempty"`

	// IP address of a pool member if a node by the given name does not already exist.
	Address string `json:"address,omitempty"`

	// Specifies the ratio weight that you want to assign to the pool member. The default value is 1.
	Ratio int64 `json:"ratio,omitempty"`

	// Displays the health monitors that are configured to monitor the pool member, and the status of each monitor. The default value is default.
	Monitor string `json:"monitor,omitempty"`

	// Specifies whether the pool member inherits the encapsulation profile from the parent pool. The default value is enabled. If you disable inheritance, no encapsulation takes place, unless you specify another encapsulation profile for the pool member using the profiles attribute.
	InheritProfile string `json:"inheritProfile,omitempty"`

	// Specifies whether the pool member's monitor(s) actions will be logged. Logs are stored in /var/log/monitors/
	Logging string `json:"logging,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	Ephemeral string `json:"ephemeral,omitempty"`

	// user-down forces the pool member offline, overriding monitors. user-up reverts the user-down. When user-up, this displays the monitor state.
	State string `json:"state,omitempty"`

	// Specifies the maximum number of connections per second allowed for a pool member. The default value is 'disabled'.
	RateLimit string `json:"rateLimit,omitempty"`

	// Specifies the priority group within the pool for this pool member. The priority group number specifies that traffic is directed to that member before being directed to a member of a lower priority. The default value is 0.
	PriorityGroup int64 `json:"priorityGroup,omitempty"`
}

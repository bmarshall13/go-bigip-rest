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
type LtmMonitorVirtualLocation struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	Destination string `json:"destination,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval.
	UpInterval int64 `json:"upInterval,omitempty"`

	// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown. The default value is 5 seconds.
	Interval int64 `json:"interval,omitempty"`

	// Displays the administrative partition within which the monitor resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the amount of time, in seconds, after the first successful response before a node is marked up. A value of 0 (zero) causes a node to be marked up immediately after a valid response is received from the node. The default value is 0 (zero).
	TimeUntilUp int64 `json:"timeUntilUp,omitempty"`

	// Specifies the pool for the target pool member. This is a required argument.
	Pool string `json:"pool,omitempty"`

	// Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 16 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire.
	Timeout int64 `json:"timeout,omitempty"`

	// Specifies whether the monitor sends error messages and additional information to a log file created and labeled specifically for this monitor. The default setting is no. You can use the log information to help diagnose and troubleshoot unsuccessful health checks. The options are no (specifies that the system does not redirect error messages and additional information related to this monitor.) and yes (specifies that the system redirects error messages and additional information to the /var/log/monitors/ monitor_name - node_name - port .log file.)
	Debug string `json:"debug,omitempty"`

	// Specifies the existing monitor from which the system imports settings for the new monitor.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

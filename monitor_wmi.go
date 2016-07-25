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

type MonitorWmi struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the user name if the monitored target requires authentication.
	Username string `json:"username,omitempty"`

	// Specifies the frequency at which the system issues the monitor check. The default value is 5 seconds.
	Interval int64 `json:"interval,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Displays the agent for the monitor. The default agent is Mozilla/4.0 (compatible: MSIE 5.0; Windows NT). You cannot modify the agent.
	Agent string `json:"agent,omitempty"`

	// Specifies the performance metrics that the commands collect from the target. The default value is LoadPercentage, DiskUsage, PhysicalMemoryUsage:1.5, VirtualMemoryUsage:2.0
	Metrics string `json:"metrics,omitempty"`

	// Specifies the existing monitor from which the system imports settings for the new monitor.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the mechanism that the monitor uses for posting. The default value is RespFormat=HTML. You cannot change the post format for WMI monitors.
	Post string `json:"post,omitempty"`

	// Specifies the password if the monitored target requires authentication.
	Password string `json:"password,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the URL that the monitor uses. The default value is /scripts/f5Isapi.dll.
	Url string `json:"url,omitempty"`

	// Specifies the command that the system uses to obtain the metrics from the resource. See the documentation for this resource for information on available commands.
	TmCommand string `json:"tmCommand,omitempty"`

	// Specifies the IP address of the resource that is the destination of this monitor. Possible values are:  * (Specifies to perform a health check on the IP address of the node.), and  IP  (Specifies to perform a health check on the IP address that you specify, route the check through the IP address of the associated node, and mark the IP address of the associated node up or down accordingly.).
	Destination string `json:"destination,omitempty"`

	// Displays the administrative partition within which the monitor resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0.
	TimeUntilUp int64 `json:"timeUntilUp,omitempty"`

	// Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 16 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire.
	Timeout int64 `json:"timeout,omitempty"`

	// Displays the GET method. You cannot modify the method.
	Method string `json:"method,omitempty"`
}

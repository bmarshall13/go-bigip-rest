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

type MonitorMysql struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the user name if the monitored target requires authentication.
	Username string `json:"username,omitempty"`

	// Displays the administrative partition within which the monitor resides.
	Partition string `json:"partition,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval.
	UpInterval int64 `json:"upInterval,omitempty"`

	// Specifies whether the system automatically changes the status of a resource to enabled at the next successful monitor check. The default value is disabled. If you set this option to enabled, you must manually re-enable the resource before the system can use it for load balancing connections.
	ManualResume string `json:"manualResume,omitempty"`

	// Specifies the row in the database where the system expects the specified Receive String to be located. This is an optional setting, and is applicable only if you configure the send and recv options.
	RecvRow string `json:"recvRow,omitempty"`

	// Specifies the existing monitor from which the system imports settings for the new monitor.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the IP address and service port of the resource that is the destination of this monitor. Possible values are:  *:* (Specifies to perform a health check on the IP address and port supplied by a pool member), *:port (Specifies to perform a health check on the server with the IP address supplied by the pool member and the port you specify.), and  IP : port  (Specifies to mark a pool member up or down based on the response of the server at the IP address and port you specify.).
	Destination string `json:"destination,omitempty"`

	// Specifies the password if the monitored target requires authentication.
	Password string `json:"password,omitempty"`

	// Specifies the text string that the monitor looks for in the returned resource. The most common receive expressions contain a text string that is included in a field in your database. If you do not specify both a Send String and a Receive String, the monitor performs a simple service check and connect only.
	Recv string `json:"recv,omitempty"`

	// Specifies the number of monitor probes after which the connection to the database will be terminated. Count value of zero indicates that the connection will never be terminated. The default value is zero.
	Count string `json:"count,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the name of the database that the monitor tries to access.
	Database string `json:"database,omitempty"`

	// Specifies the frequency at which the system issues the monitor check. The default value is 30 seconds.
	Interval int64 `json:"interval,omitempty"`

	// Specifies the column in the database where the system expects the specified Receive String to be located. This is an optional setting, and is applicable only if you configure the send and recv options.
	RecvColumn string `json:"recvColumn,omitempty"`

	// Specifies the SQL query that the monitor sends to the target object. For example, SELECT count(*) FROM mytable returns the number of rows in mytable. Since the string may have special characters, the system may require that the string be enclosed with single quotation marks. If this value is null, then a valid connection suffices to determine that the service is up. In this case, the system does not need the recv, recv-row, and recv-column options and ignores the options even if not null.
	Send string `json:"send,omitempty"`

	// Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0.
	TimeUntilUp int64 `json:"timeUntilUp,omitempty"`

	// Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 91 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire.
	Timeout int64 `json:"timeout,omitempty"`

	// Specifies whether the monitor sends error messages and additional information to a log file created and labeled specifically for this monitor. The default setting is no. You can use the log information to help diagnose and troubleshoot unsuccessful health checks. The options are 'no' (specifies that the system does not redirect error messages and additional information related to this monitor to the log file) and 'yes' (specifies that the system redirects error messages and additional information to the /var/log/monitors/ monitor_name - node_name - port .log file).
	Debug string `json:"debug,omitempty"`
}

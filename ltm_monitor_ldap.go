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
type LtmMonitorLdap struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the user name if the monitored target requires authentication.
	Username string `json:"username,omitempty"`

	// Specifies whether the system will query the LDAP servers pointed to by any referrals in the query results.
	ChaseReferrals string `json:"chaseReferrals,omitempty"`

	// Displays the administrative partition within which the monitor resides.
	Partition string `json:"partition,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies how often in seconds that the system issues the monitor check when the node is up. The default value is the same as the (down) interval.
	UpInterval int64 `json:"upInterval,omitempty"`

	// Specifies whether the target must include attributes in its response to be considered up. The options are no (Specifies that the system performs only a one-level search (based on the Filter setting), and does not require that the target returns any attributes.) and yes (Specifies that the system performs a sub-tree search, and if the target returns no attributes, the target is considered down.)
	MandatoryAttributes string `json:"mandatoryAttributes,omitempty"`

	// Specifies whether the system automatically changes the status of a resource to enabled at the next successful monitor check. The default value is disabled. If you set this option to enabled, you must manually re-enable the resource before the system can use it for load balancing connections.
	ManualResume string `json:"manualResume,omitempty"`

	// Specifies the location in the LDAP tree from which the monitor starts the health check. A sample value is dc=bigip-test,dc=net.
	Base string `json:"base,omitempty"`

	// Specifies the existing monitor from which the system imports settings for the new monitor.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the IP address and service port of the resource that is the destination of this monitor. Possible values are:  *:* (Specifies to perform a health check on the IP address and port supplied by a pool member), *:port (Specifies to perform a health check on the server with the IP address supplied by the pool member and the port you specify.), and  IP : port  (Specifies to mark a pool member up or down based on the response of the server at the IP address and port you specify.).
	Destination string `json:"destination,omitempty"`

	// Specifies the password if the monitored target requires authentication.
	Password string `json:"password,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the frequency at which the system issues the monitor check. The default value is 10 seconds.
	Interval int64 `json:"interval,omitempty"`

	// Specifies an LDAP key for which the monitor searches. A sample value is: objectclass=*.
	Filter string `json:"filter,omitempty"`

	// Specifies the amount of time in seconds after the first successful response before a node will be marked up.  A value of 0 will cause a node to be marked up immediately after a valid  response is received from the node. The default setting is 0.
	TimeUntilUp int64 `json:"timeUntilUp,omitempty"`

	// Specifies the number of seconds the target has in which to respond to the monitor request. The default value is 31 seconds. If the target responds within the set time period, it is considered up. If the target does not respond within the set time period, it is considered down. Also, if the target responds with a RESET packet, the system immediately flags the target as down without waiting for the timeout interval to expire.
	Timeout int64 `json:"timeout,omitempty"`

	// Specifies whether the monitor sends error messages and additional information to a log file created and labeled specifically for this monitor. The default setting is no. You can use the log information to help diagnose and troubleshoot unsuccessful health checks. The options are 'no' (specifies that the system does not redirect error messages and additional information related to this monitor to the log file) and 'yes' (specifies that the system redirects error messages and additional information to the /var/log/monitors/ monitor_name - node_name - port .log file).
	Debug string `json:"debug,omitempty"`

	// Specifies the secure communications protocol that the monitor uses to communicate with the target. The options are none (Specifies that the system does not use a security protocol for communications with the target.), ssl (Specifies that the system uses the SSL protocol for communications with the target.), and tls (Specifies that the system uses the TLS protocol for communications with the target.)
	Security string `json:"security,omitempty"`
}

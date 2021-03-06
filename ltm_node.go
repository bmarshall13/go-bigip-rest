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
type LtmNode struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the name of the monitor or monitor rule that you want to associate with the node.
	Monitor string `json:"monitor,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies whether the node's monitor(s) actions will be logged. Logs are stored in /var/log/monitors/
	Logging string `json:"logging,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Displays the administrative partition within which this node resides.
	Partition string `json:"partition,omitempty"`

	Ephemeral string `json:"ephemeral,omitempty"`

	// Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
	RateLimit string `json:"rateLimit,omitempty"`

	// Marks the node up or down. The default value is user-up.
	State string `json:"state,omitempty"`

	// Enables or disables the node for new sessions. The default value is user-enabled.
	Session string `json:"session,omitempty"`

	// Sets the dynamic ratio number for the node. Used for dynamic ratio load balancing. The ratio weights are based on continuous monitoring of the servers and are therefore continually changing. Dynamic Ratio load balancing may currently be implemented on RealNetworks RealServer platforms, on Windows platforms equipped with Windows Management Instrumentation (WMI), or on a server equipped with either the UC Davis SNMP agent or Windows 2000 Server SNMP agent.
	DynamicRatio int64 `json:"dynamicRatio,omitempty"`

	// Specifies the maximum number of connections allowed for the node or node address.
	ConnectionLimit int64 `json:"connectionLimit,omitempty"`

	// IP address of the node. This is an optional field; if empty, the name needs to be of the form [ip address]
	Address string `json:"address,omitempty"`

	// Specifies the fixed ratio value used for a node during ratio load balancing.
	Ratio int64 `json:"ratio,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

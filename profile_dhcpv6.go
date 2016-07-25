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

type ProfileDhcpv6 struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Provides the default value in seconds of DHCPv6 lease time in case it was missing in the client-server exchange. The default is 86400.
	DefaultLeaseTime int64 `json:"defaultLeaseTime,omitempty"`

	// Specifies the DHCPv6 transaction timeout, in seconds. The transaction should complete within the timeout specified. If a transaction does not complete for any reason, it is removed. The default value is 45 seconds.
	TransactionTimeout int64 `json:"transactionTimeout,omitempty"`

	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 60 seconds.
	IdleTimeout string `json:"idleTimeout,omitempty"`

	// Specifies the operation mode of the DHCP virtual. If the virtual to run in relay mode, then it is acting as a standard DHCPv6 relay agent. The forwarding mode is similar to relay except that the virtual will not modify the standard fields, instead it will forward the message from client to server and vice-versa. The default is relay.
	Mode string `json:"mode,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

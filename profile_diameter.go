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

type ProfileDiameter struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// When enabled and the message is a capabilities exchange request or capabilities exchange answer, rewrite the host-ip-address attribute with the system's egress IP address.
	HostIpRewrite string `json:"hostIpRewrite,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the retransmit timeout in seconds. This setting specifies the number of seconds the BIG-IP waits to retransmit the request messages if it does not receive the corresponding answer messages. The default value is 10
	RetransmitTimeout int64 `json:"retransmitTimeout,omitempty"`

	// Specifies the watchdog timeout in seconds. This setting specifies the number of seconds that a connection is idle before the device watchdog request is sent. The default value is 0, which means BIG-IP will not send a device watchdog request to either client or server side.
	WatchdogTimeout int64 `json:"watchdogTimeout,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified. The default value is diameter.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the origin host to client of the BIG-IP. The origin-host-to-client is used to overwrite the server's actual origin host attribute when it responds to the client. It can be an ASCII string such as an FQDN. See RFC 3588 section 6.3. A value of \"none\" indicates that origin-host-to-client is disabled. The default value is \"none\".
	OriginHostToClient string `json:"originHostToClient,omitempty"`

	// Specifies the maximum number of retransmit attempts. This setting specifies the maximum number of attempts BIG-IP will take to retransmit the request messages if it does not receive the answer messages. If retransmit is unsuccessful, after the maximum number of attempts, BIG-IP will send error response. The default value is 1.
	MaxRetransmitAttempts int64 `json:"maxRetransmitAttempts,omitempty"`

	// Specifies the maximum number of device watchdog failures that the traffic management system can take before it tears down the connection. After the system receives this number of device watchdog failures, it closes the connection. The default value is 10.
	MaxWatchdogFailure int64 `json:"maxWatchdogFailure,omitempty"`

	// When enabled, and the system receives a capabilities exchange request from the client, the system will establish connections and perform handshaking with all the servers prior to sending the capabilities exchange answer to the client.
	ConnectionPrime string `json:"connectionPrime,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the name of the Diameter attribute that is used to indicate if the persist-avp is embedded in a grouped avp. It can be an ASCII string or numeric ID in the range 1 to 4294967295. Acceptable strings are from RFC 3588 section 4.5. A value of \"none\" indicates that the persist-avp value is not embedded in a grouped avp. The default value is \"none\".
	ParentAvp string `json:"parentAvp,omitempty"`

	// Specifies the realm to which the message will be routed. The destination-realm is used to perform message routing decisions. It can be an ASCII string such as an FQDN. See RFC 3588 section 6.6. A value of \"none\" indicates that destination-realm is disabled. The default value is \"none\".
	DestinationRealm string `json:"destinationRealm,omitempty"`

	// Specifies the origin realm to client of the BIG-IP. The origin-realm-to-client is used to overwrite the server's actual origin realm attribute when it responds to the client. It can be an ASCII string such as an FQDN. See RFC 3588 section 6.4. A value of \"none\" indicates that origin-realm-to-client is disabled. The default value is \"none\".
	OriginRealmToClient string `json:"originRealmToClient,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// When it is enabled and the watchdog failures exceed the max watchdog failure, the system resets the connection. The default value is enabled.
	ResetOnTimeout string `json:"resetOnTimeout,omitempty"`

	// Specifies the handshake timeout in seconds. This setting specifies the maximum number of seconds that a connection can be idle after the capabilities exchange request was sent to the server. The default value is 10. The system will reset the connection after it has timed out.
	HandshakeTimeout int64 `json:"handshakeTimeout,omitempty"`

	// Specifies the name of the Diameter attribute that will be used to persist on. It can be an ASCII string or numeric ID in the range 1 to 4294967295. Acceptable strings are defined in RFC 3588 section 4.5. A value of \"none\" indicates persistence is disabled. The default value is \"session-id\".
	PersistAvp string `json:"persistAvp,omitempty"`

	// Specifies the origin realm to server of the BIG-IP. The origin-realm-to-server is used to overwrite the client's actual origin realm attribute when it responds to the server. It can be an ASCII string such as an FQDN. See RFC 3588 section 6.4. A value of \"none\" indicates that origin-realm-to-server is disabled. The default value is \"none\".
	OriginRealmToServer string `json:"originRealmToServer,omitempty"`

	// Specifies the origin host to server of the BIG-IP. The origin-host-to-server is used to overwrite the client's actual origin host attribute when it responds to the server. It can be an ASCII string such as an FQDN. See RFC 3588 section 6.3. A value of \"none\" indicates that origin-host-to-server is disabled. The default value is \"none\".
	OriginHostToServer string `json:"originHostToServer,omitempty"`
}
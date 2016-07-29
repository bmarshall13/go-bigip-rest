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
type LtmProfileSip struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the maximum number of registrations, the maximum allowable REGISTER messages can be recorded that the BIG-IP system accepts. The default value is 100.
	MaxRegistrations int64 `json:"maxRegistrations,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Indicates whether SIP ALG feature is enabled. The default value is disabled
	AlgEnable string `json:"algEnable,omitempty"`

	// Indicates the timeout value for sip registration. The default value is 3600 seconds
	RegistrationTimeout int64 `json:"registrationTimeout,omitempty"`

	// Enables or disables the termination of a connection when a BYE transaction finishes. Use this parameter with UDP connections only, not with TCP connections. The default value is enabled.
	TerminateOnBye string `json:"terminateOnBye,omitempty"`

	// Indicates the timeout value for sip session. The default value is 300 seconds
	SipSessionTimeout int64 `json:"sipSessionTimeout,omitempty"`

	// Enables or disables the insertion of a Record-Route header, which indicates the next hop for the following SIP request messages. The default value is disabled.
	InsertRecordRouteHeader string `json:"insertRecordRouteHeader,omitempty"`

	// A string that groups SIP profiles together if dialog-aware is enabled.
	Community string `json:"community,omitempty"`

	// Specifies the maximum SIP message size that the BIG-IP system accepts. The default value is 65535 bytes.
	MaxSize int64 `json:"maxSize,omitempty"`

	// Enables or disables the insertion of a Secure Via header, which indicates where the message originated. When you are using SSL/TLS (over TCP) to create a secure channel with the server node, use this setting to configure the BIG-IP system to insert a Secure Via header into SIP requests. The default value is disabled.
	SecureViaHeader string `json:"secureViaHeader,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// The header that is inserted if insert-via-header is enabled.
	UserViaHeader string `json:"userViaHeader,omitempty"`

	// Specifies the maximum number of SDP media sessions that the BIG-IP system accepts. The default value is 6.
	MaxMediaSessions int64 `json:"maxMediaSessions,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the maximum number of calls or sessions can be made by a user for a single registration that the BIG-IP system accepts. The default value is 50.
	MaxSessionsPerRegistration int64 `json:"maxSessionsPerRegistration,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Configures the log publisher that handles events logging for this profile.
	LogPublisher string `json:"logPublisher,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// Indicates the style in which the RTP will proxy the data. When a dialog is established, the necessary SDP data needs to know where RTP flows are directed. It is provided with three options 1.) Use of a bidirectional related flow (symmetric, the default option), 2.) Use of ephemeral listeners to support fixed client IP, listener is restricted to connections coming from a particular source (restricted-by-ip-address option) and 3.) Use of ephemeral listeners to support wildcard, connections are allowed to come from anyway (any-location option) .
	RtpProxyStyle string `json:"rtpProxyStyle,omitempty"`

	// Indicates the timeout value for dialog establishment. The default value is 10 seconds.
	DialogEstablishmentTimeout int64 `json:"dialogEstablishmentTimeout,omitempty"`

	// Indicates whether snooping of SIP dialogs is enabled.
	DialogAware string `json:"dialogAware,omitempty"`

	// Enables or disables the insertion of a Via header, which indicates where the message originated. The response message uses this routing information. The default value is disabled.
	InsertViaHeader string `json:"insertViaHeader,omitempty"`

	// Enables the use of enhanced SIP security checking.
	Security string `json:"security,omitempty"`

	// Indicates whether SIP firewall capability is enabled. Default value is no
	EnableSipFirewall string `json:"enableSipFirewall,omitempty"`

	// Configures the ALG log profile that controls logging.
	LogProfile string `json:"logProfile,omitempty"`
}
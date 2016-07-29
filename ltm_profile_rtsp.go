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
type LtmProfileRtsp struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// When enabled, the client can select the destination to which to stream data. The default value is disabled.
	MulticastRedirect string `json:"multicastRedirect,omitempty"`

	// Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
	IdleTimeout string `json:"idleTimeout,omitempty"`

	// The Real Time Control Protocol (RTCP) allows monitoring of the real-time data delivery. Specifies the number of the port to use for the RTCP service.
	RtcpPort int64 `json:"rtcpPort,omitempty"`

	// When enabled the system uses the source attribute in the transport header to establish the target address of the RTP stream, and before the response is forwarded to the client, updates the value of the source attribute to be the virtual address of the BIG-IP system. When disabled the system does not change the source attribute. The default value is enabled.
	CheckSource string `json:"checkSource,omitempty"`

	// Specifies whether the RTSP filter is associated with an RTSP proxy configuration. The default value is none.
	Proxy string `json:"proxy,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// The Real Time Protocol (RTP) provides data transport functions suitable for applications transmitting real-time data. Specifies the number of the port to use for the RTP service.
	RtpPort int64 `json:"rtpPort,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// When enabled, the client can select the destination to which to stream data. The default value is disabled.
	UnicastRedirect string `json:"unicastRedirect,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Configures the log publisher that handles events logging for this profile.
	LogPublisher string `json:"logPublisher,omitempty"`

	// When enabled, the RTSP filter persists the control connection, which is being resumed, to the correct server. The default value is disabled.
	SessionReconnect string `json:"sessionReconnect,omitempty"`

	// When the proxy option is set, specifies the name of the header in the RTSP proxy configuration that is passed from the client-side virtual server to the server-side virtual server. Note that the name of the header must begin with X-.
	ProxyHeader string `json:"proxyHeader,omitempty"`

	// Specifies the maximum amount of data that the RTSP filter buffers before dropping the connection. The default value is 32768 bytes.
	MaxQueuedData int64 `json:"maxQueuedData,omitempty"`

	// Configures the ALG log profile that controls logging.
	LogProfile string `json:"logProfile,omitempty"`

	// When enabled, the RTSP filter automatically persists Real Networks RTSP over HTTP using the RTSP port. The default value is enabled. If you disable this parameter, you can override the default behavior with an iRule.
	RealHttpPersistence string `json:"realHttpPersistence,omitempty"`

	// Specifies the maximum size of an RTSP request or response header that the RTSP filter allows before dropping the connection. The default value is 4096 bytes.
	MaxHeaderSize int64 `json:"maxHeaderSize,omitempty"`
}

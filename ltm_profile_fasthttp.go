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
type LtmProfileFasthttp struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
	ConnpoolMaxReuse int64 `json:"connpoolMaxReuse,omitempty"`

	// Specifies the maximum number of requests that the system can receive on a client-side connection, before the system closes the connection. A setting of 0 specifies that requests are not limited. The default value is 0 (zero).
	MaxRequests int64 `json:"maxRequests,omitempty"`

	// Specifies a string that the system inserts as a header in an HTTP request. If the header exists already, the system does not replace it.
	HeaderInsert string `json:"headerInsert,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option.
	ConnpoolReplenish string `json:"connpoolReplenish,omitempty"`

	// Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
	ConnpoolMinSize int64 `json:"connpoolMinSize,omitempty"`

	// Specifies whether the system inserts the XForwarded For header in an HTTP request with the client IP address, to use with connection pooling. Enabling this specifies that the system inserts the XForwarded For header with the client IP address. Disabling this specifies that the system does not insert the XForwarded For header.
	InsertXforwardedFor string `json:"insertXforwardedFor,omitempty"`

	// Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
	ConnpoolStep int64 `json:"connpoolStep,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
	ConnpoolIdleTimeoutOverride int64 `json:"connpoolIdleTimeoutOverride,omitempty"`

	// Specifies whether to support server sack option in cookie response by default. The default is disabled.
	ServerSack string `json:"serverSack,omitempty"`

	// Specifies the number of seconds after which the system closes a client connection, when the system either receives a client FIN packet or sends a FIN packet. This setting overrides the idle timeout setting. The default value is 5.
	ClientCloseTimeout int64 `json:"clientCloseTimeout,omitempty"`

	// Specifies whether to support server timestamp option in cookie response by default. The default is disabled.
	ServerTimestamp string `json:"serverTimestamp,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
	ForceHttp10Response string `json:"forceHttp_10Response,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
	ConnpoolMaxSize int64 `json:"connpoolMaxSize,omitempty"`

	// Specifies whether or not to use hardware SYN Cookie when cross system limit. The default is disabled.
	HardwareSynCookie string `json:"hardwareSynCookie,omitempty"`

	// Specifies how the system handles closing a connection. The default value is enabled, which allows unclean shutdown of a client connection. Use disabled to prevent unclean shutdown of a client connection. Fast specifies that the system sends a RESET packet to close the connection only if the client attempts to send further data after the response has completed.
	UncleanShutdown string `json:"uncleanShutdown,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the number of seconds after which a connection is eligible for deletion, when the connection has no traffic. The default value is 300 seconds.
	IdleTimeout int64 `json:"idleTimeout,omitempty"`

	// Enables or disables HTTP 1.1 close workarounds. The default value is disabled.
	Http11CloseWorkarounds string `json:"http_11CloseWorkarounds,omitempty"`

	// When enabled, the system parses HTTP data in the stream. Disable this setting if you want to use the performance HTTP profile to shield against denial-of-service attacks against non-HTTP protocols. The default value is enabled.
	Layer7 string `json:"layer_7,omitempty"`

	// When enabled, the system sends a TCP RESET packet when a connection times out, and deletes the connection. The default value is enabled.
	ResetOnTimeout string `json:"resetOnTimeout,omitempty"`

	// Specifies the size of the receive window, in bytes. The minimum and default value is 65535 bytes without scale.
	ReceiveWindowSize int64 `json:"receiveWindowSize,omitempty"`

	// Specifies the number of seconds after which the system closes a client connection, when the system either receives a client FIN packet or send a FIN packet. This setting overrides the idle timeout setting. The default value is 5.
	ServerCloseTimeout int64 `json:"serverCloseTimeout,omitempty"`

	// Specifies a maximum segment size (MSS) override for server-side connections. The default value is 0 (zero), which corresponds to an MSS of 1460. You can specify any integer between 536 and 1460.
	MssOverride int64 `json:"mssOverride,omitempty"`

	// Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
	MaxHeaderSize int64 `json:"maxHeaderSize,omitempty"`
}
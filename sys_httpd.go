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
type SysHttpd struct {

	// The list of SSL Protocols to accept on the management console.
	SslProtocol string `json:"sslProtocol,omitempty"`

	// Specifies the minimum httpd message level to include in the system log. The default value is warn.
	LogLevel string `json:"logLevel,omitempty"`

	// Specifies, in bytes per second, the minimum average rate at which the request body must be received. The default value is 500.
	RequestBodyMinRate int64 `json:"requestBodyMinRate,omitempty"`

	// Specifies the maximum allowable time skew in seconds for OCSP response validation. The default is 300 seconds.
	SslOcspResponseTimeSkew int64 `json:"sslOcspResponseTimeSkew,omitempty"`

	// Specifies the maximum number of concurrent connections to the GUI. The default value is 10.
	MaxClients int64 `json:"maxClients,omitempty"`

	// Specifies the name of the file that contains the Certificate Authority (CA) certificate file. The default id none.
	SslCaCertFile string `json:"sslCaCertFile,omitempty"`

	SslPort int64 `json:"sslPort,omitempty"`

	// Specifies, in seconds, the maximum time allowed to receive all of the request headers, if the request-header-min-rate option is used, in which case the timeout is extended as more data arrives. Ignored if request-header-min-rate is not used. A value of 0 means no limit. The default value is 40.
	RequestHeaderMaxTimeout int64 `json:"requestHeaderMaxTimeout,omitempty"`

	// Specifies the ciphers that the system uses.
	SslCiphersuite string `json:"sslCiphersuite,omitempty"`

	// Specifies if the client certificate needs to be verified for SSL session establishment. The default is none.
	SslVerifyClient string `json:"sslVerifyClient,omitempty"`

	// Specifies the seconds before FastCGI timeout.
	FastcgiTimeout int64 `json:"fastcgiTimeout,omitempty"`

	// Specifies OCSP validation of the client certificate chain. The default is off.
	SslOcspEnable string `json:"sslOcspEnable,omitempty"`

	// Specifies the default responder URI for OCSP validation. The default is http://localhost.localdomain. The default responder value should always be preceded with http://
	SslOcspDefaultResponder string `json:"sslOcspDefaultResponder,omitempty"`

	// Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk.
	Include string `json:"include,omitempty"`

	// Specifies the force use of default responder URI for OCSP validation. The default is off.
	SslOcspOverrideResponder string `json:"sslOcspOverrideResponder,omitempty"`

	// Specifies whether the check for consistent inbound IP for the entire web session is enforced or not.
	AuthPamValidateIp string `json:"authPamValidateIp,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the maximum allowable time in seconds for OCSP responses. The default value is 300 seconds.
	SslOcspResponderTimeout int64 `json:"sslOcspResponderTimeout,omitempty"`

	// Specifies, in seconds, the time allowed to receive all of the request headers. This time includes completion of the SSL handshake. A value of 0 means no limit. If you use the request-header-min-rate option, this represents the initial value for the timeout, which will be extended as more data arrives. The default value is 20.
	RequestHeaderTimeout int64 `json:"requestHeaderTimeout,omitempty"`

	// Specifies the name of the file that contains the SSL certificate chain. The default is none.
	SslCertchainfile string `json:"sslCertchainfile,omitempty"`

	// Specifies the maximum allowable age in seconds for OCSP responses. A value of -1 specifies that a maximum age is not enforced. The default value is -1.
	SslOcspResponseMaxAge int64 `json:"sslOcspResponseMaxAge,omitempty"`

	// Specifies whether the system should redirect HTTP requests targeted at the configuration utility to HTTPS. The default value is disabled.
	RedirectHttpToHttps string `json:"redirectHttpToHttps,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies, in seconds, the time allowed to receive all of the request body. A value of 0 means no limit. If you use the request-body-min-rate option, this represents the initial value for the timeout, which will be extended as more data arrives. The default value is 60.
	RequestBodyTimeout int64 `json:"requestBodyTimeout,omitempty"`

	// Specifies, in bytes per second, the minimum average rate at which the request headers must be received. The default value is 500.
	RequestHeaderMinRate int64 `json:"requestHeaderMinRate,omitempty"`

	// Specifies whether to lookup hostname or not. The default value is off.
	HostnameLookup string `json:"hostnameLookup,omitempty"`

	// Specifies, in seconds, the maximum time allowed to receive all of the request body, if the request-body-min-rate option is used, in which case the timeout is extended as more data arrives. Ignored if request-body-min-rate is not used. A value of 0 means no limit. The default value is 0.
	RequestBodyMaxTimeout int64 `json:"requestBodyMaxTimeout,omitempty"`

	// Specifies whether browser session timeout occurs when the dashboard is running. The default value is disabled.
	AuthPamDashboardTimeout string `json:"authPamDashboardTimeout,omitempty"`

	// Specifies the seconds before GUI session timeout.
	AuthPamIdleTimeout int64 `json:"authPamIdleTimeout,omitempty"`

	// Adds or deletes IP addresses, partial IP addresses, and IP address ranges, hostnames, partial hostnames, domain names, partial domain names, and network and netmask pairs for the HTTP clients from which the httpd daemon accepts requests. The default value is all. Warning Using the value none resets the httpd daemon to allow all HTTP clients access to the system. F5 recommends that you do not use the value none with the httpd command.
	Allow string `json:"allow,omitempty"`

	// Specifies the name for the authentication realm. The default value is BIG-IP.
	AuthName string `json:"authName,omitempty"`

	// Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk.
	SslInclude string `json:"sslInclude,omitempty"`

	// Specifies the name of the file that contains the SSL certificate. The default value is /etc/httpd/conf/ssl.crt/server.crt. Note that the path to the file must start with /etc/httpd/conf/ssl.crt/ or /config/httpd/conf/ssl.crt/ unless the path is a relative path.  If the path is a relative path, then it must start with conf/ssl.crt/.
	SslCertfile string `json:"sslCertfile,omitempty"`

	// Specifies maximum depth of CA certificates in client certificate verification. The default is 10.
	SslVerifyDepth int64 `json:"sslVerifyDepth,omitempty"`

	// Specifies the name of the file that contains the SSL certificate key. The default value is /etc/httpd/conf/ssl.key/server.key. Note that the path to the file must start with /etc/httpd/conf/ssl.key/ or /config/httpd/conf/ssl.key/ unless the path is a relative path. If the path is a relative path, then it must start with conf/ssl.key/. When you change the key file, you must also change the certificate file. In other words, the following command does not work to change the key: bigpipe httpd sslcertkeyfile  string . Instead, you must use this command: { bigpipe httpd sslcertfile  string  sslcerkeyfile  string  }.
	SslCertkeyfile string `json:"sslCertkeyfile,omitempty"`
}

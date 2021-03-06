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
type LtmPolicyStrategyOperands struct {

	// Match everything to the strategy.
	All bool `json:"all,omitempty"`

	// Match on a code.
	Code bool `json:"code,omitempty"`

	// Match on a given protocol.
	Protocol bool `json:"protocol,omitempty"`

	// Match on an HTTP URI.
	HttpUri bool `json:"httpUri,omitempty"`

	// Match on a major.
	Major bool `json:"major,omitempty"`

	// Match on a TCP stream.
	Tcp bool `json:"tcp,omitempty"`

	// Match on round trip time.
	Rtt bool `json:"rtt,omitempty"`

	Org bool `json:"org,omitempty"`

	Continent bool `json:"continent,omitempty"`

	Last1min bool `json:"last_1min,omitempty"`

	// Match on an http basic authorization.
	HttpBasicAuth bool `json:"httpBasicAuth,omitempty"`

	BrowserType bool `json:"browserType,omitempty"`

	// Match on maximum segment size.
	Mss bool `json:"mss,omitempty"`

	// Use this strategy on a response.
	Response bool `json:"response,omitempty"`

	RegionCode bool `json:"regionCode,omitempty"`

	// Match on cipher bits.
	CipherBits bool `json:"cipherBits,omitempty"`

	// Match on text.
	Text bool `json:"text,omitempty"`

	DeviceMake bool `json:"deviceMake,omitempty"`

	// Match on a scheme used.
	Scheme bool `json:"scheme,omitempty"`

	// Match on an HTTP header.
	HttpHeader bool `json:"httpHeader,omitempty"`

	// Match on a minor.
	Minor bool `json:"minor,omitempty"`

	// Match on a given username.
	Username bool `json:"username,omitempty"`

	// Match on a particular VLAN id.
	VlanId bool `json:"vlanId,omitempty"`

	CountryName bool `json:"countryName,omitempty"`

	// Match on an HTTP method.
	HttpMethod bool `json:"httpMethod,omitempty"`

	// Match on a given name.
	TmName bool `json:"tmName,omitempty"`

	// Match on a particular route domain.
	RouteDomain bool `json:"routeDomain,omitempty"`

	// Match on a given password.
	Password bool `json:"password,omitempty"`

	// Match on an HTTP version.
	HttpVersion bool `json:"httpVersion,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Match on a geographic property of an IP address.
	Geoip bool `json:"geoip,omitempty"`

	// Match on an HTTP status.
	HttpStatus bool `json:"httpStatus,omitempty"`

	// Match on a given value.
	Value bool `json:"value,omitempty"`

	// Match on an HTTP cookie.
	HttpCookie bool `json:"httpCookie,omitempty"`

	// Match on an unnamed query parameter.
	UnnamedQueryParameter bool `json:"unnamedQueryParameter,omitempty"`

	// Match on the specified domain.
	Domain bool `json:"domain,omitempty"`

	CountryCode bool `json:"countryCode,omitempty"`

	// Match on a given path.
	PathSegment bool `json:"pathSegment,omitempty"`

	// Match on a particular port.
	Port bool `json:"port,omitempty"`

	// Match on an HTTP set cookie.
	HttpSetCookie bool `json:"httpSetCookie,omitempty"`

	DeviceModel bool `json:"deviceModel,omitempty"`

	// Match on a version.
	Version bool `json:"version,omitempty"`

	RegionName bool `json:"regionName,omitempty"`

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Match on a particular VLAN.
	Vlan bool `json:"vlan,omitempty"`

	// Match on a given query string.
	QueryString bool `json:"queryString,omitempty"`

	// Match on cpu usage percentage for the past 15secs, 1min or 5min intervals.
	CpuUsage bool `json:"cpuUsage,omitempty"`

	// Match on a particular host.
	Host bool `json:"host,omitempty"`

	// Match on a specific cipher.
	Cipher bool `json:"cipher,omitempty"`

	// Match on a particular path.
	Path bool `json:"path,omitempty"`

	// Match on a given expiry.
	Expiry bool `json:"expiry,omitempty"`

	Last15secs bool `json:"last_15secs,omitempty"`

	Last5mins bool `json:"last_5mins,omitempty"`

	// Match on a client SSL stream.
	ClientSsl bool `json:"clientSsl,omitempty"`

	BrowserVersion bool `json:"browserVersion,omitempty"`

	// Match on a given extension.
	Extension bool `json:"extension,omitempty"`

	UserAgentToken bool `json:"userAgentToken,omitempty"`

	Isp bool `json:"isp,omitempty"`

	// Use this strategy on a request.
	Request bool `json:"request,omitempty"`

	// Match on a given query parameter.
	QueryParameter bool `json:"queryParameter,omitempty"`

	// Match on an HTTP referer.
	HttpReferer bool `json:"httpReferer,omitempty"`

	// Match on an HTTP user agent substring or version.
	HttpUserAgent bool `json:"httpUserAgent,omitempty"`

	// Match on an HTTP host.
	HttpHost bool `json:"httpHost,omitempty"`
}

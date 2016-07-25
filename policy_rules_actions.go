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

type PolicyRulesActions struct {

	DestinationAddress bool `json:"destinationAddress,omitempty"`

	// This action has the following code.
	Code int64 `json:"code,omitempty"`

	// This action will use the following protocol.
	Protocol string `json:"protocol,omitempty"`

	Facility string `json:"facility,omitempty"`

	// This specifies a text value.
	Text string `json:"text,omitempty"`

	// This action will apply this next hop policy.
	Nexthop string `json:"nexthop,omitempty"`

	// This name of a value used in an action.
	TmName string `json:"tmName,omitempty"`

	// This action will replace a value.
	Replace bool `json:"replace,omitempty"`

	// This action will clone the stream and send it to this pool.
	ClonePool string `json:"clonePool,omitempty"`

	// This action will invoke a tcl script.
	Tcl bool `json:"tcl,omitempty"`

	CookieRewrite bool `json:"cookieRewrite,omitempty"`

	SslServerHandshake bool `json:"sslServerHandshake,omitempty"`

	// This action will create this message.
	Message string `json:"message,omitempty"`

	// This action will use following internal virtual.
	InternalVirtual string `json:"internalVirtual,omitempty"`

	// This action will select an option.
	Select_ bool `json:"select,omitempty"`

	// This action will redirect a request.
	Redirect bool `json:"redirect,omitempty"`

	// This action will insert/replace content.
	Content string `json:"content,omitempty"`

	// This action will use an AM policy.
	Wam bool `json:"wam,omitempty"`

	// This action will log a value.
	Log bool `json:"log,omitempty"`

	// This action will run this script.
	Script string `json:"script,omitempty"`

	// This policy action is performed on basic http authentication.
	HttpBasicAuth bool `json:"httpBasicAuth,omitempty"`

	// This action will modify server-side SSL behavior.
	ServerSsl bool `json:"serverSsl,omitempty"`

	// This action will affect caching.
	Cache bool `json:"cache,omitempty"`

	// This action will direct the stream to this node.
	Node string `json:"node,omitempty"`

	Priority string `json:"priority,omitempty"`

	// This action will come from the given location.
	Location string `json:"location,omitempty"`

	Persist bool `json:"persist,omitempty"`

	// This action will use this policy.
	Policy string `json:"policy,omitempty"`

	// This action will use the specified scheme.
	Scheme string `json:"scheme,omitempty"`

	// This action will use this snat rule.
	Snat string `json:"snat,omitempty"`

	// This action will act on a given HTTP header.
	HttpHeader bool `json:"httpHeader,omitempty"`

	// This action will enable a profile.
	Enable bool `json:"enable,omitempty"`

	// This action will affect the vlan with this id.
	VlanId int64 `json:"vlanId,omitempty"`

	// This action will affect forwarding.
	Forward bool `json:"forward,omitempty"`

	// Specify an IPv4/IPv6 netmask, e.g. 192.168.10.0/24 or 192.168.10.0/255.255.255.0
	Netmask string `json:"netmask,omitempty"`

	// This action will affect a reply to a given HTTP request.
	HttpReply bool `json:"httpReply,omitempty"`

	// This action will use the following application.
	Application string `json:"application,omitempty"`

	// This action will reset a connection.
	Reset bool `json:"reset,omitempty"`

	// This action will insert a value.
	Insert bool `json:"insert,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// This action will enable or disable tcp nagle.
	TcpNagle bool `json:"tcpNagle,omitempty"`

	// This action will remove a value.
	Remove bool `json:"remove,omitempty"`

	SslSessionId bool `json:"sslSessionId,omitempty"`

	// Timeout value in seconds
	Timeout int64 `json:"timeout,omitempty"`

	// This action will on a given HTTP cookie.
	HttpCookie bool `json:"httpCookie,omitempty"`

	SourceAddress bool `json:"sourceAddress,omitempty"`

	// This action will modify request adaptation.
	RequestAdapt bool `json:"requestAdapt,omitempty"`

	// This action will use the specified profile.
	Profile string `json:"profile,omitempty"`

	// This action will act on the domain.
	Domain string `json:"domain,omitempty"`

	CookieHash bool `json:"cookieHash,omitempty"`

	// This action will affect a pin.
	Pin bool `json:"pin,omitempty"`

	// This action will direct the stream to this pool.
	Pool string `json:"pool,omitempty"`

	// This policy action is performed on connection requests.
	Request bool `json:"request,omitempty"`

	Universal bool `json:"universal,omitempty"`

	// This action will come from a profile.
	FromProfile string `json:"fromProfile,omitempty"`

	// This action will use AVR
	Avr bool `json:"avr,omitempty"`

	// This action will direct the stream to this member.
	Member string `json:"member,omitempty"`

	// This action will use the following value.
	Value string `json:"value,omitempty"`

	// This action will act on HTTP streams.
	Http bool `json:"http,omitempty"`

	// This action will act on the following port.
	Port int64 `json:"port,omitempty"`

	// This action will set an HTTP cookie.
	HttpSetCookie bool `json:"httpSetCookie,omitempty"`

	// This action will apply to the following category.
	Category string `json:"category,omitempty"`

	CookieInsert bool `json:"cookieInsert,omitempty"`

	// This action will use this snatpool.
	Snatpool string `json:"snatpool,omitempty"`

	SslServerHello bool `json:"sslServerHello,omitempty"`

	// This action will use this virtual server.
	Virtual string `json:"virtual,omitempty"`

	// This action will write a value.
	Write bool `json:"write,omitempty"`

	// This action will affect PEM.
	Pem bool `json:"pem,omitempty"`

	// This action will modify response adaptation.
	ResponseAdapt bool `json:"responseAdapt,omitempty"`

	// This action will set a variable.
	SetVariable bool `json:"setVariable,omitempty"`

	// This action will classify a stream.
	Classify bool `json:"classify,omitempty"`

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	Hash bool `json:"hash,omitempty"`

	ExpirySecs int64 `json:"expirySecs,omitempty"`

	SslClientHello bool `json:"sslClientHello,omitempty"`

	// This action will affect compression.
	Compress bool `json:"compress,omitempty"`

	Carp bool `json:"carp,omitempty"`

	// This action will affect this vlan.
	Vlan string `json:"vlan,omitempty"`

	// This action will modify the query string of a URL.
	QueryString string `json:"queryString,omitempty"`

	// This action will use the given iFile.
	Ifile string `json:"ifile,omitempty"`

	// This action will have the following expiry.
	Expiry string `json:"expiry,omitempty"`

	// This action will affect the host.
	Host string `json:"host,omitempty"`

	// This action will disable a profile.
	Disable bool `json:"disable,omitempty"`

	Key string `json:"key,omitempty"`

	// Specify offset parameter
	Offset int64 `json:"offset,omitempty"`

	// This action will use the following path.
	Path string `json:"path,omitempty"`

	LtmPolicy bool `json:"ltmPolicy,omitempty"`

	// This action will affect decompression.
	Decompress bool `json:"decompress,omitempty"`

	CookiePassive bool `json:"cookiePassive,omitempty"`

	// This policy action is performed on connection responses.
	Response bool `json:"response,omitempty"`

	// This action will use an ASM policy.
	Asm bool `json:"asm,omitempty"`

	// This action will defer an action.
	Defer_ bool `json:"defer,omitempty"`

	// This action will use the specified extension.
	Extension string `json:"extension,omitempty"`

	IpAddress string `json:"ipAddress,omitempty"`

	// This action will act on given HTTP URIs.
	HttpUri bool `json:"httpUri,omitempty"`

	Length int64 `json:"length,omitempty"`

	// This action will set a status.
	Status int64 `json:"status,omitempty"`

	// This action will act on a given HTTP host.
	HttpHost bool `json:"httpHost,omitempty"`

	// This action will act on a given HTTP referer.
	HttpReferer bool `json:"httpReferer,omitempty"`

	Uie bool `json:"uie,omitempty"`

	// This action will evaluate this expression.
	Expression string `json:"expression,omitempty"`

	// This action will use a layer 7 DOS policy.
	L7dos bool `json:"l7dos,omitempty"`

	// This action will apply this rate class policy.
	Rateclass string `json:"rateclass,omitempty"`
}

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

type DnsCacheValidatingResolver struct {

	AppService string `json:"appService,omitempty"`

	// Fetch DNSKEY early in validation process. The default value is yes.
	PrefetchKey string `json:"prefetchKey,omitempty"`

	// The time allowed for a query to stay in the queue before replaced by a new query when the number of concurrent distinct queries exceeds the limit. The default value is 200 milliseconds.
	AllowedQueryTime int64 `json:"allowedQueryTime,omitempty"`

	// Number of bytes allocated for the DNSKEY cache. The default value is 1m
	KeyCacheSize int64 `json:"keyCacheSize,omitempty"`

	// Maximum number of concurrent UDP flows used by the resolver. The default value is 8192.
	MaxConcurrentUdp int64 `json:"maxConcurrentUdp,omitempty"`

	// Number of DNS nameservers to cache. The default value is 16k.
	NameserverCacheCount int64 `json:"nameserverCacheCount,omitempty"`

	// Enables resolver to randomize the case of query names. The default value is yes.
	RandomizeQueryNameCase string `json:"randomizeQueryNameCase,omitempty"`

	// Number of bytes allocated for the resource record set cache. The default value is 10m.
	RrsetCacheSize int64 `json:"rrsetCacheSize,omitempty"`

	// Route domain for resolver outbound traffic. The default value is the default route domain.
	RouteDomain string `json:"routeDomain,omitempty"`

	// Enables resolver to issue udp queries. The default value is yes.
	UseUdp string `json:"useUdp,omitempty"`

	// List of DNSKEY or DS resource records used to establish DNSSEC validator trust with a DLV registry.  Specified in string form (e.g. dig or drill format). The default is none.
	DlvAnchors string `json:"dlvAnchors,omitempty"`

	// Ignore client queries setting of checking-disabled. Perform validation anyway and only return secure answers. The default value is no.
	IgnoreCd string `json:"ignoreCd,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Maximum number of concurrent TCP flows used by the resolver. The default value is 20.
	MaxConcurrentTcp int64 `json:"maxConcurrentTcp,omitempty"`

	// Maximum number of concurrent queries used by the resolver. The default value is 1024.
	MaxConcurrentQueries int64 `json:"maxConcurrentQueries,omitempty"`

	// Enables resolver to issue tcp queries. The default value is yes.
	UseTcp string `json:"useTcp,omitempty"`

	// Enables resolver to issue IPv6 queries. The default value is yes.
	UseIpv6 string `json:"useIpv6,omitempty"`

	Idx float32 `json:"idx,omitempty"`

	// Enables resolver to issue IPv4 queries. The default value is yes.
	UseIpv4 string `json:"useIpv4,omitempty"`

	Type_ string `json:"type,omitempty"`

	Partition string `json:"partition,omitempty"`

	// Answer queries for default zones: localhost, reverse 127.0.0.1 and ::1, and AS112 zones. The default value is no.
	AnswerDefaultZones string `json:"answerDefaultZones,omitempty"`

	// The threshold count of unsolicited query replies which triggers an alert (potential DOS attack underway). The default value is 0 (or off).
	UnwantedQueryReplyThreshold int64 `json:"unwantedQueryReplyThreshold,omitempty"`

	// Zones and associated resource records for which the cache will provide Authoritative answers.
	LocalZones string `json:"localZones,omitempty"`

	// List of DNSKEY or DS resource records used to establish DNSSEC validator trust.  Specified in string form (e.g. dig or drill format). The default is none.
	TrustAnchors string `json:"trustAnchors,omitempty"`

	// List of IP addresses to use for root name servers. Defaults are known Internet root servers.
	RootHints string `json:"rootHints,omitempty"`

	// Number of bytes allocated for the message cache. The default value is 1m
	MsgCacheSize int64 `json:"msgCacheSize,omitempty"`
}
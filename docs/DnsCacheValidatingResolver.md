# DnsCacheValidatingResolver

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**PrefetchKey** | **string** | Fetch DNSKEY early in validation process. The default value is yes. | [optional] [default to null]
**AllowedQueryTime** | **int64** | The time allowed for a query to stay in the queue before replaced by a new query when the number of concurrent distinct queries exceeds the limit. The default value is 200 milliseconds. | [optional] [default to 200]
**KeyCacheSize** | **int64** | Number of bytes allocated for the DNSKEY cache. The default value is 1m | [optional] [default to 1048576]
**MaxConcurrentUdp** | **int64** | Maximum number of concurrent UDP flows used by the resolver. The default value is 8192. | [optional] [default to 8192]
**NameserverCacheCount** | **int64** | Number of DNS nameservers to cache. The default value is 16k. | [optional] [default to 16536]
**RandomizeQueryNameCase** | **string** | Enables resolver to randomize the case of query names. The default value is yes. | [optional] [default to null]
**RrsetCacheSize** | **int64** | Number of bytes allocated for the resource record set cache. The default value is 10m. | [optional] [default to 10485760]
**RouteDomain** | **string** | Route domain for resolver outbound traffic. The default value is the default route domain. | [optional] [default to null]
**UseUdp** | **string** | Enables resolver to issue udp queries. The default value is yes. | [optional] [default to null]
**DlvAnchors** | **string** | List of DNSKEY or DS resource records used to establish DNSSEC validator trust with a DLV registry.  Specified in string form (e.g. dig or drill format). The default is none. | [optional] [default to null]
**IgnoreCd** | **string** | Ignore client queries setting of checking-disabled. Perform validation anyway and only return secure answers. The default value is no. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**MaxConcurrentTcp** | **int64** | Maximum number of concurrent TCP flows used by the resolver. The default value is 20. | [optional] [default to 20]
**MaxConcurrentQueries** | **int64** | Maximum number of concurrent queries used by the resolver. The default value is 1024. | [optional] [default to 1024]
**UseTcp** | **string** | Enables resolver to issue tcp queries. The default value is yes. | [optional] [default to null]
**UseIpv6** | **string** | Enables resolver to issue IPv6 queries. The default value is yes. | [optional] [default to null]
**Idx** | **float32** |  | [optional] [default to null]
**UseIpv4** | **string** | Enables resolver to issue IPv4 queries. The default value is yes. | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Partition** | **string** |  | [optional] [default to null]
**AnswerDefaultZones** | **string** | Answer queries for default zones: localhost, reverse 127.0.0.1 and ::1, and AS112 zones. The default value is no. | [optional] [default to null]
**UnwantedQueryReplyThreshold** | **int64** | The threshold count of unsolicited query replies which triggers an alert (potential DOS attack underway). The default value is 0 (or off). | [optional] [default to 0]
**LocalZones** | **string** | Zones and associated resource records for which the cache will provide Authoritative answers. | [optional] [default to null]
**TrustAnchors** | **string** | List of DNSKEY or DS resource records used to establish DNSSEC validator trust.  Specified in string form (e.g. dig or drill format). The default is none. | [optional] [default to null]
**RootHints** | **string** | List of IP addresses to use for root name servers. Defaults are known Internet root servers. | [optional] [default to null]
**MsgCacheSize** | **int64** | Number of bytes allocated for the message cache. The default value is 1m | [optional] [default to 1048576]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



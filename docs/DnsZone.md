# DnsZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**DnsExpressNotifyAction** | **string** | Action to take when a NOTIFY query is received for a configured zone.  Options are consume, bypass, and repeat. Default is consume, meaning the NOTIFY query is seen only by DNS Express. bypass means the query will NOT go to DNS Express, but any backend DNS resource (subject to DNS profile unhandled-query-action).  repeat means the NOTIFY will go to both DNS Express and any backend DNS resource.  If TSIG is configured, the signature is only validated for consume and repeat actions. NOTIFY responses are assumed to be sent by the backend DNS resource, except when the action is Consume and DNS Express will generate a response. | [optional] [default to null]
**DnsExpressServer** | **string** | Specifies the server from which to retrieve zone information for DNS Express. | [optional] [default to null]
**DnsExpressAllowNotify** | **string** | Specifies a list of IP addresses, in addition to the DNS Zone&#39;s DNS-Express Server address, which are allowed to notify the BIGIP of DNS Zone changes. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DnsExpressEnabled** | **string** | Specifies whether DNS Express is enabled to process queries for this zone. The default value is yes. | [optional] [default to null]
**ServerTsigKey** | **string** | Specifies the TSIG key associated with the DNS zone. | [optional] [default to null]
**DnsExpressNotifyTsigVerify** | **string** | Verify NOTIFY query TSIG for a DNS Express zone. Default is yes. | [optional] [default to null]
**ResponsePolicy** | **string** | Specifies if the zone contains response policy resource records. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



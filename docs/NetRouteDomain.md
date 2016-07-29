# NetRouteDomain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**FlowEvictionPolicy** | **string** | Specifies the flow eviction policy for the route domain to use when the connection limit is approached. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Parent** | **string** | Specifies the name of the route domain from which this route domain inherits settings. The default value is empty, which indicates no parent. | [optional] [default to null]
**BwcPolicy** | **string** | The bandwidth control policy used for traffic flow. | [optional] [default to null]
**ServicePolicy** | **string** | Specifies the service policy to use. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the route-domain resides. | [optional] [default to null]
**Strict** | **string** | Specifies whether the system allows a connection to span route domains. | [optional] [default to null]
**FwStagedPolicy** | **string** | Staged firewall policy. | [optional] [default to null]
**FwEnforcedPolicy** | **string** | Enforced firewall policy. | [optional] [default to null]
**ConnectionLimit** | **int64** | Specifies the maximum number of concurrent connections you want to allow for the route domain. | [optional] [default to 0]
**IpIntelligencePolicy** | **string** | Name of the IP Intelligence (Dynamic White/Black List) policy that is attached. | [optional] [default to null]
**Id** | **int64** | Numerical ID of the route domain | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



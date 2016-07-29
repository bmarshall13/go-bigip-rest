# NetStpTrunks

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**Priority** | **int64** | Specifies the trunk priority number. Each network trunk has an associated priority within each spanning tree instance. The relative values of the trunk priorities influence which trunks are chosen to carry network traffic. All other things being equal, trunks with numerically lower priority values are favored to carry traffic. Trunk priorities take values in the range 0-240 in steps of 16. The default trunk priority is 128, the middle of the valid range. | [optional] [default to 128]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ExternalPathCost** | **int64** | The relative cost of send network traffic through the trunk. The external path cost only applies to STP 0 (zero). See the help page for a detailed description. | [optional] [default to 0]
**InternalPathCost** | **int64** | The relative cost of send network traffic through the trunk. See the help page for a detailed description. | [optional] [default to 0]
**Type_** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetPacketFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Logging** | **string** | Enables or disables packet filter logging. If you omit this value, no logging is performed. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**BwcPolicy** | **string** | Specifies the name of a bwc policy. The value for the bwc policy association is the name of any existing bwc policy. If omitted, no bwc filter is applied. | [optional] [default to null]
**Vlan** | **string** | Specifies the VLAN to which the packet filter rule should apply. The value for this option is any VLAN name currently in existence. If you omit this value, the rule applies to all VLANs. | [optional] [default to null]
**Rule** | **string** | Specifies the BPF expression to match. The filter is mandatory, however you can leave it empty. If empty, the packet filter rule matches all packets. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Action** | **string** | Specifies the action that the packet filter rule should take. There is no default value; you must specify a value when you create a packet filter rule. | [default to null]
**Order** | **int64** | Specifies a sort order. The values for the sort order are integers greater than 0 (zero). No two rules may have the same sort order. | [default to 0]
**RateClass** | **string** | Specifies the name of a rate class. The value for the rate class association is the name of any existing rate class. If omitted, no rate filter is applied. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



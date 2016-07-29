# SysFolder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**DeviceGroup** | **string** | Associate this folder with a device failover group or device sync group. &#39;default&#39; to associate this folder with its parent&#39;s device group. &#39;non-default&#39; to leave this field&#39;s value untouched but disassociate this folder from its parent. | [optional] [default to null]
**Description** | **string** | User-defined description of the folder | [optional] [default to null]
**TrafficGroup** | **string** | Associate this folder with a network failover group. &#39;default&#39; to associate this folder with its parent&#39;s device group. &#39;non-default&#39; to leave this field&#39;s value untouched but disassociate this folder from its parent. | [optional] [default to null]
**NoRefCheck** | **string** | Specifies whether strict device group reference validation is performed during sync behavior on items in this folder | [optional] [default to null]
**InheritedDevicegroup** | **string** | Read-only. Shows whether this folder will automatically remain with the same device-group as its parent folder. Use &#39;device-group default&#39; or &#39;device-group non-default&#39; to set this. | [optional] [default to null]
**Hidden** | **string** | Specifies if this folder will be hidden.  If set to &#39;true&#39;, this folder will be hidden from standard command usage.  Administrators can display, modify or remove hidden folders using the &#39;-hidden&#39; option. | [optional] [default to null]
**InheritedTrafficGroup** | **string** | Read-only. Shows whether this folder will automatically remain with the same traffic-group as its parent folder. Use &#39;traffic-group default&#39; or &#39;traffic-group non-default&#39; to set this. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



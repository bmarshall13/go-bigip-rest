# SysDiskLogicalDisk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**VgReserved** | **int64** | Specifies the reserved logical disk space (MiB). (Reserved space can only be used by certain system utilities for base OS growth - specifically it will not be included in the disk space available to resource provisioning) | [optional] [default to null]
**VgFree** | **int64** | Specifies the usable free space (MiB) available in Logical Disk. | [optional] [default to null]
**Size** | **int64** | Specifies the size (MiB) of the Logical Disk. | [optional] [default to null]
**VgInUse** | **int64** | Specifies the total logical disk space (MiB) in use. | [optional] [default to null]
**Mode** | **string** | The mode is how the disk is used - it can be dedicated to a specific application or set for multiple uses (mixed mode). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



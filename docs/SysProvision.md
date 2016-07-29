# SysProvision

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Level** | **string** | Specifies the level of resources that you want to provision for a module. \&quot;minimum\&quot; specifies that you want to provision the minimal amount of resources for the module you are provisioning. \&quot;nominal\&quot; specifies that you want to share all of the available resources equally among all of the modules that are licensed of the unit. \&quot;dedicated\&quot; specifies that all resources are dedicated to the module you are provisioning. For all other modules, the level option must be set to none. \&quot;none\&quot; specifies that you do not want to provision any resources for this module. \&quot;custom\&quot; F5 does not recommend that you specify this level. | [optional] [default to null]
**MemoryRatio** | **int64** | Use this option only when the level option is set to custom. F5 recommends that you do not modify this option. The default value is zero. | [optional] [default to 0]
**CpuRatio** | **int64** | Use this option only when the level option is set to custom. F5 recommends that you do not modify this option. The default value is zero. | [optional] [default to 0]
**DiskRatio** | **int64** | Use this option only when the level option is set to custom. F5 recommends that you do not modify this option. The default value is zero. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



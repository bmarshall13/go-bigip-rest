# SysApplicationService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the application resides. | [optional] [default to null]
**DeviceGroup** | **string** | The name of the device group that the application service is assigned to. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**ExecuteAction** | **string** | Run the specified template action associated with the application. | [optional] [default to null]
**TemplateModified** | **string** | Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes. | [optional] [default to null]
**TrafficGroup** | **string** | The name of the traffic group that the application service is assigned to. | [optional] [default to null]
**InheritedDevicegroup** | **string** | Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use &#39;device-group default&#39; or &#39;device-group non-default&#39; to set this. | [optional] [default to null]
**TemplatePrerequisiteErrors** | **string** | Indicates any missing prerequisites associated with the template that defines this application. | [optional] [default to null]
**Template** | **string** | The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template. | [optional] [default to null]
**InheritedTrafficGroup** | **string** | Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use &#39;traffic-group default&#39; or &#39;traffic-group non-default&#39; to set this. | [optional] [default to null]
**StrictUpdates** | **string** | Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system&#39;s application management interfaces. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



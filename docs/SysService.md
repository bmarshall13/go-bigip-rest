# SysService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Enable** | **bool** | Specifies to enable the services. Please note that disabled services do not run until you enable, and then start the services again. | [optional] [default to null]
**Force** | **bool** | Ends the services. Use this option as a last resort when services do not terminate automatically. Note that this option does not operate on the default set of services, you must specify the services that you want to end. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**Memstat** | **bool** | Provides a status on the pre-set memory limit of the services and how much memory each service is currently using. | [optional] [default to null]
**Reinit** | **bool** | Specifies to reinitialize the services or re-read the config files for the services. You must specify the services that you want to reinitialize. | [optional] [default to null]
**Remove** | **bool** | Specifies to remove a service so that the service does not start at boot time or with the default service set. | [optional] [default to null]
**Add** | **bool** | Specifies to add a service so that the service starts at boot time, or with the default service set. | [optional] [default to null]
**Disable** | **bool** | Specifies to manually disable the services. After you disable services, you must enable, and then start the services before the services run again. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



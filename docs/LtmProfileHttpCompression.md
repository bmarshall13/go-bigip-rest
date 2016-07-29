# LtmProfileHttpCompression

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**GzipMemoryLevel** | **int64** | Specifies amount of memory (in kilobytes) that the system uses when compressing a server response. The value will be rounded up to the nearest power of 2. The default value is 8k. The maximum value is 256k. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ContentTypeInclude** | **string** | Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress. | [optional] [default to null]
**BrowserWorkarounds** | **string** | Enables or disables browser workarounds. The default value is disabled. | [optional] [default to null]
**Selective** | **string** | Enables or disables selective compression mode. | [optional] [default to null]
**DefaultsFrom** | **string** | Specifies the profile that you want to use as the parent profile. | [optional] [default to null]
**BufferSize** | **int64** | Specifies the maximum number of uncompressed bytes that the system buffers before determining whether or not to compress the response. Useful when the headers of a server response do not specify the length of the response content. The default value is 4096. | [optional] [default to null]
**GzipWindowSize** | **int64** | Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The value will be rounded up to the nearest power of 2. The default is 16k. The maximum value is 128k. | [optional] [default to null]
**CpuSaverLow** | **int64** | Specifies the percent CPU usage at which the system resumes content compression at the user-defined rates. The default value is 75 percent. | [optional] [default to 75]
**MethodPrefer** | **string** | Specifies the type of compression that is preferred by the system. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**UriInclude** | **string** | Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress. | [optional] [default to null]
**VaryHeader** | **string** | Enables or disables the insertion of a Vary header into cacheable server responses. The default value is enabled. | [optional] [default to null]
**GzipLevel** | **int64** | Specifies a value that determines the amount of memory that the system uses when compressing a server response. The default is 1. | [optional] [default to null]
**ContentTypeExclude** | **string** | Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which this profile resides. | [optional] [default to null]
**CpuSaverHigh** | **int64** | Specifies the percent of CPU usage at which the system starts automatically decreasing the amount of content being compressed, as well as the amount of compression which the system is applying. The default value is 90 percent. | [optional] [default to 90]
**CpuSaver** | **string** | Specifies the CPU saver setting. When the CPU saver is enabled, the system monitors the percent of CPU usage and adjusts compression rates automatically when the CPU usage reaches the percentage defined in the cpu saver low or the cpu saver high options. The default value is enabled. | [optional] [default to null]
**MinSize** | **int64** | Specifies the minimum length in bytes of a server response that is acceptable for compressing that response. The length in bytes applies to content length only, not headers. The default value is 1024. | [optional] [default to null]
**KeepAcceptEncoding** | **string** | Enables or disables keep accept encoding. When enabled, causes the target server, rather than the BIG-IP, to perform the data compression. | [optional] [default to null]
**AllowHttp10** | **string** | Enables or disables compression of HTTP/1.0 server responses. | [optional] [default to null]
**UriExclude** | **string** | Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



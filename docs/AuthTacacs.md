# AuthTacacs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Protocol** | **string** | Specifies the protocol associated with the value specified in the service option, which is a subset of the associated service being used for client authorization or system accounting. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Service** | **string** | Specifies the name of the service that the user is requesting to be authenticated to use. Identifying the service enables the TACACS+ server to behave differently for different types of authentication requests. This option is required. | [optional] [default to null]
**Encryption** | **string** | Enables or disables encryption of TACACS+ packets. Recommended for normal use. The default value is enabled. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**Authentication** | **string** | Specifies the process the system employs when sending authentication requests. The default is use-first-server. use-first-server specifies that the system sends authentication requests to only the first server in the list. use-all-servers specifies that the system sends an authentication request to each server until authentication succeeds, or until the system has sent a request to all servers in the list. | [optional] [default to null]
**Servers** | **string** | Specifies a host name or IP address for the TACACS+ server. This option is required. Possible values are a user-specified string, and none. You must specify a server when you create a TACACS+ configuration object. | [default to null]
**Secret** | **string** | Sets the secret key used to encrypt and decrypt packets sent or received from the server. This option is required. | [default to null]
**Debug** | **string** | Enables syslog-ng debugging information at LOG DEBUG level. Not recommended for normal use. The default value is disabled. | [optional] [default to null]
**Accounting** | **string** | If multiple TACACS+ servers are defined and pluggable authentication module (PAM) session accounting is enabled, sends accounting start and stop packets to the first available server or to all servers. The default value is send-to-first-server. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



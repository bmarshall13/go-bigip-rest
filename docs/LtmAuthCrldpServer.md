# LtmAuthCrldpServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**ReverseDn** | **string** | Specifies in which order the system attempts to match the base-dn value to the value of the X509v3 attribute crlDistributionPoints. When enabled, the system matches the base-dn value from left to right, or from the beginning of the DN string, to accommodate dirName strings in certificates such as C&#x3D;US,ST&#x3D;WA,L&#x3D;SEA,OU&#x3D;F5,CN&#x3D;xxx. The default value is disabled. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**BaseDn** | **string** | Specifies the LDAP base directory name for certificates that specify the CRL distribution point in directory name format (dirName). Use this option when the value of the X509v3 attribute crlDistributionPoints is of type dirName. In this case, the BIG-IP system attempts to match the value of the crlDistributionPoints attribute to the base-dn value. An example of a base-dn value is cn&#x3D;lxxx,dc&#x3D;f5,dc&#x3D;com. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Partition** | **string** | Displays the partition within which the server resides. | [optional] [default to null]
**Host** | **string** | Specifies an IP address for the CRLDP server. This option is required. | [optional] [default to null]
**Port** | **int64** | Specifies the port for CRLDP authentication traffic. The default port is 389. | [optional] [default to 389]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



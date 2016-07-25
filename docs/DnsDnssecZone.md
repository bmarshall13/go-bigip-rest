# DnsDnssecZone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Seps** | **string** | Secure Entry Point(s) (DS and DNSKEY resource records used as client trust anchors) of the zone. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Nsec3Iterations** | **int64** | Specifies the number of times to hash the Next Secure (NSEC3) names. The default value is 1. | [optional] [default to 1]
**Enabled** | **bool** | Specifies that the DNSSEC zone will be signed. | [optional] [default to null]
**XfrPrimarySoaSerial** | **float32** | The learned zone SOA serial number from the primary server. | [optional] [default to null]
**Disabled** | **bool** | Specifies that the DNSSEC zone will not be signed. | [optional] [default to null]
**Nsec3Algorithm** | **string** | Specifies the hash algorithm to use when creating the Next Secure (NSEC3) resource record. The default value is SHA1. | [optional] [default to null]
**XfrSoaSerial** | **float32** | The advertised zone SOA serial number to all clients. | [optional] [default to null]
**DsAlgorithm** | **string** | Specifies the hash algorithm to use when creating the Delegation Signer (DS) resource record. The default value is SHA256 | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SysSnmp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | User defined description. | [optional] [default to null]
**AgentTrap** | **string** | Enables or disables traps about the agent status. | [optional] [default to null]
**AllowedAddresses** | **string** | Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow. | [optional] [default to null]
**AgentAddresses** | **string** | A list of protocol/address combinations that the agent listens for traffic on. | [optional] [default to null]
**AuthTrap** | **string** | Enable or disable generating traps on authentication failures | [optional] [default to null]
**SysContact** | **string** | Specifies the contact information for the system administrator. | [optional] [default to null]
**L2forwardVlan** | **string** | Enables l2forwarding statistics for specified VLANs. | [optional] [default to null]
**TrapCommunity** | **string** | Default community name used when sending traps. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**LoadMax1** | **int64** | One minute load threshold before SNMPd sends a trap. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**LoadMax5** | **int64** | Five minute load threshold before SNMPd sends a trap. | [optional] [default to null]
**TrapSource** | **string** | Source used for sending traps. | [optional] [default to null]
**SysLocation** | **string** | Describes the system&#39;s physical location. | [optional] [default to null]
**SysServices** | **int64** | Sets the value of the system.sysServices.0 object. | [optional] [default to null]
**BigipTraps** | **string** | Enables or disables traps about device warnings. | [optional] [default to null]
**Include** | **string** | Adds information to the snmpd.conf file. Use with extreme caution or under the guidance of F5 support. | [optional] [default to null]
**LoadMax15** | **int64** | Fifteen minute load threshold before SNMPd sends a trap. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



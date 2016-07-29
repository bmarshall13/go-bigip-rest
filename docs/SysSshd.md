# SysSshd

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**LogLevel** | **string** | Specifies the minimum sshd message level to include in the system log. You must enter the following values: -- debug, debug1, debug2, debug3, which indicates that the minimum sshd message level that the system logs is the specified debugging level of messages. -- error, which indicates that the minimum sshd message level that the system logs is error. -- fatal, which indicates that the minimum sshd message level that the system logs is fatal. -- info, which indicates that the minimum sshd message level that the system logs is informational. -- quiet, which indicates that the system does not log sshd messages.-- verbose, which indicates that the system logs all sshd messages. The default is info. | [optional] [default to null]
**InactivityTimeout** | **int64** | Specifies the number of seconds before inactivity causes an SSH session to log out. The default value is 0 (zero) seconds, which indicates that inactivity timeout is disabled. | [optional] [default to null]
**BannerText** | **string** | When banner is enabled, specifies the text to include in the banner that displays when a user attempts to login to the system. | [optional] [default to null]
**Allow** | **string** | Adds a server to or removes a server from the /etc/hosts.allow file. Use this option to either add servers to the BIG-IP system that are allowed to access the system, or delete these servers from the system. Specify \&quot;none\&quot; to disallow ssh access to the system. Specify \&quot;replace-all-with { ALL }\&quot; to allow ssh access from any server. The default value is \&quot;replace-all-with { ALL }\&quot;. | [optional] [default to null]
**Login** | **string** | Enables or disables SSH logins to the system. The default is enabled. | [optional] [default to null]
**Include** | **string** | Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk. | [optional] [default to null]
**Banner** | **string** | Enables or disables the display of the banner text field when a user logs in to the system using SSH. The default value is disabled. | [optional] [default to null]
**Port** | **int64** | Specifies the TCP port to run SSHD | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



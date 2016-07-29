# SysGlobalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MgmtDhcp** | **string** | Specifies whether or not to enable DHCP client on the management interface. | [optional] [default to null]
**UsernamePrompt** | **string** | Specifies the text to use on the GUI login form above the field for the username. This allows you to give the user a hint as to which username to use if the BIG-IP has been configured to use an external user directory such as Active Directory or LDAP. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**AwsApiMaxConcurrency** | **int64** | Maximum concurrent connections allowed while making Amazon Web Service(AWS) api calls. | [optional] [default to 1]
**GuiSecurityBanner** | **string** | Specifies whether or not show GUI login security banner. | [optional] [default to null]
**AwsAccessKey** | **string** | Amazon Web Service(AWS) supplied access key needed to make secure requests to AWS. | [optional] [default to null]
**PasswordPrompt** | **string** | Specifies the text to use on the GUI login form above the field for the password. This allows you to give the user a hint as to which password to use if the BIG-IP has been configured to use an external user directory such as Active Directory or LDAP. | [optional] [default to null]
**FileLocalPathPrefix** | **string** | Specifies a list of folder prefixes that can be applied for file objects. This is a space separated list of folder prefixes, contained in curly braces. Example: \&quot;{file:///shared/}\&quot; or \&quot;{file:///file/object/folder/} {/shared/}\&quot;. By default the folders are \&quot;/shared/\&quot; and \&quot;/tmp/\&quot;, represented as \&quot;{/shared/} {/tmp/}\&quot;. | [optional] [default to null]
**CustomAddr** | **string** | Indicates a user-specified IP address for the system.  The default value is none. It is important to note that you must set the host-addr-mode option to custom, if you want to specify an IP address using custom addr. For more information, see the host-addr-mode option, below. | [optional] [default to null]
**HostsAllowInclude** | **string** | Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk. | [optional] [default to null]
**LcdDisplay** | **string** | Enables or disables the LCD display on the front of the system. The default is enabled. | [optional] [default to null]
**AwsSecretKey** | **string** | Amazon Web Service(AWS) supplied secret key needed to make secure requests to AWS. | [optional] [default to null]
**FailsafeAction** | **string** | Specifies the action that the system takes when the switch board fails. The default value is go offline abort tm. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**NetReboot** | **string** | Enables or disables the network reboot feature. The default is disabled.  If you enable this feature and then reboot the system, the system boots from an ISO image on the network, rather than from an internal media drive. Use this option only when you want to install software on the system, for example, for an upgrade or a re-installation. Note that this setting reverts to disabled after you reboot the system a second time. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**ConsoleInactivityTimeout** | **int64** | Specifies the number of seconds of inactivity before the system logs off a user that is logged on. The default value is 0. This means that no timeout is set. | [optional] [default to null]
**HostAddrMode** | **string** | Specifies the type of host address assigned to the system.  The default value is mgmt, which indicates that the host address is the management port of the system. If you use the state-mirror option, then the host address of the system is shared by the other system in a redundant pair. In case of system failure, the traffic to the other system is routed to this system. If you use the custom option, you must specify a custom IP address for the system using the custom addr option. For more information, see the custom addr option, above. | [optional] [default to null]
**Hostname** | **string** | Specifies a local name for the system. The default value is bigip1. | [optional] [default to null]
**GuiSetup** | **string** | Enables or disables the Setup utility in the browser-based Configuration utility. The default value is enabled. Note that when you configure a system using the command line interface, you should disable this option. Disabling the gui setup option of the system command allows your system administrators to use the browser-based Configuration utility without having to run the Setup Utility. | [optional] [default to null]
**GuiSecurityBannerText** | **string** | Specifies the text of GUI login security banner. | [optional] [default to null]
**QuietBoot** | **string** | Enables or disables the quiet boot feature. The default is enabled. If you enable this feature, the system suppresses informational text on the console during the boot cycle. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



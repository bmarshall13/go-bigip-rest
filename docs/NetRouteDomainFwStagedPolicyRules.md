# NetRouteDomainFwStagedPolicyRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** | The application service that the object belongs to. | [optional] [default to null]
**IruleSampleRate** | **int64** | Specifies the rate at which iRule will be triggered if the packet matches this firewall rule. | [optional] [default to 1]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**PlaceAfter** | **string** | Specifies that a new rule should be placed after another rule, first or last. If individual rules are being added (as opposed to specifying replace-all-with) then place-before or place-after must be specified. | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**Schedule** | **string** | Specifies a schedule for the rule. See security firewall schedule.  If the rule refers to a rule-list the rule-list will be enabled according to the schedule. When the rule list is enabled, the schedules defined within the rule-list will be honored. | [optional] [default to null]
**ServicePolicy** | **string** | Specifies the service policy to use. | [optional] [default to null]
**Irule** | **string** | Specifies the name of the iRule which will be triggered if the packet matches this firewall rule. | [optional] [default to null]
**Status** | **string** | Specifies whether the rule is enabled, disabled or scheduled. A rule that is enabled is always checked. A rule that is disabled is never checked. A rule that is scheduled is checked according to the corresponding schedule configuration. A rule that is scheduled must have an associated schedule configuration. | [optional] [default to null]
**Log** | **string** | Specifies whether the packet will be logged if it matches the rule.  Logging must also be enabled in the security log profile global-network configuration. Note that the statistics counter is always incremented when a packet matches a rule. | [optional] [default to null]
**RuleList** | **string** | Specifies a list of rules to evaluate. See security firewall rule-list. If a rule-list is specified then only the schedule and status properties effect the rule. | [optional] [default to null]
**Action** | **string** | Specifies the action that the system takes when a rule is matched. | [optional] [default to null]
**IpProtocol** | **string** | Specifies the IP protocol against which the packet will be compared. | [optional] [default to null]
**PlaceBefore** | **string** | Specifies that a new rule should be placed before another rule, first or last. If individual rules are being added (as opposed to specifying replace-all-with) then place-before or place-after must be specified. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



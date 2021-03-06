/*
 * BigIP iControl REST
 *
 * REST API for F5 BigIP. List of operations is not complete, nor known to be accurate.
 *
 * OpenAPI spec version: 12.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package f5api

// This describes a message sent to or received from some operations
type NetRouteDomainFwActive struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Specifies the rate at which iRule will be triggered if the packet matches this firewall rule.
	IruleSampleRate int64 `json:"iruleSampleRate,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies that a new rule should be placed after another rule, first or last. If individual rules are being added (as opposed to specifying replace-all-with) then place-before or place-after must be specified.
	PlaceAfter string `json:"placeAfter,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies a schedule for the rule. See security firewall schedule.  If the rule refers to a rule-list the rule-list will be enabled according to the schedule. When the rule list is enabled, the schedules defined within the rule-list will be honored.
	Schedule string `json:"schedule,omitempty"`

	// Specifies the service policy to use.
	ServicePolicy string `json:"servicePolicy,omitempty"`

	// Specifies the name of the iRule which will be triggered if the packet matches this firewall rule.
	Irule string `json:"irule,omitempty"`

	// Specifies whether the rule is enabled, disabled or scheduled. A rule that is enabled is always checked. A rule that is disabled is never checked. A rule that is scheduled is checked according to the corresponding schedule configuration. A rule that is scheduled must have an associated schedule configuration.
	Status string `json:"status,omitempty"`

	// Specifies whether the packet will be logged if it matches the rule.  Logging must also be enabled in the security log profile global-network configuration. Note that the statistics counter is always incremented when a packet matches a rule.
	Log string `json:"log,omitempty"`

	// Specifies a list of rules to evaluate. See security firewall rule-list. If a rule-list is specified then only the schedule and status properties effect the rule.
	RuleList string `json:"ruleList,omitempty"`

	// Specifies the action that the system takes when a rule is matched.
	Action string `json:"action,omitempty"`

	// Specifies the IP protocol against which the packet will be compared.
	IpProtocol string `json:"ipProtocol,omitempty"`

	// Specifies that a new rule should be placed before another rule, first or last. If individual rules are being added (as opposed to specifying replace-all-with) then place-before or place-after must be specified.
	PlaceBefore string `json:"placeBefore,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

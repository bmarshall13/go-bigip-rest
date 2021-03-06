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
type NetPacketFilter struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Enables or disables packet filter logging. If you omit this value, no logging is performed.
	Logging string `json:"logging,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the name of a bwc policy. The value for the bwc policy association is the name of any existing bwc policy. If omitted, no bwc filter is applied.
	BwcPolicy string `json:"bwcPolicy,omitempty"`

	// Specifies the VLAN to which the packet filter rule should apply. The value for this option is any VLAN name currently in existence. If you omit this value, the rule applies to all VLANs.
	Vlan string `json:"vlan,omitempty"`

	// Specifies the BPF expression to match. The filter is mandatory, however you can leave it empty. If empty, the packet filter rule matches all packets.
	Rule string `json:"rule,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the action that the packet filter rule should take. There is no default value; you must specify a value when you create a packet filter rule.
	Action string `json:"action,omitempty"`

	// Specifies a sort order. The values for the sort order are integers greater than 0 (zero). No two rules may have the same sort order.
	Order int64 `json:"order,omitempty"`

	// Specifies the name of a rate class. The value for the rate class association is the name of any existing rate class. If omitted, no rate filter is applied.
	RateClass string `json:"rateClass,omitempty"`
}

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
type NetBwcPolicy struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Policy Priority Traffic group to use as enforcement during congestion along with policy rate configured.
	TrafficPriorityMap string `json:"trafficPriorityMap,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Description for the policy.
	Description string `json:"description,omitempty"`

	// Specifies the Type of Service (ToS) level to use when outgoing packets cross threshold (Note: in such case, bwc just does only marking and not shaping). The default value is pass-through.
	IpTos string `json:"ipTos,omitempty"`

	// Specifies the logging publisher. A new log-publisher object can be created via TMSH command tmsh create sys log-config publisher.
	LogPublisher string `json:"logPublisher,omitempty"`

	// Displays the administrative partition within which the bwc policy resides.
	Partition string `json:"partition,omitempty"`

	// Policy type: If dynamic enabled or disabled. Policy type change modification is a disallowed configuration.
	Dynamic string `json:"dynamic,omitempty"`

	// Specifies the Quality of Service (QoS) level to use when outgoing packets cross threshold (Note: in such case, bwc just does only marking and not shaping). The default value is pass-through.
	LinkQos string `json:"linkQos,omitempty"`

	// Maximum rate for this policy class.
	MaxRate float32 `json:"maxRate,omitempty"`

	// Maximum user pps for this policy class.
	MaxUserRatePps float32 `json:"maxUserRatePps,omitempty"`

	// Enables/Disables bandwidth measurement for this policy. Available for only dynamic policies.
	Measure string `json:"measure,omitempty"`

	// Specifies the frequency for bwc measurement to log to publisher.
	LogPeriod int64 `json:"logPeriod,omitempty"`

	// Maximum per user rate for this policy class.
	MaxUserRate float32 `json:"maxUserRate,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}
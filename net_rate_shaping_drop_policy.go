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
type NetRateShapingDropPolicy struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the queue length above which packets are not dropped. The default value is 0.
	MinThreshold int64 `json:"minThreshold,omitempty"`

	// Specifies the average MTU (maximum transmission unit) size in the range of 0 to 10000 bytes. The default value is 0.
	AveragePacketSize int64 `json:"averagePacketSize,omitempty"`

	// Specifies the hard drop limit in the range of 0 to 400. The default value is 0. Setting this to a small value does not change the hard drop limit, but a higher number increases the limit.
	FredMaxDrop int64 `json:"fredMaxDrop,omitempty"`

	// Specifies the maximum percentage probability in the range of 0 to 100 according to which packets are dropped when the average queue length is between the minimum and maximum thresholds. The default value is 0.
	MaxProbability int64 `json:"maxProbability,omitempty"`

	// Specifies the maximum queue size in kilobytes or megabytes. Additional packets are dropped. The default value is 0. This option applies only to the red type.
	RedHardLimit int64 `json:"redHardLimit,omitempty"`

	// Specifies the maximum number of flows that can be active for each queue. The range is 0 to 10000. The default value is 0, which disables active flow limitation.
	FredMaxActive int64 `json:"fredMaxActive,omitempty"`

	// Specifies the weight used to calculate the average queue length. Valid values are 0, 64, 128, 256, 512, and 1024. The default value is 0.
	InverseWeight int64 `json:"inverseWeight,omitempty"`

	// Specifies the queue length below which packets are not dropped. The default value is 0.
	MaxThreshold int64 `json:"maxThreshold,omitempty"`

	// Specifies the type of drop policy. The available settings are tail (drops the end of the traffic stream), red (randomly drops packets), and fred (drops packets according to the type of traffic in the flow). The default value is red. Although you could create a drop policy based on tail, that is already the default value for drop policy in both the shaping policy and rate class commands.
	Type_ string `json:"type,omitempty"`

	// Specifies the hard no drop limit in the range of 0 to 100. The default value is 0. Setting this to a large value prevents packets from being dropped.
	FredMinDrop int64 `json:"fredMinDrop,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

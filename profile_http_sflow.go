/* 
 * BigIP iControl REST
 *
 * REST API for F5 BigIP. Only LTM is included, and most endpoints are untested.
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

type ProfileHttpSflow struct {

	// Specifies the ratio of packets observed to the samples generated. For example, a sampling rate of 2000 specifies that 1 sample will be randomly generated for every 2000 packets observed. To enable this setting, you must also set the sampling-rate-global setting to no.
	SamplingRate int64 `json:"samplingRate,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies whether the global HTTP sampling-rate setting overrides the object-level sampling-rate setting. The default value is yes.
	SamplingRateGlobal string `json:"samplingRateGlobal,omitempty"`

	// Specifies the maximum interval in seconds between two pollings. To enable this setting, you must also set the poll-interval-global setting to no.
	PollInterval int64 `json:"pollInterval,omitempty"`

	// Specifies whether the global HTTP poll-interval setting overrides the object-level poll-interval setting. The default value is yes.
	PollIntervalGlobal string `json:"pollIntervalGlobal,omitempty"`
}

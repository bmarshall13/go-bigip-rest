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
type SysDb struct {

	// The type of value that may be entered. integer (min and/or max value), unsigned integer (min and/or max value), string (min and/or max length), ip address, management address, or a list of valid values. This is a read-only setting.
	ValueRange string `json:"valueRange,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// The system supplied default value. This is read-only.
	DefaultValue string `json:"defaultValue,omitempty"`

	ScfConfig string `json:"scfConfig,omitempty"`

	// The current value of the property. It must conform to the value-range.
	Value string `json:"value,omitempty"`
}

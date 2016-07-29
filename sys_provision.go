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
type SysProvision struct {

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the level of resources that you want to provision for a module. \"minimum\" specifies that you want to provision the minimal amount of resources for the module you are provisioning. \"nominal\" specifies that you want to share all of the available resources equally among all of the modules that are licensed of the unit. \"dedicated\" specifies that all resources are dedicated to the module you are provisioning. For all other modules, the level option must be set to none. \"none\" specifies that you do not want to provision any resources for this module. \"custom\" F5 does not recommend that you specify this level.
	Level string `json:"level,omitempty"`

	// Use this option only when the level option is set to custom. F5 recommends that you do not modify this option. The default value is zero.
	MemoryRatio int64 `json:"memoryRatio,omitempty"`

	// Use this option only when the level option is set to custom. F5 recommends that you do not modify this option. The default value is zero.
	CpuRatio int64 `json:"cpuRatio,omitempty"`

	// Use this option only when the level option is set to custom. F5 recommends that you do not modify this option. The default value is zero.
	DiskRatio int64 `json:"diskRatio,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

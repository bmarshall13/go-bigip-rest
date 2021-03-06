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
type SysSoftware struct {

	// Status for software volumes
	Status string `json:"status,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Manage software images
	Image string `json:"image,omitempty"`

	// Manage block device software images
	BlockDeviceImage string `json:"blockDeviceImage,omitempty"`

	// Manage software update checking
	Update string `json:"update,omitempty"`

	// Manage software volumes, or reboot to a specific volume
	Volume string `json:"volume,omitempty"`

	// Manage software signatures
	Signature string `json:"signature,omitempty"`

	// Manage block device hotfix images
	BlockDeviceHotfix string `json:"blockDeviceHotfix,omitempty"`

	// Manage hotfix images
	Hotfix string `json:"hotfix,omitempty"`

	// Display update check results
	UpdateStatus string `json:"updateStatus,omitempty"`
}

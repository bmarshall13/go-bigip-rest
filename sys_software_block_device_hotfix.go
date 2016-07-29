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
type SysSoftwareBlockDeviceHotfix struct {
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Set to 'yes' if the hotfix is an authentic hotfix
	Verified string `json:"verified,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// A textual description of the hotfix.
	Title string `json:"title,omitempty"`

	// The checksum of the hotfix. Used to verify hotfix integrity.
	Checksum string `json:"checksum,omitempty"`

	// The F5 product this hotfix applies to.
	Product string `json:"product,omitempty"`

	// The provider/handler of this type of hotfix
	DeviceAgent string `json:"deviceAgent,omitempty"`

	Volume string `json:"volume,omitempty"`

	// The dotted version of the product this hotfix applies to
	Version string `json:"version,omitempty"`

	// The hotfix build number.
	Build string `json:"build,omitempty"`

	// The resource id of the hotfix used for global identification
	ResourceId string `json:"resourceId,omitempty"`

	// The sequential ID of the hotfix, e.g. HF1, HF2, HF3.
	Id string `json:"id,omitempty"`
}

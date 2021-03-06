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
type SysDiskLogicalDisk struct {

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the reserved logical disk space (MiB). (Reserved space can only be used by certain system utilities for base OS growth - specifically it will not be included in the disk space available to resource provisioning)
	VgReserved int64 `json:"vgReserved,omitempty"`

	// Specifies the usable free space (MiB) available in Logical Disk.
	VgFree int64 `json:"vgFree,omitempty"`

	// Specifies the size (MiB) of the Logical Disk.
	Size int64 `json:"size,omitempty"`

	// Specifies the total logical disk space (MiB) in use.
	VgInUse int64 `json:"vgInUse,omitempty"`

	// The mode is how the disk is used - it can be dedicated to a specific application or set for multiple uses (mixed mode).
	Mode string `json:"mode,omitempty"`
}

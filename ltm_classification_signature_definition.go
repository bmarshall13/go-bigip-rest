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
type LtmClassificationSignatureDefinition struct {

	// Indicates whether the last attempt to update the signature file was done manually or automatically by the system.
	LastAttemptAutomaticMode string `json:"lastAttemptAutomaticMode,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Indicates the user who did the last successful signature file update.
	LastUpdateUser string `json:"lastUpdateUser,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Indicates the date and time of the last successful signature file update.
	LastUpdateDatetime string `json:"lastUpdateDatetime,omitempty"`

	// Indicates the progress status when attempting to update the signature file.
	ProgressStatus string `json:"progressStatus,omitempty"`

	// Displays the administrative partition.
	Partition string `json:"partition,omitempty"`

	// Indicates the user who is the last one to attempt to update the signature file.
	LastAttemptUser string `json:"lastAttemptUser,omitempty"`

	// Indicates whether the last successful signature file update was done manually or automatically by the system.
	LastUpdateAutomaticMode string `json:"lastUpdateAutomaticMode,omitempty"`

	// Indicates the error message when it fails to attempt to update the signature file.
	Message string `json:"message,omitempty"`

	// Indicates the date and time of the last attempt to update the signature file.
	LastAttemptDatetime string `json:"lastAttemptDatetime,omitempty"`
}

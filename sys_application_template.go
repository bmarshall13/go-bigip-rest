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
type SysApplicationTemplate struct {

	// User defined description.
	Description string `json:"description,omitempty"`

	SigningKey string `json:"signingKey,omitempty"`

	TotalSigningStatus string `json:"totalSigningStatus,omitempty"`

	// Set to true to temporarily stop the verification of signature or checksum. Signature or checksum is still retained.
	IgnoreVerification string `json:"ignoreVerification,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Computes a checksum for the script.
	TmplChecksum string `json:"tmplChecksum,omitempty"`

	SigningAction string `json:"signingAction,omitempty"`

	// Displays the administrative partition within which the application template resides.
	Partition string `json:"partition,omitempty"`

	// Displays a list of the errors with the prerequisites.  If blank, the template is valid.
	PrerequisiteErrors string `json:"prerequisiteErrors,omitempty"`

	// Specifies the maximum version of BIG-IP required by this template.
	RequiresBigipVersionMax string `json:"requiresBigipVersionMax,omitempty"`

	// Adds, deletes, or replaces the list of modules that are required to be provisioned or enabled for this template to work.
	RequiresModules string `json:"requiresModules,omitempty"`

	// Specifies the minimum version of BIG-IP required by this template.
	RequiresBigipVersionMin string `json:"requiresBigipVersionMin,omitempty"`

	// Sign the script using the specified private key and corresponding public certificate.
	TmplSignature string `json:"tmplSignature,omitempty"`

	VerificationStatus string `json:"verificationStatus,omitempty"`

	PublicCert string `json:"publicCert,omitempty"`
}

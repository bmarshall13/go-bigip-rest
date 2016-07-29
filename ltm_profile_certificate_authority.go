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
type LtmProfileCertificateAuthority struct {

	// Displays the application service to which the object belongs. The default value is none.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the certificate revocation list file name. You can use default for the default certificate revocation file name.
	CrlFile string `json:"crlFile,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the authenticate depth. This is the client certificate chain maximum traversal depth.
	AuthenticateDepth int64 `json:"authenticateDepth,omitempty"`

	// Specifies the administrative partition within which the profile resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the name of the default certificate authority profile from which you want your custom profile to inherit settings. This setting is required.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	LocationSpecific string `json:"locationSpecific,omitempty"`

	// Automatically updates the CRL file.
	UpdateCrl string `json:"updateCrl,omitempty"`

	// Specifies the certificate authority file name or, you can use default for the default certificate authority file name. Configures certificate verification by specifying a list of client or server certificate authorities that the traffic management system trusts.
	CaFile string `json:"caFile,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

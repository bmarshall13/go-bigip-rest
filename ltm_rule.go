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
type LtmRule struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	Nowrite bool `json:"nowrite,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	Plugin string `json:"plugin,omitempty"`

	// The private key to use for signing the iRule.
	SigningKey string `json:"signingKey,omitempty"`

	// Displays the administrative partition within which this iRule resides.
	Partition string `json:"partition,omitempty"`

	DefinitionChecksum string `json:"definitionChecksum,omitempty"`

	DefinitionSignature string `json:"definitionSignature,omitempty"`

	IgnoreVerification string `json:"ignoreVerification,omitempty"`

	Action string `json:"action,omitempty"`

	Hidden bool `json:"hidden,omitempty"`

	Nodelete bool `json:"nodelete,omitempty"`

	// The public certificate corresponding to the signing-key.
	PublicCert string `json:"publicCert,omitempty"`
}

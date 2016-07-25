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

type ProfileClientSslCertKeyChain struct {

	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies or builds a certificate chain file that a client can use to authenticate the profile. To use the default chain name, specify default.
	Chain string `json:"chain,omitempty"`

	// Specifies the name of the certificate installed on the traffic management system for the purpose of terminating or initiating an SSL connection.
	Cert string `json:"cert,omitempty"`

	// Specifies the name of a key file that you generated and installed on the system. The default key name is default.key.
	Key string `json:"key,omitempty"`

	// Specifies the key passphrase if required.
	Passphrase string `json:"passphrase,omitempty"`

	// Specifies the OCSP Stapling Parameters object associated with this cert-key-chain object in a clientssl profile.
	OcspStaplingParams string `json:"ocspStaplingParams,omitempty"`
}

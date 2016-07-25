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

type DnsDnssecKey struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the length of time before the key expires. The default value is 0 (zero).  This value must be greater than the value of the rollover-period option.
	ExpirationPeriod string `json:"expirationPeriod,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the length of the key you want to generate. The default value is 1024. If the key is manually managed, MCPD will derive this value from the key file.
	BitWidth int64 `json:"bitWidth,omitempty"`

	// Specifies that the DNSSEC key is disabled.
	Disabled bool `json:"disabled,omitempty"`

	// Specifies the type of FIPS-compliant hardware security module to use when storing, and signing with, the private key. The default value is none.  Optionally external or internal.
	UseFips string `json:"useFips,omitempty"`

	// Specifies the number of seconds that a DNS server can cache the key. The default value is 86400.
	Ttl int64 `json:"ttl,omitempty"`

	// Specifies the file containing the private key.  Fields certificate-file and key-file are required for manual DNSSEC key import.
	KeyFile string `json:"keyFile,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the length of time (seconds) that the signature of the key is valid.  The default value is 604800.
	SignatureValidPeriod string `json:"signatureValidPeriod,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the algorithm to use to generate the key.  The default value is RSASHA1.
	Algorithm string `json:"algorithm,omitempty"`

	// Specifies the file containing the public key.  Fields certificate-file and key-file are required for manual DNSSEC key import.
	CertificateFile string `json:"certificateFile,omitempty"`

	// Specifies whether the key is of type KSK or ZSK.  The default value is ZSK.
	KeyType string `json:"keyType,omitempty"`

	// Specifies that the DNSSEC key is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// Specifies the length of time before we create a new signature.  The default value is 403200.
	SignaturePubPeriod string `json:"signaturePubPeriod,omitempty"`

	// Specifies the length of time before the key changes to another key. The default value is 0 (zero).  This value must be less than the value of the expiration-period option.
	RolloverPeriod string `json:"rolloverPeriod,omitempty"`
}

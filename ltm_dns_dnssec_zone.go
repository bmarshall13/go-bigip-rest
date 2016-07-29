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
type LtmDnsDnssecZone struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Secure Entry Point(s) (DS and DNSKEY resource records used as client trust anchors) of the zone.
	Seps string `json:"seps,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the number of times to hash the Next Secure (NSEC3) names. The default value is 1.
	Nsec3Iterations int64 `json:"nsec3Iterations,omitempty"`

	// Specifies that the DNSSEC zone will be signed.
	Enabled bool `json:"enabled,omitempty"`

	// The learned zone SOA serial number from the primary server.
	XfrPrimarySoaSerial float32 `json:"xfrPrimarySoaSerial,omitempty"`

	// Specifies that the DNSSEC zone will not be signed.
	Disabled bool `json:"disabled,omitempty"`

	// Specifies the hash algorithm to use when creating the Next Secure (NSEC3) resource record. The default value is SHA1.
	Nsec3Algorithm string `json:"nsec3Algorithm,omitempty"`

	// The advertised zone SOA serial number to all clients.
	XfrSoaSerial float32 `json:"xfrSoaSerial,omitempty"`

	// Specifies the hash algorithm to use when creating the Delegation Signer (DS) resource record. The default value is SHA256
	DsAlgorithm string `json:"dsAlgorithm,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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
type NetIpsecManualSecurityAssociation struct {
	AppService string `json:"appService,omitempty"`

	// Displays the encrypted key for the authentication algorithm.
	AuthKeyEncrypted string `json:"authKeyEncrypted,omitempty"`

	// Specifies the destination of the security association.
	DestinationAddress string `json:"destinationAddress,omitempty"`

	// Displays the encrypted key for the encryption algorithm.
	EncryptKeyEncrypted string `json:"encryptKeyEncrypted,omitempty"`

	// Specifies an authentication algorithm.
	AuthAlgorithm string `json:"authAlgorithm,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies an encryption algorithm.
	EncryptAlgorithm string `json:"encryptAlgorithm,omitempty"`

	// Specifies the ipsec-policy associated with this manual-security-association.
	IpsecPolicy string `json:"ipsecPolicy,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the key for the authentication algorithm.
	AuthKey string `json:"authKey,omitempty"`

	// Specifies the Security Parameters Index. If this is the Security Association(SA) for the outbound traffic, make sure it matches the SPI of the inbound SA configured on the remote site and vice versa. SPI values between 0 and 255 are reserved for the future use by IANA and cannot be used.
	Spi int64 `json:"spi,omitempty"`

	// Specifies the IPsec protocol: Encapsulating Security Payload (ESP) or Authentication Header (AH).
	Protocol string `json:"protocol,omitempty"`

	// Specifies the key for the encryption algorithm.
	EncryptKey string `json:"encryptKey,omitempty"`

	// Specifies the source address of the security association.
	SourceAddress string `json:"sourceAddress,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}
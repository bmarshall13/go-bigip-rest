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
type NetIpsecIpsecSa struct {

	// Specifies the source address of the security associations
	SrcAddr string `json:"srcAddr,omitempty"`

	// Specifies the SPI of the security associations
	Spi int64 `json:"spi,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the name of the traffic selector
	TrafficSelector string `json:"trafficSelector,omitempty"`

	// Specifies route domain used for traffic. The default value is the default route domain.
	RouteDomain int64 `json:"routeDomain,omitempty"`

	// Specifies the destination address of the security associations
	DstAddr string `json:"dstAddr,omitempty"`
}
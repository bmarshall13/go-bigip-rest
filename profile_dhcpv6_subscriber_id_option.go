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

type ProfileDhcpv6SubscriberIdOption struct {

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies if the user wants the DHCP relay agent to insert option 38 or not. Default is false.
	Add string `json:"add,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// A string to represent the subscriber-id option value.
	Value string `json:"value,omitempty"`

	// Specifies if the user wants the DHCP relay agent to remove option 38 or not. Default is false.
	Remove string `json:"remove,omitempty"`
}
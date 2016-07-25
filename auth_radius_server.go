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

type AuthRadiusServer struct {

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Displays the partition within which the server resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the host name or IP address of the RADIUS server. This option is required.
	Server string `json:"server,omitempty"`

	// Sets the secret key used to encrypt and decrypt packets sent or received from the server. This option is required.
	Secret string `json:"secret,omitempty"`

	// Specifies the timeout value in seconds. The default value is 3 seconds.
	Timeout int64 `json:"timeout,omitempty"`

	// Specifies the port for RADIUS authentication traffic. The default value is port 1812.
	Port int64 `json:"port,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}

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
type LtmProfileMblb struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Delays sending FIN when BIGIP receives the first FIN packet from either the client or the server. Value of 0 means send FIN immediately otherwise the minimum of tcp idle timeout and shutdown timeout is used. The default value is 5 seconds
	ShutdownTimeout int64 `json:"shutdownTimeout,omitempty"`

	// whether to isolate abort event propagation
	IsolateAbort string `json:"isolateAbort,omitempty"`

	Description string `json:"description,omitempty"`

	// specifies the TTL (time to live) for TAGs
	TagTtl int64 `json:"tagTtl,omitempty"`

	// specifies the low water mark for ingress message queue
	IngressLow int64 `json:"ingressLow,omitempty"`

	// whether to isolate clientside shutdown event propagation
	IsolateClient string `json:"isolateClient,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// specifies the high water mark for ingress message queue
	IngressHigh int64 `json:"ingressHigh,omitempty"`

	// specifies the low water mark for egress message queue
	EgressLow int64 `json:"egressLow,omitempty"`

	// whether to isolate serverside shutdown event propagation
	IsolateServer string `json:"isolateServer,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// specifies the high water mark for egress message queue
	EgressHigh int64 `json:"egressHigh,omitempty"`

	// specifies the minimum number of serverside connections
	MinConn int64 `json:"minConn,omitempty"`

	// whether to isolate expire event propagation
	IsolateExpire string `json:"isolateExpire,omitempty"`
}

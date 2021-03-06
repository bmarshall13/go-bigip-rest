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
type LtmLsnPoolPortBlockAllocation struct {

	// Configures the number of blocks that can be assigned to a single subscriber IP address. The default value is 1.
	ClientBlockLimit int64 `json:"clientBlockLimit,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Configures the timeout after which the block is no longer used for new port allocations. The block becomes a zombie block. The default is 0 which corresponds to an infinite timeout.
	BlockLifetime int64 `json:"blockLifetime,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Configures the number of ports in a block. The default value is 64.
	BlockSize int64 `json:"blockSize,omitempty"`

	// Configures the timeout after which connections using the zombie block are killed. After connections are killed zombie block is freed after port-block-allocation.block-idle-timeout. This parameter is unused unless the port-block-allocation.block-lifetime is set. The default value is 0 which corresponds to infinite timeout.
	ZombieTimeout int64 `json:"zombieTimeout,omitempty"`

	// Configures the time after the last connection using the block is freed that the block assignment expires. The default value is 3600 seconds.
	BlockIdleTimeout int64 `json:"blockIdleTimeout,omitempty"`
}

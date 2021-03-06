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
type SysLtcfgClass struct {
	AppService string `json:"appService,omitempty"`

	Singleton string `json:"singleton,omitempty"`

	ManPage string `json:"manPage,omitempty"`

	GuiHints string `json:"guiHints,omitempty"`

	Clustered string `json:"clustered,omitempty"`

	CliHints string `json:"cliHints,omitempty"`

	FullPath string `json:"fullPath,omitempty"`

	IcontrolHints string `json:"icontrolHints,omitempty"`

	DefaultModel string `json:"defaultModel,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	Category string `json:"category,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	MinRole string `json:"minRole,omitempty"`

	Stats string `json:"stats,omitempty"`

	FieldCnt int64 `json:"fieldCnt,omitempty"`

	NestedClasses string `json:"nestedClasses,omitempty"`

	Partition string `json:"partition,omitempty"`

	ConfigLevel string `json:"configLevel,omitempty"`

	ToplevelCmd string `json:"toplevelCmd,omitempty"`

	Configsyncd string `json:"configsyncd,omitempty"`

	Partitioned string `json:"partitioned,omitempty"`

	Hidden string `json:"hidden,omitempty"`

	Constraints string `json:"constraints,omitempty"`
}

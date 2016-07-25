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

type ProfileXml struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Enables or disables multiple matches for a single XPath query.
	MultipleQueryMatches string `json:"multipleQueryMatches,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Displays the administrative partition within which the profile resides.
	Partition string `json:"partition,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies a list of mappings between namespaces and prefixes to be used in the profile context.
	NamespaceMappings string `json:"namespaceMappings,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the list of XPath queries that are used by the profile. A match of any of the queries will trigger an iRule event.
	XpathQueries string `json:"xpathQueries,omitempty"`
}

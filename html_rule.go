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

type HtmlRule struct {

	TagAppendHtml string `json:"tagAppendHtml,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	TagPrependHtml string `json:"tagPrependHtml,omitempty"`

	TagRemoveAttribute string `json:"tagRemoveAttribute,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	TagRemove string `json:"tagRemove,omitempty"`

	CommentRemove string `json:"commentRemove,omitempty"`

	CommentRaiseEvent string `json:"commentRaiseEvent,omitempty"`

	TagRaiseEvent string `json:"tagRaiseEvent,omitempty"`
}
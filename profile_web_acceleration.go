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

type ProfileWebAcceleration struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies the maximum number of entries that can be in the cache. The default value is 0 (zero), which means that the system does not limit the maximum entries.
	CacheMaxEntries int64 `json:"cacheMaxEntries,omitempty"`

	// Specifies the largest object that the system considers eligible for caching. The default value is 50000 bytes.
	CacheObjectMaxSize int64 `json:"cacheObjectMaxSize,omitempty"`

	// Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is 9.
	CacheAgingRate int64 `json:"cacheAgingRate,omitempty"`

	// Specifies the maximum size, in megabytes, of the metadata cache. This does not include the content cache, and only applies when a WebAccelerator application is on the profile.
	MetadataCacheMaxSize int64 `json:"metadataCacheMaxSize,omitempty"`

	// Specifies how long the system considers the cached content to be valid. The default value is 3600 seconds.
	CacheMaxAge int64 `json:"cacheMaxAge,omitempty"`

	// Configures a list of URIs to exclude from the cache. The default value of none specifies no URIs are excluded.
	CacheUriExclude string `json:"cacheUriExclude,omitempty"`

	// Configures a list of applications assigned to this profile. Assigning at least one application enables WA functionality. The default value of none specifies that WA is not enabled.
	Applications string `json:"applications,omitempty"`

	// Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs.
	CacheUriPinned string `json:"cacheUriPinned,omitempty"`

	// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies which cache disabling headers sent by clients the system ignores. The default value is all.
	CacheClientCacheControlMode string `json:"cacheClientCacheControlMode,omitempty"`

	// Specifies the smallest object that the system considers eligible for caching. The default value is 500 bytes.
	CacheObjectMinSize int64 `json:"cacheObjectMinSize,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache.
	CacheUriIncludeOverride string `json:"cacheUriIncludeOverride,omitempty"`

	// Displays the administrative partition within which the profile resides.
	Partition string `json:"partition,omitempty"`

	// Inserts Age and Date headers in the response. The default value is enabled.
	CacheInsertAgeHeader string `json:"cacheInsertAgeHeader,omitempty"`

	// Configures a list of URIs to include in the cache. The default value of .* specifies that all URIs are cacheable.
	CacheUriInclude string `json:"cacheUriInclude,omitempty"`

	// Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is 100 megabytes.
	CacheSize int64 `json:"cacheSize,omitempty"`
}

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
type LtmProfileHttpCompression struct {

	// The application service to which the object belongs.
	AppService string `json:"appService,omitempty"`

	// Specifies amount of memory (in kilobytes) that the system uses when compressing a server response. The value will be rounded up to the nearest power of 2. The default value is 8k. The maximum value is 256k.
	GzipMemoryLevel int64 `json:"gzipMemoryLevel,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeInclude string `json:"contentTypeInclude,omitempty"`

	// Enables or disables browser workarounds. The default value is disabled.
	BrowserWorkarounds string `json:"browserWorkarounds,omitempty"`

	// Enables or disables selective compression mode.
	Selective string `json:"selective,omitempty"`

	// Specifies the profile that you want to use as the parent profile.
	DefaultsFrom string `json:"defaultsFrom,omitempty"`

	// Specifies the maximum number of uncompressed bytes that the system buffers before determining whether or not to compress the response. Useful when the headers of a server response do not specify the length of the response content. The default value is 4096.
	BufferSize int64 `json:"bufferSize,omitempty"`

	// Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The value will be rounded up to the nearest power of 2. The default is 16k. The maximum value is 128k.
	GzipWindowSize int64 `json:"gzipWindowSize,omitempty"`

	// Specifies the percent CPU usage at which the system resumes content compression at the user-defined rates. The default value is 75 percent.
	CpuSaverLow int64 `json:"cpuSaverLow,omitempty"`

	// Specifies the type of compression that is preferred by the system.
	MethodPrefer string `json:"methodPrefer,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
	UriInclude string `json:"uriInclude,omitempty"`

	// Enables or disables the insertion of a Vary header into cacheable server responses. The default value is enabled.
	VaryHeader string `json:"varyHeader,omitempty"`

	// Specifies a value that determines the amount of memory that the system uses when compressing a server response. The default is 1.
	GzipLevel int64 `json:"gzipLevel,omitempty"`

	// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
	ContentTypeExclude string `json:"contentTypeExclude,omitempty"`

	// Displays the administrative partition within which this profile resides.
	Partition string `json:"partition,omitempty"`

	// Specifies the percent of CPU usage at which the system starts automatically decreasing the amount of content being compressed, as well as the amount of compression which the system is applying. The default value is 90 percent.
	CpuSaverHigh int64 `json:"cpuSaverHigh,omitempty"`

	// Specifies the CPU saver setting. When the CPU saver is enabled, the system monitors the percent of CPU usage and adjusts compression rates automatically when the CPU usage reaches the percentage defined in the cpu saver low or the cpu saver high options. The default value is enabled.
	CpuSaver string `json:"cpuSaver,omitempty"`

	// Specifies the minimum length in bytes of a server response that is acceptable for compressing that response. The length in bytes applies to content length only, not headers. The default value is 1024.
	MinSize int64 `json:"minSize,omitempty"`

	// Enables or disables keep accept encoding. When enabled, causes the target server, rather than the BIG-IP, to perform the data compression.
	KeepAcceptEncoding string `json:"keepAcceptEncoding,omitempty"`

	// Enables or disables compression of HTTP/1.0 server responses.
	AllowHttp10 string `json:"allowHttp_10,omitempty"`

	// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
	UriExclude string `json:"uriExclude,omitempty"`
}

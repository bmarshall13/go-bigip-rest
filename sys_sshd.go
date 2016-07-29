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
type SysSshd struct {

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the minimum sshd message level to include in the system log. You must enter the following values: -- debug, debug1, debug2, debug3, which indicates that the minimum sshd message level that the system logs is the specified debugging level of messages. -- error, which indicates that the minimum sshd message level that the system logs is error. -- fatal, which indicates that the minimum sshd message level that the system logs is fatal. -- info, which indicates that the minimum sshd message level that the system logs is informational. -- quiet, which indicates that the system does not log sshd messages.-- verbose, which indicates that the system logs all sshd messages. The default is info.
	LogLevel string `json:"logLevel,omitempty"`

	// Specifies the number of seconds before inactivity causes an SSH session to log out. The default value is 0 (zero) seconds, which indicates that inactivity timeout is disabled.
	InactivityTimeout int64 `json:"inactivityTimeout,omitempty"`

	// When banner is enabled, specifies the text to include in the banner that displays when a user attempts to login to the system.
	BannerText string `json:"bannerText,omitempty"`

	// Adds a server to or removes a server from the /etc/hosts.allow file. Use this option to either add servers to the BIG-IP system that are allowed to access the system, or delete these servers from the system. Specify \"none\" to disallow ssh access to the system. Specify \"replace-all-with { ALL }\" to allow ssh access from any server. The default value is \"replace-all-with { ALL }\".
	Allow string `json:"allow,omitempty"`

	// Enables or disables SSH logins to the system. The default is enabled.
	Login string `json:"login,omitempty"`

	// Warning: Do not use this parameter without assistance from the F5 Technical Support team. The system does not validate the commands issued using the include parameter. If you use this parameter incorrectly, you put the functionality of the system at risk.
	Include string `json:"include,omitempty"`

	// Enables or disables the display of the banner text field when a user logs in to the system using SSH. The default value is disabled.
	Banner string `json:"banner,omitempty"`

	// Specifies the TCP port to run SSHD
	Port int64 `json:"port,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`
}
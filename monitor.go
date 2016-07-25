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

type Monitor struct {

	// Diameter monitor configuration
	Diameter string `json:"diameter,omitempty"`

	// SOAP monitor configuration.
	Soap string `json:"soap,omitempty"`

	// MySQL monitor configuration.
	Mysql string `json:"mysql,omitempty"`

	// SMTP monitor configuration.
	Smtp string `json:"smtp,omitempty"`

	// TCP monitor configuration.
	Tcp string `json:"tcp,omitempty"`

	// RPC monitor configuration.
	Rpc string `json:"rpc,omitempty"`

	// Radius monitor configuration.
	Radius string `json:"radius,omitempty"`

	// Firepass monitor configuration
	Firepass string `json:"firepass,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// FTP monitor configuration.
	Ftp string `json:"ftp,omitempty"`

	// MSSQL monitor configuration.
	Mssql string `json:"mssql,omitempty"`

	// SIP monitor configuration.
	Sip string `json:"sip,omitempty"`

	// IMAP monitor configuration.
	Imap string `json:"imap,omitempty"`

	// WMI monitor configuration.
	Wmi string `json:"wmi,omitempty"`

	// HTTPS monitor configuration.
	Https string `json:"https,omitempty"`

	// LDAP monitor configuration.
	Ldap string `json:"ldap,omitempty"`

	// SNMP monitor configuration.
	SnmpDcaBase string `json:"snmpDcaBase,omitempty"`

	// NNTP monitor configuration.
	Nntp string `json:"nntp,omitempty"`

	// SMB monitor configuration.
	Smb string `json:"smb,omitempty"`

	// Scripted monitor configuration.
	Scripted string `json:"scripted,omitempty"`

	// WAP monitor configuration.
	Wap string `json:"wap,omitempty"`

	// UDP monitor configuration.
	Udp string `json:"udp,omitempty"`

	// Virtual Location monitor configuration
	VirtualLocation string `json:"virtualLocation,omitempty"`

	// HTTP monitor configuration.
	Http string `json:"http,omitempty"`

	// DNS monitor configuration.
	Dns string `json:"dns,omitempty"`

	// Inband monitor configuration.
	Inband string `json:"inband,omitempty"`

	// Radius accounting monitor configuration.
	RadiusAccounting string `json:"radiusAccounting,omitempty"`

	// POP3 monitor configuration.
	Pop3 string `json:"pop3,omitempty"`

	// SNMP DCA monitor configuration.
	SnmpDca string `json:"snmpDca,omitempty"`

	// External monitor configuration.
	External string `json:"external,omitempty"`

	// Gateway ICMP monitor configuration.
	GatewayIcmp string `json:"gatewayIcmp,omitempty"`

	// ICMP monitor configuration.
	Icmp string `json:"icmp,omitempty"`

	// TCP Echo monitor configuration.
	TcpEcho string `json:"tcpEcho,omitempty"`

	// Real Server monitor configuration.
	RealServer string `json:"realServer,omitempty"`

	// The NULL monitor.
	None string `json:"none,omitempty"`

	// Postgresql monitor configuration.
	Postgresql string `json:"postgresql,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// SASP monitor configuration.
	Sasp string `json:"sasp,omitempty"`

	// TCP Half Open monitor configuration.
	TcpHalfOpen string `json:"tcpHalfOpen,omitempty"`

	// Module score monitor configuration.
	ModuleScore string `json:"moduleScore,omitempty"`

	// Oracle monitor configuration.
	Oracle string `json:"oracle,omitempty"`
}

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

type AuthSslCcLdap struct {

	// Specifies whether the system uses the client certificate's subject or serial number (in conjunction with the certificate's issuer) when trying to match an entry in the certificate map subtree. A value of yes uses the serial number. A value of no uses the subject. The default value is no.
	CertmapUserSerial string `json:"certmapUserSerial,omitempty"`

	// The application service that the object belongs to.
	AppService string `json:"appService,omitempty"`

	// Specifies the search base for the subtree used by the user and cert search methods. A typical search base is: ou=people,dc=company,dc=com. Possible values are a user-specified string, and none. You must specify a user base when you create an SSL client certificate configuration object.
	UserBase string `json:"userBase,omitempty"`

	// Displays the partition within which the server resides.
	Partition string `json:"partition,omitempty"`

	// Enables or disables an attempt to use secure LDAP (LDAP over SSL). The alternative to using secure LDAP is to use insecure (clear text) LDAP. Secure LDAP is a consideration when the connection between the BIG-IP system and the LDAP server cannot be trusted. The default value is disabled.
	Secure string `json:"secure,omitempty"`

	// Specifies the password for the admin account. See the admin dn option above. Possible values are a user-specified string, and none.
	AdminPassword string `json:"adminPassword,omitempty"`

	// Specifies the search base for the subtree used by the certmap search method. A typical search base is: ou=people,dc=company,dc=com. Possible values are a user-specified string, and none.
	CertmapBase string `json:"certmapBase,omitempty"`

	// Specifies a space-delimited list specifying the names of groups that the client must belong to in order to be authorized (matches against the group key in the group subtree). The client needs to be a member of only one of the groups in the list. Possible values are a user-specified string, or none.
	ValidGroups string `json:"validGroups,omitempty"`

	// Specifies a list of LDAP servers you want to search. Possible values are a user-specified list of servers, and none. You must specify a server when you create an SSL client certificate configuration object.
	Servers string `json:"servers,omitempty"`

	// Specifies the number of usable lifetime seconds of negotiable SSL session IDs. When this time expires, a client must negotiate a new session. Allowed values are:  number , immediate, and indefinite. The default value is 300 seconds.
	CacheTimeout int64 `json:"cacheTimeout,omitempty"`

	// Specifies the key that denotes a user ID in the LDAP database (for example, the common key for the user option is uid). Possible values are a user-specified string, and none. You must always specify a user key when you create an SSL client certificate configuration object.
	UserKey string `json:"userKey,omitempty"`

	// Kind of entity
	Kind string `json:"kind,omitempty"`

	// Specifies the object class in the LDAP database to which the user must belong in order to be authenticated.
	UserClass string `json:"userClass,omitempty"`

	// User defined description.
	Description string `json:"description,omitempty"`

	// Specifies the distinguished name of an account to which to bind, in order to perform searches. This search account is a read-only account used to do searches. The admin account can also be used as the search account. If no admin DN is specified, then no bind is attempted. This parameter is required only when an LDAP database does not allow anonymous searches. Possible values are a user-specified string, and none.
	AdminDn string `json:"adminDn,omitempty"`

	// Specifies the name of the attribute in the LDAP database that specifies a user's authorization roles. This key is used only with the valid roles option. A typical role key might be authorizationRole. Possible values are a user-specified string, and none.
	RoleKey string `json:"roleKey,omitempty"`

	// Specifies the name of the certificate map found in the LDAP database. Used by the certmap search method. Possible values are a user-specified string, and none.
	CertmapKey string `json:"certmapKey,omitempty"`

	// Name of entity
	Name string `json:"name,omitempty"`

	// Specifies the name of the attribute in the LDAP database that specifies the group name in the group subtree. An example of a typical key is cn (common name for the group). Possible values are a user-specified string, and none.
	GroupKey string `json:"groupKey,omitempty"`

	// Specifies the type of LDAP search that is performed based on the client's certificate.
	SearchType string `json:"searchType,omitempty"`

	// Specifies the search base for the subtree used by group searches. This parameter is only used when specifying the valid groups option. The typical search base is similar to: ou=groups,dc=company,dc=com. Possible values are a user-specified string, and none.
	GroupBase string `json:"groupBase,omitempty"`

	// Specifies the name of the attribute in the LDAP database that specifies members (DNs) of a group. A typical key would be member. Possible values are a user-specified string, and none.
	GroupMemberKey string `json:"groupMemberKey,omitempty"`

	// Specifies the maximum size, in bytes, allowed for the SSL session cache. Setting this value to 0 disallows SSL session caching. The default value is 20000 bytes (that is 20KB).
	CacheSize int64 `json:"cacheSize,omitempty"`

	// Specifies a space-delimited list specifying the valid roles that clients must have in order to be authorized. Possible values are a user-specified string, and none.
	ValidRoles string `json:"validRoles,omitempty"`
}

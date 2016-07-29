# LtmPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinUpMembersChecking** | **string** | Enables or disables the min-up-members feature. If you enable this feature, you must also specify a value for both the min-up-members and min-up-members-action options. | [optional] [default to null]
**AppService** | **string** | The application service to which the object belongs. | [optional] [default to null]
**LinkQosToServer** | **string** | Specifies the Quality of Service (QoS) level to use when sending packets to a server. The default value is 65535 (pass-through). | [optional] [default to null]
**Description** | **string** | User defined description. | [optional] [default to null]
**QueueTimeLimit** | **int64** | Specifies the maximum time, in milliseconds, a connection will remain enqueued. The default is zero which indicates there is no limit. | [optional] [default to 0]
**MinActiveMembers** | **int64** | Specifies the minimum number of members that must be up for traffic to be confined to a priority group when using priority-based activation. The default value is 0 (zero). An active member is a member that is up (not marked down) and is handling fewer connections than its connection limit. | [optional] [default to 0]
**IpTosToServer** | **string** | Specifies the Type of Service (ToS) level to use when sending packets to a server. 65534 (mimic) specifies that the system sets the ToS level of outgoing packets to the same ToS level of the most-recently received incoming packet. The default value is 65535 (pass-through). | [optional] [default to null]
**LoadBalancingMode** | **string** | Specifies the modes that the system uses to load balance name resolution requests among the members of this pool. See \&quot;help pool\&quot; for a description of each loading balancing mode. | [optional] [default to null]
**GatewayFailsafeDevice** | **string** | Specifies that the pool is a gateway failsafe pool in a redundant configuration. This string identifies the device that will failover when the monitor reports the pool member down. By default the device string is empty. | [optional] [default to null]
**AllowNat** | **string** | Specifies whether the pool can load balance NAT connections. The default value is yes. | [optional] [default to null]
**AutoscaleGroupId** | **string** | autoscale-group id to which pool members belong to. | [optional] [default to null]
**QueueDepthLimit** | **int64** | Specifies the maximum number of connections that may simultaneously be queued to go to any member of this pool. The default is zero which indicates there is no limit. | [optional] [default to 0]
**QueueOnConnectionLimit** | **string** | Enable or disable queuing connections when pool member or node connection limits are reached. When queuing is not enabled, new connections are reset when connection limits are met. | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]
**MinUpMembers** | **int64** | Specifies the minimum number of pool members that must be up; otherwise, the system takes the action specified in the min-up-members-action option. Use this option for gateway pools in a redundant system where a unit number is applied to a pool. This indicates that the pool is only configured on the specified unit. | [optional] [default to 0]
**IpTosToClient** | **string** | Specifies the Type of Service (ToS) level to use when sending packets to a client. 65534 (mimic) specifies that the system sets the ToS level of outgoing packets to the same ToS level of the most-recently received incoming packet. The default value is 65535 (pass-through). | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Monitor** | **string** | Specifies the health monitors that the system uses to determine whether it can use this pool for load balancing. The monitor marks the pool up or down based on whether the monitor(s) are successful. You can specify a single monitor, multiple monitors \&quot;http and https\&quot;, or a \&quot;min\&quot; rule, \&quot;min 1 of { http https }\&quot;. You may remove the monitor by specifying \&quot;none\&quot;. | [optional] [default to null]
**Partition** | **string** | Displays the administrative partition within which the pool resides. | [optional] [default to null]
**SlowRampTime** | **int64** | Specifies, in seconds, the ramp time for the pool. This provides the ability to cause a pool member that has just been enabled, or marked up, to receive proportionally less traffic than other members in the pool. The proportion of traffic the member accepts is determined by how long the member has been up in comparison to the slow-ramp-time setting for the pool.For example, if the load-balancing-mode of a pool is round-robin, and it has a slow-ramp-time of 60 seconds, when a pool member has been up for only 30 seconds, the pool member receives approximately half the amount of new traffic as other pool members that have been up for more than 60 seconds. After the pool member has been up for 45 seconds, it receives approximately three quarters of the new traffic.The slow ramp time is particularly useful when used with the least-connections-member load balancing mode. The default value is 10. | [optional] [default to 10]
**AllowSnat** | **string** | Specifies whether the pool can load balance SNAT connections. The default value is yes. | [optional] [default to null]
**MinUpMembersAction** | **string** | Specifies the action to take if the min-up-members-checking is enabled and the number of active pool members falls below the number specified in min-up-members. The default value is failover. | [optional] [default to null]
**ReselectTries** | **int64** | Specifies the number of times the system tries to contact a pool member after a passive failure. A passive failure consists of a server-connect failure or a failure to receive a data response within a user-specified interval. The default is 0 (zero), which indicates no reselect attempts. | [optional] [default to 0]
**ServiceDownAction** | **string** | Specifies the action to take if the service specified in the pool is marked down. The default value is none. | [optional] [default to null]
**IgnorePersistedWeight** | **string** | Do not count the weight of persisted connections on pool members when making load balancing decisions. | [optional] [default to null]
**LinkQosToClient** | **string** | Specifies the Quality of Service (QoS) level to use when sending packets to a client. The default value is 65535 (pass-through). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LtmLsnPoolPortBlockAllocation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientBlockLimit** | **int64** | Configures the number of blocks that can be assigned to a single subscriber IP address. The default value is 1. | [optional] [default to 1]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**BlockLifetime** | **int64** | Configures the timeout after which the block is no longer used for new port allocations. The block becomes a zombie block. The default is 0 which corresponds to an infinite timeout. | [optional] [default to 0]
**Name** | **string** | Name of entity | [optional] [default to null]
**BlockSize** | **int64** | Configures the number of ports in a block. The default value is 64. | [optional] [default to 64]
**ZombieTimeout** | **int64** | Configures the timeout after which connections using the zombie block are killed. After connections are killed zombie block is freed after port-block-allocation.block-idle-timeout. This parameter is unused unless the port-block-allocation.block-lifetime is set. The default value is 0 which corresponds to infinite timeout. | [optional] [default to 0]
**BlockIdleTimeout** | **int64** | Configures the time after the last connection using the block is freed that the block assignment expires. The default value is 3600 seconds. | [optional] [default to 3600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



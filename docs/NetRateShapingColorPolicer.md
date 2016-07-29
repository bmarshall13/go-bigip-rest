# NetRateShapingColorPolicer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | **string** |  | [optional] [default to null]
**CommittedInformationRate** | **float32** | The committed information rate (CIR) that is given to an entity being metered. | [optional] [default to null]
**ExcessBurstSize** | **float32** | An additional burst size allowed beyond committed-information-rate and committed-burst-size. Packets which are part of a burst which is larger than the CBS but less than CBS + EBS are counted as Yellow, but those which are part of a burst that exceeds this total are counted as Red. | [optional] [default to null]
**Action** | **string** | The action to be applied to out of profile (Red) traffic. | [optional] [default to null]
**CommittedBurstSize** | **float32** | The committed burst size (CBS) above the committed-information-rate. Packets which fall within the CIR plus such a burst are counted as Green. | [optional] [default to null]
**Kind** | **string** | Kind of entity | [optional] [default to null]
**Name** | **string** | Name of entity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



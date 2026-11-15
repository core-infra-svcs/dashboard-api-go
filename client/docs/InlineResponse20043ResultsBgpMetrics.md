# InlineResponse20043ResultsBgpMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPreference** | Pointer to **int32** | Value reported by the device to choose preferred exit paths within an autonomous system. Higher values are preferred. | [optional] 
**MultiExitDiscriminator** | Pointer to **int32** | Value reported by the device to influence inbound traffic from neighboring autonomous systems. Lower values are preferred. | [optional] 
**Weight** | Pointer to **int32** | Cisco-proprietary value reported by the device to choose a local route path. Higher values are preferred. | [optional] 

## Methods

### NewInlineResponse20043ResultsBgpMetrics

`func NewInlineResponse20043ResultsBgpMetrics() *InlineResponse20043ResultsBgpMetrics`

NewInlineResponse20043ResultsBgpMetrics instantiates a new InlineResponse20043ResultsBgpMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043ResultsBgpMetricsWithDefaults

`func NewInlineResponse20043ResultsBgpMetricsWithDefaults() *InlineResponse20043ResultsBgpMetrics`

NewInlineResponse20043ResultsBgpMetricsWithDefaults instantiates a new InlineResponse20043ResultsBgpMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPreference

`func (o *InlineResponse20043ResultsBgpMetrics) GetLocalPreference() int32`

GetLocalPreference returns the LocalPreference field if non-nil, zero value otherwise.

### GetLocalPreferenceOk

`func (o *InlineResponse20043ResultsBgpMetrics) GetLocalPreferenceOk() (*int32, bool)`

GetLocalPreferenceOk returns a tuple with the LocalPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPreference

`func (o *InlineResponse20043ResultsBgpMetrics) SetLocalPreference(v int32)`

SetLocalPreference sets LocalPreference field to given value.

### HasLocalPreference

`func (o *InlineResponse20043ResultsBgpMetrics) HasLocalPreference() bool`

HasLocalPreference returns a boolean if a field has been set.

### GetMultiExitDiscriminator

`func (o *InlineResponse20043ResultsBgpMetrics) GetMultiExitDiscriminator() int32`

GetMultiExitDiscriminator returns the MultiExitDiscriminator field if non-nil, zero value otherwise.

### GetMultiExitDiscriminatorOk

`func (o *InlineResponse20043ResultsBgpMetrics) GetMultiExitDiscriminatorOk() (*int32, bool)`

GetMultiExitDiscriminatorOk returns a tuple with the MultiExitDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiExitDiscriminator

`func (o *InlineResponse20043ResultsBgpMetrics) SetMultiExitDiscriminator(v int32)`

SetMultiExitDiscriminator sets MultiExitDiscriminator field to given value.

### HasMultiExitDiscriminator

`func (o *InlineResponse20043ResultsBgpMetrics) HasMultiExitDiscriminator() bool`

HasMultiExitDiscriminator returns a boolean if a field has been set.

### GetWeight

`func (o *InlineResponse20043ResultsBgpMetrics) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *InlineResponse20043ResultsBgpMetrics) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *InlineResponse20043ResultsBgpMetrics) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *InlineResponse20043ResultsBgpMetrics) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



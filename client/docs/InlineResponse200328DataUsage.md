# InlineResponse200328DataUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total amount of data sent and received (in kilobytes). | [optional] 
**Upstream** | Pointer to **int32** | The amount of data sent (in kilobytes). | [optional] 
**Downstream** | Pointer to **int32** | The amount of data received (in kilobytes). | [optional] 

## Methods

### NewInlineResponse200328DataUsage

`func NewInlineResponse200328DataUsage() *InlineResponse200328DataUsage`

NewInlineResponse200328DataUsage instantiates a new InlineResponse200328DataUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200328DataUsageWithDefaults

`func NewInlineResponse200328DataUsageWithDefaults() *InlineResponse200328DataUsage`

NewInlineResponse200328DataUsageWithDefaults instantiates a new InlineResponse200328DataUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200328DataUsage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200328DataUsage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200328DataUsage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200328DataUsage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpstream

`func (o *InlineResponse200328DataUsage) GetUpstream() int32`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *InlineResponse200328DataUsage) GetUpstreamOk() (*int32, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *InlineResponse200328DataUsage) SetUpstream(v int32)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *InlineResponse200328DataUsage) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetDownstream

`func (o *InlineResponse200328DataUsage) GetDownstream() int32`

GetDownstream returns the Downstream field if non-nil, zero value otherwise.

### GetDownstreamOk

`func (o *InlineResponse200328DataUsage) GetDownstreamOk() (*int32, bool)`

GetDownstreamOk returns a tuple with the Downstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstream

`func (o *InlineResponse200328DataUsage) SetDownstream(v int32)`

SetDownstream sets Downstream field to given value.

### HasDownstream

`func (o *InlineResponse200328DataUsage) HasDownstream() bool`

HasDownstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



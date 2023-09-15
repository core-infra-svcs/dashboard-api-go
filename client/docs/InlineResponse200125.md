# InlineResponse200125

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp for the bandwidth usage snapshot. | [optional] 
**Total** | Pointer to **int32** | Total bandwidth usage, in mbps. | [optional] 
**Upstream** | Pointer to **int32** | Uploaded data, in mbps. | [optional] 
**Downstream** | Pointer to **int32** | Downloaded data, in mbps. | [optional] 

## Methods

### NewInlineResponse200125

`func NewInlineResponse200125() *InlineResponse200125`

NewInlineResponse200125 instantiates a new InlineResponse200125 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200125WithDefaults

`func NewInlineResponse200125WithDefaults() *InlineResponse200125`

NewInlineResponse200125WithDefaults instantiates a new InlineResponse200125 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200125) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200125) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200125) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200125) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200125) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200125) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200125) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200125) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpstream

`func (o *InlineResponse200125) GetUpstream() int32`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *InlineResponse200125) GetUpstreamOk() (*int32, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *InlineResponse200125) SetUpstream(v int32)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *InlineResponse200125) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetDownstream

`func (o *InlineResponse200125) GetDownstream() int32`

GetDownstream returns the Downstream field if non-nil, zero value otherwise.

### GetDownstreamOk

`func (o *InlineResponse200125) GetDownstreamOk() (*int32, bool)`

GetDownstreamOk returns a tuple with the Downstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstream

`func (o *InlineResponse200125) SetDownstream(v int32)`

SetDownstream sets Downstream field to given value.

### HasDownstream

`func (o *InlineResponse200125) HasDownstream() bool`

HasDownstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



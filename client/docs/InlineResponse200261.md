# InlineResponse200261

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp for the bandwidth usage snapshot. | [optional] 
**Total** | Pointer to **int32** | Total bandwidth usage, in mbps. | [optional] 
**Upstream** | Pointer to **int32** | Uploaded data, in mbps. | [optional] 
**Downstream** | Pointer to **int32** | Downloaded data, in mbps. | [optional] 

## Methods

### NewInlineResponse200261

`func NewInlineResponse200261() *InlineResponse200261`

NewInlineResponse200261 instantiates a new InlineResponse200261 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200261WithDefaults

`func NewInlineResponse200261WithDefaults() *InlineResponse200261`

NewInlineResponse200261WithDefaults instantiates a new InlineResponse200261 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200261) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200261) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200261) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200261) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200261) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200261) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200261) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200261) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpstream

`func (o *InlineResponse200261) GetUpstream() int32`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *InlineResponse200261) GetUpstreamOk() (*int32, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *InlineResponse200261) SetUpstream(v int32)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *InlineResponse200261) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetDownstream

`func (o *InlineResponse200261) GetDownstream() int32`

GetDownstream returns the Downstream field if non-nil, zero value otherwise.

### GetDownstreamOk

`func (o *InlineResponse200261) GetDownstreamOk() (*int32, bool)`

GetDownstreamOk returns a tuple with the Downstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstream

`func (o *InlineResponse200261) SetDownstream(v int32)`

SetDownstream sets Downstream field to given value.

### HasDownstream

`func (o *InlineResponse200261) HasDownstream() bool`

HasDownstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



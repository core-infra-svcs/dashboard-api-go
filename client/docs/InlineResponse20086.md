# InlineResponse20086

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | The timestamp | [optional] 
**Total** | Pointer to **float32** | The total traffic over a time range for clients on a network | [optional] 
**Upstream** | Pointer to **float32** | The upstream traffic over a time range for clients on a network | [optional] 
**Downstream** | Pointer to **float32** | The downstream traffic over a time range for clients on a network | [optional] 

## Methods

### NewInlineResponse20086

`func NewInlineResponse20086() *InlineResponse20086`

NewInlineResponse20086 instantiates a new InlineResponse20086 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20086WithDefaults

`func NewInlineResponse20086WithDefaults() *InlineResponse20086`

NewInlineResponse20086WithDefaults instantiates a new InlineResponse20086 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse20086) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse20086) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse20086) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse20086) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse20086) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse20086) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse20086) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse20086) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpstream

`func (o *InlineResponse20086) GetUpstream() float32`

GetUpstream returns the Upstream field if non-nil, zero value otherwise.

### GetUpstreamOk

`func (o *InlineResponse20086) GetUpstreamOk() (*float32, bool)`

GetUpstreamOk returns a tuple with the Upstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstream

`func (o *InlineResponse20086) SetUpstream(v float32)`

SetUpstream sets Upstream field to given value.

### HasUpstream

`func (o *InlineResponse20086) HasUpstream() bool`

HasUpstream returns a boolean if a field has been set.

### GetDownstream

`func (o *InlineResponse20086) GetDownstream() float32`

GetDownstream returns the Downstream field if non-nil, zero value otherwise.

### GetDownstreamOk

`func (o *InlineResponse20086) GetDownstreamOk() (*float32, bool)`

GetDownstreamOk returns a tuple with the Downstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstream

`func (o *InlineResponse20086) SetDownstream(v float32)`

SetDownstream sets Downstream field to given value.

### HasDownstream

`func (o *InlineResponse20086) HasDownstream() bool`

HasDownstream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



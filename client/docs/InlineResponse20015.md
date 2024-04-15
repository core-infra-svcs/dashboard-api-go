# InlineResponse20015

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]InlineResponse20015Sims**](InlineResponse20015Sims.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies which SIMs to use for primary and secondary. Required for devices with 3 or more SIMs. | [optional] 
**SimFailover** | Pointer to [**InlineResponse20015SimFailover**](InlineResponse20015SimFailover.md) |  | [optional] 

## Methods

### NewInlineResponse20015

`func NewInlineResponse20015() *InlineResponse20015`

NewInlineResponse20015 instantiates a new InlineResponse20015 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20015WithDefaults

`func NewInlineResponse20015WithDefaults() *InlineResponse20015`

NewInlineResponse20015WithDefaults instantiates a new InlineResponse20015 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *InlineResponse20015) GetSims() []InlineResponse20015Sims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineResponse20015) GetSimsOk() (*[]InlineResponse20015Sims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineResponse20015) SetSims(v []InlineResponse20015Sims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineResponse20015) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *InlineResponse20015) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *InlineResponse20015) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *InlineResponse20015) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *InlineResponse20015) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *InlineResponse20015) GetSimFailover() InlineResponse20015SimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *InlineResponse20015) GetSimFailoverOk() (*InlineResponse20015SimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *InlineResponse20015) SetSimFailover(v InlineResponse20015SimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *InlineResponse20015) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



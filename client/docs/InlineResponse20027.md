# InlineResponse20027

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]InlineResponse20027Sims**](InlineResponse20027Sims.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It&#39;s required for devices with 3 or more SIMs and can be used in place of &#39;isPrimary&#39; for dual-SIM devices. Use the raw eSIM slot value for the device, such as &#39;sim2&#39; or &#39;sim3&#39;. Sim failover will occur only between primary and secondary sim slots. | [optional] 
**SimFailover** | Pointer to [**InlineResponse20027SimFailover**](InlineResponse20027SimFailover.md) |  | [optional] 

## Methods

### NewInlineResponse20027

`func NewInlineResponse20027() *InlineResponse20027`

NewInlineResponse20027 instantiates a new InlineResponse20027 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027WithDefaults

`func NewInlineResponse20027WithDefaults() *InlineResponse20027`

NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *InlineResponse20027) GetSims() []InlineResponse20027Sims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineResponse20027) GetSimsOk() (*[]InlineResponse20027Sims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineResponse20027) SetSims(v []InlineResponse20027Sims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineResponse20027) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *InlineResponse20027) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *InlineResponse20027) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *InlineResponse20027) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *InlineResponse20027) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *InlineResponse20027) GetSimFailover() InlineResponse20027SimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *InlineResponse20027) GetSimFailoverOk() (*InlineResponse20027SimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *InlineResponse20027) SetSimFailover(v InlineResponse20027SimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *InlineResponse20027) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



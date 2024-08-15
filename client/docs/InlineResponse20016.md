# InlineResponse20016

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]InlineResponse20016Sims**](InlineResponse20016Sims.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It&#39;s required for devices with 3 or more SIMs and can be used in place of &#39;isPrimary&#39; for dual-SIM devices. To indicate eSIM, use &#39;sim3&#39;. Sim failover will occur only between primary and secondary sim slots. | [optional] 
**SimFailover** | Pointer to [**InlineResponse20016SimFailover**](InlineResponse20016SimFailover.md) |  | [optional] 

## Methods

### NewInlineResponse20016

`func NewInlineResponse20016() *InlineResponse20016`

NewInlineResponse20016 instantiates a new InlineResponse20016 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20016WithDefaults

`func NewInlineResponse20016WithDefaults() *InlineResponse20016`

NewInlineResponse20016WithDefaults instantiates a new InlineResponse20016 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *InlineResponse20016) GetSims() []InlineResponse20016Sims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineResponse20016) GetSimsOk() (*[]InlineResponse20016Sims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineResponse20016) SetSims(v []InlineResponse20016Sims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineResponse20016) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *InlineResponse20016) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *InlineResponse20016) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *InlineResponse20016) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *InlineResponse20016) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *InlineResponse20016) GetSimFailover() InlineResponse20016SimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *InlineResponse20016) GetSimFailoverOk() (*InlineResponse20016SimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *InlineResponse20016) SetSimFailover(v InlineResponse20016SimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *InlineResponse20016) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineObject14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]DevicesSerialCellularSimsSims**](DevicesSerialCellularSimsSims.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It&#39;s required for devices with 3 or more SIMs and can be used in place of &#39;isPrimary&#39; for dual-SIM devices. To indicate eSIM, use &#39;sim3&#39;. Sim failover will occur only between primary and secondary sim slots. | [optional] 
**SimFailover** | Pointer to [**DevicesSerialCellularSimsSimFailover**](DevicesSerialCellularSimsSimFailover.md) |  | [optional] 

## Methods

### NewInlineObject14

`func NewInlineObject14() *InlineObject14`

NewInlineObject14 instantiates a new InlineObject14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject14WithDefaults

`func NewInlineObject14WithDefaults() *InlineObject14`

NewInlineObject14WithDefaults instantiates a new InlineObject14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *InlineObject14) GetSims() []DevicesSerialCellularSimsSims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineObject14) GetSimsOk() (*[]DevicesSerialCellularSimsSims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineObject14) SetSims(v []DevicesSerialCellularSimsSims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineObject14) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *InlineObject14) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *InlineObject14) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *InlineObject14) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *InlineObject14) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *InlineObject14) GetSimFailover() DevicesSerialCellularSimsSimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *InlineObject14) GetSimFailoverOk() (*DevicesSerialCellularSimsSimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *InlineObject14) SetSimFailover(v DevicesSerialCellularSimsSimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *InlineObject14) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



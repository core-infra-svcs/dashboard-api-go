# InlineObject16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]DevicesSerialCellularSimsSims**](DevicesSerialCellularSimsSims.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It&#39;s required for devices with 3 or more SIMs and can be used in place of &#39;isPrimary&#39; for dual-SIM devices. Use the raw eSIM slot value for the device, such as &#39;sim2&#39; or &#39;sim3&#39;. Sim failover will occur only between primary and secondary sim slots. | [optional] 
**SimFailover** | Pointer to [**DevicesSerialCellularSimsSimFailover**](DevicesSerialCellularSimsSimFailover.md) |  | [optional] 

## Methods

### NewInlineObject16

`func NewInlineObject16() *InlineObject16`

NewInlineObject16 instantiates a new InlineObject16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject16WithDefaults

`func NewInlineObject16WithDefaults() *InlineObject16`

NewInlineObject16WithDefaults instantiates a new InlineObject16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *InlineObject16) GetSims() []DevicesSerialCellularSimsSims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineObject16) GetSimsOk() (*[]DevicesSerialCellularSimsSims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineObject16) SetSims(v []DevicesSerialCellularSimsSims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineObject16) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *InlineObject16) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *InlineObject16) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *InlineObject16) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *InlineObject16) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *InlineObject16) GetSimFailover() DevicesSerialCellularSimsSimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *InlineObject16) GetSimFailoverOk() (*DevicesSerialCellularSimsSimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *InlineObject16) SetSimFailover(v DevicesSerialCellularSimsSimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *InlineObject16) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



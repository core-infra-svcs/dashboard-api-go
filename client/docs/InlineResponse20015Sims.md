# InlineResponse20015Sims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is used for boot. Must be true on single-SIM devices. | [optional] [default to false]
**Apns** | Pointer to [**[]InlineResponse20015Apns**](InlineResponse20015Apns.md) | APN configurations. If empty, the default APN will be used. | [optional] 

## Methods

### NewInlineResponse20015Sims

`func NewInlineResponse20015Sims() *InlineResponse20015Sims`

NewInlineResponse20015Sims instantiates a new InlineResponse20015Sims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20015SimsWithDefaults

`func NewInlineResponse20015SimsWithDefaults() *InlineResponse20015Sims`

NewInlineResponse20015SimsWithDefaults instantiates a new InlineResponse20015Sims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse20015Sims) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse20015Sims) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse20015Sims) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse20015Sims) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsPrimary

`func (o *InlineResponse20015Sims) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *InlineResponse20015Sims) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *InlineResponse20015Sims) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *InlineResponse20015Sims) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetApns

`func (o *InlineResponse20015Sims) GetApns() []InlineResponse20015Apns`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *InlineResponse20015Sims) GetApnsOk() (*[]InlineResponse20015Apns, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *InlineResponse20015Sims) SetApns(v []InlineResponse20015Apns)`

SetApns sets Apns field to given value.

### HasApns

`func (o *InlineResponse20015Sims) HasApns() bool`

HasApns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



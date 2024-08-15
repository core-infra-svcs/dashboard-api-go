# InlineResponse20016Sims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. Use &#39;sim3&#39; for eSIM. | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using &#39;simOrdering&#39;. | [optional] [default to false]
**Apns** | Pointer to [**[]InlineResponse20016Apns**](InlineResponse20016Apns.md) | APN configurations. If empty, the default APN will be used. | [optional] 

## Methods

### NewInlineResponse20016Sims

`func NewInlineResponse20016Sims() *InlineResponse20016Sims`

NewInlineResponse20016Sims instantiates a new InlineResponse20016Sims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20016SimsWithDefaults

`func NewInlineResponse20016SimsWithDefaults() *InlineResponse20016Sims`

NewInlineResponse20016SimsWithDefaults instantiates a new InlineResponse20016Sims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse20016Sims) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse20016Sims) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse20016Sims) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse20016Sims) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsPrimary

`func (o *InlineResponse20016Sims) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *InlineResponse20016Sims) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *InlineResponse20016Sims) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *InlineResponse20016Sims) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetApns

`func (o *InlineResponse20016Sims) GetApns() []InlineResponse20016Apns`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *InlineResponse20016Sims) GetApnsOk() (*[]InlineResponse20016Apns, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *InlineResponse20016Sims) SetApns(v []InlineResponse20016Apns)`

SetApns sets Apns field to given value.

### HasApns

`func (o *InlineResponse20016Sims) HasApns() bool`

HasApns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



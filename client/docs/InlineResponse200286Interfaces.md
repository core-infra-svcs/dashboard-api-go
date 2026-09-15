# InlineResponse200286Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Full interface name for this interface | [optional] 
**Slot** | Pointer to **NullableInt32** | Optional slot number for this interface, if applicable. | [optional] 
**Subslot** | Pointer to **NullableInt32** | Optional subslot number for this interface, if applicable. | [optional] 
**Number** | Pointer to **int32** | Leaf interface number for this interface | [optional] 
**ByType** | Pointer to [**[]InlineResponse200286ByType**](InlineResponse200286ByType.md) | Packet counter overviews for this interface, grouped by packet type, with counts and average rates over the requested timespan. Note: results may be approximate because packet counter aggregated records may not align exactly with the requested timespan. Recent/current data may be incomplete because aggregated data is refreshed periodically, which can introduce latency. | [optional] 

## Methods

### NewInlineResponse200286Interfaces

`func NewInlineResponse200286Interfaces() *InlineResponse200286Interfaces`

NewInlineResponse200286Interfaces instantiates a new InlineResponse200286Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200286InterfacesWithDefaults

`func NewInlineResponse200286InterfacesWithDefaults() *InlineResponse200286Interfaces`

NewInlineResponse200286InterfacesWithDefaults instantiates a new InlineResponse200286Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200286Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200286Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200286Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200286Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200286Interfaces) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200286Interfaces) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200286Interfaces) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200286Interfaces) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### SetSlotNil

`func (o *InlineResponse200286Interfaces) SetSlotNil(b bool)`

 SetSlotNil sets the value for Slot to be an explicit nil

### UnsetSlot
`func (o *InlineResponse200286Interfaces) UnsetSlot()`

UnsetSlot ensures that no value is present for Slot, not even an explicit nil
### GetSubslot

`func (o *InlineResponse200286Interfaces) GetSubslot() int32`

GetSubslot returns the Subslot field if non-nil, zero value otherwise.

### GetSubslotOk

`func (o *InlineResponse200286Interfaces) GetSubslotOk() (*int32, bool)`

GetSubslotOk returns a tuple with the Subslot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubslot

`func (o *InlineResponse200286Interfaces) SetSubslot(v int32)`

SetSubslot sets Subslot field to given value.

### HasSubslot

`func (o *InlineResponse200286Interfaces) HasSubslot() bool`

HasSubslot returns a boolean if a field has been set.

### SetSubslotNil

`func (o *InlineResponse200286Interfaces) SetSubslotNil(b bool)`

 SetSubslotNil sets the value for Subslot to be an explicit nil

### UnsetSubslot
`func (o *InlineResponse200286Interfaces) UnsetSubslot()`

UnsetSubslot ensures that no value is present for Subslot, not even an explicit nil
### GetNumber

`func (o *InlineResponse200286Interfaces) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200286Interfaces) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200286Interfaces) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200286Interfaces) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetByType

`func (o *InlineResponse200286Interfaces) GetByType() []InlineResponse200286ByType`

GetByType returns the ByType field if non-nil, zero value otherwise.

### GetByTypeOk

`func (o *InlineResponse200286Interfaces) GetByTypeOk() (*[]InlineResponse200286ByType, bool)`

GetByTypeOk returns a tuple with the ByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByType

`func (o *InlineResponse200286Interfaces) SetByType(v []InlineResponse200286ByType)`

SetByType sets ByType field to given value.

### HasByType

`func (o *InlineResponse200286Interfaces) HasByType() bool`

HasByType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



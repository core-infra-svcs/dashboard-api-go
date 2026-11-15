# InlineResponse200283PortInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The full interface identifier for this port | [optional] 
**Slot** | Pointer to **NullableInt32** | The slot number for this port. Range depends on the appliance model. Null when the port does not use a slotted interface identifier. | [optional] 
**Subslot** | Pointer to **NullableInt32** | The subslot number for this port. Range depends on the appliance model. Null when the port does not use a subslot interface identifier. | [optional] 
**Number** | Pointer to **int32** | The leaf interface number for this port. Range depends on the appliance model. | [optional] 

## Methods

### NewInlineResponse200283PortInterface

`func NewInlineResponse200283PortInterface() *InlineResponse200283PortInterface`

NewInlineResponse200283PortInterface instantiates a new InlineResponse200283PortInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200283PortInterfaceWithDefaults

`func NewInlineResponse200283PortInterfaceWithDefaults() *InlineResponse200283PortInterface`

NewInlineResponse200283PortInterfaceWithDefaults instantiates a new InlineResponse200283PortInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200283PortInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200283PortInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200283PortInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200283PortInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlot

`func (o *InlineResponse200283PortInterface) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200283PortInterface) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200283PortInterface) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200283PortInterface) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### SetSlotNil

`func (o *InlineResponse200283PortInterface) SetSlotNil(b bool)`

 SetSlotNil sets the value for Slot to be an explicit nil

### UnsetSlot
`func (o *InlineResponse200283PortInterface) UnsetSlot()`

UnsetSlot ensures that no value is present for Slot, not even an explicit nil
### GetSubslot

`func (o *InlineResponse200283PortInterface) GetSubslot() int32`

GetSubslot returns the Subslot field if non-nil, zero value otherwise.

### GetSubslotOk

`func (o *InlineResponse200283PortInterface) GetSubslotOk() (*int32, bool)`

GetSubslotOk returns a tuple with the Subslot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubslot

`func (o *InlineResponse200283PortInterface) SetSubslot(v int32)`

SetSubslot sets Subslot field to given value.

### HasSubslot

`func (o *InlineResponse200283PortInterface) HasSubslot() bool`

HasSubslot returns a boolean if a field has been set.

### SetSubslotNil

`func (o *InlineResponse200283PortInterface) SetSubslotNil(b bool)`

 SetSubslotNil sets the value for Subslot to be an explicit nil

### UnsetSubslot
`func (o *InlineResponse200283PortInterface) UnsetSubslot()`

UnsetSubslot ensures that no value is present for Subslot, not even an explicit nil
### GetNumber

`func (o *InlineResponse200283PortInterface) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200283PortInterface) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200283PortInterface) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200283PortInterface) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



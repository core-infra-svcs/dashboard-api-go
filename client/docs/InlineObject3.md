# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of a device | [optional] 
**Tags** | Pointer to **[]string** | The list of tags of a device | [optional] 
**Lat** | Pointer to **float32** | The latitude of a device | [optional] 
**Lng** | Pointer to **float32** | The longitude of a device | [optional] 
**Address** | Pointer to **NullableString** | The address of a device | [optional] 
**Notes** | Pointer to **NullableString** | The notes for the device. String. Limited to 255 characters. | [optional] 
**MoveMapMarker** | Pointer to **bool** | Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified. | [optional] 
**SwitchProfileId** | Pointer to **NullableString** | The ID of a switch template to bind to the device (for available switch templates, see the &#39;Switch Templates&#39; endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template. | [optional] 
**FloorPlanId** | Pointer to **NullableString** | The floor plan to associate to this device. null disassociates the device from the floorplan. | [optional] 

## Methods

### NewInlineObject3

`func NewInlineObject3() *InlineObject3`

NewInlineObject3 instantiates a new InlineObject3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject3WithDefaults

`func NewInlineObject3WithDefaults() *InlineObject3`

NewInlineObject3WithDefaults instantiates a new InlineObject3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject3) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InlineObject3) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InlineObject3) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTags

`func (o *InlineObject3) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject3) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject3) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLat

`func (o *InlineObject3) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *InlineObject3) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *InlineObject3) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *InlineObject3) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *InlineObject3) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *InlineObject3) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *InlineObject3) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *InlineObject3) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *InlineObject3) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineObject3) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineObject3) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineObject3) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *InlineObject3) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *InlineObject3) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetNotes

`func (o *InlineObject3) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineObject3) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineObject3) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineObject3) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InlineObject3) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InlineObject3) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetMoveMapMarker

`func (o *InlineObject3) GetMoveMapMarker() bool`

GetMoveMapMarker returns the MoveMapMarker field if non-nil, zero value otherwise.

### GetMoveMapMarkerOk

`func (o *InlineObject3) GetMoveMapMarkerOk() (*bool, bool)`

GetMoveMapMarkerOk returns a tuple with the MoveMapMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveMapMarker

`func (o *InlineObject3) SetMoveMapMarker(v bool)`

SetMoveMapMarker sets MoveMapMarker field to given value.

### HasMoveMapMarker

`func (o *InlineObject3) HasMoveMapMarker() bool`

HasMoveMapMarker returns a boolean if a field has been set.

### GetSwitchProfileId

`func (o *InlineObject3) GetSwitchProfileId() string`

GetSwitchProfileId returns the SwitchProfileId field if non-nil, zero value otherwise.

### GetSwitchProfileIdOk

`func (o *InlineObject3) GetSwitchProfileIdOk() (*string, bool)`

GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfileId

`func (o *InlineObject3) SetSwitchProfileId(v string)`

SetSwitchProfileId sets SwitchProfileId field to given value.

### HasSwitchProfileId

`func (o *InlineObject3) HasSwitchProfileId() bool`

HasSwitchProfileId returns a boolean if a field has been set.

### SetSwitchProfileIdNil

`func (o *InlineObject3) SetSwitchProfileIdNil(b bool)`

 SetSwitchProfileIdNil sets the value for SwitchProfileId to be an explicit nil

### UnsetSwitchProfileId
`func (o *InlineObject3) UnsetSwitchProfileId()`

UnsetSwitchProfileId ensures that no value is present for SwitchProfileId, not even an explicit nil
### GetFloorPlanId

`func (o *InlineObject3) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineObject3) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineObject3) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineObject3) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### SetFloorPlanIdNil

`func (o *InlineObject3) SetFloorPlanIdNil(b bool)`

 SetFloorPlanIdNil sets the value for FloorPlanId to be an explicit nil

### UnsetFloorPlanId
`func (o *InlineObject3) UnsetFloorPlanId()`

UnsetFloorPlanId ensures that no value is present for FloorPlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



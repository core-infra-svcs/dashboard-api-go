# InlineResponse2008InterfacesWan2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the interface. | [optional] 
**VlanTagging** | Pointer to [**InlineResponse2008InterfacesWan1VlanTagging**](InlineResponse2008InterfacesWan1VlanTagging.md) |  | [optional] 
**Svis** | Pointer to [**InlineResponse2008InterfacesWan1Svis**](InlineResponse2008InterfacesWan1Svis.md) |  | [optional] 
**Pppoe** | Pointer to [**InlineResponse2008InterfacesWan1Pppoe**](InlineResponse2008InterfacesWan1Pppoe.md) |  | [optional] 

## Methods

### NewInlineResponse2008InterfacesWan2

`func NewInlineResponse2008InterfacesWan2() *InlineResponse2008InterfacesWan2`

NewInlineResponse2008InterfacesWan2 instantiates a new InlineResponse2008InterfacesWan2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008InterfacesWan2WithDefaults

`func NewInlineResponse2008InterfacesWan2WithDefaults() *InlineResponse2008InterfacesWan2`

NewInlineResponse2008InterfacesWan2WithDefaults instantiates a new InlineResponse2008InterfacesWan2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse2008InterfacesWan2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse2008InterfacesWan2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse2008InterfacesWan2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse2008InterfacesWan2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanTagging

`func (o *InlineResponse2008InterfacesWan2) GetVlanTagging() InlineResponse2008InterfacesWan1VlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *InlineResponse2008InterfacesWan2) GetVlanTaggingOk() (*InlineResponse2008InterfacesWan1VlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *InlineResponse2008InterfacesWan2) SetVlanTagging(v InlineResponse2008InterfacesWan1VlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *InlineResponse2008InterfacesWan2) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetSvis

`func (o *InlineResponse2008InterfacesWan2) GetSvis() InlineResponse2008InterfacesWan1Svis`

GetSvis returns the Svis field if non-nil, zero value otherwise.

### GetSvisOk

`func (o *InlineResponse2008InterfacesWan2) GetSvisOk() (*InlineResponse2008InterfacesWan1Svis, bool)`

GetSvisOk returns a tuple with the Svis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvis

`func (o *InlineResponse2008InterfacesWan2) SetSvis(v InlineResponse2008InterfacesWan1Svis)`

SetSvis sets Svis field to given value.

### HasSvis

`func (o *InlineResponse2008InterfacesWan2) HasSvis() bool`

HasSvis returns a boolean if a field has been set.

### GetPppoe

`func (o *InlineResponse2008InterfacesWan2) GetPppoe() InlineResponse2008InterfacesWan1Pppoe`

GetPppoe returns the Pppoe field if non-nil, zero value otherwise.

### GetPppoeOk

`func (o *InlineResponse2008InterfacesWan2) GetPppoeOk() (*InlineResponse2008InterfacesWan1Pppoe, bool)`

GetPppoeOk returns a tuple with the Pppoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPppoe

`func (o *InlineResponse2008InterfacesWan2) SetPppoe(v InlineResponse2008InterfacesWan1Pppoe)`

SetPppoe sets Pppoe field to given value.

### HasPppoe

`func (o *InlineResponse2008InterfacesWan2) HasPppoe() bool`

HasPppoe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



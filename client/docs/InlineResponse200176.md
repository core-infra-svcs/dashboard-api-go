# InlineResponse200176

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iname** | Pointer to **string** | IName of the VLAN profile | [optional] 
**Name** | Pointer to **string** | Name of the profile, string length must be from 1 to 255 characters | [optional] 
**IsDefault** | Pointer to **bool** | Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned | [optional] 
**VlanNames** | Pointer to [**[]NetworksNetworkIdVlanProfilesVlanNames**](NetworksNetworkIdVlanProfilesVlanNames.md) | An array of named VLANs | [optional] 
**VlanGroups** | Pointer to [**[]NetworksNetworkIdVlanProfilesVlanGroups**](NetworksNetworkIdVlanProfilesVlanGroups.md) | An array of named VLANs | [optional] 

## Methods

### NewInlineResponse200176

`func NewInlineResponse200176() *InlineResponse200176`

NewInlineResponse200176 instantiates a new InlineResponse200176 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200176WithDefaults

`func NewInlineResponse200176WithDefaults() *InlineResponse200176`

NewInlineResponse200176WithDefaults instantiates a new InlineResponse200176 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIname

`func (o *InlineResponse200176) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *InlineResponse200176) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *InlineResponse200176) SetIname(v string)`

SetIname sets Iname field to given value.

### HasIname

`func (o *InlineResponse200176) HasIname() bool`

HasIname returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200176) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200176) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200176) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200176) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *InlineResponse200176) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InlineResponse200176) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InlineResponse200176) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *InlineResponse200176) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetVlanNames

`func (o *InlineResponse200176) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *InlineResponse200176) GetVlanNamesOk() (*[]NetworksNetworkIdVlanProfilesVlanNames, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *InlineResponse200176) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames)`

SetVlanNames sets VlanNames field to given value.

### HasVlanNames

`func (o *InlineResponse200176) HasVlanNames() bool`

HasVlanNames returns a boolean if a field has been set.

### GetVlanGroups

`func (o *InlineResponse200176) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *InlineResponse200176) GetVlanGroupsOk() (*[]NetworksNetworkIdVlanProfilesVlanGroups, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *InlineResponse200176) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups)`

SetVlanGroups sets VlanGroups field to given value.

### HasVlanGroups

`func (o *InlineResponse200176) HasVlanGroups() bool`

HasVlanGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



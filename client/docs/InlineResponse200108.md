# InlineResponse200108

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iname** | Pointer to **string** | IName of the VLAN profile | [optional] 
**Name** | Pointer to **string** | Name of the profile, string length must be from 1 to 255 characters | [optional] 
**IsDefault** | Pointer to **bool** | Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned | [optional] 
**VlanNames** | Pointer to [**[]NetworksNetworkIdVlanProfilesVlanNames**](NetworksNetworkIdVlanProfilesVlanNames.md) | An array of named VLANs | [optional] 
**VlanGroups** | Pointer to [**[]NetworksNetworkIdVlanProfilesVlanGroups**](NetworksNetworkIdVlanProfilesVlanGroups.md) | An array of named VLANs | [optional] 

## Methods

### NewInlineResponse200108

`func NewInlineResponse200108() *InlineResponse200108`

NewInlineResponse200108 instantiates a new InlineResponse200108 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108WithDefaults

`func NewInlineResponse200108WithDefaults() *InlineResponse200108`

NewInlineResponse200108WithDefaults instantiates a new InlineResponse200108 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIname

`func (o *InlineResponse200108) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *InlineResponse200108) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *InlineResponse200108) SetIname(v string)`

SetIname sets Iname field to given value.

### HasIname

`func (o *InlineResponse200108) HasIname() bool`

HasIname returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200108) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200108) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200108) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200108) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *InlineResponse200108) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InlineResponse200108) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InlineResponse200108) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *InlineResponse200108) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetVlanNames

`func (o *InlineResponse200108) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *InlineResponse200108) GetVlanNamesOk() (*[]NetworksNetworkIdVlanProfilesVlanNames, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *InlineResponse200108) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames)`

SetVlanNames sets VlanNames field to given value.

### HasVlanNames

`func (o *InlineResponse200108) HasVlanNames() bool`

HasVlanNames returns a boolean if a field has been set.

### GetVlanGroups

`func (o *InlineResponse200108) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *InlineResponse200108) GetVlanGroupsOk() (*[]NetworksNetworkIdVlanProfilesVlanGroups, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *InlineResponse200108) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups)`

SetVlanGroups sets VlanGroups field to given value.

### HasVlanGroups

`func (o *InlineResponse200108) HasVlanGroups() bool`

HasVlanGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



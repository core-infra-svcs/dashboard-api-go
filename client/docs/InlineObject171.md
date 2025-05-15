# InlineObject171

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the profile, string length must be from 1 to 255 characters | 
**VlanNames** | [**[]NetworksNetworkIdVlanProfilesVlanNames1**](NetworksNetworkIdVlanProfilesVlanNames1.md) | An array of named VLANs | 
**VlanGroups** | [**[]NetworksNetworkIdVlanProfilesVlanGroups1**](NetworksNetworkIdVlanProfilesVlanGroups1.md) | An array of VLAN groups | 

## Methods

### NewInlineObject171

`func NewInlineObject171(name string, vlanNames []NetworksNetworkIdVlanProfilesVlanNames1, vlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1, ) *InlineObject171`

NewInlineObject171 instantiates a new InlineObject171 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject171WithDefaults

`func NewInlineObject171WithDefaults() *InlineObject171`

NewInlineObject171WithDefaults instantiates a new InlineObject171 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject171) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject171) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject171) SetName(v string)`

SetName sets Name field to given value.


### GetVlanNames

`func (o *InlineObject171) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames1`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *InlineObject171) GetVlanNamesOk() (*[]NetworksNetworkIdVlanProfilesVlanNames1, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *InlineObject171) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames1)`

SetVlanNames sets VlanNames field to given value.


### GetVlanGroups

`func (o *InlineObject171) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups1`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *InlineObject171) GetVlanGroupsOk() (*[]NetworksNetworkIdVlanProfilesVlanGroups1, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *InlineObject171) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups1)`

SetVlanGroups sets VlanGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



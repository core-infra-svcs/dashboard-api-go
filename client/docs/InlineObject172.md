# InlineObject172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the profile, string length must be from 1 to 255 characters | 
**VlanNames** | [**[]NetworksNetworkIdVlanProfilesVlanNames1**](NetworksNetworkIdVlanProfilesVlanNames1.md) | An array of named VLANs | 
**VlanGroups** | [**[]NetworksNetworkIdVlanProfilesVlanGroups1**](NetworksNetworkIdVlanProfilesVlanGroups1.md) | An array of VLAN groups | 
**Iname** | **string** | IName of the profile | 

## Methods

### NewInlineObject172

`func NewInlineObject172(name string, vlanNames []NetworksNetworkIdVlanProfilesVlanNames1, vlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1, iname string, ) *InlineObject172`

NewInlineObject172 instantiates a new InlineObject172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject172WithDefaults

`func NewInlineObject172WithDefaults() *InlineObject172`

NewInlineObject172WithDefaults instantiates a new InlineObject172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject172) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject172) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject172) SetName(v string)`

SetName sets Name field to given value.


### GetVlanNames

`func (o *InlineObject172) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames1`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *InlineObject172) GetVlanNamesOk() (*[]NetworksNetworkIdVlanProfilesVlanNames1, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *InlineObject172) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames1)`

SetVlanNames sets VlanNames field to given value.


### GetVlanGroups

`func (o *InlineObject172) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups1`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *InlineObject172) GetVlanGroupsOk() (*[]NetworksNetworkIdVlanProfilesVlanGroups1, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *InlineObject172) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups1)`

SetVlanGroups sets VlanGroups field to given value.


### GetIname

`func (o *InlineObject172) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *InlineObject172) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *InlineObject172) SetIname(v string)`

SetIname sets Iname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



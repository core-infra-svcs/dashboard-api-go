# InlineObject194

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the profile, string length must be from 1 to 255 characters | 
**ActiveVlans** | Pointer to **string** | The active VLANs for the VLAN profile. Only applicable to trunk ports. The given range must be inclusive of all named VLANs. | [optional] 
**VlanNames** | [**[]NetworksNetworkIdVlanProfilesVlanNames1**](NetworksNetworkIdVlanProfilesVlanNames1.md) | An array of named VLANs | 
**VlanGroups** | [**[]NetworksNetworkIdVlanProfilesVlanGroups1**](NetworksNetworkIdVlanProfilesVlanGroups1.md) | An array of VLAN groups | 
**Iname** | **string** | IName of the profile | 

## Methods

### NewInlineObject194

`func NewInlineObject194(name string, vlanNames []NetworksNetworkIdVlanProfilesVlanNames1, vlanGroups []NetworksNetworkIdVlanProfilesVlanGroups1, iname string, ) *InlineObject194`

NewInlineObject194 instantiates a new InlineObject194 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject194WithDefaults

`func NewInlineObject194WithDefaults() *InlineObject194`

NewInlineObject194WithDefaults instantiates a new InlineObject194 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject194) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject194) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject194) SetName(v string)`

SetName sets Name field to given value.


### GetActiveVlans

`func (o *InlineObject194) GetActiveVlans() string`

GetActiveVlans returns the ActiveVlans field if non-nil, zero value otherwise.

### GetActiveVlansOk

`func (o *InlineObject194) GetActiveVlansOk() (*string, bool)`

GetActiveVlansOk returns a tuple with the ActiveVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVlans

`func (o *InlineObject194) SetActiveVlans(v string)`

SetActiveVlans sets ActiveVlans field to given value.

### HasActiveVlans

`func (o *InlineObject194) HasActiveVlans() bool`

HasActiveVlans returns a boolean if a field has been set.

### GetVlanNames

`func (o *InlineObject194) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames1`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *InlineObject194) GetVlanNamesOk() (*[]NetworksNetworkIdVlanProfilesVlanNames1, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *InlineObject194) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames1)`

SetVlanNames sets VlanNames field to given value.


### GetVlanGroups

`func (o *InlineObject194) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups1`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *InlineObject194) GetVlanGroupsOk() (*[]NetworksNetworkIdVlanProfilesVlanGroups1, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *InlineObject194) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups1)`

SetVlanGroups sets VlanGroups field to given value.


### GetIname

`func (o *InlineObject194) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *InlineObject194) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *InlineObject194) SetIname(v string)`

SetIname sets Iname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200163

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iname** | Pointer to **string** | IName of the VLAN profile | [optional] 
**Name** | Pointer to **string** | Name of the profile, string length must be from 1 to 255 characters | [optional] 
**IsDefault** | Pointer to **bool** | Boolean indicating the default VLAN Profile for any device that does not have a profile explicitly assigned | [optional] 
**VlanNames** | Pointer to [**[]NetworksNetworkIdVlanProfilesVlanNames**](NetworksNetworkIdVlanProfilesVlanNames.md) | An array of named VLANs | [optional] 
**VlanGroups** | Pointer to [**[]NetworksNetworkIdVlanProfilesVlanGroups**](NetworksNetworkIdVlanProfilesVlanGroups.md) | An array of named VLANs | [optional] 

## Methods

### NewInlineResponse200163

`func NewInlineResponse200163() *InlineResponse200163`

NewInlineResponse200163 instantiates a new InlineResponse200163 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200163WithDefaults

`func NewInlineResponse200163WithDefaults() *InlineResponse200163`

NewInlineResponse200163WithDefaults instantiates a new InlineResponse200163 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIname

`func (o *InlineResponse200163) GetIname() string`

GetIname returns the Iname field if non-nil, zero value otherwise.

### GetInameOk

`func (o *InlineResponse200163) GetInameOk() (*string, bool)`

GetInameOk returns a tuple with the Iname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIname

`func (o *InlineResponse200163) SetIname(v string)`

SetIname sets Iname field to given value.

### HasIname

`func (o *InlineResponse200163) HasIname() bool`

HasIname returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200163) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200163) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200163) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200163) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *InlineResponse200163) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InlineResponse200163) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InlineResponse200163) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *InlineResponse200163) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetVlanNames

`func (o *InlineResponse200163) GetVlanNames() []NetworksNetworkIdVlanProfilesVlanNames`

GetVlanNames returns the VlanNames field if non-nil, zero value otherwise.

### GetVlanNamesOk

`func (o *InlineResponse200163) GetVlanNamesOk() (*[]NetworksNetworkIdVlanProfilesVlanNames, bool)`

GetVlanNamesOk returns a tuple with the VlanNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNames

`func (o *InlineResponse200163) SetVlanNames(v []NetworksNetworkIdVlanProfilesVlanNames)`

SetVlanNames sets VlanNames field to given value.

### HasVlanNames

`func (o *InlineResponse200163) HasVlanNames() bool`

HasVlanNames returns a boolean if a field has been set.

### GetVlanGroups

`func (o *InlineResponse200163) GetVlanGroups() []NetworksNetworkIdVlanProfilesVlanGroups`

GetVlanGroups returns the VlanGroups field if non-nil, zero value otherwise.

### GetVlanGroupsOk

`func (o *InlineResponse200163) GetVlanGroupsOk() (*[]NetworksNetworkIdVlanProfilesVlanGroups, bool)`

GetVlanGroupsOk returns a tuple with the VlanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanGroups

`func (o *InlineResponse200163) SetVlanGroups(v []NetworksNetworkIdVlanProfilesVlanGroups)`

SetVlanGroups sets VlanGroups field to given value.

### HasVlanGroups

`func (o *InlineResponse200163) HasVlanGroups() bool`

HasVlanGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



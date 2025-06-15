# InlineObject172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanProfile** | Pointer to [**NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile**](NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile.md) |  | [optional] 
**Serials** | **[]string** | Array of Device Serials | 
**StackIds** | **[]string** | Array of Switch Stack IDs | 

## Methods

### NewInlineObject172

`func NewInlineObject172(serials []string, stackIds []string, ) *InlineObject172`

NewInlineObject172 instantiates a new InlineObject172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject172WithDefaults

`func NewInlineObject172WithDefaults() *InlineObject172`

NewInlineObject172WithDefaults instantiates a new InlineObject172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanProfile

`func (o *InlineObject172) GetVlanProfile() NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile`

GetVlanProfile returns the VlanProfile field if non-nil, zero value otherwise.

### GetVlanProfileOk

`func (o *InlineObject172) GetVlanProfileOk() (*NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile, bool)`

GetVlanProfileOk returns a tuple with the VlanProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanProfile

`func (o *InlineObject172) SetVlanProfile(v NetworksNetworkIdVlanProfilesAssignmentsReassignVlanProfile)`

SetVlanProfile sets VlanProfile field to given value.

### HasVlanProfile

`func (o *InlineObject172) HasVlanProfile() bool`

HasVlanProfile returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject172) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject172) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject172) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetStackIds

`func (o *InlineObject172) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *InlineObject172) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *InlineObject172) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



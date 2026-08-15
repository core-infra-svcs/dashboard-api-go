# InlineResponse200217

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanProfile** | Pointer to [**InlineResponse200217VlanProfile**](InlineResponse200217VlanProfile.md) |  | [optional] 
**Serials** | Pointer to **[]string** | Array of Device Serials | [optional] 
**StackIds** | Pointer to **[]string** | Array of Switch Stack IDs | [optional] 

## Methods

### NewInlineResponse200217

`func NewInlineResponse200217() *InlineResponse200217`

NewInlineResponse200217 instantiates a new InlineResponse200217 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200217WithDefaults

`func NewInlineResponse200217WithDefaults() *InlineResponse200217`

NewInlineResponse200217WithDefaults instantiates a new InlineResponse200217 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanProfile

`func (o *InlineResponse200217) GetVlanProfile() InlineResponse200217VlanProfile`

GetVlanProfile returns the VlanProfile field if non-nil, zero value otherwise.

### GetVlanProfileOk

`func (o *InlineResponse200217) GetVlanProfileOk() (*InlineResponse200217VlanProfile, bool)`

GetVlanProfileOk returns a tuple with the VlanProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanProfile

`func (o *InlineResponse200217) SetVlanProfile(v InlineResponse200217VlanProfile)`

SetVlanProfile sets VlanProfile field to given value.

### HasVlanProfile

`func (o *InlineResponse200217) HasVlanProfile() bool`

HasVlanProfile returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200217) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200217) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200217) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200217) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetStackIds

`func (o *InlineResponse200217) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *InlineResponse200217) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *InlineResponse200217) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.

### HasStackIds

`func (o *InlineResponse200217) HasStackIds() bool`

HasStackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



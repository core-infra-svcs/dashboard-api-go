# InlineResponse200207

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanProfile** | Pointer to [**InlineResponse200207VlanProfile**](InlineResponse200207VlanProfile.md) |  | [optional] 
**Serials** | Pointer to **[]string** | Array of Device Serials | [optional] 
**StackIds** | Pointer to **[]string** | Array of Switch Stack IDs | [optional] 

## Methods

### NewInlineResponse200207

`func NewInlineResponse200207() *InlineResponse200207`

NewInlineResponse200207 instantiates a new InlineResponse200207 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200207WithDefaults

`func NewInlineResponse200207WithDefaults() *InlineResponse200207`

NewInlineResponse200207WithDefaults instantiates a new InlineResponse200207 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanProfile

`func (o *InlineResponse200207) GetVlanProfile() InlineResponse200207VlanProfile`

GetVlanProfile returns the VlanProfile field if non-nil, zero value otherwise.

### GetVlanProfileOk

`func (o *InlineResponse200207) GetVlanProfileOk() (*InlineResponse200207VlanProfile, bool)`

GetVlanProfileOk returns a tuple with the VlanProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanProfile

`func (o *InlineResponse200207) SetVlanProfile(v InlineResponse200207VlanProfile)`

SetVlanProfile sets VlanProfile field to given value.

### HasVlanProfile

`func (o *InlineResponse200207) HasVlanProfile() bool`

HasVlanProfile returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200207) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200207) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200207) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200207) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetStackIds

`func (o *InlineResponse200207) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *InlineResponse200207) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *InlineResponse200207) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.

### HasStackIds

`func (o *InlineResponse200207) HasStackIds() bool`

HasStackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



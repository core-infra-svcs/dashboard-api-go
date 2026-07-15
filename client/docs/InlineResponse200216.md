# InlineResponse200216

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanProfile** | Pointer to [**InlineResponse200216VlanProfile**](InlineResponse200216VlanProfile.md) |  | [optional] 
**Serials** | Pointer to **[]string** | Array of Device Serials | [optional] 
**StackIds** | Pointer to **[]string** | Array of Switch Stack IDs | [optional] 

## Methods

### NewInlineResponse200216

`func NewInlineResponse200216() *InlineResponse200216`

NewInlineResponse200216 instantiates a new InlineResponse200216 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200216WithDefaults

`func NewInlineResponse200216WithDefaults() *InlineResponse200216`

NewInlineResponse200216WithDefaults instantiates a new InlineResponse200216 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanProfile

`func (o *InlineResponse200216) GetVlanProfile() InlineResponse200216VlanProfile`

GetVlanProfile returns the VlanProfile field if non-nil, zero value otherwise.

### GetVlanProfileOk

`func (o *InlineResponse200216) GetVlanProfileOk() (*InlineResponse200216VlanProfile, bool)`

GetVlanProfileOk returns a tuple with the VlanProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanProfile

`func (o *InlineResponse200216) SetVlanProfile(v InlineResponse200216VlanProfile)`

SetVlanProfile sets VlanProfile field to given value.

### HasVlanProfile

`func (o *InlineResponse200216) HasVlanProfile() bool`

HasVlanProfile returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200216) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200216) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200216) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200216) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetStackIds

`func (o *InlineResponse200216) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *InlineResponse200216) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *InlineResponse200216) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.

### HasStackIds

`func (o *InlineResponse200216) HasStackIds() bool`

HasStackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



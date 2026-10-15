# InlineObject181

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the switch stack | [optional] 
**Members** | Pointer to [**[]NetworksNetworkIdSwitchStacksSwitchStackIdMembers**](NetworksNetworkIdSwitchStacksSwitchStackIdMembers.md) | The complete list of switches that should be in the stack. Minimum 2 and maximum 8 members. Omitting this field leaves stack membership unchanged. | [optional] 

## Methods

### NewInlineObject181

`func NewInlineObject181() *InlineObject181`

NewInlineObject181 instantiates a new InlineObject181 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject181WithDefaults

`func NewInlineObject181WithDefaults() *InlineObject181`

NewInlineObject181WithDefaults instantiates a new InlineObject181 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject181) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject181) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject181) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject181) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *InlineObject181) GetMembers() []NetworksNetworkIdSwitchStacksSwitchStackIdMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineObject181) GetMembersOk() (*[]NetworksNetworkIdSwitchStacksSwitchStackIdMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineObject181) SetMembers(v []NetworksNetworkIdSwitchStacksSwitchStackIdMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineObject181) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



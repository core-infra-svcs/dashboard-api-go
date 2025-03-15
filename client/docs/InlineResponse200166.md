# InlineResponse200166

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the Switch stack | [optional] 
**Name** | Pointer to **string** | Name of the Switch stack | [optional] 
**Serials** | Pointer to **[]string** | Serials of the switches in the switch stack | [optional] 
**IsMonitorOnly** | Pointer to **bool** | Tells if stack is Monitored Stack. | [optional] 
**Members** | Pointer to [**[]NetworksNetworkIdSwitchStacksMembers**](NetworksNetworkIdSwitchStacksMembers.md) | Members of the Stack | [optional] 

## Methods

### NewInlineResponse200166

`func NewInlineResponse200166() *InlineResponse200166`

NewInlineResponse200166 instantiates a new InlineResponse200166 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200166WithDefaults

`func NewInlineResponse200166WithDefaults() *InlineResponse200166`

NewInlineResponse200166WithDefaults instantiates a new InlineResponse200166 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200166) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200166) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200166) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200166) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200166) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200166) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200166) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200166) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200166) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200166) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200166) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200166) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetIsMonitorOnly

`func (o *InlineResponse200166) GetIsMonitorOnly() bool`

GetIsMonitorOnly returns the IsMonitorOnly field if non-nil, zero value otherwise.

### GetIsMonitorOnlyOk

`func (o *InlineResponse200166) GetIsMonitorOnlyOk() (*bool, bool)`

GetIsMonitorOnlyOk returns a tuple with the IsMonitorOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitorOnly

`func (o *InlineResponse200166) SetIsMonitorOnly(v bool)`

SetIsMonitorOnly sets IsMonitorOnly field to given value.

### HasIsMonitorOnly

`func (o *InlineResponse200166) HasIsMonitorOnly() bool`

HasIsMonitorOnly returns a boolean if a field has been set.

### GetMembers

`func (o *InlineResponse200166) GetMembers() []NetworksNetworkIdSwitchStacksMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineResponse200166) GetMembersOk() (*[]NetworksNetworkIdSwitchStacksMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineResponse200166) SetMembers(v []NetworksNetworkIdSwitchStacksMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineResponse200166) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



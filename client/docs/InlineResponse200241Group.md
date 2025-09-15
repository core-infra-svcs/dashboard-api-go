# InlineResponse200241Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | [optional] Represents the ordering of primary and backup tunnels group. primary and backup tunnels are grouped by this number. Peers containing same group number belongs to same group. | [optional] 
**Failover** | Pointer to [**InlineResponse200241GroupFailover**](InlineResponse200241GroupFailover.md) |  | [optional] 
**ActiveActiveTunnel** | Pointer to **bool** | [optional] Both primary and backup tunnels are active. | [optional] 

## Methods

### NewInlineResponse200241Group

`func NewInlineResponse200241Group() *InlineResponse200241Group`

NewInlineResponse200241Group instantiates a new InlineResponse200241Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200241GroupWithDefaults

`func NewInlineResponse200241GroupWithDefaults() *InlineResponse200241Group`

NewInlineResponse200241GroupWithDefaults instantiates a new InlineResponse200241Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse200241Group) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200241Group) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200241Group) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200241Group) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetFailover

`func (o *InlineResponse200241Group) GetFailover() InlineResponse200241GroupFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineResponse200241Group) GetFailoverOk() (*InlineResponse200241GroupFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineResponse200241Group) SetFailover(v InlineResponse200241GroupFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineResponse200241Group) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetActiveActiveTunnel

`func (o *InlineResponse200241Group) GetActiveActiveTunnel() bool`

GetActiveActiveTunnel returns the ActiveActiveTunnel field if non-nil, zero value otherwise.

### GetActiveActiveTunnelOk

`func (o *InlineResponse200241Group) GetActiveActiveTunnelOk() (*bool, bool)`

GetActiveActiveTunnelOk returns a tuple with the ActiveActiveTunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActiveTunnel

`func (o *InlineResponse200241Group) SetActiveActiveTunnel(v bool)`

SetActiveActiveTunnel sets ActiveActiveTunnel field to given value.

### HasActiveActiveTunnel

`func (o *InlineResponse200241Group) HasActiveActiveTunnel() bool`

HasActiveActiveTunnel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



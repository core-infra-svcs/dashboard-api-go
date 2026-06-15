# InlineResponse20118Devices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of the device | [optional] 
**MemberId** | Pointer to **string** | Member number assigned to the device within cluster - This is a read only attribute and cannot be altered using post/put operations | [optional] 
**Uplinks** | Pointer to [**[]InlineResponse20118Uplinks1**](InlineResponse20118Uplinks1.md) | Uplink settings of the device | [optional] 
**Tunnels** | Pointer to [**[]InlineResponse20118Tunnels1**](InlineResponse20118Tunnels1.md) | Tunnel settings of the device | [optional] 

## Methods

### NewInlineResponse20118Devices

`func NewInlineResponse20118Devices() *InlineResponse20118Devices`

NewInlineResponse20118Devices instantiates a new InlineResponse20118Devices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20118DevicesWithDefaults

`func NewInlineResponse20118DevicesWithDefaults() *InlineResponse20118Devices`

NewInlineResponse20118DevicesWithDefaults instantiates a new InlineResponse20118Devices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20118Devices) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20118Devices) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20118Devices) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20118Devices) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMemberId

`func (o *InlineResponse20118Devices) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *InlineResponse20118Devices) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *InlineResponse20118Devices) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *InlineResponse20118Devices) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse20118Devices) GetUplinks() []InlineResponse20118Uplinks1`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse20118Devices) GetUplinksOk() (*[]InlineResponse20118Uplinks1, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse20118Devices) SetUplinks(v []InlineResponse20118Uplinks1)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse20118Devices) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetTunnels

`func (o *InlineResponse20118Devices) GetTunnels() []InlineResponse20118Tunnels1`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *InlineResponse20118Devices) GetTunnelsOk() (*[]InlineResponse20118Tunnels1, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *InlineResponse20118Devices) SetTunnels(v []InlineResponse20118Tunnels1)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *InlineResponse20118Devices) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



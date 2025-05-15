# InlineResponse200258ClientVpnConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteIp** | Pointer to **string** | The IP address of the VPN the client last connected to | [optional] 
**ConnectedAt** | Pointer to **int32** | The time the client last connected to the VPN | [optional] 
**DisconnectedAt** | Pointer to **int32** | The time the client last disconnected from the VPN | [optional] 

## Methods

### NewInlineResponse200258ClientVpnConnections

`func NewInlineResponse200258ClientVpnConnections() *InlineResponse200258ClientVpnConnections`

NewInlineResponse200258ClientVpnConnections instantiates a new InlineResponse200258ClientVpnConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200258ClientVpnConnectionsWithDefaults

`func NewInlineResponse200258ClientVpnConnectionsWithDefaults() *InlineResponse200258ClientVpnConnections`

NewInlineResponse200258ClientVpnConnectionsWithDefaults instantiates a new InlineResponse200258ClientVpnConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteIp

`func (o *InlineResponse200258ClientVpnConnections) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *InlineResponse200258ClientVpnConnections) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *InlineResponse200258ClientVpnConnections) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *InlineResponse200258ClientVpnConnections) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetConnectedAt

`func (o *InlineResponse200258ClientVpnConnections) GetConnectedAt() int32`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *InlineResponse200258ClientVpnConnections) GetConnectedAtOk() (*int32, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *InlineResponse200258ClientVpnConnections) SetConnectedAt(v int32)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *InlineResponse200258ClientVpnConnections) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetDisconnectedAt

`func (o *InlineResponse200258ClientVpnConnections) GetDisconnectedAt() int32`

GetDisconnectedAt returns the DisconnectedAt field if non-nil, zero value otherwise.

### GetDisconnectedAtOk

`func (o *InlineResponse200258ClientVpnConnections) GetDisconnectedAtOk() (*int32, bool)`

GetDisconnectedAtOk returns a tuple with the DisconnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedAt

`func (o *InlineResponse200258ClientVpnConnections) SetDisconnectedAt(v int32)`

SetDisconnectedAt sets DisconnectedAt field to given value.

### HasDisconnectedAt

`func (o *InlineResponse200258ClientVpnConnections) HasDisconnectedAt() bool`

HasDisconnectedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineObject176

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceIp** | **string** | The IP address of the interface where the RP needs to be created. | 
**MulticastGroup** | **string** | &#39;Any&#39;, or the IP address of a multicast group | 
**Vrf** | Pointer to [**NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf**](NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf.md) |  | [optional] 

## Methods

### NewInlineObject176

`func NewInlineObject176(interfaceIp string, multicastGroup string, ) *InlineObject176`

NewInlineObject176 instantiates a new InlineObject176 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject176WithDefaults

`func NewInlineObject176WithDefaults() *InlineObject176`

NewInlineObject176WithDefaults instantiates a new InlineObject176 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceIp

`func (o *InlineObject176) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineObject176) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineObject176) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.


### GetMulticastGroup

`func (o *InlineObject176) GetMulticastGroup() string`

GetMulticastGroup returns the MulticastGroup field if non-nil, zero value otherwise.

### GetMulticastGroupOk

`func (o *InlineObject176) GetMulticastGroupOk() (*string, bool)`

GetMulticastGroupOk returns a tuple with the MulticastGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastGroup

`func (o *InlineObject176) SetMulticastGroup(v string)`

SetMulticastGroup sets MulticastGroup field to given value.


### GetVrf

`func (o *InlineObject176) GetVrf() NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject176) GetVrfOk() (*NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject176) SetVrf(v NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject176) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



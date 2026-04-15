# InlineObject158

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceIp** | **string** | The IP address of the interface where the RP needs to be created. | 
**MulticastGroup** | **string** | &#39;Any&#39;, or the IP address of a multicast group | 
**Vrf** | Pointer to [**NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf**](NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf.md) |  | [optional] 

## Methods

### NewInlineObject158

`func NewInlineObject158(interfaceIp string, multicastGroup string, ) *InlineObject158`

NewInlineObject158 instantiates a new InlineObject158 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject158WithDefaults

`func NewInlineObject158WithDefaults() *InlineObject158`

NewInlineObject158WithDefaults instantiates a new InlineObject158 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceIp

`func (o *InlineObject158) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineObject158) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineObject158) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.


### GetMulticastGroup

`func (o *InlineObject158) GetMulticastGroup() string`

GetMulticastGroup returns the MulticastGroup field if non-nil, zero value otherwise.

### GetMulticastGroupOk

`func (o *InlineObject158) GetMulticastGroupOk() (*string, bool)`

GetMulticastGroupOk returns a tuple with the MulticastGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastGroup

`func (o *InlineObject158) SetMulticastGroup(v string)`

SetMulticastGroup sets MulticastGroup field to given value.


### GetVrf

`func (o *InlineObject158) GetVrf() NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject158) GetVrfOk() (*NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject158) SetVrf(v NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject158) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



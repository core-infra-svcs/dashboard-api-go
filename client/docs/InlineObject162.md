# InlineObject162

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceIp** | **string** | The IP address of the interface where the RP needs to be created. | 
**MulticastGroup** | **string** | &#39;Any&#39;, or the IP address of a multicast group | 
**Vrf** | Pointer to [**NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf**](NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf.md) |  | [optional] 

## Methods

### NewInlineObject162

`func NewInlineObject162(interfaceIp string, multicastGroup string, ) *InlineObject162`

NewInlineObject162 instantiates a new InlineObject162 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject162WithDefaults

`func NewInlineObject162WithDefaults() *InlineObject162`

NewInlineObject162WithDefaults instantiates a new InlineObject162 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceIp

`func (o *InlineObject162) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineObject162) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineObject162) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.


### GetMulticastGroup

`func (o *InlineObject162) GetMulticastGroup() string`

GetMulticastGroup returns the MulticastGroup field if non-nil, zero value otherwise.

### GetMulticastGroupOk

`func (o *InlineObject162) GetMulticastGroupOk() (*string, bool)`

GetMulticastGroupOk returns a tuple with the MulticastGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastGroup

`func (o *InlineObject162) SetMulticastGroup(v string)`

SetMulticastGroup sets MulticastGroup field to given value.


### GetVrf

`func (o *InlineObject162) GetVrf() NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject162) GetVrfOk() (*NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject162) SetVrf(v NetworksNetworkIdSwitchRoutingMulticastRendezvousPointsVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject162) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



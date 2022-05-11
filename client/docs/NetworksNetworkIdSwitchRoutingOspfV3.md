# NetworksNetworkIdSwitchRoutingOspfV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default. | [optional] 
**HelloTimerInSeconds** | Pointer to **int32** | Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds. | [optional] 
**DeadTimerInSeconds** | Pointer to **int32** | Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535 | [optional] 
**Areas** | Pointer to [**[]NetworksNetworkIdSwitchRoutingOspfAreas**](NetworksNetworkIdSwitchRoutingOspfAreas.md) | OSPF v3 areas | [optional] 

## Methods

### NewNetworksNetworkIdSwitchRoutingOspfV3

`func NewNetworksNetworkIdSwitchRoutingOspfV3() *NetworksNetworkIdSwitchRoutingOspfV3`

NewNetworksNetworkIdSwitchRoutingOspfV3 instantiates a new NetworksNetworkIdSwitchRoutingOspfV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchRoutingOspfV3WithDefaults

`func NewNetworksNetworkIdSwitchRoutingOspfV3WithDefaults() *NetworksNetworkIdSwitchRoutingOspfV3`

NewNetworksNetworkIdSwitchRoutingOspfV3WithDefaults instantiates a new NetworksNetworkIdSwitchRoutingOspfV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHelloTimerInSeconds

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetHelloTimerInSeconds() int32`

GetHelloTimerInSeconds returns the HelloTimerInSeconds field if non-nil, zero value otherwise.

### GetHelloTimerInSecondsOk

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetHelloTimerInSecondsOk() (*int32, bool)`

GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTimerInSeconds

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetHelloTimerInSeconds(v int32)`

SetHelloTimerInSeconds sets HelloTimerInSeconds field to given value.

### HasHelloTimerInSeconds

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasHelloTimerInSeconds() bool`

HasHelloTimerInSeconds returns a boolean if a field has been set.

### GetDeadTimerInSeconds

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetDeadTimerInSeconds() int32`

GetDeadTimerInSeconds returns the DeadTimerInSeconds field if non-nil, zero value otherwise.

### GetDeadTimerInSecondsOk

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetDeadTimerInSecondsOk() (*int32, bool)`

GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTimerInSeconds

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetDeadTimerInSeconds(v int32)`

SetDeadTimerInSeconds sets DeadTimerInSeconds field to given value.

### HasDeadTimerInSeconds

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasDeadTimerInSeconds() bool`

HasDeadTimerInSeconds returns a boolean if a field has been set.

### GetAreas

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetAreas() []NetworksNetworkIdSwitchRoutingOspfAreas`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) GetAreasOk() (*[]NetworksNetworkIdSwitchRoutingOspfAreas, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) SetAreas(v []NetworksNetworkIdSwitchRoutingOspfAreas)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *NetworksNetworkIdSwitchRoutingOspfV3) HasAreas() bool`

HasAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



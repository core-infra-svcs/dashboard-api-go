# InlineObject32

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name or description for layer 3 static route | [optional] 
**Subnet** | **string** | The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24) | 
**NextHopIp** | **string** | IP address of the next hop device to which the device sends its traffic for the subnet | 
**AdvertiseViaOspfEnabled** | Pointer to **bool** | Option to advertise static route via OSPF | [optional] 
**PreferOverOspfRoutesEnabled** | Pointer to **bool** | Option to prefer static route over OSPF routes | [optional] 
**Vrf** | Pointer to [**DevicesSerialSwitchRoutingStaticRoutesVrf1**](DevicesSerialSwitchRoutingStaticRoutesVrf1.md) |  | [optional] 

## Methods

### NewInlineObject32

`func NewInlineObject32(subnet string, nextHopIp string, ) *InlineObject32`

NewInlineObject32 instantiates a new InlineObject32 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject32WithDefaults

`func NewInlineObject32WithDefaults() *InlineObject32`

NewInlineObject32WithDefaults instantiates a new InlineObject32 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject32) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject32) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject32) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject32) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject32) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject32) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject32) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetNextHopIp

`func (o *InlineObject32) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineObject32) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineObject32) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.


### GetAdvertiseViaOspfEnabled

`func (o *InlineObject32) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *InlineObject32) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *InlineObject32) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *InlineObject32) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *InlineObject32) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *InlineObject32) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *InlineObject32) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *InlineObject32) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject32) GetVrf() DevicesSerialSwitchRoutingStaticRoutesVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject32) GetVrfOk() (*DevicesSerialSwitchRoutingStaticRoutesVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject32) SetVrf(v DevicesSerialSwitchRoutingStaticRoutesVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject32) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



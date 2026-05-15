# InlineObject172

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name or description for layer 3 static route | [optional] 
**Subnet** | Pointer to **string** | The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24) | [optional] 
**NextHopIp** | Pointer to **string** | IP address of the next hop device to which the device sends its traffic for the subnet | [optional] 
**ManagementNextHop** | Pointer to **NullableString** | Optional fallback IP address for management traffic | [optional] 
**AdvertiseViaOspfEnabled** | Pointer to **bool** | Option to advertise static route via OSPF | [optional] 
**PreferOverOspfRoutesEnabled** | Pointer to **bool** | Option to prefer static route over OSPF routes | [optional] 
**Vrf** | Pointer to [**DevicesSerialSwitchRoutingStaticRoutesVrf1**](DevicesSerialSwitchRoutingStaticRoutesVrf1.md) |  | [optional] 

## Methods

### NewInlineObject172

`func NewInlineObject172() *InlineObject172`

NewInlineObject172 instantiates a new InlineObject172 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject172WithDefaults

`func NewInlineObject172WithDefaults() *InlineObject172`

NewInlineObject172WithDefaults instantiates a new InlineObject172 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject172) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject172) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject172) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject172) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject172) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject172) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject172) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject172) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNextHopIp

`func (o *InlineObject172) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineObject172) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineObject172) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *InlineObject172) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetManagementNextHop

`func (o *InlineObject172) GetManagementNextHop() string`

GetManagementNextHop returns the ManagementNextHop field if non-nil, zero value otherwise.

### GetManagementNextHopOk

`func (o *InlineObject172) GetManagementNextHopOk() (*string, bool)`

GetManagementNextHopOk returns a tuple with the ManagementNextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNextHop

`func (o *InlineObject172) SetManagementNextHop(v string)`

SetManagementNextHop sets ManagementNextHop field to given value.

### HasManagementNextHop

`func (o *InlineObject172) HasManagementNextHop() bool`

HasManagementNextHop returns a boolean if a field has been set.

### SetManagementNextHopNil

`func (o *InlineObject172) SetManagementNextHopNil(b bool)`

 SetManagementNextHopNil sets the value for ManagementNextHop to be an explicit nil

### UnsetManagementNextHop
`func (o *InlineObject172) UnsetManagementNextHop()`

UnsetManagementNextHop ensures that no value is present for ManagementNextHop, not even an explicit nil
### GetAdvertiseViaOspfEnabled

`func (o *InlineObject172) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *InlineObject172) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *InlineObject172) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *InlineObject172) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *InlineObject172) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *InlineObject172) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *InlineObject172) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *InlineObject172) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject172) GetVrf() DevicesSerialSwitchRoutingStaticRoutesVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject172) GetVrfOk() (*DevicesSerialSwitchRoutingStaticRoutesVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject172) SetVrf(v DevicesSerialSwitchRoutingStaticRoutesVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject172) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



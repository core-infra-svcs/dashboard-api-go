# InlineObject33

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

### NewInlineObject33

`func NewInlineObject33() *InlineObject33`

NewInlineObject33 instantiates a new InlineObject33 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject33WithDefaults

`func NewInlineObject33WithDefaults() *InlineObject33`

NewInlineObject33WithDefaults instantiates a new InlineObject33 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject33) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject33) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject33) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject33) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject33) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject33) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject33) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject33) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNextHopIp

`func (o *InlineObject33) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineObject33) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineObject33) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *InlineObject33) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetManagementNextHop

`func (o *InlineObject33) GetManagementNextHop() string`

GetManagementNextHop returns the ManagementNextHop field if non-nil, zero value otherwise.

### GetManagementNextHopOk

`func (o *InlineObject33) GetManagementNextHopOk() (*string, bool)`

GetManagementNextHopOk returns a tuple with the ManagementNextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNextHop

`func (o *InlineObject33) SetManagementNextHop(v string)`

SetManagementNextHop sets ManagementNextHop field to given value.

### HasManagementNextHop

`func (o *InlineObject33) HasManagementNextHop() bool`

HasManagementNextHop returns a boolean if a field has been set.

### SetManagementNextHopNil

`func (o *InlineObject33) SetManagementNextHopNil(b bool)`

 SetManagementNextHopNil sets the value for ManagementNextHop to be an explicit nil

### UnsetManagementNextHop
`func (o *InlineObject33) UnsetManagementNextHop()`

UnsetManagementNextHop ensures that no value is present for ManagementNextHop, not even an explicit nil
### GetAdvertiseViaOspfEnabled

`func (o *InlineObject33) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *InlineObject33) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *InlineObject33) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *InlineObject33) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *InlineObject33) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *InlineObject33) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *InlineObject33) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *InlineObject33) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject33) GetVrf() DevicesSerialSwitchRoutingStaticRoutesVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject33) GetVrfOk() (*DevicesSerialSwitchRoutingStaticRoutesVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject33) SetVrf(v DevicesSerialSwitchRoutingStaticRoutesVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject33) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



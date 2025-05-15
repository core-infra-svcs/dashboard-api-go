# InlineResponse20041

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticRouteId** | **string** | The identifier of a layer 3 static route | 
**Name** | Pointer to **string** | The name or description of the layer 3 static route | [optional] 
**Subnet** | **string** | The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24) | 
**NextHopIp** | **string** | The IP address of the router to which traffic for this destination network should be sent | 
**ManagementNextHop** | Pointer to **string** | Optional fallback IP address for management traffic | [optional] 
**AdvertiseViaOspfEnabled** | Pointer to **bool** | Option to advertise static routes via OSPF | [optional] 
**PreferOverOspfRoutesEnabled** | Pointer to **bool** | Option to prefer static routes over OSPF routes | [optional] 

## Methods

### NewInlineResponse20041

`func NewInlineResponse20041(staticRouteId string, subnet string, nextHopIp string, ) *InlineResponse20041`

NewInlineResponse20041 instantiates a new InlineResponse20041 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20041WithDefaults

`func NewInlineResponse20041WithDefaults() *InlineResponse20041`

NewInlineResponse20041WithDefaults instantiates a new InlineResponse20041 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticRouteId

`func (o *InlineResponse20041) GetStaticRouteId() string`

GetStaticRouteId returns the StaticRouteId field if non-nil, zero value otherwise.

### GetStaticRouteIdOk

`func (o *InlineResponse20041) GetStaticRouteIdOk() (*string, bool)`

GetStaticRouteIdOk returns a tuple with the StaticRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRouteId

`func (o *InlineResponse20041) SetStaticRouteId(v string)`

SetStaticRouteId sets StaticRouteId field to given value.


### GetName

`func (o *InlineResponse20041) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20041) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20041) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20041) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20041) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20041) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20041) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetNextHopIp

`func (o *InlineResponse20041) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *InlineResponse20041) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *InlineResponse20041) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.


### GetManagementNextHop

`func (o *InlineResponse20041) GetManagementNextHop() string`

GetManagementNextHop returns the ManagementNextHop field if non-nil, zero value otherwise.

### GetManagementNextHopOk

`func (o *InlineResponse20041) GetManagementNextHopOk() (*string, bool)`

GetManagementNextHopOk returns a tuple with the ManagementNextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNextHop

`func (o *InlineResponse20041) SetManagementNextHop(v string)`

SetManagementNextHop sets ManagementNextHop field to given value.

### HasManagementNextHop

`func (o *InlineResponse20041) HasManagementNextHop() bool`

HasManagementNextHop returns a boolean if a field has been set.

### GetAdvertiseViaOspfEnabled

`func (o *InlineResponse20041) GetAdvertiseViaOspfEnabled() bool`

GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field if non-nil, zero value otherwise.

### GetAdvertiseViaOspfEnabledOk

`func (o *InlineResponse20041) GetAdvertiseViaOspfEnabledOk() (*bool, bool)`

GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseViaOspfEnabled

`func (o *InlineResponse20041) SetAdvertiseViaOspfEnabled(v bool)`

SetAdvertiseViaOspfEnabled sets AdvertiseViaOspfEnabled field to given value.

### HasAdvertiseViaOspfEnabled

`func (o *InlineResponse20041) HasAdvertiseViaOspfEnabled() bool`

HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.

### GetPreferOverOspfRoutesEnabled

`func (o *InlineResponse20041) GetPreferOverOspfRoutesEnabled() bool`

GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field if non-nil, zero value otherwise.

### GetPreferOverOspfRoutesEnabledOk

`func (o *InlineResponse20041) GetPreferOverOspfRoutesEnabledOk() (*bool, bool)`

GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferOverOspfRoutesEnabled

`func (o *InlineResponse20041) SetPreferOverOspfRoutesEnabled(v bool)`

SetPreferOverOspfRoutesEnabled sets PreferOverOspfRoutesEnabled field to given value.

### HasPreferOverOspfRoutesEnabled

`func (o *InlineResponse20041) HasPreferOverOspfRoutesEnabled() bool`

HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



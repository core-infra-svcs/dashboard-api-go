# InlineResponse20013InterfacesWan1Transition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the IPoE transition IPv4 to IPv6 is enabled. | [optional] 
**Mode** | Pointer to **string** | Mode selector for the IPoE transition IPv4 to IPv6 is enabled. | [optional] 
**Authentication** | Pointer to [**InlineResponse20013InterfacesWan1TransitionAuthentication**](InlineResponse20013InterfacesWan1TransitionAuthentication.md) |  | [optional] 
**Ipv4** | Pointer to [**InlineResponse20013InterfacesWan1TransitionIpv4**](InlineResponse20013InterfacesWan1TransitionIpv4.md) |  | [optional] 
**Ipv6** | Pointer to [**InlineResponse20013InterfacesWan1TransitionIpv6**](InlineResponse20013InterfacesWan1TransitionIpv6.md) |  | [optional] 
**Isp** | Pointer to [**InlineResponse20013InterfacesWan1TransitionIsp**](InlineResponse20013InterfacesWan1TransitionIsp.md) |  | [optional] 
**Map** | Pointer to [**InlineResponse20013InterfacesWan1TransitionMap**](InlineResponse20013InterfacesWan1TransitionMap.md) |  | [optional] 

## Methods

### NewInlineResponse20013InterfacesWan1Transition

`func NewInlineResponse20013InterfacesWan1Transition() *InlineResponse20013InterfacesWan1Transition`

NewInlineResponse20013InterfacesWan1Transition instantiates a new InlineResponse20013InterfacesWan1Transition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013InterfacesWan1TransitionWithDefaults

`func NewInlineResponse20013InterfacesWan1TransitionWithDefaults() *InlineResponse20013InterfacesWan1Transition`

NewInlineResponse20013InterfacesWan1TransitionWithDefaults instantiates a new InlineResponse20013InterfacesWan1Transition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20013InterfacesWan1Transition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20013InterfacesWan1Transition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20013InterfacesWan1Transition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20013InterfacesWan1Transition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse20013InterfacesWan1Transition) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20013InterfacesWan1Transition) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20013InterfacesWan1Transition) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20013InterfacesWan1Transition) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse20013InterfacesWan1Transition) GetAuthentication() InlineResponse20013InterfacesWan1TransitionAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse20013InterfacesWan1Transition) GetAuthenticationOk() (*InlineResponse20013InterfacesWan1TransitionAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse20013InterfacesWan1Transition) SetAuthentication(v InlineResponse20013InterfacesWan1TransitionAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse20013InterfacesWan1Transition) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetIpv4

`func (o *InlineResponse20013InterfacesWan1Transition) GetIpv4() InlineResponse20013InterfacesWan1TransitionIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineResponse20013InterfacesWan1Transition) GetIpv4Ok() (*InlineResponse20013InterfacesWan1TransitionIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineResponse20013InterfacesWan1Transition) SetIpv4(v InlineResponse20013InterfacesWan1TransitionIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *InlineResponse20013InterfacesWan1Transition) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20013InterfacesWan1Transition) GetIpv6() InlineResponse20013InterfacesWan1TransitionIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20013InterfacesWan1Transition) GetIpv6Ok() (*InlineResponse20013InterfacesWan1TransitionIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20013InterfacesWan1Transition) SetIpv6(v InlineResponse20013InterfacesWan1TransitionIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20013InterfacesWan1Transition) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIsp

`func (o *InlineResponse20013InterfacesWan1Transition) GetIsp() InlineResponse20013InterfacesWan1TransitionIsp`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *InlineResponse20013InterfacesWan1Transition) GetIspOk() (*InlineResponse20013InterfacesWan1TransitionIsp, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *InlineResponse20013InterfacesWan1Transition) SetIsp(v InlineResponse20013InterfacesWan1TransitionIsp)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *InlineResponse20013InterfacesWan1Transition) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetMap

`func (o *InlineResponse20013InterfacesWan1Transition) GetMap() InlineResponse20013InterfacesWan1TransitionMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *InlineResponse20013InterfacesWan1Transition) GetMapOk() (*InlineResponse20013InterfacesWan1TransitionMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *InlineResponse20013InterfacesWan1Transition) SetMap(v InlineResponse20013InterfacesWan1TransitionMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *InlineResponse20013InterfacesWan1Transition) HasMap() bool`

HasMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



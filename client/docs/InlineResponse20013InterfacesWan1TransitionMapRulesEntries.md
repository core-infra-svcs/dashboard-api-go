# InlineResponse20013InterfacesWan1TransitionMapRulesEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv6** | Pointer to [**InlineResponse20013InterfacesWan1TransitionMapRulesIpv6**](InlineResponse20013InterfacesWan1TransitionMapRulesIpv6.md) |  | [optional] 
**Ipv4** | Pointer to [**InlineResponse20013InterfacesWan1TransitionMapRulesIpv4**](InlineResponse20013InterfacesWan1TransitionMapRulesIpv4.md) |  | [optional] 
**BorderRelay** | Pointer to [**InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay**](InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay.md) |  | [optional] 
**EaBitLength** | Pointer to **int32** | Embedded address bit length (0-48). | [optional] 
**PsIdOffset** | Pointer to **int32** | PSID offset (0-15). | [optional] 

## Methods

### NewInlineResponse20013InterfacesWan1TransitionMapRulesEntries

`func NewInlineResponse20013InterfacesWan1TransitionMapRulesEntries() *InlineResponse20013InterfacesWan1TransitionMapRulesEntries`

NewInlineResponse20013InterfacesWan1TransitionMapRulesEntries instantiates a new InlineResponse20013InterfacesWan1TransitionMapRulesEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013InterfacesWan1TransitionMapRulesEntriesWithDefaults

`func NewInlineResponse20013InterfacesWan1TransitionMapRulesEntriesWithDefaults() *InlineResponse20013InterfacesWan1TransitionMapRulesEntries`

NewInlineResponse20013InterfacesWan1TransitionMapRulesEntriesWithDefaults instantiates a new InlineResponse20013InterfacesWan1TransitionMapRulesEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv6

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetIpv6() InlineResponse20013InterfacesWan1TransitionMapRulesIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetIpv6Ok() (*InlineResponse20013InterfacesWan1TransitionMapRulesIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) SetIpv6(v InlineResponse20013InterfacesWan1TransitionMapRulesIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetIpv4

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetIpv4() InlineResponse20013InterfacesWan1TransitionMapRulesIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetIpv4Ok() (*InlineResponse20013InterfacesWan1TransitionMapRulesIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) SetIpv4(v InlineResponse20013InterfacesWan1TransitionMapRulesIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetBorderRelay

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetBorderRelay() InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay`

GetBorderRelay returns the BorderRelay field if non-nil, zero value otherwise.

### GetBorderRelayOk

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetBorderRelayOk() (*InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay, bool)`

GetBorderRelayOk returns a tuple with the BorderRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderRelay

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) SetBorderRelay(v InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay)`

SetBorderRelay sets BorderRelay field to given value.

### HasBorderRelay

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) HasBorderRelay() bool`

HasBorderRelay returns a boolean if a field has been set.

### GetEaBitLength

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetEaBitLength() int32`

GetEaBitLength returns the EaBitLength field if non-nil, zero value otherwise.

### GetEaBitLengthOk

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetEaBitLengthOk() (*int32, bool)`

GetEaBitLengthOk returns a tuple with the EaBitLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaBitLength

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) SetEaBitLength(v int32)`

SetEaBitLength sets EaBitLength field to given value.

### HasEaBitLength

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) HasEaBitLength() bool`

HasEaBitLength returns a boolean if a field has been set.

### GetPsIdOffset

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetPsIdOffset() int32`

GetPsIdOffset returns the PsIdOffset field if non-nil, zero value otherwise.

### GetPsIdOffsetOk

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) GetPsIdOffsetOk() (*int32, bool)`

GetPsIdOffsetOk returns a tuple with the PsIdOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsIdOffset

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) SetPsIdOffset(v int32)`

SetPsIdOffset sets PsIdOffset field to given value.

### HasPsIdOffset

`func (o *InlineResponse20013InterfacesWan1TransitionMapRulesEntries) HasPsIdOffset() bool`

HasPsIdOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



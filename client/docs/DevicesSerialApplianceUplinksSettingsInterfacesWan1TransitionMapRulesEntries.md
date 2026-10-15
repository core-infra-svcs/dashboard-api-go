# DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv6** | [**DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv6**](DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv6.md) |  | 
**Ipv4** | [**DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv4**](DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv4.md) |  | 
**BorderRelay** | [**InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay**](InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay.md) |  | 
**EaBitLength** | **int32** | Embedded address bit length (0-48). | 
**PsIdOffset** | **int32** | PSID offset (0-15). | 

## Methods

### NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries

`func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries(ipv6 DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv6, ipv4 DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv4, borderRelay InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay, eaBitLength int32, psIdOffset int32, ) *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries`

NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntriesWithDefaults

`func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntriesWithDefaults() *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries`

NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntriesWithDefaults instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv6

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetIpv6() DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetIpv6Ok() (*DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) SetIpv6(v DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv6)`

SetIpv6 sets Ipv6 field to given value.


### GetIpv4

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetIpv4() DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetIpv4Ok() (*DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) SetIpv4(v DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesIpv4)`

SetIpv4 sets Ipv4 field to given value.


### GetBorderRelay

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetBorderRelay() InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay`

GetBorderRelay returns the BorderRelay field if non-nil, zero value otherwise.

### GetBorderRelayOk

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetBorderRelayOk() (*InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay, bool)`

GetBorderRelayOk returns a tuple with the BorderRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderRelay

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) SetBorderRelay(v InlineResponse20013InterfacesWan1TransitionMapRulesBorderRelay)`

SetBorderRelay sets BorderRelay field to given value.


### GetEaBitLength

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetEaBitLength() int32`

GetEaBitLength returns the EaBitLength field if non-nil, zero value otherwise.

### GetEaBitLengthOk

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetEaBitLengthOk() (*int32, bool)`

GetEaBitLengthOk returns a tuple with the EaBitLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaBitLength

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) SetEaBitLength(v int32)`

SetEaBitLength sets EaBitLength field to given value.


### GetPsIdOffset

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetPsIdOffset() int32`

GetPsIdOffset returns the PsIdOffset field if non-nil, zero value otherwise.

### GetPsIdOffsetOk

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) GetPsIdOffsetOk() (*int32, bool)`

GetPsIdOffsetOk returns a tuple with the PsIdOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsIdOffset

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries) SetPsIdOffset(v int32)`

SetPsIdOffset sets PsIdOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



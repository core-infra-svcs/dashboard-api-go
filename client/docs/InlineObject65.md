# InlineObject65

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the SSID. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the SSID is enabled. | [optional] 
**DefaultVlanId** | Pointer to **int32** | The VLAN ID of the VLAN associated to this SSID. This parameter is only valid if the network is in routed mode. | [optional] 
**AuthMode** | Pointer to **string** | The association control method for the SSID (&#39;open&#39;, &#39;psk&#39;, &#39;8021x-meraki&#39; or &#39;8021x-radius&#39;). | [optional] 
**Psk** | Pointer to **string** | The passkey for the SSID. This param is only valid if the authMode is &#39;psk&#39;. | [optional] 
**RadiusServers** | Pointer to [**[]NetworksNetworkIdApplianceSsidsNumberRadiusServers**](NetworksNetworkIdApplianceSsidsNumberRadiusServers.md) | The RADIUS 802.1x servers to be used for authentication. This param is only valid if the authMode is &#39;8021x-radius&#39;. | [optional] 
**EncryptionMode** | Pointer to **string** | The psk encryption mode for the SSID (&#39;wep&#39; or &#39;wpa&#39;). This param is only valid if the authMode is &#39;psk&#39;. | [optional] 
**WpaEncryptionMode** | Pointer to **string** | The types of WPA encryption. (&#39;WPA1 and WPA2&#39;, &#39;WPA2 only&#39;, &#39;WPA3 Transition Mode&#39; or &#39;WPA3 only&#39;). This param is only valid if (1) the authMode is &#39;psk&#39; &amp; the encryptionMode is &#39;wpa&#39; OR (2) the authMode is &#39;8021x-meraki&#39; OR (3) the authMode is &#39;8021x-radius&#39; | [optional] 
**Visible** | Pointer to **bool** | Boolean indicating whether the MX should advertise or hide this SSID. | [optional] 
**DhcpEnforcedDeauthentication** | Pointer to [**NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication**](NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication.md) |  | [optional] 
**Dot11w** | Pointer to [**NetworksNetworkIdApplianceSsidsNumberDot11w**](NetworksNetworkIdApplianceSsidsNumberDot11w.md) |  | [optional] 

## Methods

### NewInlineObject65

`func NewInlineObject65() *InlineObject65`

NewInlineObject65 instantiates a new InlineObject65 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject65WithDefaults

`func NewInlineObject65WithDefaults() *InlineObject65`

NewInlineObject65WithDefaults instantiates a new InlineObject65 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject65) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject65) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject65) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject65) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineObject65) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject65) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject65) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject65) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultVlanId

`func (o *InlineObject65) GetDefaultVlanId() int32`

GetDefaultVlanId returns the DefaultVlanId field if non-nil, zero value otherwise.

### GetDefaultVlanIdOk

`func (o *InlineObject65) GetDefaultVlanIdOk() (*int32, bool)`

GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlanId

`func (o *InlineObject65) SetDefaultVlanId(v int32)`

SetDefaultVlanId sets DefaultVlanId field to given value.

### HasDefaultVlanId

`func (o *InlineObject65) HasDefaultVlanId() bool`

HasDefaultVlanId returns a boolean if a field has been set.

### GetAuthMode

`func (o *InlineObject65) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InlineObject65) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InlineObject65) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *InlineObject65) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetPsk

`func (o *InlineObject65) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *InlineObject65) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *InlineObject65) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *InlineObject65) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### GetRadiusServers

`func (o *InlineObject65) GetRadiusServers() []NetworksNetworkIdApplianceSsidsNumberRadiusServers`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *InlineObject65) GetRadiusServersOk() (*[]NetworksNetworkIdApplianceSsidsNumberRadiusServers, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *InlineObject65) SetRadiusServers(v []NetworksNetworkIdApplianceSsidsNumberRadiusServers)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *InlineObject65) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.

### GetEncryptionMode

`func (o *InlineObject65) GetEncryptionMode() string`

GetEncryptionMode returns the EncryptionMode field if non-nil, zero value otherwise.

### GetEncryptionModeOk

`func (o *InlineObject65) GetEncryptionModeOk() (*string, bool)`

GetEncryptionModeOk returns a tuple with the EncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionMode

`func (o *InlineObject65) SetEncryptionMode(v string)`

SetEncryptionMode sets EncryptionMode field to given value.

### HasEncryptionMode

`func (o *InlineObject65) HasEncryptionMode() bool`

HasEncryptionMode returns a boolean if a field has been set.

### GetWpaEncryptionMode

`func (o *InlineObject65) GetWpaEncryptionMode() string`

GetWpaEncryptionMode returns the WpaEncryptionMode field if non-nil, zero value otherwise.

### GetWpaEncryptionModeOk

`func (o *InlineObject65) GetWpaEncryptionModeOk() (*string, bool)`

GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpaEncryptionMode

`func (o *InlineObject65) SetWpaEncryptionMode(v string)`

SetWpaEncryptionMode sets WpaEncryptionMode field to given value.

### HasWpaEncryptionMode

`func (o *InlineObject65) HasWpaEncryptionMode() bool`

HasWpaEncryptionMode returns a boolean if a field has been set.

### GetVisible

`func (o *InlineObject65) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineObject65) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineObject65) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InlineObject65) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetDhcpEnforcedDeauthentication

`func (o *InlineObject65) GetDhcpEnforcedDeauthentication() NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication`

GetDhcpEnforcedDeauthentication returns the DhcpEnforcedDeauthentication field if non-nil, zero value otherwise.

### GetDhcpEnforcedDeauthenticationOk

`func (o *InlineObject65) GetDhcpEnforcedDeauthenticationOk() (*NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication, bool)`

GetDhcpEnforcedDeauthenticationOk returns a tuple with the DhcpEnforcedDeauthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnforcedDeauthentication

`func (o *InlineObject65) SetDhcpEnforcedDeauthentication(v NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication)`

SetDhcpEnforcedDeauthentication sets DhcpEnforcedDeauthentication field to given value.

### HasDhcpEnforcedDeauthentication

`func (o *InlineObject65) HasDhcpEnforcedDeauthentication() bool`

HasDhcpEnforcedDeauthentication returns a boolean if a field has been set.

### GetDot11w

`func (o *InlineObject65) GetDot11w() NetworksNetworkIdApplianceSsidsNumberDot11w`

GetDot11w returns the Dot11w field if non-nil, zero value otherwise.

### GetDot11wOk

`func (o *InlineObject65) GetDot11wOk() (*NetworksNetworkIdApplianceSsidsNumberDot11w, bool)`

GetDot11wOk returns a tuple with the Dot11w field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11w

`func (o *InlineObject65) SetDot11w(v NetworksNetworkIdApplianceSsidsNumberDot11w)`

SetDot11w sets Dot11w field to given value.

### HasDot11w

`func (o *InlineObject65) HasDot11w() bool`

HasDot11w returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



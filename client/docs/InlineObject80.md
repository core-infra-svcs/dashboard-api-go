# InlineObject80

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStatusPageEnabled** | Pointer to **bool** | Enables / disables the local device status pages (&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://my.meraki.com/&#39;&gt;my.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://ap.meraki.com/&#39;&gt;ap.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://switch.meraki.com/&#39;&gt;switch.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://wired.meraki.com/&#39;&gt;wired.meraki.com&lt;/a&gt;). Optional (defaults to false) | [optional] 
**RemoteStatusPageEnabled** | Pointer to **bool** | Enables / disables access to the device status page (&lt;a target&#x3D;&#39;_blank&#39;&gt;http://[device&#39;s LAN IP])&lt;/a&gt;. Optional. Can only be set if localStatusPageEnabled is set to true | [optional] 
**SecureConnect** | Pointer to [**NetworksNetworkIdSettingsSecureConnect**](NetworksNetworkIdSettingsSecureConnect.md) |  | [optional] 
**LocalStatusPage** | Pointer to [**NetworksNetworkIdSettingsLocalStatusPage**](NetworksNetworkIdSettingsLocalStatusPage.md) |  | [optional] 

## Methods

### NewInlineObject80

`func NewInlineObject80() *InlineObject80`

NewInlineObject80 instantiates a new InlineObject80 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject80WithDefaults

`func NewInlineObject80WithDefaults() *InlineObject80`

NewInlineObject80WithDefaults instantiates a new InlineObject80 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStatusPageEnabled

`func (o *InlineObject80) GetLocalStatusPageEnabled() bool`

GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field if non-nil, zero value otherwise.

### GetLocalStatusPageEnabledOk

`func (o *InlineObject80) GetLocalStatusPageEnabledOk() (*bool, bool)`

GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPageEnabled

`func (o *InlineObject80) SetLocalStatusPageEnabled(v bool)`

SetLocalStatusPageEnabled sets LocalStatusPageEnabled field to given value.

### HasLocalStatusPageEnabled

`func (o *InlineObject80) HasLocalStatusPageEnabled() bool`

HasLocalStatusPageEnabled returns a boolean if a field has been set.

### GetRemoteStatusPageEnabled

`func (o *InlineObject80) GetRemoteStatusPageEnabled() bool`

GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field if non-nil, zero value otherwise.

### GetRemoteStatusPageEnabledOk

`func (o *InlineObject80) GetRemoteStatusPageEnabledOk() (*bool, bool)`

GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStatusPageEnabled

`func (o *InlineObject80) SetRemoteStatusPageEnabled(v bool)`

SetRemoteStatusPageEnabled sets RemoteStatusPageEnabled field to given value.

### HasRemoteStatusPageEnabled

`func (o *InlineObject80) HasRemoteStatusPageEnabled() bool`

HasRemoteStatusPageEnabled returns a boolean if a field has been set.

### GetSecureConnect

`func (o *InlineObject80) GetSecureConnect() NetworksNetworkIdSettingsSecureConnect`

GetSecureConnect returns the SecureConnect field if non-nil, zero value otherwise.

### GetSecureConnectOk

`func (o *InlineObject80) GetSecureConnectOk() (*NetworksNetworkIdSettingsSecureConnect, bool)`

GetSecureConnectOk returns a tuple with the SecureConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureConnect

`func (o *InlineObject80) SetSecureConnect(v NetworksNetworkIdSettingsSecureConnect)`

SetSecureConnect sets SecureConnect field to given value.

### HasSecureConnect

`func (o *InlineObject80) HasSecureConnect() bool`

HasSecureConnect returns a boolean if a field has been set.

### GetLocalStatusPage

`func (o *InlineObject80) GetLocalStatusPage() NetworksNetworkIdSettingsLocalStatusPage`

GetLocalStatusPage returns the LocalStatusPage field if non-nil, zero value otherwise.

### GetLocalStatusPageOk

`func (o *InlineObject80) GetLocalStatusPageOk() (*NetworksNetworkIdSettingsLocalStatusPage, bool)`

GetLocalStatusPageOk returns a tuple with the LocalStatusPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPage

`func (o *InlineObject80) SetLocalStatusPage(v NetworksNetworkIdSettingsLocalStatusPage)`

SetLocalStatusPage sets LocalStatusPage field to given value.

### HasLocalStatusPage

`func (o *InlineObject80) HasLocalStatusPage() bool`

HasLocalStatusPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



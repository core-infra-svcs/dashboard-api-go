# InlineResponse200121

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStatusPageEnabled** | Pointer to **bool** | Enables / disables the local device status pages (&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://my.meraki.com/&#39;&gt;my.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://ap.meraki.com/&#39;&gt;ap.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://switch.meraki.com/&#39;&gt;switch.meraki.com, &lt;/a&gt;&lt;a target&#x3D;&#39;_blank&#39; href&#x3D;&#39;http://wired.meraki.com/&#39;&gt;wired.meraki.com&lt;/a&gt;). Optional (defaults to false) | [optional] 
**RemoteStatusPageEnabled** | Pointer to **bool** | Enables / disables access to the device status page (&lt;a target&#x3D;&#39;_blank&#39;&gt;http://[device&#39;s LAN IP])&lt;/a&gt;. Optional. Can only be set if localStatusPageEnabled is set to true | [optional] 
**LocalStatusPage** | Pointer to [**InlineResponse200121LocalStatusPage**](InlineResponse200121LocalStatusPage.md) |  | [optional] 
**SecurePort** | Pointer to [**InlineResponse200121SecurePort**](InlineResponse200121SecurePort.md) |  | [optional] 
**Fips** | Pointer to [**InlineResponse200121Fips**](InlineResponse200121Fips.md) |  | [optional] 
**NamedVlans** | Pointer to [**InlineResponse200121NamedVlans**](InlineResponse200121NamedVlans.md) |  | [optional] 

## Methods

### NewInlineResponse200121

`func NewInlineResponse200121() *InlineResponse200121`

NewInlineResponse200121 instantiates a new InlineResponse200121 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200121WithDefaults

`func NewInlineResponse200121WithDefaults() *InlineResponse200121`

NewInlineResponse200121WithDefaults instantiates a new InlineResponse200121 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStatusPageEnabled

`func (o *InlineResponse200121) GetLocalStatusPageEnabled() bool`

GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field if non-nil, zero value otherwise.

### GetLocalStatusPageEnabledOk

`func (o *InlineResponse200121) GetLocalStatusPageEnabledOk() (*bool, bool)`

GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPageEnabled

`func (o *InlineResponse200121) SetLocalStatusPageEnabled(v bool)`

SetLocalStatusPageEnabled sets LocalStatusPageEnabled field to given value.

### HasLocalStatusPageEnabled

`func (o *InlineResponse200121) HasLocalStatusPageEnabled() bool`

HasLocalStatusPageEnabled returns a boolean if a field has been set.

### GetRemoteStatusPageEnabled

`func (o *InlineResponse200121) GetRemoteStatusPageEnabled() bool`

GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field if non-nil, zero value otherwise.

### GetRemoteStatusPageEnabledOk

`func (o *InlineResponse200121) GetRemoteStatusPageEnabledOk() (*bool, bool)`

GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStatusPageEnabled

`func (o *InlineResponse200121) SetRemoteStatusPageEnabled(v bool)`

SetRemoteStatusPageEnabled sets RemoteStatusPageEnabled field to given value.

### HasRemoteStatusPageEnabled

`func (o *InlineResponse200121) HasRemoteStatusPageEnabled() bool`

HasRemoteStatusPageEnabled returns a boolean if a field has been set.

### GetLocalStatusPage

`func (o *InlineResponse200121) GetLocalStatusPage() InlineResponse200121LocalStatusPage`

GetLocalStatusPage returns the LocalStatusPage field if non-nil, zero value otherwise.

### GetLocalStatusPageOk

`func (o *InlineResponse200121) GetLocalStatusPageOk() (*InlineResponse200121LocalStatusPage, bool)`

GetLocalStatusPageOk returns a tuple with the LocalStatusPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStatusPage

`func (o *InlineResponse200121) SetLocalStatusPage(v InlineResponse200121LocalStatusPage)`

SetLocalStatusPage sets LocalStatusPage field to given value.

### HasLocalStatusPage

`func (o *InlineResponse200121) HasLocalStatusPage() bool`

HasLocalStatusPage returns a boolean if a field has been set.

### GetSecurePort

`func (o *InlineResponse200121) GetSecurePort() InlineResponse200121SecurePort`

GetSecurePort returns the SecurePort field if non-nil, zero value otherwise.

### GetSecurePortOk

`func (o *InlineResponse200121) GetSecurePortOk() (*InlineResponse200121SecurePort, bool)`

GetSecurePortOk returns a tuple with the SecurePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePort

`func (o *InlineResponse200121) SetSecurePort(v InlineResponse200121SecurePort)`

SetSecurePort sets SecurePort field to given value.

### HasSecurePort

`func (o *InlineResponse200121) HasSecurePort() bool`

HasSecurePort returns a boolean if a field has been set.

### GetFips

`func (o *InlineResponse200121) GetFips() InlineResponse200121Fips`

GetFips returns the Fips field if non-nil, zero value otherwise.

### GetFipsOk

`func (o *InlineResponse200121) GetFipsOk() (*InlineResponse200121Fips, bool)`

GetFipsOk returns a tuple with the Fips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFips

`func (o *InlineResponse200121) SetFips(v InlineResponse200121Fips)`

SetFips sets Fips field to given value.

### HasFips

`func (o *InlineResponse200121) HasFips() bool`

HasFips returns a boolean if a field has been set.

### GetNamedVlans

`func (o *InlineResponse200121) GetNamedVlans() InlineResponse200121NamedVlans`

GetNamedVlans returns the NamedVlans field if non-nil, zero value otherwise.

### GetNamedVlansOk

`func (o *InlineResponse200121) GetNamedVlansOk() (*InlineResponse200121NamedVlans, bool)`

GetNamedVlansOk returns a tuple with the NamedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlans

`func (o *InlineResponse200121) SetNamedVlans(v InlineResponse200121NamedVlans)`

SetNamedVlans sets NamedVlans field to given value.

### HasNamedVlans

`func (o *InlineResponse200121) HasNamedVlans() bool`

HasNamedVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



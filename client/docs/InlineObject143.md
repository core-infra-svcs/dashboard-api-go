# InlineObject143

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplashUrl** | Pointer to **string** | [optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see &#39;useSplashUrl&#39; | [optional] 
**UseSplashUrl** | Pointer to **bool** | [optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID&#39;s access control settings, it may not be possible to use the custom splash URL. | [optional] 
**SplashTimeout** | Pointer to **int32** | Splash timeout in minutes. This will determine how often users will see the splash page. | [optional] 
**RedirectUrl** | Pointer to **string** | The custom redirect URL where the users will go after the splash page. | [optional] 
**UseRedirectUrl** | Pointer to **bool** | The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true. | [optional] 
**WelcomeMessage** | Pointer to **string** | The welcome message for the users on the splash page. | [optional] 
**SplashLogo** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo**](NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage**](NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront**](NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. Can be either &#39;open&#39;, &#39;restricted&#39;, or &#39;default&#39;. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**GuestSponsorship** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship**](NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship.md) |  | [optional] 
**Billing** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling**](NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment**](NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment.md) |  | [optional] 

## Methods

### NewInlineObject143

`func NewInlineObject143() *InlineObject143`

NewInlineObject143 instantiates a new InlineObject143 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject143WithDefaults

`func NewInlineObject143WithDefaults() *InlineObject143`

NewInlineObject143WithDefaults instantiates a new InlineObject143 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplashUrl

`func (o *InlineObject143) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *InlineObject143) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *InlineObject143) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *InlineObject143) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *InlineObject143) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *InlineObject143) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *InlineObject143) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *InlineObject143) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineObject143) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineObject143) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineObject143) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineObject143) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineObject143) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineObject143) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineObject143) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineObject143) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *InlineObject143) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *InlineObject143) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *InlineObject143) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *InlineObject143) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *InlineObject143) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *InlineObject143) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *InlineObject143) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *InlineObject143) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetSplashLogo

`func (o *InlineObject143) GetSplashLogo() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *InlineObject143) GetSplashLogoOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *InlineObject143) SetSplashLogo(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *InlineObject143) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *InlineObject143) GetSplashImage() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *InlineObject143) GetSplashImageOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *InlineObject143) SetSplashImage(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *InlineObject143) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *InlineObject143) GetSplashPrepaidFront() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *InlineObject143) GetSplashPrepaidFrontOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *InlineObject143) SetSplashPrepaidFront(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *InlineObject143) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *InlineObject143) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *InlineObject143) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *InlineObject143) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *InlineObject143) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *InlineObject143) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *InlineObject143) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *InlineObject143) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *InlineObject143) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *InlineObject143) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *InlineObject143) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *InlineObject143) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *InlineObject143) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *InlineObject143) GetGuestSponsorship() NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *InlineObject143) GetGuestSponsorshipOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *InlineObject143) SetGuestSponsorship(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *InlineObject143) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBilling

`func (o *InlineObject143) GetBilling() NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InlineObject143) GetBillingOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InlineObject143) SetBilling(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *InlineObject143) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *InlineObject143) GetSentryEnrollment() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *InlineObject143) GetSentryEnrollmentOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *InlineObject143) SetSentryEnrollment(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *InlineObject143) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


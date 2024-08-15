# InlineResponse200194

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | SSID number | [optional] 
**SplashPage** | Pointer to **string** | The type of splash page for this SSID | [optional] 
**UseSplashUrl** | Pointer to **bool** | Boolean indicating whether the users will be redirected to the custom splash url | [optional] 
**SplashUrl** | Pointer to **string** | The custom splash URL of the click-through splash page. | [optional] 
**SplashTimeout** | Pointer to **int32** | Splash timeout in minutes. | [optional] 
**RedirectUrl** | Pointer to **string** | The custom redirect URL where the users will go after the splash page. | [optional] 
**UseRedirectUrl** | Pointer to **bool** | The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. | [optional] 
**WelcomeMessage** | Pointer to **string** | The welcome message for the users on the splash page. | [optional] 
**ThemeId** | Pointer to **string** | The id of the selected splash theme. | [optional] 
**SplashLogo** | Pointer to [**InlineResponse200194SplashLogo**](InlineResponse200194SplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**InlineResponse200194SplashImage**](InlineResponse200194SplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**InlineResponse200194SplashPrepaidFront**](InlineResponse200194SplashPrepaidFront.md) |  | [optional] 
**GuestSponsorship** | Pointer to [**InlineResponse200194GuestSponsorship**](InlineResponse200194GuestSponsorship.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**Billing** | Pointer to [**InlineResponse200194Billing**](InlineResponse200194Billing.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**InlineResponse200194SentryEnrollment**](InlineResponse200194SentryEnrollment.md) |  | [optional] 
**SelfRegistration** | Pointer to [**InlineResponse200194SelfRegistration**](InlineResponse200194SelfRegistration.md) |  | [optional] 

## Methods

### NewInlineResponse200194

`func NewInlineResponse200194() *InlineResponse200194`

NewInlineResponse200194 instantiates a new InlineResponse200194 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200194WithDefaults

`func NewInlineResponse200194WithDefaults() *InlineResponse200194`

NewInlineResponse200194WithDefaults instantiates a new InlineResponse200194 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse200194) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200194) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200194) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200194) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineResponse200194) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineResponse200194) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineResponse200194) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *InlineResponse200194) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *InlineResponse200194) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *InlineResponse200194) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *InlineResponse200194) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *InlineResponse200194) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashUrl

`func (o *InlineResponse200194) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *InlineResponse200194) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *InlineResponse200194) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *InlineResponse200194) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineResponse200194) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineResponse200194) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineResponse200194) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineResponse200194) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineResponse200194) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineResponse200194) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineResponse200194) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineResponse200194) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *InlineResponse200194) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *InlineResponse200194) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *InlineResponse200194) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *InlineResponse200194) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *InlineResponse200194) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *InlineResponse200194) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *InlineResponse200194) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *InlineResponse200194) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetThemeId

`func (o *InlineResponse200194) GetThemeId() string`

GetThemeId returns the ThemeId field if non-nil, zero value otherwise.

### GetThemeIdOk

`func (o *InlineResponse200194) GetThemeIdOk() (*string, bool)`

GetThemeIdOk returns a tuple with the ThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeId

`func (o *InlineResponse200194) SetThemeId(v string)`

SetThemeId sets ThemeId field to given value.

### HasThemeId

`func (o *InlineResponse200194) HasThemeId() bool`

HasThemeId returns a boolean if a field has been set.

### GetSplashLogo

`func (o *InlineResponse200194) GetSplashLogo() InlineResponse200194SplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *InlineResponse200194) GetSplashLogoOk() (*InlineResponse200194SplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *InlineResponse200194) SetSplashLogo(v InlineResponse200194SplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *InlineResponse200194) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *InlineResponse200194) GetSplashImage() InlineResponse200194SplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *InlineResponse200194) GetSplashImageOk() (*InlineResponse200194SplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *InlineResponse200194) SetSplashImage(v InlineResponse200194SplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *InlineResponse200194) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *InlineResponse200194) GetSplashPrepaidFront() InlineResponse200194SplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *InlineResponse200194) GetSplashPrepaidFrontOk() (*InlineResponse200194SplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *InlineResponse200194) SetSplashPrepaidFront(v InlineResponse200194SplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *InlineResponse200194) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *InlineResponse200194) GetGuestSponsorship() InlineResponse200194GuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *InlineResponse200194) GetGuestSponsorshipOk() (*InlineResponse200194GuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *InlineResponse200194) SetGuestSponsorship(v InlineResponse200194GuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *InlineResponse200194) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200194) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *InlineResponse200194) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200194) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200194) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *InlineResponse200194) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *InlineResponse200194) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *InlineResponse200194) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *InlineResponse200194) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *InlineResponse200194) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *InlineResponse200194) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *InlineResponse200194) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *InlineResponse200194) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetBilling

`func (o *InlineResponse200194) GetBilling() InlineResponse200194Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InlineResponse200194) GetBillingOk() (*InlineResponse200194Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InlineResponse200194) SetBilling(v InlineResponse200194Billing)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *InlineResponse200194) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *InlineResponse200194) GetSentryEnrollment() InlineResponse200194SentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *InlineResponse200194) GetSentryEnrollmentOk() (*InlineResponse200194SentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *InlineResponse200194) SetSentryEnrollment(v InlineResponse200194SentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *InlineResponse200194) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.

### GetSelfRegistration

`func (o *InlineResponse200194) GetSelfRegistration() InlineResponse200194SelfRegistration`

GetSelfRegistration returns the SelfRegistration field if non-nil, zero value otherwise.

### GetSelfRegistrationOk

`func (o *InlineResponse200194) GetSelfRegistrationOk() (*InlineResponse200194SelfRegistration, bool)`

GetSelfRegistrationOk returns a tuple with the SelfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegistration

`func (o *InlineResponse200194) SetSelfRegistration(v InlineResponse200194SelfRegistration)`

SetSelfRegistration sets SelfRegistration field to given value.

### HasSelfRegistration

`func (o *InlineResponse200194) HasSelfRegistration() bool`

HasSelfRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



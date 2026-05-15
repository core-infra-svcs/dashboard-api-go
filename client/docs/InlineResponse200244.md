# InlineResponse200244

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
**SplashLogo** | Pointer to [**InlineResponse200244SplashLogo**](InlineResponse200244SplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**InlineResponse200244SplashImage**](InlineResponse200244SplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**InlineResponse200244SplashPrepaidFront**](InlineResponse200244SplashPrepaidFront.md) |  | [optional] 
**GuestSponsorship** | Pointer to [**InlineResponse200244GuestSponsorship**](InlineResponse200244GuestSponsorship.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**Billing** | Pointer to [**InlineResponse200244Billing**](InlineResponse200244Billing.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**InlineResponse200244SentryEnrollment**](InlineResponse200244SentryEnrollment.md) |  | [optional] 
**SelfRegistration** | Pointer to [**InlineResponse200244SelfRegistration**](InlineResponse200244SelfRegistration.md) |  | [optional] 

## Methods

### NewInlineResponse200244

`func NewInlineResponse200244() *InlineResponse200244`

NewInlineResponse200244 instantiates a new InlineResponse200244 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200244WithDefaults

`func NewInlineResponse200244WithDefaults() *InlineResponse200244`

NewInlineResponse200244WithDefaults instantiates a new InlineResponse200244 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse200244) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200244) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200244) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200244) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineResponse200244) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineResponse200244) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineResponse200244) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *InlineResponse200244) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *InlineResponse200244) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *InlineResponse200244) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *InlineResponse200244) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *InlineResponse200244) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashUrl

`func (o *InlineResponse200244) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *InlineResponse200244) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *InlineResponse200244) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *InlineResponse200244) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineResponse200244) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineResponse200244) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineResponse200244) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineResponse200244) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineResponse200244) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineResponse200244) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineResponse200244) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineResponse200244) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *InlineResponse200244) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *InlineResponse200244) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *InlineResponse200244) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *InlineResponse200244) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *InlineResponse200244) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *InlineResponse200244) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *InlineResponse200244) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *InlineResponse200244) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetThemeId

`func (o *InlineResponse200244) GetThemeId() string`

GetThemeId returns the ThemeId field if non-nil, zero value otherwise.

### GetThemeIdOk

`func (o *InlineResponse200244) GetThemeIdOk() (*string, bool)`

GetThemeIdOk returns a tuple with the ThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeId

`func (o *InlineResponse200244) SetThemeId(v string)`

SetThemeId sets ThemeId field to given value.

### HasThemeId

`func (o *InlineResponse200244) HasThemeId() bool`

HasThemeId returns a boolean if a field has been set.

### GetSplashLogo

`func (o *InlineResponse200244) GetSplashLogo() InlineResponse200244SplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *InlineResponse200244) GetSplashLogoOk() (*InlineResponse200244SplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *InlineResponse200244) SetSplashLogo(v InlineResponse200244SplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *InlineResponse200244) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *InlineResponse200244) GetSplashImage() InlineResponse200244SplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *InlineResponse200244) GetSplashImageOk() (*InlineResponse200244SplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *InlineResponse200244) SetSplashImage(v InlineResponse200244SplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *InlineResponse200244) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *InlineResponse200244) GetSplashPrepaidFront() InlineResponse200244SplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *InlineResponse200244) GetSplashPrepaidFrontOk() (*InlineResponse200244SplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *InlineResponse200244) SetSplashPrepaidFront(v InlineResponse200244SplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *InlineResponse200244) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *InlineResponse200244) GetGuestSponsorship() InlineResponse200244GuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *InlineResponse200244) GetGuestSponsorshipOk() (*InlineResponse200244GuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *InlineResponse200244) SetGuestSponsorship(v InlineResponse200244GuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *InlineResponse200244) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200244) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *InlineResponse200244) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200244) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200244) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *InlineResponse200244) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *InlineResponse200244) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *InlineResponse200244) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *InlineResponse200244) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *InlineResponse200244) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *InlineResponse200244) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *InlineResponse200244) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *InlineResponse200244) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetBilling

`func (o *InlineResponse200244) GetBilling() InlineResponse200244Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InlineResponse200244) GetBillingOk() (*InlineResponse200244Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InlineResponse200244) SetBilling(v InlineResponse200244Billing)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *InlineResponse200244) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *InlineResponse200244) GetSentryEnrollment() InlineResponse200244SentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *InlineResponse200244) GetSentryEnrollmentOk() (*InlineResponse200244SentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *InlineResponse200244) SetSentryEnrollment(v InlineResponse200244SentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *InlineResponse200244) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.

### GetSelfRegistration

`func (o *InlineResponse200244) GetSelfRegistration() InlineResponse200244SelfRegistration`

GetSelfRegistration returns the SelfRegistration field if non-nil, zero value otherwise.

### GetSelfRegistrationOk

`func (o *InlineResponse200244) GetSelfRegistrationOk() (*InlineResponse200244SelfRegistration, bool)`

GetSelfRegistrationOk returns a tuple with the SelfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegistration

`func (o *InlineResponse200244) SetSelfRegistration(v InlineResponse200244SelfRegistration)`

SetSelfRegistration sets SelfRegistration field to given value.

### HasSelfRegistration

`func (o *InlineResponse200244) HasSelfRegistration() bool`

HasSelfRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



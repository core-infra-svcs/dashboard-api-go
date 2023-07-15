# InlineResponse200101

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
**SplashLogo** | Pointer to [**InlineResponse200101SplashLogo**](InlineResponse200101SplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**InlineResponse200101SplashImage**](InlineResponse200101SplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**InlineResponse200101SplashPrepaidFront**](InlineResponse200101SplashPrepaidFront.md) |  | [optional] 
**GuestSponsorship** | Pointer to [**InlineResponse200101GuestSponsorship**](InlineResponse200101GuestSponsorship.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**Billing** | Pointer to [**InlineResponse200101Billing**](InlineResponse200101Billing.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**InlineResponse200101SentryEnrollment**](InlineResponse200101SentryEnrollment.md) |  | [optional] 
**SelfRegistration** | Pointer to [**InlineResponse200101SelfRegistration**](InlineResponse200101SelfRegistration.md) |  | [optional] 

## Methods

### NewInlineResponse200101

`func NewInlineResponse200101() *InlineResponse200101`

NewInlineResponse200101 instantiates a new InlineResponse200101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101WithDefaults

`func NewInlineResponse200101WithDefaults() *InlineResponse200101`

NewInlineResponse200101WithDefaults instantiates a new InlineResponse200101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse200101) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse200101) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse200101) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse200101) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineResponse200101) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineResponse200101) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineResponse200101) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *InlineResponse200101) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *InlineResponse200101) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *InlineResponse200101) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *InlineResponse200101) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *InlineResponse200101) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashUrl

`func (o *InlineResponse200101) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *InlineResponse200101) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *InlineResponse200101) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *InlineResponse200101) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineResponse200101) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineResponse200101) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineResponse200101) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineResponse200101) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineResponse200101) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineResponse200101) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineResponse200101) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineResponse200101) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *InlineResponse200101) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *InlineResponse200101) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *InlineResponse200101) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *InlineResponse200101) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *InlineResponse200101) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *InlineResponse200101) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *InlineResponse200101) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *InlineResponse200101) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetSplashLogo

`func (o *InlineResponse200101) GetSplashLogo() InlineResponse200101SplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *InlineResponse200101) GetSplashLogoOk() (*InlineResponse200101SplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *InlineResponse200101) SetSplashLogo(v InlineResponse200101SplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *InlineResponse200101) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *InlineResponse200101) GetSplashImage() InlineResponse200101SplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *InlineResponse200101) GetSplashImageOk() (*InlineResponse200101SplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *InlineResponse200101) SetSplashImage(v InlineResponse200101SplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *InlineResponse200101) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *InlineResponse200101) GetSplashPrepaidFront() InlineResponse200101SplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *InlineResponse200101) GetSplashPrepaidFrontOk() (*InlineResponse200101SplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *InlineResponse200101) SetSplashPrepaidFront(v InlineResponse200101SplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *InlineResponse200101) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *InlineResponse200101) GetGuestSponsorship() InlineResponse200101GuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *InlineResponse200101) GetGuestSponsorshipOk() (*InlineResponse200101GuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *InlineResponse200101) SetGuestSponsorship(v InlineResponse200101GuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *InlineResponse200101) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200101) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *InlineResponse200101) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200101) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *InlineResponse200101) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *InlineResponse200101) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *InlineResponse200101) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *InlineResponse200101) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *InlineResponse200101) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *InlineResponse200101) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *InlineResponse200101) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *InlineResponse200101) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *InlineResponse200101) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetBilling

`func (o *InlineResponse200101) GetBilling() InlineResponse200101Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InlineResponse200101) GetBillingOk() (*InlineResponse200101Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InlineResponse200101) SetBilling(v InlineResponse200101Billing)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *InlineResponse200101) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *InlineResponse200101) GetSentryEnrollment() InlineResponse200101SentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *InlineResponse200101) GetSentryEnrollmentOk() (*InlineResponse200101SentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *InlineResponse200101) SetSentryEnrollment(v InlineResponse200101SentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *InlineResponse200101) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.

### GetSelfRegistration

`func (o *InlineResponse200101) GetSelfRegistration() InlineResponse200101SelfRegistration`

GetSelfRegistration returns the SelfRegistration field if non-nil, zero value otherwise.

### GetSelfRegistrationOk

`func (o *InlineResponse200101) GetSelfRegistrationOk() (*InlineResponse200101SelfRegistration, bool)`

GetSelfRegistrationOk returns a tuple with the SelfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegistration

`func (o *InlineResponse200101) SetSelfRegistration(v InlineResponse200101SelfRegistration)`

SetSelfRegistration sets SelfRegistration field to given value.

### HasSelfRegistration

`func (o *InlineResponse200101) HasSelfRegistration() bool`

HasSelfRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



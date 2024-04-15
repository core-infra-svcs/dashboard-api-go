/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200189 struct for InlineResponse200189
type InlineResponse200189 struct {
	// SSID number
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// The type of splash page for this SSID
	SplashPage *string `json:"splashPage,omitempty"`
	// Boolean indicating whether the users will be redirected to the custom splash url
	UseSplashUrl *bool `json:"useSplashUrl,omitempty"`
	// The custom splash URL of the click-through splash page.
	SplashUrl *string `json:"splashUrl,omitempty"`
	// Splash timeout in minutes.
	SplashTimeout *int32 `json:"splashTimeout,omitempty"`
	// The custom redirect URL where the users will go after the splash page.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.
	UseRedirectUrl *bool `json:"useRedirectUrl,omitempty"`
	// The welcome message for the users on the splash page.
	WelcomeMessage *string `json:"welcomeMessage,omitempty"`
	// The id of the selected splash theme.
	ThemeId *string `json:"themeId,omitempty"`
	SplashLogo *InlineResponse200189SplashLogo `json:"splashLogo,omitempty"`
	SplashImage *InlineResponse200189SplashImage `json:"splashImage,omitempty"`
	SplashPrepaidFront *InlineResponse200189SplashPrepaidFront `json:"splashPrepaidFront,omitempty"`
	GuestSponsorship *InlineResponse200189GuestSponsorship `json:"guestSponsorship,omitempty"`
	// How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	BlockAllTrafficBeforeSignOn *bool `json:"blockAllTrafficBeforeSignOn,omitempty"`
	// How login attempts should be handled when the controller is unreachable.
	ControllerDisconnectionBehavior *string `json:"controllerDisconnectionBehavior,omitempty"`
	// Whether or not to allow simultaneous logins from different devices.
	AllowSimultaneousLogins *bool `json:"allowSimultaneousLogins,omitempty"`
	Billing *InlineResponse200189Billing `json:"billing,omitempty"`
	SentryEnrollment *InlineResponse200189SentryEnrollment `json:"sentryEnrollment,omitempty"`
	SelfRegistration *InlineResponse200189SelfRegistration `json:"selfRegistration,omitempty"`
}

// NewInlineResponse200189 instantiates a new InlineResponse200189 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200189() *InlineResponse200189 {
	this := InlineResponse200189{}
	return &this
}

// NewInlineResponse200189WithDefaults instantiates a new InlineResponse200189 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200189WithDefaults() *InlineResponse200189 {
	this := InlineResponse200189{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse200189) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetSplashPage returns the SplashPage field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSplashPage() string {
	if o == nil || isNil(o.SplashPage) {
		var ret string
		return ret
	}
	return *o.SplashPage
}

// GetSplashPageOk returns a tuple with the SplashPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSplashPageOk() (*string, bool) {
	if o == nil || isNil(o.SplashPage) {
    return nil, false
	}
	return o.SplashPage, true
}

// HasSplashPage returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSplashPage() bool {
	if o != nil && !isNil(o.SplashPage) {
		return true
	}

	return false
}

// SetSplashPage gets a reference to the given string and assigns it to the SplashPage field.
func (o *InlineResponse200189) SetSplashPage(v string) {
	o.SplashPage = &v
}

// GetUseSplashUrl returns the UseSplashUrl field value if set, zero value otherwise.
func (o *InlineResponse200189) GetUseSplashUrl() bool {
	if o == nil || isNil(o.UseSplashUrl) {
		var ret bool
		return ret
	}
	return *o.UseSplashUrl
}

// GetUseSplashUrlOk returns a tuple with the UseSplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetUseSplashUrlOk() (*bool, bool) {
	if o == nil || isNil(o.UseSplashUrl) {
    return nil, false
	}
	return o.UseSplashUrl, true
}

// HasUseSplashUrl returns a boolean if a field has been set.
func (o *InlineResponse200189) HasUseSplashUrl() bool {
	if o != nil && !isNil(o.UseSplashUrl) {
		return true
	}

	return false
}

// SetUseSplashUrl gets a reference to the given bool and assigns it to the UseSplashUrl field.
func (o *InlineResponse200189) SetUseSplashUrl(v bool) {
	o.UseSplashUrl = &v
}

// GetSplashUrl returns the SplashUrl field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSplashUrl() string {
	if o == nil || isNil(o.SplashUrl) {
		var ret string
		return ret
	}
	return *o.SplashUrl
}

// GetSplashUrlOk returns a tuple with the SplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSplashUrlOk() (*string, bool) {
	if o == nil || isNil(o.SplashUrl) {
    return nil, false
	}
	return o.SplashUrl, true
}

// HasSplashUrl returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSplashUrl() bool {
	if o != nil && !isNil(o.SplashUrl) {
		return true
	}

	return false
}

// SetSplashUrl gets a reference to the given string and assigns it to the SplashUrl field.
func (o *InlineResponse200189) SetSplashUrl(v string) {
	o.SplashUrl = &v
}

// GetSplashTimeout returns the SplashTimeout field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSplashTimeout() int32 {
	if o == nil || isNil(o.SplashTimeout) {
		var ret int32
		return ret
	}
	return *o.SplashTimeout
}

// GetSplashTimeoutOk returns a tuple with the SplashTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSplashTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.SplashTimeout) {
    return nil, false
	}
	return o.SplashTimeout, true
}

// HasSplashTimeout returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSplashTimeout() bool {
	if o != nil && !isNil(o.SplashTimeout) {
		return true
	}

	return false
}

// SetSplashTimeout gets a reference to the given int32 and assigns it to the SplashTimeout field.
func (o *InlineResponse200189) SetSplashTimeout(v int32) {
	o.SplashTimeout = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *InlineResponse200189) GetRedirectUrl() string {
	if o == nil || isNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetRedirectUrlOk() (*string, bool) {
	if o == nil || isNil(o.RedirectUrl) {
    return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *InlineResponse200189) HasRedirectUrl() bool {
	if o != nil && !isNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *InlineResponse200189) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetUseRedirectUrl returns the UseRedirectUrl field value if set, zero value otherwise.
func (o *InlineResponse200189) GetUseRedirectUrl() bool {
	if o == nil || isNil(o.UseRedirectUrl) {
		var ret bool
		return ret
	}
	return *o.UseRedirectUrl
}

// GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetUseRedirectUrlOk() (*bool, bool) {
	if o == nil || isNil(o.UseRedirectUrl) {
    return nil, false
	}
	return o.UseRedirectUrl, true
}

// HasUseRedirectUrl returns a boolean if a field has been set.
func (o *InlineResponse200189) HasUseRedirectUrl() bool {
	if o != nil && !isNil(o.UseRedirectUrl) {
		return true
	}

	return false
}

// SetUseRedirectUrl gets a reference to the given bool and assigns it to the UseRedirectUrl field.
func (o *InlineResponse200189) SetUseRedirectUrl(v bool) {
	o.UseRedirectUrl = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise.
func (o *InlineResponse200189) GetWelcomeMessage() string {
	if o == nil || isNil(o.WelcomeMessage) {
		var ret string
		return ret
	}
	return *o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetWelcomeMessageOk() (*string, bool) {
	if o == nil || isNil(o.WelcomeMessage) {
    return nil, false
	}
	return o.WelcomeMessage, true
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *InlineResponse200189) HasWelcomeMessage() bool {
	if o != nil && !isNil(o.WelcomeMessage) {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given string and assigns it to the WelcomeMessage field.
func (o *InlineResponse200189) SetWelcomeMessage(v string) {
	o.WelcomeMessage = &v
}

// GetThemeId returns the ThemeId field value if set, zero value otherwise.
func (o *InlineResponse200189) GetThemeId() string {
	if o == nil || isNil(o.ThemeId) {
		var ret string
		return ret
	}
	return *o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetThemeIdOk() (*string, bool) {
	if o == nil || isNil(o.ThemeId) {
    return nil, false
	}
	return o.ThemeId, true
}

// HasThemeId returns a boolean if a field has been set.
func (o *InlineResponse200189) HasThemeId() bool {
	if o != nil && !isNil(o.ThemeId) {
		return true
	}

	return false
}

// SetThemeId gets a reference to the given string and assigns it to the ThemeId field.
func (o *InlineResponse200189) SetThemeId(v string) {
	o.ThemeId = &v
}

// GetSplashLogo returns the SplashLogo field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSplashLogo() InlineResponse200189SplashLogo {
	if o == nil || isNil(o.SplashLogo) {
		var ret InlineResponse200189SplashLogo
		return ret
	}
	return *o.SplashLogo
}

// GetSplashLogoOk returns a tuple with the SplashLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSplashLogoOk() (*InlineResponse200189SplashLogo, bool) {
	if o == nil || isNil(o.SplashLogo) {
    return nil, false
	}
	return o.SplashLogo, true
}

// HasSplashLogo returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSplashLogo() bool {
	if o != nil && !isNil(o.SplashLogo) {
		return true
	}

	return false
}

// SetSplashLogo gets a reference to the given InlineResponse200189SplashLogo and assigns it to the SplashLogo field.
func (o *InlineResponse200189) SetSplashLogo(v InlineResponse200189SplashLogo) {
	o.SplashLogo = &v
}

// GetSplashImage returns the SplashImage field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSplashImage() InlineResponse200189SplashImage {
	if o == nil || isNil(o.SplashImage) {
		var ret InlineResponse200189SplashImage
		return ret
	}
	return *o.SplashImage
}

// GetSplashImageOk returns a tuple with the SplashImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSplashImageOk() (*InlineResponse200189SplashImage, bool) {
	if o == nil || isNil(o.SplashImage) {
    return nil, false
	}
	return o.SplashImage, true
}

// HasSplashImage returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSplashImage() bool {
	if o != nil && !isNil(o.SplashImage) {
		return true
	}

	return false
}

// SetSplashImage gets a reference to the given InlineResponse200189SplashImage and assigns it to the SplashImage field.
func (o *InlineResponse200189) SetSplashImage(v InlineResponse200189SplashImage) {
	o.SplashImage = &v
}

// GetSplashPrepaidFront returns the SplashPrepaidFront field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSplashPrepaidFront() InlineResponse200189SplashPrepaidFront {
	if o == nil || isNil(o.SplashPrepaidFront) {
		var ret InlineResponse200189SplashPrepaidFront
		return ret
	}
	return *o.SplashPrepaidFront
}

// GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSplashPrepaidFrontOk() (*InlineResponse200189SplashPrepaidFront, bool) {
	if o == nil || isNil(o.SplashPrepaidFront) {
    return nil, false
	}
	return o.SplashPrepaidFront, true
}

// HasSplashPrepaidFront returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSplashPrepaidFront() bool {
	if o != nil && !isNil(o.SplashPrepaidFront) {
		return true
	}

	return false
}

// SetSplashPrepaidFront gets a reference to the given InlineResponse200189SplashPrepaidFront and assigns it to the SplashPrepaidFront field.
func (o *InlineResponse200189) SetSplashPrepaidFront(v InlineResponse200189SplashPrepaidFront) {
	o.SplashPrepaidFront = &v
}

// GetGuestSponsorship returns the GuestSponsorship field value if set, zero value otherwise.
func (o *InlineResponse200189) GetGuestSponsorship() InlineResponse200189GuestSponsorship {
	if o == nil || isNil(o.GuestSponsorship) {
		var ret InlineResponse200189GuestSponsorship
		return ret
	}
	return *o.GuestSponsorship
}

// GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetGuestSponsorshipOk() (*InlineResponse200189GuestSponsorship, bool) {
	if o == nil || isNil(o.GuestSponsorship) {
    return nil, false
	}
	return o.GuestSponsorship, true
}

// HasGuestSponsorship returns a boolean if a field has been set.
func (o *InlineResponse200189) HasGuestSponsorship() bool {
	if o != nil && !isNil(o.GuestSponsorship) {
		return true
	}

	return false
}

// SetGuestSponsorship gets a reference to the given InlineResponse200189GuestSponsorship and assigns it to the GuestSponsorship field.
func (o *InlineResponse200189) SetGuestSponsorship(v InlineResponse200189GuestSponsorship) {
	o.GuestSponsorship = &v
}

// GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field value if set, zero value otherwise.
func (o *InlineResponse200189) GetBlockAllTrafficBeforeSignOn() bool {
	if o == nil || isNil(o.BlockAllTrafficBeforeSignOn) {
		var ret bool
		return ret
	}
	return *o.BlockAllTrafficBeforeSignOn
}

// GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool) {
	if o == nil || isNil(o.BlockAllTrafficBeforeSignOn) {
    return nil, false
	}
	return o.BlockAllTrafficBeforeSignOn, true
}

// HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.
func (o *InlineResponse200189) HasBlockAllTrafficBeforeSignOn() bool {
	if o != nil && !isNil(o.BlockAllTrafficBeforeSignOn) {
		return true
	}

	return false
}

// SetBlockAllTrafficBeforeSignOn gets a reference to the given bool and assigns it to the BlockAllTrafficBeforeSignOn field.
func (o *InlineResponse200189) SetBlockAllTrafficBeforeSignOn(v bool) {
	o.BlockAllTrafficBeforeSignOn = &v
}

// GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field value if set, zero value otherwise.
func (o *InlineResponse200189) GetControllerDisconnectionBehavior() string {
	if o == nil || isNil(o.ControllerDisconnectionBehavior) {
		var ret string
		return ret
	}
	return *o.ControllerDisconnectionBehavior
}

// GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetControllerDisconnectionBehaviorOk() (*string, bool) {
	if o == nil || isNil(o.ControllerDisconnectionBehavior) {
    return nil, false
	}
	return o.ControllerDisconnectionBehavior, true
}

// HasControllerDisconnectionBehavior returns a boolean if a field has been set.
func (o *InlineResponse200189) HasControllerDisconnectionBehavior() bool {
	if o != nil && !isNil(o.ControllerDisconnectionBehavior) {
		return true
	}

	return false
}

// SetControllerDisconnectionBehavior gets a reference to the given string and assigns it to the ControllerDisconnectionBehavior field.
func (o *InlineResponse200189) SetControllerDisconnectionBehavior(v string) {
	o.ControllerDisconnectionBehavior = &v
}

// GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field value if set, zero value otherwise.
func (o *InlineResponse200189) GetAllowSimultaneousLogins() bool {
	if o == nil || isNil(o.AllowSimultaneousLogins) {
		var ret bool
		return ret
	}
	return *o.AllowSimultaneousLogins
}

// GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetAllowSimultaneousLoginsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowSimultaneousLogins) {
    return nil, false
	}
	return o.AllowSimultaneousLogins, true
}

// HasAllowSimultaneousLogins returns a boolean if a field has been set.
func (o *InlineResponse200189) HasAllowSimultaneousLogins() bool {
	if o != nil && !isNil(o.AllowSimultaneousLogins) {
		return true
	}

	return false
}

// SetAllowSimultaneousLogins gets a reference to the given bool and assigns it to the AllowSimultaneousLogins field.
func (o *InlineResponse200189) SetAllowSimultaneousLogins(v bool) {
	o.AllowSimultaneousLogins = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *InlineResponse200189) GetBilling() InlineResponse200189Billing {
	if o == nil || isNil(o.Billing) {
		var ret InlineResponse200189Billing
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetBillingOk() (*InlineResponse200189Billing, bool) {
	if o == nil || isNil(o.Billing) {
    return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *InlineResponse200189) HasBilling() bool {
	if o != nil && !isNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given InlineResponse200189Billing and assigns it to the Billing field.
func (o *InlineResponse200189) SetBilling(v InlineResponse200189Billing) {
	o.Billing = &v
}

// GetSentryEnrollment returns the SentryEnrollment field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSentryEnrollment() InlineResponse200189SentryEnrollment {
	if o == nil || isNil(o.SentryEnrollment) {
		var ret InlineResponse200189SentryEnrollment
		return ret
	}
	return *o.SentryEnrollment
}

// GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSentryEnrollmentOk() (*InlineResponse200189SentryEnrollment, bool) {
	if o == nil || isNil(o.SentryEnrollment) {
    return nil, false
	}
	return o.SentryEnrollment, true
}

// HasSentryEnrollment returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSentryEnrollment() bool {
	if o != nil && !isNil(o.SentryEnrollment) {
		return true
	}

	return false
}

// SetSentryEnrollment gets a reference to the given InlineResponse200189SentryEnrollment and assigns it to the SentryEnrollment field.
func (o *InlineResponse200189) SetSentryEnrollment(v InlineResponse200189SentryEnrollment) {
	o.SentryEnrollment = &v
}

// GetSelfRegistration returns the SelfRegistration field value if set, zero value otherwise.
func (o *InlineResponse200189) GetSelfRegistration() InlineResponse200189SelfRegistration {
	if o == nil || isNil(o.SelfRegistration) {
		var ret InlineResponse200189SelfRegistration
		return ret
	}
	return *o.SelfRegistration
}

// GetSelfRegistrationOk returns a tuple with the SelfRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200189) GetSelfRegistrationOk() (*InlineResponse200189SelfRegistration, bool) {
	if o == nil || isNil(o.SelfRegistration) {
    return nil, false
	}
	return o.SelfRegistration, true
}

// HasSelfRegistration returns a boolean if a field has been set.
func (o *InlineResponse200189) HasSelfRegistration() bool {
	if o != nil && !isNil(o.SelfRegistration) {
		return true
	}

	return false
}

// SetSelfRegistration gets a reference to the given InlineResponse200189SelfRegistration and assigns it to the SelfRegistration field.
func (o *InlineResponse200189) SetSelfRegistration(v InlineResponse200189SelfRegistration) {
	o.SelfRegistration = &v
}

func (o InlineResponse200189) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.SplashPage) {
		toSerialize["splashPage"] = o.SplashPage
	}
	if !isNil(o.UseSplashUrl) {
		toSerialize["useSplashUrl"] = o.UseSplashUrl
	}
	if !isNil(o.SplashUrl) {
		toSerialize["splashUrl"] = o.SplashUrl
	}
	if !isNil(o.SplashTimeout) {
		toSerialize["splashTimeout"] = o.SplashTimeout
	}
	if !isNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !isNil(o.UseRedirectUrl) {
		toSerialize["useRedirectUrl"] = o.UseRedirectUrl
	}
	if !isNil(o.WelcomeMessage) {
		toSerialize["welcomeMessage"] = o.WelcomeMessage
	}
	if !isNil(o.ThemeId) {
		toSerialize["themeId"] = o.ThemeId
	}
	if !isNil(o.SplashLogo) {
		toSerialize["splashLogo"] = o.SplashLogo
	}
	if !isNil(o.SplashImage) {
		toSerialize["splashImage"] = o.SplashImage
	}
	if !isNil(o.SplashPrepaidFront) {
		toSerialize["splashPrepaidFront"] = o.SplashPrepaidFront
	}
	if !isNil(o.GuestSponsorship) {
		toSerialize["guestSponsorship"] = o.GuestSponsorship
	}
	if !isNil(o.BlockAllTrafficBeforeSignOn) {
		toSerialize["blockAllTrafficBeforeSignOn"] = o.BlockAllTrafficBeforeSignOn
	}
	if !isNil(o.ControllerDisconnectionBehavior) {
		toSerialize["controllerDisconnectionBehavior"] = o.ControllerDisconnectionBehavior
	}
	if !isNil(o.AllowSimultaneousLogins) {
		toSerialize["allowSimultaneousLogins"] = o.AllowSimultaneousLogins
	}
	if !isNil(o.Billing) {
		toSerialize["billing"] = o.Billing
	}
	if !isNil(o.SentryEnrollment) {
		toSerialize["sentryEnrollment"] = o.SentryEnrollment
	}
	if !isNil(o.SelfRegistration) {
		toSerialize["selfRegistration"] = o.SelfRegistration
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200189 struct {
	value *InlineResponse200189
	isSet bool
}

func (v NullableInlineResponse200189) Get() *InlineResponse200189 {
	return v.value
}

func (v *NullableInlineResponse200189) Set(val *InlineResponse200189) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200189) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200189) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200189(val *InlineResponse200189) *NullableInlineResponse200189 {
	return &NullableInlineResponse200189{value: val, isSet: true}
}

func (v NullableInlineResponse200189) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200189) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



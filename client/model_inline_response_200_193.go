/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200193 struct for InlineResponse200193
type InlineResponse200193 struct {
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
	SplashLogo *InlineResponse200193SplashLogo `json:"splashLogo,omitempty"`
	SplashImage *InlineResponse200193SplashImage `json:"splashImage,omitempty"`
	SplashPrepaidFront *InlineResponse200193SplashPrepaidFront `json:"splashPrepaidFront,omitempty"`
	GuestSponsorship *InlineResponse200193GuestSponsorship `json:"guestSponsorship,omitempty"`
	// How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	BlockAllTrafficBeforeSignOn *bool `json:"blockAllTrafficBeforeSignOn,omitempty"`
	// How login attempts should be handled when the controller is unreachable.
	ControllerDisconnectionBehavior *string `json:"controllerDisconnectionBehavior,omitempty"`
	// Whether or not to allow simultaneous logins from different devices.
	AllowSimultaneousLogins *bool `json:"allowSimultaneousLogins,omitempty"`
	Billing *InlineResponse200193Billing `json:"billing,omitempty"`
	SentryEnrollment *InlineResponse200193SentryEnrollment `json:"sentryEnrollment,omitempty"`
	SelfRegistration *InlineResponse200193SelfRegistration `json:"selfRegistration,omitempty"`
}

// NewInlineResponse200193 instantiates a new InlineResponse200193 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200193() *InlineResponse200193 {
	this := InlineResponse200193{}
	return &this
}

// NewInlineResponse200193WithDefaults instantiates a new InlineResponse200193 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200193WithDefaults() *InlineResponse200193 {
	this := InlineResponse200193{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse200193) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetSplashPage returns the SplashPage field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSplashPage() string {
	if o == nil || isNil(o.SplashPage) {
		var ret string
		return ret
	}
	return *o.SplashPage
}

// GetSplashPageOk returns a tuple with the SplashPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSplashPageOk() (*string, bool) {
	if o == nil || isNil(o.SplashPage) {
    return nil, false
	}
	return o.SplashPage, true
}

// HasSplashPage returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSplashPage() bool {
	if o != nil && !isNil(o.SplashPage) {
		return true
	}

	return false
}

// SetSplashPage gets a reference to the given string and assigns it to the SplashPage field.
func (o *InlineResponse200193) SetSplashPage(v string) {
	o.SplashPage = &v
}

// GetUseSplashUrl returns the UseSplashUrl field value if set, zero value otherwise.
func (o *InlineResponse200193) GetUseSplashUrl() bool {
	if o == nil || isNil(o.UseSplashUrl) {
		var ret bool
		return ret
	}
	return *o.UseSplashUrl
}

// GetUseSplashUrlOk returns a tuple with the UseSplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetUseSplashUrlOk() (*bool, bool) {
	if o == nil || isNil(o.UseSplashUrl) {
    return nil, false
	}
	return o.UseSplashUrl, true
}

// HasUseSplashUrl returns a boolean if a field has been set.
func (o *InlineResponse200193) HasUseSplashUrl() bool {
	if o != nil && !isNil(o.UseSplashUrl) {
		return true
	}

	return false
}

// SetUseSplashUrl gets a reference to the given bool and assigns it to the UseSplashUrl field.
func (o *InlineResponse200193) SetUseSplashUrl(v bool) {
	o.UseSplashUrl = &v
}

// GetSplashUrl returns the SplashUrl field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSplashUrl() string {
	if o == nil || isNil(o.SplashUrl) {
		var ret string
		return ret
	}
	return *o.SplashUrl
}

// GetSplashUrlOk returns a tuple with the SplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSplashUrlOk() (*string, bool) {
	if o == nil || isNil(o.SplashUrl) {
    return nil, false
	}
	return o.SplashUrl, true
}

// HasSplashUrl returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSplashUrl() bool {
	if o != nil && !isNil(o.SplashUrl) {
		return true
	}

	return false
}

// SetSplashUrl gets a reference to the given string and assigns it to the SplashUrl field.
func (o *InlineResponse200193) SetSplashUrl(v string) {
	o.SplashUrl = &v
}

// GetSplashTimeout returns the SplashTimeout field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSplashTimeout() int32 {
	if o == nil || isNil(o.SplashTimeout) {
		var ret int32
		return ret
	}
	return *o.SplashTimeout
}

// GetSplashTimeoutOk returns a tuple with the SplashTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSplashTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.SplashTimeout) {
    return nil, false
	}
	return o.SplashTimeout, true
}

// HasSplashTimeout returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSplashTimeout() bool {
	if o != nil && !isNil(o.SplashTimeout) {
		return true
	}

	return false
}

// SetSplashTimeout gets a reference to the given int32 and assigns it to the SplashTimeout field.
func (o *InlineResponse200193) SetSplashTimeout(v int32) {
	o.SplashTimeout = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *InlineResponse200193) GetRedirectUrl() string {
	if o == nil || isNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetRedirectUrlOk() (*string, bool) {
	if o == nil || isNil(o.RedirectUrl) {
    return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *InlineResponse200193) HasRedirectUrl() bool {
	if o != nil && !isNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *InlineResponse200193) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetUseRedirectUrl returns the UseRedirectUrl field value if set, zero value otherwise.
func (o *InlineResponse200193) GetUseRedirectUrl() bool {
	if o == nil || isNil(o.UseRedirectUrl) {
		var ret bool
		return ret
	}
	return *o.UseRedirectUrl
}

// GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetUseRedirectUrlOk() (*bool, bool) {
	if o == nil || isNil(o.UseRedirectUrl) {
    return nil, false
	}
	return o.UseRedirectUrl, true
}

// HasUseRedirectUrl returns a boolean if a field has been set.
func (o *InlineResponse200193) HasUseRedirectUrl() bool {
	if o != nil && !isNil(o.UseRedirectUrl) {
		return true
	}

	return false
}

// SetUseRedirectUrl gets a reference to the given bool and assigns it to the UseRedirectUrl field.
func (o *InlineResponse200193) SetUseRedirectUrl(v bool) {
	o.UseRedirectUrl = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise.
func (o *InlineResponse200193) GetWelcomeMessage() string {
	if o == nil || isNil(o.WelcomeMessage) {
		var ret string
		return ret
	}
	return *o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetWelcomeMessageOk() (*string, bool) {
	if o == nil || isNil(o.WelcomeMessage) {
    return nil, false
	}
	return o.WelcomeMessage, true
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *InlineResponse200193) HasWelcomeMessage() bool {
	if o != nil && !isNil(o.WelcomeMessage) {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given string and assigns it to the WelcomeMessage field.
func (o *InlineResponse200193) SetWelcomeMessage(v string) {
	o.WelcomeMessage = &v
}

// GetThemeId returns the ThemeId field value if set, zero value otherwise.
func (o *InlineResponse200193) GetThemeId() string {
	if o == nil || isNil(o.ThemeId) {
		var ret string
		return ret
	}
	return *o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetThemeIdOk() (*string, bool) {
	if o == nil || isNil(o.ThemeId) {
    return nil, false
	}
	return o.ThemeId, true
}

// HasThemeId returns a boolean if a field has been set.
func (o *InlineResponse200193) HasThemeId() bool {
	if o != nil && !isNil(o.ThemeId) {
		return true
	}

	return false
}

// SetThemeId gets a reference to the given string and assigns it to the ThemeId field.
func (o *InlineResponse200193) SetThemeId(v string) {
	o.ThemeId = &v
}

// GetSplashLogo returns the SplashLogo field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSplashLogo() InlineResponse200193SplashLogo {
	if o == nil || isNil(o.SplashLogo) {
		var ret InlineResponse200193SplashLogo
		return ret
	}
	return *o.SplashLogo
}

// GetSplashLogoOk returns a tuple with the SplashLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSplashLogoOk() (*InlineResponse200193SplashLogo, bool) {
	if o == nil || isNil(o.SplashLogo) {
    return nil, false
	}
	return o.SplashLogo, true
}

// HasSplashLogo returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSplashLogo() bool {
	if o != nil && !isNil(o.SplashLogo) {
		return true
	}

	return false
}

// SetSplashLogo gets a reference to the given InlineResponse200193SplashLogo and assigns it to the SplashLogo field.
func (o *InlineResponse200193) SetSplashLogo(v InlineResponse200193SplashLogo) {
	o.SplashLogo = &v
}

// GetSplashImage returns the SplashImage field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSplashImage() InlineResponse200193SplashImage {
	if o == nil || isNil(o.SplashImage) {
		var ret InlineResponse200193SplashImage
		return ret
	}
	return *o.SplashImage
}

// GetSplashImageOk returns a tuple with the SplashImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSplashImageOk() (*InlineResponse200193SplashImage, bool) {
	if o == nil || isNil(o.SplashImage) {
    return nil, false
	}
	return o.SplashImage, true
}

// HasSplashImage returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSplashImage() bool {
	if o != nil && !isNil(o.SplashImage) {
		return true
	}

	return false
}

// SetSplashImage gets a reference to the given InlineResponse200193SplashImage and assigns it to the SplashImage field.
func (o *InlineResponse200193) SetSplashImage(v InlineResponse200193SplashImage) {
	o.SplashImage = &v
}

// GetSplashPrepaidFront returns the SplashPrepaidFront field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSplashPrepaidFront() InlineResponse200193SplashPrepaidFront {
	if o == nil || isNil(o.SplashPrepaidFront) {
		var ret InlineResponse200193SplashPrepaidFront
		return ret
	}
	return *o.SplashPrepaidFront
}

// GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSplashPrepaidFrontOk() (*InlineResponse200193SplashPrepaidFront, bool) {
	if o == nil || isNil(o.SplashPrepaidFront) {
    return nil, false
	}
	return o.SplashPrepaidFront, true
}

// HasSplashPrepaidFront returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSplashPrepaidFront() bool {
	if o != nil && !isNil(o.SplashPrepaidFront) {
		return true
	}

	return false
}

// SetSplashPrepaidFront gets a reference to the given InlineResponse200193SplashPrepaidFront and assigns it to the SplashPrepaidFront field.
func (o *InlineResponse200193) SetSplashPrepaidFront(v InlineResponse200193SplashPrepaidFront) {
	o.SplashPrepaidFront = &v
}

// GetGuestSponsorship returns the GuestSponsorship field value if set, zero value otherwise.
func (o *InlineResponse200193) GetGuestSponsorship() InlineResponse200193GuestSponsorship {
	if o == nil || isNil(o.GuestSponsorship) {
		var ret InlineResponse200193GuestSponsorship
		return ret
	}
	return *o.GuestSponsorship
}

// GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetGuestSponsorshipOk() (*InlineResponse200193GuestSponsorship, bool) {
	if o == nil || isNil(o.GuestSponsorship) {
    return nil, false
	}
	return o.GuestSponsorship, true
}

// HasGuestSponsorship returns a boolean if a field has been set.
func (o *InlineResponse200193) HasGuestSponsorship() bool {
	if o != nil && !isNil(o.GuestSponsorship) {
		return true
	}

	return false
}

// SetGuestSponsorship gets a reference to the given InlineResponse200193GuestSponsorship and assigns it to the GuestSponsorship field.
func (o *InlineResponse200193) SetGuestSponsorship(v InlineResponse200193GuestSponsorship) {
	o.GuestSponsorship = &v
}

// GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field value if set, zero value otherwise.
func (o *InlineResponse200193) GetBlockAllTrafficBeforeSignOn() bool {
	if o == nil || isNil(o.BlockAllTrafficBeforeSignOn) {
		var ret bool
		return ret
	}
	return *o.BlockAllTrafficBeforeSignOn
}

// GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool) {
	if o == nil || isNil(o.BlockAllTrafficBeforeSignOn) {
    return nil, false
	}
	return o.BlockAllTrafficBeforeSignOn, true
}

// HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.
func (o *InlineResponse200193) HasBlockAllTrafficBeforeSignOn() bool {
	if o != nil && !isNil(o.BlockAllTrafficBeforeSignOn) {
		return true
	}

	return false
}

// SetBlockAllTrafficBeforeSignOn gets a reference to the given bool and assigns it to the BlockAllTrafficBeforeSignOn field.
func (o *InlineResponse200193) SetBlockAllTrafficBeforeSignOn(v bool) {
	o.BlockAllTrafficBeforeSignOn = &v
}

// GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field value if set, zero value otherwise.
func (o *InlineResponse200193) GetControllerDisconnectionBehavior() string {
	if o == nil || isNil(o.ControllerDisconnectionBehavior) {
		var ret string
		return ret
	}
	return *o.ControllerDisconnectionBehavior
}

// GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetControllerDisconnectionBehaviorOk() (*string, bool) {
	if o == nil || isNil(o.ControllerDisconnectionBehavior) {
    return nil, false
	}
	return o.ControllerDisconnectionBehavior, true
}

// HasControllerDisconnectionBehavior returns a boolean if a field has been set.
func (o *InlineResponse200193) HasControllerDisconnectionBehavior() bool {
	if o != nil && !isNil(o.ControllerDisconnectionBehavior) {
		return true
	}

	return false
}

// SetControllerDisconnectionBehavior gets a reference to the given string and assigns it to the ControllerDisconnectionBehavior field.
func (o *InlineResponse200193) SetControllerDisconnectionBehavior(v string) {
	o.ControllerDisconnectionBehavior = &v
}

// GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field value if set, zero value otherwise.
func (o *InlineResponse200193) GetAllowSimultaneousLogins() bool {
	if o == nil || isNil(o.AllowSimultaneousLogins) {
		var ret bool
		return ret
	}
	return *o.AllowSimultaneousLogins
}

// GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetAllowSimultaneousLoginsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowSimultaneousLogins) {
    return nil, false
	}
	return o.AllowSimultaneousLogins, true
}

// HasAllowSimultaneousLogins returns a boolean if a field has been set.
func (o *InlineResponse200193) HasAllowSimultaneousLogins() bool {
	if o != nil && !isNil(o.AllowSimultaneousLogins) {
		return true
	}

	return false
}

// SetAllowSimultaneousLogins gets a reference to the given bool and assigns it to the AllowSimultaneousLogins field.
func (o *InlineResponse200193) SetAllowSimultaneousLogins(v bool) {
	o.AllowSimultaneousLogins = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *InlineResponse200193) GetBilling() InlineResponse200193Billing {
	if o == nil || isNil(o.Billing) {
		var ret InlineResponse200193Billing
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetBillingOk() (*InlineResponse200193Billing, bool) {
	if o == nil || isNil(o.Billing) {
    return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *InlineResponse200193) HasBilling() bool {
	if o != nil && !isNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given InlineResponse200193Billing and assigns it to the Billing field.
func (o *InlineResponse200193) SetBilling(v InlineResponse200193Billing) {
	o.Billing = &v
}

// GetSentryEnrollment returns the SentryEnrollment field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSentryEnrollment() InlineResponse200193SentryEnrollment {
	if o == nil || isNil(o.SentryEnrollment) {
		var ret InlineResponse200193SentryEnrollment
		return ret
	}
	return *o.SentryEnrollment
}

// GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSentryEnrollmentOk() (*InlineResponse200193SentryEnrollment, bool) {
	if o == nil || isNil(o.SentryEnrollment) {
    return nil, false
	}
	return o.SentryEnrollment, true
}

// HasSentryEnrollment returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSentryEnrollment() bool {
	if o != nil && !isNil(o.SentryEnrollment) {
		return true
	}

	return false
}

// SetSentryEnrollment gets a reference to the given InlineResponse200193SentryEnrollment and assigns it to the SentryEnrollment field.
func (o *InlineResponse200193) SetSentryEnrollment(v InlineResponse200193SentryEnrollment) {
	o.SentryEnrollment = &v
}

// GetSelfRegistration returns the SelfRegistration field value if set, zero value otherwise.
func (o *InlineResponse200193) GetSelfRegistration() InlineResponse200193SelfRegistration {
	if o == nil || isNil(o.SelfRegistration) {
		var ret InlineResponse200193SelfRegistration
		return ret
	}
	return *o.SelfRegistration
}

// GetSelfRegistrationOk returns a tuple with the SelfRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193) GetSelfRegistrationOk() (*InlineResponse200193SelfRegistration, bool) {
	if o == nil || isNil(o.SelfRegistration) {
    return nil, false
	}
	return o.SelfRegistration, true
}

// HasSelfRegistration returns a boolean if a field has been set.
func (o *InlineResponse200193) HasSelfRegistration() bool {
	if o != nil && !isNil(o.SelfRegistration) {
		return true
	}

	return false
}

// SetSelfRegistration gets a reference to the given InlineResponse200193SelfRegistration and assigns it to the SelfRegistration field.
func (o *InlineResponse200193) SetSelfRegistration(v InlineResponse200193SelfRegistration) {
	o.SelfRegistration = &v
}

func (o InlineResponse200193) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200193 struct {
	value *InlineResponse200193
	isSet bool
}

func (v NullableInlineResponse200193) Get() *InlineResponse200193 {
	return v.value
}

func (v *NullableInlineResponse200193) Set(val *InlineResponse200193) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200193) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200193) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200193(val *InlineResponse200193) *NullableInlineResponse200193 {
	return &NullableInlineResponse200193{value: val, isSet: true}
}

func (v NullableInlineResponse200193) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200193) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



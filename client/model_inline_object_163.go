/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject163 struct for InlineObject163
type InlineObject163 struct {
	// [optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see 'useSplashUrl'
	SplashUrl *string `json:"splashUrl,omitempty"`
	// [optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID's access control settings, it may not be possible to use the custom splash URL.
	UseSplashUrl *bool `json:"useSplashUrl,omitempty"`
	// Splash timeout in minutes. This will determine how often users will see the splash page.
	SplashTimeout *int32 `json:"splashTimeout,omitempty"`
	// The custom redirect URL where the users will go after the splash page.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true.
	UseRedirectUrl *bool `json:"useRedirectUrl,omitempty"`
	// The welcome message for the users on the splash page.
	WelcomeMessage *string `json:"welcomeMessage,omitempty"`
	SplashLogo *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo `json:"splashLogo,omitempty"`
	SplashImage *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage `json:"splashImage,omitempty"`
	SplashPrepaidFront *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront `json:"splashPrepaidFront,omitempty"`
	// How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.
	BlockAllTrafficBeforeSignOn *bool `json:"blockAllTrafficBeforeSignOn,omitempty"`
	// How login attempts should be handled when the controller is unreachable. Can be either 'open', 'restricted', or 'default'.
	ControllerDisconnectionBehavior *string `json:"controllerDisconnectionBehavior,omitempty"`
	// Whether or not to allow simultaneous logins from different devices.
	AllowSimultaneousLogins *bool `json:"allowSimultaneousLogins,omitempty"`
	GuestSponsorship *NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship `json:"guestSponsorship,omitempty"`
	Billing *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling `json:"billing,omitempty"`
	SentryEnrollment *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment `json:"sentryEnrollment,omitempty"`
}

// NewInlineObject163 instantiates a new InlineObject163 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject163() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// NewInlineObject163WithDefaults instantiates a new InlineObject163 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject163WithDefaults() *InlineObject163 {
	this := InlineObject163{}
	return &this
}

// GetSplashUrl returns the SplashUrl field value if set, zero value otherwise.
func (o *InlineObject163) GetSplashUrl() string {
	if o == nil || o.SplashUrl == nil {
		var ret string
		return ret
	}
	return *o.SplashUrl
}

// GetSplashUrlOk returns a tuple with the SplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetSplashUrlOk() (*string, bool) {
	if o == nil || o.SplashUrl == nil {
		return nil, false
	}
	return o.SplashUrl, true
}

// HasSplashUrl returns a boolean if a field has been set.
func (o *InlineObject163) HasSplashUrl() bool {
	if o != nil && o.SplashUrl != nil {
		return true
	}

	return false
}

// SetSplashUrl gets a reference to the given string and assigns it to the SplashUrl field.
func (o *InlineObject163) SetSplashUrl(v string) {
	o.SplashUrl = &v
}

// GetUseSplashUrl returns the UseSplashUrl field value if set, zero value otherwise.
func (o *InlineObject163) GetUseSplashUrl() bool {
	if o == nil || o.UseSplashUrl == nil {
		var ret bool
		return ret
	}
	return *o.UseSplashUrl
}

// GetUseSplashUrlOk returns a tuple with the UseSplashUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetUseSplashUrlOk() (*bool, bool) {
	if o == nil || o.UseSplashUrl == nil {
		return nil, false
	}
	return o.UseSplashUrl, true
}

// HasUseSplashUrl returns a boolean if a field has been set.
func (o *InlineObject163) HasUseSplashUrl() bool {
	if o != nil && o.UseSplashUrl != nil {
		return true
	}

	return false
}

// SetUseSplashUrl gets a reference to the given bool and assigns it to the UseSplashUrl field.
func (o *InlineObject163) SetUseSplashUrl(v bool) {
	o.UseSplashUrl = &v
}

// GetSplashTimeout returns the SplashTimeout field value if set, zero value otherwise.
func (o *InlineObject163) GetSplashTimeout() int32 {
	if o == nil || o.SplashTimeout == nil {
		var ret int32
		return ret
	}
	return *o.SplashTimeout
}

// GetSplashTimeoutOk returns a tuple with the SplashTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetSplashTimeoutOk() (*int32, bool) {
	if o == nil || o.SplashTimeout == nil {
		return nil, false
	}
	return o.SplashTimeout, true
}

// HasSplashTimeout returns a boolean if a field has been set.
func (o *InlineObject163) HasSplashTimeout() bool {
	if o != nil && o.SplashTimeout != nil {
		return true
	}

	return false
}

// SetSplashTimeout gets a reference to the given int32 and assigns it to the SplashTimeout field.
func (o *InlineObject163) SetSplashTimeout(v int32) {
	o.SplashTimeout = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *InlineObject163) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *InlineObject163) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *InlineObject163) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetUseRedirectUrl returns the UseRedirectUrl field value if set, zero value otherwise.
func (o *InlineObject163) GetUseRedirectUrl() bool {
	if o == nil || o.UseRedirectUrl == nil {
		var ret bool
		return ret
	}
	return *o.UseRedirectUrl
}

// GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetUseRedirectUrlOk() (*bool, bool) {
	if o == nil || o.UseRedirectUrl == nil {
		return nil, false
	}
	return o.UseRedirectUrl, true
}

// HasUseRedirectUrl returns a boolean if a field has been set.
func (o *InlineObject163) HasUseRedirectUrl() bool {
	if o != nil && o.UseRedirectUrl != nil {
		return true
	}

	return false
}

// SetUseRedirectUrl gets a reference to the given bool and assigns it to the UseRedirectUrl field.
func (o *InlineObject163) SetUseRedirectUrl(v bool) {
	o.UseRedirectUrl = &v
}

// GetWelcomeMessage returns the WelcomeMessage field value if set, zero value otherwise.
func (o *InlineObject163) GetWelcomeMessage() string {
	if o == nil || o.WelcomeMessage == nil {
		var ret string
		return ret
	}
	return *o.WelcomeMessage
}

// GetWelcomeMessageOk returns a tuple with the WelcomeMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetWelcomeMessageOk() (*string, bool) {
	if o == nil || o.WelcomeMessage == nil {
		return nil, false
	}
	return o.WelcomeMessage, true
}

// HasWelcomeMessage returns a boolean if a field has been set.
func (o *InlineObject163) HasWelcomeMessage() bool {
	if o != nil && o.WelcomeMessage != nil {
		return true
	}

	return false
}

// SetWelcomeMessage gets a reference to the given string and assigns it to the WelcomeMessage field.
func (o *InlineObject163) SetWelcomeMessage(v string) {
	o.WelcomeMessage = &v
}

// GetSplashLogo returns the SplashLogo field value if set, zero value otherwise.
func (o *InlineObject163) GetSplashLogo() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo {
	if o == nil || o.SplashLogo == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo
		return ret
	}
	return *o.SplashLogo
}

// GetSplashLogoOk returns a tuple with the SplashLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetSplashLogoOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo, bool) {
	if o == nil || o.SplashLogo == nil {
		return nil, false
	}
	return o.SplashLogo, true
}

// HasSplashLogo returns a boolean if a field has been set.
func (o *InlineObject163) HasSplashLogo() bool {
	if o != nil && o.SplashLogo != nil {
		return true
	}

	return false
}

// SetSplashLogo gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo and assigns it to the SplashLogo field.
func (o *InlineObject163) SetSplashLogo(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) {
	o.SplashLogo = &v
}

// GetSplashImage returns the SplashImage field value if set, zero value otherwise.
func (o *InlineObject163) GetSplashImage() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage {
	if o == nil || o.SplashImage == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage
		return ret
	}
	return *o.SplashImage
}

// GetSplashImageOk returns a tuple with the SplashImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetSplashImageOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage, bool) {
	if o == nil || o.SplashImage == nil {
		return nil, false
	}
	return o.SplashImage, true
}

// HasSplashImage returns a boolean if a field has been set.
func (o *InlineObject163) HasSplashImage() bool {
	if o != nil && o.SplashImage != nil {
		return true
	}

	return false
}

// SetSplashImage gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage and assigns it to the SplashImage field.
func (o *InlineObject163) SetSplashImage(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) {
	o.SplashImage = &v
}

// GetSplashPrepaidFront returns the SplashPrepaidFront field value if set, zero value otherwise.
func (o *InlineObject163) GetSplashPrepaidFront() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront {
	if o == nil || o.SplashPrepaidFront == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront
		return ret
	}
	return *o.SplashPrepaidFront
}

// GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetSplashPrepaidFrontOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront, bool) {
	if o == nil || o.SplashPrepaidFront == nil {
		return nil, false
	}
	return o.SplashPrepaidFront, true
}

// HasSplashPrepaidFront returns a boolean if a field has been set.
func (o *InlineObject163) HasSplashPrepaidFront() bool {
	if o != nil && o.SplashPrepaidFront != nil {
		return true
	}

	return false
}

// SetSplashPrepaidFront gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront and assigns it to the SplashPrepaidFront field.
func (o *InlineObject163) SetSplashPrepaidFront(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFront) {
	o.SplashPrepaidFront = &v
}

// GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field value if set, zero value otherwise.
func (o *InlineObject163) GetBlockAllTrafficBeforeSignOn() bool {
	if o == nil || o.BlockAllTrafficBeforeSignOn == nil {
		var ret bool
		return ret
	}
	return *o.BlockAllTrafficBeforeSignOn
}

// GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool) {
	if o == nil || o.BlockAllTrafficBeforeSignOn == nil {
		return nil, false
	}
	return o.BlockAllTrafficBeforeSignOn, true
}

// HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.
func (o *InlineObject163) HasBlockAllTrafficBeforeSignOn() bool {
	if o != nil && o.BlockAllTrafficBeforeSignOn != nil {
		return true
	}

	return false
}

// SetBlockAllTrafficBeforeSignOn gets a reference to the given bool and assigns it to the BlockAllTrafficBeforeSignOn field.
func (o *InlineObject163) SetBlockAllTrafficBeforeSignOn(v bool) {
	o.BlockAllTrafficBeforeSignOn = &v
}

// GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field value if set, zero value otherwise.
func (o *InlineObject163) GetControllerDisconnectionBehavior() string {
	if o == nil || o.ControllerDisconnectionBehavior == nil {
		var ret string
		return ret
	}
	return *o.ControllerDisconnectionBehavior
}

// GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetControllerDisconnectionBehaviorOk() (*string, bool) {
	if o == nil || o.ControllerDisconnectionBehavior == nil {
		return nil, false
	}
	return o.ControllerDisconnectionBehavior, true
}

// HasControllerDisconnectionBehavior returns a boolean if a field has been set.
func (o *InlineObject163) HasControllerDisconnectionBehavior() bool {
	if o != nil && o.ControllerDisconnectionBehavior != nil {
		return true
	}

	return false
}

// SetControllerDisconnectionBehavior gets a reference to the given string and assigns it to the ControllerDisconnectionBehavior field.
func (o *InlineObject163) SetControllerDisconnectionBehavior(v string) {
	o.ControllerDisconnectionBehavior = &v
}

// GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field value if set, zero value otherwise.
func (o *InlineObject163) GetAllowSimultaneousLogins() bool {
	if o == nil || o.AllowSimultaneousLogins == nil {
		var ret bool
		return ret
	}
	return *o.AllowSimultaneousLogins
}

// GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetAllowSimultaneousLoginsOk() (*bool, bool) {
	if o == nil || o.AllowSimultaneousLogins == nil {
		return nil, false
	}
	return o.AllowSimultaneousLogins, true
}

// HasAllowSimultaneousLogins returns a boolean if a field has been set.
func (o *InlineObject163) HasAllowSimultaneousLogins() bool {
	if o != nil && o.AllowSimultaneousLogins != nil {
		return true
	}

	return false
}

// SetAllowSimultaneousLogins gets a reference to the given bool and assigns it to the AllowSimultaneousLogins field.
func (o *InlineObject163) SetAllowSimultaneousLogins(v bool) {
	o.AllowSimultaneousLogins = &v
}

// GetGuestSponsorship returns the GuestSponsorship field value if set, zero value otherwise.
func (o *InlineObject163) GetGuestSponsorship() NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship {
	if o == nil || o.GuestSponsorship == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship
		return ret
	}
	return *o.GuestSponsorship
}

// GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetGuestSponsorshipOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship, bool) {
	if o == nil || o.GuestSponsorship == nil {
		return nil, false
	}
	return o.GuestSponsorship, true
}

// HasGuestSponsorship returns a boolean if a field has been set.
func (o *InlineObject163) HasGuestSponsorship() bool {
	if o != nil && o.GuestSponsorship != nil {
		return true
	}

	return false
}

// SetGuestSponsorship gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship and assigns it to the GuestSponsorship field.
func (o *InlineObject163) SetGuestSponsorship(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsGuestSponsorship) {
	o.GuestSponsorship = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *InlineObject163) GetBilling() NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling {
	if o == nil || o.Billing == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetBillingOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling, bool) {
	if o == nil || o.Billing == nil {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *InlineObject163) HasBilling() bool {
	if o != nil && o.Billing != nil {
		return true
	}

	return false
}

// SetBilling gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling and assigns it to the Billing field.
func (o *InlineObject163) SetBilling(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) {
	o.Billing = &v
}

// GetSentryEnrollment returns the SentryEnrollment field value if set, zero value otherwise.
func (o *InlineObject163) GetSentryEnrollment() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment {
	if o == nil || o.SentryEnrollment == nil {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment
		return ret
	}
	return *o.SentryEnrollment
}

// GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject163) GetSentryEnrollmentOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment, bool) {
	if o == nil || o.SentryEnrollment == nil {
		return nil, false
	}
	return o.SentryEnrollment, true
}

// HasSentryEnrollment returns a boolean if a field has been set.
func (o *InlineObject163) HasSentryEnrollment() bool {
	if o != nil && o.SentryEnrollment != nil {
		return true
	}

	return false
}

// SetSentryEnrollment gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment and assigns it to the SentryEnrollment field.
func (o *InlineObject163) SetSentryEnrollment(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) {
	o.SentryEnrollment = &v
}

func (o InlineObject163) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SplashUrl != nil {
		toSerialize["splashUrl"] = o.SplashUrl
	}
	if o.UseSplashUrl != nil {
		toSerialize["useSplashUrl"] = o.UseSplashUrl
	}
	if o.SplashTimeout != nil {
		toSerialize["splashTimeout"] = o.SplashTimeout
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if o.UseRedirectUrl != nil {
		toSerialize["useRedirectUrl"] = o.UseRedirectUrl
	}
	if o.WelcomeMessage != nil {
		toSerialize["welcomeMessage"] = o.WelcomeMessage
	}
	if o.SplashLogo != nil {
		toSerialize["splashLogo"] = o.SplashLogo
	}
	if o.SplashImage != nil {
		toSerialize["splashImage"] = o.SplashImage
	}
	if o.SplashPrepaidFront != nil {
		toSerialize["splashPrepaidFront"] = o.SplashPrepaidFront
	}
	if o.BlockAllTrafficBeforeSignOn != nil {
		toSerialize["blockAllTrafficBeforeSignOn"] = o.BlockAllTrafficBeforeSignOn
	}
	if o.ControllerDisconnectionBehavior != nil {
		toSerialize["controllerDisconnectionBehavior"] = o.ControllerDisconnectionBehavior
	}
	if o.AllowSimultaneousLogins != nil {
		toSerialize["allowSimultaneousLogins"] = o.AllowSimultaneousLogins
	}
	if o.GuestSponsorship != nil {
		toSerialize["guestSponsorship"] = o.GuestSponsorship
	}
	if o.Billing != nil {
		toSerialize["billing"] = o.Billing
	}
	if o.SentryEnrollment != nil {
		toSerialize["sentryEnrollment"] = o.SentryEnrollment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject163 struct {
	value *InlineObject163
	isSet bool
}

func (v NullableInlineObject163) Get() *InlineObject163 {
	return v.value
}

func (v *NullableInlineObject163) Set(val *InlineObject163) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject163) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject163) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject163(val *InlineObject163) *NullableInlineObject163 {
	return &NullableInlineObject163{value: val, isSet: true}
}

func (v NullableInlineObject163) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject163) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



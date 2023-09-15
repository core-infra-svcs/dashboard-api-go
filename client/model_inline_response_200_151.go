/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200151 struct for InlineResponse200151
type InlineResponse200151 struct {
	// Boolean indicating whether users are forced to change their password every X number of days.
	EnforcePasswordExpiration *bool `json:"enforcePasswordExpiration,omitempty"`
	// Number of days after which users will be forced to change their password.
	PasswordExpirationDays *int32 `json:"passwordExpirationDays,omitempty"`
	// Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
	EnforceDifferentPasswords *bool `json:"enforceDifferentPasswords,omitempty"`
	// Number of recent passwords that new password must be distinct from.
	NumDifferentPasswords *int32 `json:"numDifferentPasswords,omitempty"`
	// Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol
	EnforceStrongPasswords *bool `json:"enforceStrongPasswords,omitempty"`
	// Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
	EnforceAccountLockout *bool `json:"enforceAccountLockout,omitempty"`
	// Number of consecutive failed login attempts after which users' accounts will be locked.
	AccountLockoutAttempts *int32 `json:"accountLockoutAttempts,omitempty"`
	// Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
	EnforceIdleTimeout *bool `json:"enforceIdleTimeout,omitempty"`
	// Number of minutes users can remain idle before being logged out of their accounts.
	IdleTimeoutMinutes *int32 `json:"idleTimeoutMinutes,omitempty"`
	// Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.
	EnforceTwoFactorAuth *bool `json:"enforceTwoFactorAuth,omitempty"`
	// Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
	EnforceLoginIpRanges *bool `json:"enforceLoginIpRanges,omitempty"`
	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	LoginIpRanges []string `json:"loginIpRanges,omitempty"`
	ApiAuthentication *InlineResponse200151ApiAuthentication `json:"apiAuthentication,omitempty"`
}

// NewInlineResponse200151 instantiates a new InlineResponse200151 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200151() *InlineResponse200151 {
	this := InlineResponse200151{}
	return &this
}

// NewInlineResponse200151WithDefaults instantiates a new InlineResponse200151 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200151WithDefaults() *InlineResponse200151 {
	this := InlineResponse200151{}
	return &this
}

// GetEnforcePasswordExpiration returns the EnforcePasswordExpiration field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforcePasswordExpiration() bool {
	if o == nil || isNil(o.EnforcePasswordExpiration) {
		var ret bool
		return ret
	}
	return *o.EnforcePasswordExpiration
}

// GetEnforcePasswordExpirationOk returns a tuple with the EnforcePasswordExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforcePasswordExpirationOk() (*bool, bool) {
	if o == nil || isNil(o.EnforcePasswordExpiration) {
    return nil, false
	}
	return o.EnforcePasswordExpiration, true
}

// HasEnforcePasswordExpiration returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforcePasswordExpiration() bool {
	if o != nil && !isNil(o.EnforcePasswordExpiration) {
		return true
	}

	return false
}

// SetEnforcePasswordExpiration gets a reference to the given bool and assigns it to the EnforcePasswordExpiration field.
func (o *InlineResponse200151) SetEnforcePasswordExpiration(v bool) {
	o.EnforcePasswordExpiration = &v
}

// GetPasswordExpirationDays returns the PasswordExpirationDays field value if set, zero value otherwise.
func (o *InlineResponse200151) GetPasswordExpirationDays() int32 {
	if o == nil || isNil(o.PasswordExpirationDays) {
		var ret int32
		return ret
	}
	return *o.PasswordExpirationDays
}

// GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetPasswordExpirationDaysOk() (*int32, bool) {
	if o == nil || isNil(o.PasswordExpirationDays) {
    return nil, false
	}
	return o.PasswordExpirationDays, true
}

// HasPasswordExpirationDays returns a boolean if a field has been set.
func (o *InlineResponse200151) HasPasswordExpirationDays() bool {
	if o != nil && !isNil(o.PasswordExpirationDays) {
		return true
	}

	return false
}

// SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.
func (o *InlineResponse200151) SetPasswordExpirationDays(v int32) {
	o.PasswordExpirationDays = &v
}

// GetEnforceDifferentPasswords returns the EnforceDifferentPasswords field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforceDifferentPasswords() bool {
	if o == nil || isNil(o.EnforceDifferentPasswords) {
		var ret bool
		return ret
	}
	return *o.EnforceDifferentPasswords
}

// GetEnforceDifferentPasswordsOk returns a tuple with the EnforceDifferentPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforceDifferentPasswordsOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceDifferentPasswords) {
    return nil, false
	}
	return o.EnforceDifferentPasswords, true
}

// HasEnforceDifferentPasswords returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforceDifferentPasswords() bool {
	if o != nil && !isNil(o.EnforceDifferentPasswords) {
		return true
	}

	return false
}

// SetEnforceDifferentPasswords gets a reference to the given bool and assigns it to the EnforceDifferentPasswords field.
func (o *InlineResponse200151) SetEnforceDifferentPasswords(v bool) {
	o.EnforceDifferentPasswords = &v
}

// GetNumDifferentPasswords returns the NumDifferentPasswords field value if set, zero value otherwise.
func (o *InlineResponse200151) GetNumDifferentPasswords() int32 {
	if o == nil || isNil(o.NumDifferentPasswords) {
		var ret int32
		return ret
	}
	return *o.NumDifferentPasswords
}

// GetNumDifferentPasswordsOk returns a tuple with the NumDifferentPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetNumDifferentPasswordsOk() (*int32, bool) {
	if o == nil || isNil(o.NumDifferentPasswords) {
    return nil, false
	}
	return o.NumDifferentPasswords, true
}

// HasNumDifferentPasswords returns a boolean if a field has been set.
func (o *InlineResponse200151) HasNumDifferentPasswords() bool {
	if o != nil && !isNil(o.NumDifferentPasswords) {
		return true
	}

	return false
}

// SetNumDifferentPasswords gets a reference to the given int32 and assigns it to the NumDifferentPasswords field.
func (o *InlineResponse200151) SetNumDifferentPasswords(v int32) {
	o.NumDifferentPasswords = &v
}

// GetEnforceStrongPasswords returns the EnforceStrongPasswords field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforceStrongPasswords() bool {
	if o == nil || isNil(o.EnforceStrongPasswords) {
		var ret bool
		return ret
	}
	return *o.EnforceStrongPasswords
}

// GetEnforceStrongPasswordsOk returns a tuple with the EnforceStrongPasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforceStrongPasswordsOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceStrongPasswords) {
    return nil, false
	}
	return o.EnforceStrongPasswords, true
}

// HasEnforceStrongPasswords returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforceStrongPasswords() bool {
	if o != nil && !isNil(o.EnforceStrongPasswords) {
		return true
	}

	return false
}

// SetEnforceStrongPasswords gets a reference to the given bool and assigns it to the EnforceStrongPasswords field.
func (o *InlineResponse200151) SetEnforceStrongPasswords(v bool) {
	o.EnforceStrongPasswords = &v
}

// GetEnforceAccountLockout returns the EnforceAccountLockout field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforceAccountLockout() bool {
	if o == nil || isNil(o.EnforceAccountLockout) {
		var ret bool
		return ret
	}
	return *o.EnforceAccountLockout
}

// GetEnforceAccountLockoutOk returns a tuple with the EnforceAccountLockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforceAccountLockoutOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceAccountLockout) {
    return nil, false
	}
	return o.EnforceAccountLockout, true
}

// HasEnforceAccountLockout returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforceAccountLockout() bool {
	if o != nil && !isNil(o.EnforceAccountLockout) {
		return true
	}

	return false
}

// SetEnforceAccountLockout gets a reference to the given bool and assigns it to the EnforceAccountLockout field.
func (o *InlineResponse200151) SetEnforceAccountLockout(v bool) {
	o.EnforceAccountLockout = &v
}

// GetAccountLockoutAttempts returns the AccountLockoutAttempts field value if set, zero value otherwise.
func (o *InlineResponse200151) GetAccountLockoutAttempts() int32 {
	if o == nil || isNil(o.AccountLockoutAttempts) {
		var ret int32
		return ret
	}
	return *o.AccountLockoutAttempts
}

// GetAccountLockoutAttemptsOk returns a tuple with the AccountLockoutAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetAccountLockoutAttemptsOk() (*int32, bool) {
	if o == nil || isNil(o.AccountLockoutAttempts) {
    return nil, false
	}
	return o.AccountLockoutAttempts, true
}

// HasAccountLockoutAttempts returns a boolean if a field has been set.
func (o *InlineResponse200151) HasAccountLockoutAttempts() bool {
	if o != nil && !isNil(o.AccountLockoutAttempts) {
		return true
	}

	return false
}

// SetAccountLockoutAttempts gets a reference to the given int32 and assigns it to the AccountLockoutAttempts field.
func (o *InlineResponse200151) SetAccountLockoutAttempts(v int32) {
	o.AccountLockoutAttempts = &v
}

// GetEnforceIdleTimeout returns the EnforceIdleTimeout field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforceIdleTimeout() bool {
	if o == nil || isNil(o.EnforceIdleTimeout) {
		var ret bool
		return ret
	}
	return *o.EnforceIdleTimeout
}

// GetEnforceIdleTimeoutOk returns a tuple with the EnforceIdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforceIdleTimeoutOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceIdleTimeout) {
    return nil, false
	}
	return o.EnforceIdleTimeout, true
}

// HasEnforceIdleTimeout returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforceIdleTimeout() bool {
	if o != nil && !isNil(o.EnforceIdleTimeout) {
		return true
	}

	return false
}

// SetEnforceIdleTimeout gets a reference to the given bool and assigns it to the EnforceIdleTimeout field.
func (o *InlineResponse200151) SetEnforceIdleTimeout(v bool) {
	o.EnforceIdleTimeout = &v
}

// GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field value if set, zero value otherwise.
func (o *InlineResponse200151) GetIdleTimeoutMinutes() int32 {
	if o == nil || isNil(o.IdleTimeoutMinutes) {
		var ret int32
		return ret
	}
	return *o.IdleTimeoutMinutes
}

// GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetIdleTimeoutMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.IdleTimeoutMinutes) {
    return nil, false
	}
	return o.IdleTimeoutMinutes, true
}

// HasIdleTimeoutMinutes returns a boolean if a field has been set.
func (o *InlineResponse200151) HasIdleTimeoutMinutes() bool {
	if o != nil && !isNil(o.IdleTimeoutMinutes) {
		return true
	}

	return false
}

// SetIdleTimeoutMinutes gets a reference to the given int32 and assigns it to the IdleTimeoutMinutes field.
func (o *InlineResponse200151) SetIdleTimeoutMinutes(v int32) {
	o.IdleTimeoutMinutes = &v
}

// GetEnforceTwoFactorAuth returns the EnforceTwoFactorAuth field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforceTwoFactorAuth() bool {
	if o == nil || isNil(o.EnforceTwoFactorAuth) {
		var ret bool
		return ret
	}
	return *o.EnforceTwoFactorAuth
}

// GetEnforceTwoFactorAuthOk returns a tuple with the EnforceTwoFactorAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforceTwoFactorAuthOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceTwoFactorAuth) {
    return nil, false
	}
	return o.EnforceTwoFactorAuth, true
}

// HasEnforceTwoFactorAuth returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforceTwoFactorAuth() bool {
	if o != nil && !isNil(o.EnforceTwoFactorAuth) {
		return true
	}

	return false
}

// SetEnforceTwoFactorAuth gets a reference to the given bool and assigns it to the EnforceTwoFactorAuth field.
func (o *InlineResponse200151) SetEnforceTwoFactorAuth(v bool) {
	o.EnforceTwoFactorAuth = &v
}

// GetEnforceLoginIpRanges returns the EnforceLoginIpRanges field value if set, zero value otherwise.
func (o *InlineResponse200151) GetEnforceLoginIpRanges() bool {
	if o == nil || isNil(o.EnforceLoginIpRanges) {
		var ret bool
		return ret
	}
	return *o.EnforceLoginIpRanges
}

// GetEnforceLoginIpRangesOk returns a tuple with the EnforceLoginIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetEnforceLoginIpRangesOk() (*bool, bool) {
	if o == nil || isNil(o.EnforceLoginIpRanges) {
    return nil, false
	}
	return o.EnforceLoginIpRanges, true
}

// HasEnforceLoginIpRanges returns a boolean if a field has been set.
func (o *InlineResponse200151) HasEnforceLoginIpRanges() bool {
	if o != nil && !isNil(o.EnforceLoginIpRanges) {
		return true
	}

	return false
}

// SetEnforceLoginIpRanges gets a reference to the given bool and assigns it to the EnforceLoginIpRanges field.
func (o *InlineResponse200151) SetEnforceLoginIpRanges(v bool) {
	o.EnforceLoginIpRanges = &v
}

// GetLoginIpRanges returns the LoginIpRanges field value if set, zero value otherwise.
func (o *InlineResponse200151) GetLoginIpRanges() []string {
	if o == nil || isNil(o.LoginIpRanges) {
		var ret []string
		return ret
	}
	return o.LoginIpRanges
}

// GetLoginIpRangesOk returns a tuple with the LoginIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetLoginIpRangesOk() ([]string, bool) {
	if o == nil || isNil(o.LoginIpRanges) {
    return nil, false
	}
	return o.LoginIpRanges, true
}

// HasLoginIpRanges returns a boolean if a field has been set.
func (o *InlineResponse200151) HasLoginIpRanges() bool {
	if o != nil && !isNil(o.LoginIpRanges) {
		return true
	}

	return false
}

// SetLoginIpRanges gets a reference to the given []string and assigns it to the LoginIpRanges field.
func (o *InlineResponse200151) SetLoginIpRanges(v []string) {
	o.LoginIpRanges = v
}

// GetApiAuthentication returns the ApiAuthentication field value if set, zero value otherwise.
func (o *InlineResponse200151) GetApiAuthentication() InlineResponse200151ApiAuthentication {
	if o == nil || isNil(o.ApiAuthentication) {
		var ret InlineResponse200151ApiAuthentication
		return ret
	}
	return *o.ApiAuthentication
}

// GetApiAuthenticationOk returns a tuple with the ApiAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetApiAuthenticationOk() (*InlineResponse200151ApiAuthentication, bool) {
	if o == nil || isNil(o.ApiAuthentication) {
    return nil, false
	}
	return o.ApiAuthentication, true
}

// HasApiAuthentication returns a boolean if a field has been set.
func (o *InlineResponse200151) HasApiAuthentication() bool {
	if o != nil && !isNil(o.ApiAuthentication) {
		return true
	}

	return false
}

// SetApiAuthentication gets a reference to the given InlineResponse200151ApiAuthentication and assigns it to the ApiAuthentication field.
func (o *InlineResponse200151) SetApiAuthentication(v InlineResponse200151ApiAuthentication) {
	o.ApiAuthentication = &v
}

func (o InlineResponse200151) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnforcePasswordExpiration) {
		toSerialize["enforcePasswordExpiration"] = o.EnforcePasswordExpiration
	}
	if !isNil(o.PasswordExpirationDays) {
		toSerialize["passwordExpirationDays"] = o.PasswordExpirationDays
	}
	if !isNil(o.EnforceDifferentPasswords) {
		toSerialize["enforceDifferentPasswords"] = o.EnforceDifferentPasswords
	}
	if !isNil(o.NumDifferentPasswords) {
		toSerialize["numDifferentPasswords"] = o.NumDifferentPasswords
	}
	if !isNil(o.EnforceStrongPasswords) {
		toSerialize["enforceStrongPasswords"] = o.EnforceStrongPasswords
	}
	if !isNil(o.EnforceAccountLockout) {
		toSerialize["enforceAccountLockout"] = o.EnforceAccountLockout
	}
	if !isNil(o.AccountLockoutAttempts) {
		toSerialize["accountLockoutAttempts"] = o.AccountLockoutAttempts
	}
	if !isNil(o.EnforceIdleTimeout) {
		toSerialize["enforceIdleTimeout"] = o.EnforceIdleTimeout
	}
	if !isNil(o.IdleTimeoutMinutes) {
		toSerialize["idleTimeoutMinutes"] = o.IdleTimeoutMinutes
	}
	if !isNil(o.EnforceTwoFactorAuth) {
		toSerialize["enforceTwoFactorAuth"] = o.EnforceTwoFactorAuth
	}
	if !isNil(o.EnforceLoginIpRanges) {
		toSerialize["enforceLoginIpRanges"] = o.EnforceLoginIpRanges
	}
	if !isNil(o.LoginIpRanges) {
		toSerialize["loginIpRanges"] = o.LoginIpRanges
	}
	if !isNil(o.ApiAuthentication) {
		toSerialize["apiAuthentication"] = o.ApiAuthentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200151 struct {
	value *InlineResponse200151
	isSet bool
}

func (v NullableInlineResponse200151) Get() *InlineResponse200151 {
	return v.value
}

func (v *NullableInlineResponse200151) Set(val *InlineResponse200151) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200151) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200151) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200151(val *InlineResponse200151) *NullableInlineResponse200151 {
	return &NullableInlineResponse200151{value: val, isSet: true}
}

func (v NullableInlineResponse200151) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200151) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



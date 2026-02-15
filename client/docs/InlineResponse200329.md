# InlineResponse200329

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpId** | Pointer to **string** | ID associated with the SAML Identity Provider (IdP) | [optional] 
**ConsumerUrl** | Pointer to **string** | URL that is consuming SAML Identity Provider (IdP) | [optional] 
**VisionConsumerUrl** | Pointer to **string** | URL that is consuming SAML Identity Provider (IdP) for Meraki Vision Portal | [optional] 
**X509certSha1Fingerprint** | Pointer to **string** | Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation. | [optional] 
**SsoLoginUrl** | Pointer to **string** | Dashboard will redirect users to this URL to log in again when their sessions expire. | [optional] 
**SloLogoutUrl** | Pointer to **string** | Dashboard will redirect users to this URL when they sign out. | [optional] 

## Methods

### NewInlineResponse200329

`func NewInlineResponse200329() *InlineResponse200329`

NewInlineResponse200329 instantiates a new InlineResponse200329 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200329WithDefaults

`func NewInlineResponse200329WithDefaults() *InlineResponse200329`

NewInlineResponse200329WithDefaults instantiates a new InlineResponse200329 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpId

`func (o *InlineResponse200329) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *InlineResponse200329) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *InlineResponse200329) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *InlineResponse200329) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### GetConsumerUrl

`func (o *InlineResponse200329) GetConsumerUrl() string`

GetConsumerUrl returns the ConsumerUrl field if non-nil, zero value otherwise.

### GetConsumerUrlOk

`func (o *InlineResponse200329) GetConsumerUrlOk() (*string, bool)`

GetConsumerUrlOk returns a tuple with the ConsumerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUrl

`func (o *InlineResponse200329) SetConsumerUrl(v string)`

SetConsumerUrl sets ConsumerUrl field to given value.

### HasConsumerUrl

`func (o *InlineResponse200329) HasConsumerUrl() bool`

HasConsumerUrl returns a boolean if a field has been set.

### GetVisionConsumerUrl

`func (o *InlineResponse200329) GetVisionConsumerUrl() string`

GetVisionConsumerUrl returns the VisionConsumerUrl field if non-nil, zero value otherwise.

### GetVisionConsumerUrlOk

`func (o *InlineResponse200329) GetVisionConsumerUrlOk() (*string, bool)`

GetVisionConsumerUrlOk returns a tuple with the VisionConsumerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionConsumerUrl

`func (o *InlineResponse200329) SetVisionConsumerUrl(v string)`

SetVisionConsumerUrl sets VisionConsumerUrl field to given value.

### HasVisionConsumerUrl

`func (o *InlineResponse200329) HasVisionConsumerUrl() bool`

HasVisionConsumerUrl returns a boolean if a field has been set.

### GetX509certSha1Fingerprint

`func (o *InlineResponse200329) GetX509certSha1Fingerprint() string`

GetX509certSha1Fingerprint returns the X509certSha1Fingerprint field if non-nil, zero value otherwise.

### GetX509certSha1FingerprintOk

`func (o *InlineResponse200329) GetX509certSha1FingerprintOk() (*string, bool)`

GetX509certSha1FingerprintOk returns a tuple with the X509certSha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509certSha1Fingerprint

`func (o *InlineResponse200329) SetX509certSha1Fingerprint(v string)`

SetX509certSha1Fingerprint sets X509certSha1Fingerprint field to given value.

### HasX509certSha1Fingerprint

`func (o *InlineResponse200329) HasX509certSha1Fingerprint() bool`

HasX509certSha1Fingerprint returns a boolean if a field has been set.

### GetSsoLoginUrl

`func (o *InlineResponse200329) GetSsoLoginUrl() string`

GetSsoLoginUrl returns the SsoLoginUrl field if non-nil, zero value otherwise.

### GetSsoLoginUrlOk

`func (o *InlineResponse200329) GetSsoLoginUrlOk() (*string, bool)`

GetSsoLoginUrlOk returns a tuple with the SsoLoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLoginUrl

`func (o *InlineResponse200329) SetSsoLoginUrl(v string)`

SetSsoLoginUrl sets SsoLoginUrl field to given value.

### HasSsoLoginUrl

`func (o *InlineResponse200329) HasSsoLoginUrl() bool`

HasSsoLoginUrl returns a boolean if a field has been set.

### GetSloLogoutUrl

`func (o *InlineResponse200329) GetSloLogoutUrl() string`

GetSloLogoutUrl returns the SloLogoutUrl field if non-nil, zero value otherwise.

### GetSloLogoutUrlOk

`func (o *InlineResponse200329) GetSloLogoutUrlOk() (*string, bool)`

GetSloLogoutUrlOk returns a tuple with the SloLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloLogoutUrl

`func (o *InlineResponse200329) SetSloLogoutUrl(v string)`

SetSloLogoutUrl sets SloLogoutUrl field to given value.

### HasSloLogoutUrl

`func (o *InlineResponse200329) HasSloLogoutUrl() bool`

HasSloLogoutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



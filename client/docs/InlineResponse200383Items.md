# InlineResponse200383Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Feature** | Pointer to [**InlineResponse20031ResultFeature**](InlineResponse20031ResultFeature.md) |  | [optional] 
**Authority** | Pointer to [**InlineResponse200383Authority**](InlineResponse200383Authority.md) |  | [optional] 
**Contents** | Pointer to **string** | PEM encoded device certificate with newlines as \\n | [optional] 
**Status** | Pointer to **string** | Certificate status | [optional] 
**ExpiresAt** | Pointer to **string** | Certificate expiry timestamp | [optional] 
**IssuedAt** | Pointer to **string** | Certificate issuance timestamp | [optional] 
**CertificateSerial** | Pointer to **string** | Certificate serial in hex with colons | [optional] 

## Methods

### NewInlineResponse200383Items

`func NewInlineResponse200383Items() *InlineResponse200383Items`

NewInlineResponse200383Items instantiates a new InlineResponse200383Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200383ItemsWithDefaults

`func NewInlineResponse200383ItemsWithDefaults() *InlineResponse200383Items`

NewInlineResponse200383ItemsWithDefaults instantiates a new InlineResponse200383Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200383Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200383Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200383Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200383Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetFeature

`func (o *InlineResponse200383Items) GetFeature() InlineResponse20031ResultFeature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *InlineResponse200383Items) GetFeatureOk() (*InlineResponse20031ResultFeature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *InlineResponse200383Items) SetFeature(v InlineResponse20031ResultFeature)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *InlineResponse200383Items) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetAuthority

`func (o *InlineResponse200383Items) GetAuthority() InlineResponse200383Authority`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *InlineResponse200383Items) GetAuthorityOk() (*InlineResponse200383Authority, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *InlineResponse200383Items) SetAuthority(v InlineResponse200383Authority)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *InlineResponse200383Items) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200383Items) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200383Items) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200383Items) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200383Items) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200383Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200383Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200383Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200383Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InlineResponse200383Items) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InlineResponse200383Items) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InlineResponse200383Items) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InlineResponse200383Items) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIssuedAt

`func (o *InlineResponse200383Items) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *InlineResponse200383Items) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *InlineResponse200383Items) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *InlineResponse200383Items) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetCertificateSerial

`func (o *InlineResponse200383Items) GetCertificateSerial() string`

GetCertificateSerial returns the CertificateSerial field if non-nil, zero value otherwise.

### GetCertificateSerialOk

`func (o *InlineResponse200383Items) GetCertificateSerialOk() (*string, bool)`

GetCertificateSerialOk returns a tuple with the CertificateSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSerial

`func (o *InlineResponse200383Items) SetCertificateSerial(v string)`

SetCertificateSerial sets CertificateSerial field to given value.

### HasCertificateSerial

`func (o *InlineResponse200383Items) HasCertificateSerial() bool`

HasCertificateSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



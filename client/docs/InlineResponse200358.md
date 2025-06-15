# InlineResponse200358

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthorityId** | Pointer to **string** | The ID of the Certificate Authority | [optional] 
**Status** | Pointer to **string** | The status of the Certificate Authority | [optional] 
**Contents** | Pointer to **string** | The PEM encoded contents of the Certificate Authority - with newlines as \&quot; \&quot;. Contents can be null, if the cert has not yet been generated. | [optional] 

## Methods

### NewInlineResponse200358

`func NewInlineResponse200358() *InlineResponse200358`

NewInlineResponse200358 instantiates a new InlineResponse200358 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200358WithDefaults

`func NewInlineResponse200358WithDefaults() *InlineResponse200358`

NewInlineResponse200358WithDefaults instantiates a new InlineResponse200358 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthorityId

`func (o *InlineResponse200358) GetCertificateAuthorityId() string`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *InlineResponse200358) GetCertificateAuthorityIdOk() (*string, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *InlineResponse200358) SetCertificateAuthorityId(v string)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *InlineResponse200358) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200358) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200358) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200358) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200358) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200358) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200358) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200358) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200358) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



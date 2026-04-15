# InlineResponse200418

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthorityId** | Pointer to **string** | The ID of the Certificate Authority | [optional] 
**Status** | Pointer to **string** | The status of the Certificate Authority | [optional] 
**Contents** | Pointer to **string** | The PEM encoded contents of the Certificate Authority - with newlines as \&quot; \&quot;. Contents can be null, if the cert has not yet been generated. | [optional] 

## Methods

### NewInlineResponse200418

`func NewInlineResponse200418() *InlineResponse200418`

NewInlineResponse200418 instantiates a new InlineResponse200418 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200418WithDefaults

`func NewInlineResponse200418WithDefaults() *InlineResponse200418`

NewInlineResponse200418WithDefaults instantiates a new InlineResponse200418 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthorityId

`func (o *InlineResponse200418) GetCertificateAuthorityId() string`

GetCertificateAuthorityId returns the CertificateAuthorityId field if non-nil, zero value otherwise.

### GetCertificateAuthorityIdOk

`func (o *InlineResponse200418) GetCertificateAuthorityIdOk() (*string, bool)`

GetCertificateAuthorityIdOk returns a tuple with the CertificateAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityId

`func (o *InlineResponse200418) SetCertificateAuthorityId(v string)`

SetCertificateAuthorityId sets CertificateAuthorityId field to given value.

### HasCertificateAuthorityId

`func (o *InlineResponse200418) HasCertificateAuthorityId() bool`

HasCertificateAuthorityId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200418) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200418) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200418) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200418) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContents

`func (o *InlineResponse200418) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineResponse200418) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineResponse200418) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *InlineResponse200418) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



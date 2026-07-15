# InlineResponse200131Encryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | When true, traffic will be encrypted to the syslog server | [optional] 
**Certificate** | Pointer to [**NetworksNetworkIdDevicesSyslogServersEncryptionCertificate**](NetworksNetworkIdDevicesSyslogServersEncryptionCertificate.md) |  | [optional] 

## Methods

### NewInlineResponse200131Encryption

`func NewInlineResponse200131Encryption() *InlineResponse200131Encryption`

NewInlineResponse200131Encryption instantiates a new InlineResponse200131Encryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200131EncryptionWithDefaults

`func NewInlineResponse200131EncryptionWithDefaults() *InlineResponse200131Encryption`

NewInlineResponse200131EncryptionWithDefaults instantiates a new InlineResponse200131Encryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200131Encryption) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200131Encryption) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200131Encryption) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200131Encryption) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCertificate

`func (o *InlineResponse200131Encryption) GetCertificate() NetworksNetworkIdDevicesSyslogServersEncryptionCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *InlineResponse200131Encryption) GetCertificateOk() (*NetworksNetworkIdDevicesSyslogServersEncryptionCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *InlineResponse200131Encryption) SetCertificate(v NetworksNetworkIdDevicesSyslogServersEncryptionCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *InlineResponse200131Encryption) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



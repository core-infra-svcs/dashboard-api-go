# OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphers** | Pointer to **[]string** | Encryption ciphers used by the SSID (applicable for WPA3 only) | [optional] 
**Akms** | Pointer to **[]string** | Authentication key management systems used by the SSID (applicable for WPA3 only) | [optional] 
**Passphrase** | Pointer to **string** | Passphrase for the SSID | [optional] 
**Dot11w** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11w**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11w.md) |  | [optional] 
**Dot11r** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11r**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11r.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption

`func NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption() *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption`

NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption instantiates a new OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionWithDefaults

`func NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionWithDefaults() *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption`

NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionWithDefaults instantiates a new OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetCiphers() []string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetCiphersOk() (*[]string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) SetCiphers(v []string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetAkms

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetAkms() []string`

GetAkms returns the Akms field if non-nil, zero value otherwise.

### GetAkmsOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetAkmsOk() (*[]string, bool)`

GetAkmsOk returns a tuple with the Akms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkms

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) SetAkms(v []string)`

SetAkms sets Akms field to given value.

### HasAkms

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) HasAkms() bool`

HasAkms returns a boolean if a field has been set.

### GetPassphrase

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetDot11w

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetDot11w() OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11w`

GetDot11w returns the Dot11w field if non-nil, zero value otherwise.

### GetDot11wOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetDot11wOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11w, bool)`

GetDot11wOk returns a tuple with the Dot11w field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11w

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) SetDot11w(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11w)`

SetDot11w sets Dot11w field to given value.

### HasDot11w

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) HasDot11w() bool`

HasDot11w returns a boolean if a field has been set.

### GetDot11r

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetDot11r() OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11r`

GetDot11r returns the Dot11r field if non-nil, zero value otherwise.

### GetDot11rOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) GetDot11rOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11r, bool)`

GetDot11rOk returns a tuple with the Dot11r field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11r

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) SetDot11r(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryptionDot11r)`

SetDot11r sets Dot11r field to given value.

### HasDot11r

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidSecurityEncryption) HasDot11r() bool`

HasDot11r returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworksNetworkIdDevicesSyslogServersEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | When true, traffic will be encrypted to the syslog server. Defaults to false | [optional] 
**Certificate** | Pointer to [**NetworksNetworkIdDevicesSyslogServersEncryptionCertificate**](NetworksNetworkIdDevicesSyslogServersEncryptionCertificate.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdDevicesSyslogServersEncryption

`func NewNetworksNetworkIdDevicesSyslogServersEncryption() *NetworksNetworkIdDevicesSyslogServersEncryption`

NewNetworksNetworkIdDevicesSyslogServersEncryption instantiates a new NetworksNetworkIdDevicesSyslogServersEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdDevicesSyslogServersEncryptionWithDefaults

`func NewNetworksNetworkIdDevicesSyslogServersEncryptionWithDefaults() *NetworksNetworkIdDevicesSyslogServersEncryption`

NewNetworksNetworkIdDevicesSyslogServersEncryptionWithDefaults instantiates a new NetworksNetworkIdDevicesSyslogServersEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCertificate

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) GetCertificate() NetworksNetworkIdDevicesSyslogServersEncryptionCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) GetCertificateOk() (*NetworksNetworkIdDevicesSyslogServersEncryptionCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) SetCertificate(v NetworksNetworkIdDevicesSyslogServersEncryptionCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *NetworksNetworkIdDevicesSyslogServersEncryption) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



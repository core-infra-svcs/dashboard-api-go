# NetworksNetworkIdWirelessSsidsLocalAuthFallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, MR devices will cache authentication credentials for EAP-TLS or for MAC based authentication. | [optional] 
**CacheTimeout** | Pointer to **int32** | The duration (in seconds) for which auths are cached. | [optional] 
**ServerCaCertificate** | Pointer to [**NetworksNetworkIdWirelessSsidsLocalAuthFallbackServerCaCertificate**](NetworksNetworkIdWirelessSsidsLocalAuthFallbackServerCaCertificate.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsLocalAuthFallback

`func NewNetworksNetworkIdWirelessSsidsLocalAuthFallback() *NetworksNetworkIdWirelessSsidsLocalAuthFallback`

NewNetworksNetworkIdWirelessSsidsLocalAuthFallback instantiates a new NetworksNetworkIdWirelessSsidsLocalAuthFallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsLocalAuthFallbackWithDefaults

`func NewNetworksNetworkIdWirelessSsidsLocalAuthFallbackWithDefaults() *NetworksNetworkIdWirelessSsidsLocalAuthFallback`

NewNetworksNetworkIdWirelessSsidsLocalAuthFallbackWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsLocalAuthFallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCacheTimeout

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) GetServerCaCertificate() NetworksNetworkIdWirelessSsidsLocalAuthFallbackServerCaCertificate`

GetServerCaCertificate returns the ServerCaCertificate field if non-nil, zero value otherwise.

### GetServerCaCertificateOk

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) GetServerCaCertificateOk() (*NetworksNetworkIdWirelessSsidsLocalAuthFallbackServerCaCertificate, bool)`

GetServerCaCertificateOk returns a tuple with the ServerCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) SetServerCaCertificate(v NetworksNetworkIdWirelessSsidsLocalAuthFallbackServerCaCertificate)`

SetServerCaCertificate sets ServerCaCertificate field to given value.

### HasServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsLocalAuthFallback) HasServerCaCertificate() bool`

HasServerCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheTimeout** | Pointer to **int32** | The duration (in seconds) for which auths are cached. The timeout is measured from the user&#39;s most recent non-cached authentication to the network. Between 3600 (1 hour) and 86400 (1 day) | [optional] 
**Enabled** | Pointer to **bool** | If true, MR devices will cache authentication credentials for EAP-TLS or for MAC based authentication. | [optional] 
**ServerCaCertificate** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackServerCaCertificate**](NetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackServerCaCertificate.md) |  | [optional] 

## Methods

### NewNetworksNetworkIdWirelessSsidsNumberLocalAuthFallback

`func NewNetworksNetworkIdWirelessSsidsNumberLocalAuthFallback() *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback`

NewNetworksNetworkIdWirelessSsidsNumberLocalAuthFallback instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackWithDefaults

`func NewNetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback`

NewNetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheTimeout

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) GetServerCaCertificate() NetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackServerCaCertificate`

GetServerCaCertificate returns the ServerCaCertificate field if non-nil, zero value otherwise.

### GetServerCaCertificateOk

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) GetServerCaCertificateOk() (*NetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackServerCaCertificate, bool)`

GetServerCaCertificateOk returns a tuple with the ServerCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) SetServerCaCertificate(v NetworksNetworkIdWirelessSsidsNumberLocalAuthFallbackServerCaCertificate)`

SetServerCaCertificate sets ServerCaCertificate field to given value.

### HasServerCaCertificate

`func (o *NetworksNetworkIdWirelessSsidsNumberLocalAuthFallback) HasServerCaCertificate() bool`

HasServerCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



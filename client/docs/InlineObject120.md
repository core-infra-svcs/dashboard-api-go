# InlineObject120

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the endpoints to be rebooted. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the endpoints to be rebooted. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the endpoints to be rebooted. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be rebooted. | [optional] 
**KextPaths** | Pointer to **[]string** | The KextPaths of the endpoints to be rebooted. Available for macOS 11+ | [optional] 
**NotifyUser** | Pointer to **bool** | Whether or not to notify the user before rebooting the endpoint. Available for macOS 11.3+ | [optional] 
**RebuildKernelCache** | Pointer to **bool** | Whether or not to rebuild the kernel cache when rebooting the endpoint. Available for macOS 11+ | [optional] 
**RequestRequiresNetworkTether** | Pointer to **bool** | Whether or not the request requires network tethering. Available for macOS and supervised iOS or tvOS | [optional] 

## Methods

### NewInlineObject120

`func NewInlineObject120() *InlineObject120`

NewInlineObject120 instantiates a new InlineObject120 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject120WithDefaults

`func NewInlineObject120WithDefaults() *InlineObject120`

NewInlineObject120WithDefaults instantiates a new InlineObject120 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *InlineObject120) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *InlineObject120) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *InlineObject120) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *InlineObject120) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *InlineObject120) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InlineObject120) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InlineObject120) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *InlineObject120) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject120) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject120) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject120) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject120) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *InlineObject120) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineObject120) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineObject120) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineObject120) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetKextPaths

`func (o *InlineObject120) GetKextPaths() []string`

GetKextPaths returns the KextPaths field if non-nil, zero value otherwise.

### GetKextPathsOk

`func (o *InlineObject120) GetKextPathsOk() (*[]string, bool)`

GetKextPathsOk returns a tuple with the KextPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKextPaths

`func (o *InlineObject120) SetKextPaths(v []string)`

SetKextPaths sets KextPaths field to given value.

### HasKextPaths

`func (o *InlineObject120) HasKextPaths() bool`

HasKextPaths returns a boolean if a field has been set.

### GetNotifyUser

`func (o *InlineObject120) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *InlineObject120) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *InlineObject120) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *InlineObject120) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetRebuildKernelCache

`func (o *InlineObject120) GetRebuildKernelCache() bool`

GetRebuildKernelCache returns the RebuildKernelCache field if non-nil, zero value otherwise.

### GetRebuildKernelCacheOk

`func (o *InlineObject120) GetRebuildKernelCacheOk() (*bool, bool)`

GetRebuildKernelCacheOk returns a tuple with the RebuildKernelCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildKernelCache

`func (o *InlineObject120) SetRebuildKernelCache(v bool)`

SetRebuildKernelCache sets RebuildKernelCache field to given value.

### HasRebuildKernelCache

`func (o *InlineObject120) HasRebuildKernelCache() bool`

HasRebuildKernelCache returns a boolean if a field has been set.

### GetRequestRequiresNetworkTether

`func (o *InlineObject120) GetRequestRequiresNetworkTether() bool`

GetRequestRequiresNetworkTether returns the RequestRequiresNetworkTether field if non-nil, zero value otherwise.

### GetRequestRequiresNetworkTetherOk

`func (o *InlineObject120) GetRequestRequiresNetworkTetherOk() (*bool, bool)`

GetRequestRequiresNetworkTetherOk returns a tuple with the RequestRequiresNetworkTether field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRequiresNetworkTether

`func (o *InlineObject120) SetRequestRequiresNetworkTether(v bool)`

SetRequestRequiresNetworkTether sets RequestRequiresNetworkTether field to given value.

### HasRequestRequiresNetworkTether

`func (o *InlineObject120) HasRequestRequiresNetworkTether() bool`

HasRequestRequiresNetworkTether returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



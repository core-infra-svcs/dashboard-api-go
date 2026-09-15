# InlineResponse200126

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The identifier of the client | [optional] 
**ClientIp** | Pointer to **string** | The IP address of the client | [optional] 
**ClientMac** | Pointer to **string** | The MAC address of the client | [optional] 
**UsageHistory** | Pointer to [**[]NetworksNetworkIdClientsUsageHistoriesUsageHistory**](NetworksNetworkIdClientsUsageHistoriesUsageHistory.md) | An array of usage history records for the client | [optional] 

## Methods

### NewInlineResponse200126

`func NewInlineResponse200126() *InlineResponse200126`

NewInlineResponse200126 instantiates a new InlineResponse200126 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200126WithDefaults

`func NewInlineResponse200126WithDefaults() *InlineResponse200126`

NewInlineResponse200126WithDefaults instantiates a new InlineResponse200126 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InlineResponse200126) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse200126) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse200126) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse200126) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIp

`func (o *InlineResponse200126) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *InlineResponse200126) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *InlineResponse200126) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *InlineResponse200126) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientMac

`func (o *InlineResponse200126) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *InlineResponse200126) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *InlineResponse200126) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *InlineResponse200126) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetUsageHistory

`func (o *InlineResponse200126) GetUsageHistory() []NetworksNetworkIdClientsUsageHistoriesUsageHistory`

GetUsageHistory returns the UsageHistory field if non-nil, zero value otherwise.

### GetUsageHistoryOk

`func (o *InlineResponse200126) GetUsageHistoryOk() (*[]NetworksNetworkIdClientsUsageHistoriesUsageHistory, bool)`

GetUsageHistoryOk returns a tuple with the UsageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHistory

`func (o *InlineResponse200126) SetUsageHistory(v []NetworksNetworkIdClientsUsageHistoriesUsageHistory)`

SetUsageHistory sets UsageHistory field to given value.

### HasUsageHistory

`func (o *InlineResponse200126) HasUsageHistory() bool`

HasUsageHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



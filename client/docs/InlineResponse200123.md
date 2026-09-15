# InlineResponse200123

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The identifier of the client | [optional] 
**ClientIp** | Pointer to **string** | The IP address of the client | [optional] 
**ClientMac** | Pointer to **string** | The MAC address of the client | [optional] 
**ApplicationUsage** | Pointer to [**[]NetworksNetworkIdClientsApplicationUsageApplicationUsage**](NetworksNetworkIdClientsApplicationUsageApplicationUsage.md) | An array of application usage records for the client | [optional] 

## Methods

### NewInlineResponse200123

`func NewInlineResponse200123() *InlineResponse200123`

NewInlineResponse200123 instantiates a new InlineResponse200123 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200123WithDefaults

`func NewInlineResponse200123WithDefaults() *InlineResponse200123`

NewInlineResponse200123WithDefaults instantiates a new InlineResponse200123 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InlineResponse200123) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse200123) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse200123) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse200123) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIp

`func (o *InlineResponse200123) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *InlineResponse200123) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *InlineResponse200123) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *InlineResponse200123) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientMac

`func (o *InlineResponse200123) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *InlineResponse200123) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *InlineResponse200123) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *InlineResponse200123) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetApplicationUsage

`func (o *InlineResponse200123) GetApplicationUsage() []NetworksNetworkIdClientsApplicationUsageApplicationUsage`

GetApplicationUsage returns the ApplicationUsage field if non-nil, zero value otherwise.

### GetApplicationUsageOk

`func (o *InlineResponse200123) GetApplicationUsageOk() (*[]NetworksNetworkIdClientsApplicationUsageApplicationUsage, bool)`

GetApplicationUsageOk returns a tuple with the ApplicationUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUsage

`func (o *InlineResponse200123) SetApplicationUsage(v []NetworksNetworkIdClientsApplicationUsageApplicationUsage)`

SetApplicationUsage sets ApplicationUsage field to given value.

### HasApplicationUsage

`func (o *InlineResponse200123) HasApplicationUsage() bool`

HasApplicationUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



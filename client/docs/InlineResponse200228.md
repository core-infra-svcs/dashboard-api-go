# InlineResponse200228

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**ConnectionStats** | Pointer to [**InlineResponse200228ConnectionStats**](InlineResponse200228ConnectionStats.md) |  | [optional] 

## Methods

### NewInlineResponse200228

`func NewInlineResponse200228() *InlineResponse200228`

NewInlineResponse200228 instantiates a new InlineResponse200228 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200228WithDefaults

`func NewInlineResponse200228WithDefaults() *InlineResponse200228`

NewInlineResponse200228WithDefaults instantiates a new InlineResponse200228 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200228) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200228) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200228) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200228) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetConnectionStats

`func (o *InlineResponse200228) GetConnectionStats() InlineResponse200228ConnectionStats`

GetConnectionStats returns the ConnectionStats field if non-nil, zero value otherwise.

### GetConnectionStatsOk

`func (o *InlineResponse200228) GetConnectionStatsOk() (*InlineResponse200228ConnectionStats, bool)`

GetConnectionStatsOk returns a tuple with the ConnectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStats

`func (o *InlineResponse200228) SetConnectionStats(v InlineResponse200228ConnectionStats)`

SetConnectionStats sets ConnectionStats field to given value.

### HasConnectionStats

`func (o *InlineResponse200228) HasConnectionStats() bool`

HasConnectionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



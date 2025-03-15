# InlineResponse200183

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**ConnectionStats** | Pointer to [**InlineResponse200183ConnectionStats**](InlineResponse200183ConnectionStats.md) |  | [optional] 

## Methods

### NewInlineResponse200183

`func NewInlineResponse200183() *InlineResponse200183`

NewInlineResponse200183 instantiates a new InlineResponse200183 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200183WithDefaults

`func NewInlineResponse200183WithDefaults() *InlineResponse200183`

NewInlineResponse200183WithDefaults instantiates a new InlineResponse200183 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *InlineResponse200183) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200183) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200183) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200183) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetConnectionStats

`func (o *InlineResponse200183) GetConnectionStats() InlineResponse200183ConnectionStats`

GetConnectionStats returns the ConnectionStats field if non-nil, zero value otherwise.

### GetConnectionStatsOk

`func (o *InlineResponse200183) GetConnectionStatsOk() (*InlineResponse200183ConnectionStats, bool)`

GetConnectionStatsOk returns a tuple with the ConnectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStats

`func (o *InlineResponse200183) SetConnectionStats(v InlineResponse200183ConnectionStats)`

SetConnectionStats sets ConnectionStats field to given value.

### HasConnectionStats

`func (o *InlineResponse200183) HasConnectionStats() bool`

HasConnectionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



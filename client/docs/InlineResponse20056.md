# InlineResponse20056

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device | [optional] 
**ConnectionStats** | Pointer to [**InlineResponse20056ConnectionStats**](InlineResponse20056ConnectionStats.md) |  | [optional] 

## Methods

### NewInlineResponse20056

`func NewInlineResponse20056() *InlineResponse20056`

NewInlineResponse20056 instantiates a new InlineResponse20056 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20056WithDefaults

`func NewInlineResponse20056WithDefaults() *InlineResponse20056`

NewInlineResponse20056WithDefaults instantiates a new InlineResponse20056 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20056) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20056) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20056) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20056) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetConnectionStats

`func (o *InlineResponse20056) GetConnectionStats() InlineResponse20056ConnectionStats`

GetConnectionStats returns the ConnectionStats field if non-nil, zero value otherwise.

### GetConnectionStatsOk

`func (o *InlineResponse20056) GetConnectionStatsOk() (*InlineResponse20056ConnectionStats, bool)`

GetConnectionStatsOk returns a tuple with the ConnectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStats

`func (o *InlineResponse20056) SetConnectionStats(v InlineResponse20056ConnectionStats)`

SetConnectionStats sets ConnectionStats field to given value.

### HasConnectionStats

`func (o *InlineResponse20056) HasConnectionStats() bool`

HasConnectionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



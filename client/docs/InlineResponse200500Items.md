# InlineResponse200500Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200500Network**](InlineResponse200500Network.md) |  | [optional] 
**Serial** | Pointer to **string** | Access point Serial number | [optional] 
**Counts** | Pointer to [**InlineResponse200500Counts**](InlineResponse200500Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200500Items

`func NewInlineResponse200500Items() *InlineResponse200500Items`

NewInlineResponse200500Items instantiates a new InlineResponse200500Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200500ItemsWithDefaults

`func NewInlineResponse200500ItemsWithDefaults() *InlineResponse200500Items`

NewInlineResponse200500ItemsWithDefaults instantiates a new InlineResponse200500Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200500Items) GetNetwork() InlineResponse200500Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200500Items) GetNetworkOk() (*InlineResponse200500Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200500Items) SetNetwork(v InlineResponse200500Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200500Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200500Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200500Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200500Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200500Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200500Items) GetCounts() InlineResponse200500Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200500Items) GetCountsOk() (*InlineResponse200500Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200500Items) SetCounts(v InlineResponse200500Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200500Items) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



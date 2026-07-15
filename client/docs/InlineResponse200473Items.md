# InlineResponse200473Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200473Network**](InlineResponse200473Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200473Readings**](InlineResponse200473Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200473Items

`func NewInlineResponse200473Items() *InlineResponse200473Items`

NewInlineResponse200473Items instantiates a new InlineResponse200473Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200473ItemsWithDefaults

`func NewInlineResponse200473ItemsWithDefaults() *InlineResponse200473Items`

NewInlineResponse200473ItemsWithDefaults instantiates a new InlineResponse200473Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200473Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200473Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200473Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200473Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200473Items) GetNetwork() InlineResponse200473Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200473Items) GetNetworkOk() (*InlineResponse200473Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200473Items) SetNetwork(v InlineResponse200473Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200473Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200473Items) GetReadings() []InlineResponse200473Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200473Items) GetReadingsOk() (*[]InlineResponse200473Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200473Items) SetReadings(v []InlineResponse200473Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200473Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



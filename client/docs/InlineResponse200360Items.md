# InlineResponse200360Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200360Network**](InlineResponse200360Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200360Readings**](InlineResponse200360Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200360Items

`func NewInlineResponse200360Items() *InlineResponse200360Items`

NewInlineResponse200360Items instantiates a new InlineResponse200360Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200360ItemsWithDefaults

`func NewInlineResponse200360ItemsWithDefaults() *InlineResponse200360Items`

NewInlineResponse200360ItemsWithDefaults instantiates a new InlineResponse200360Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200360Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200360Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200360Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200360Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200360Items) GetNetwork() InlineResponse200360Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200360Items) GetNetworkOk() (*InlineResponse200360Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200360Items) SetNetwork(v InlineResponse200360Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200360Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200360Items) GetReadings() []InlineResponse200360Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200360Items) GetReadingsOk() (*[]InlineResponse200360Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200360Items) SetReadings(v []InlineResponse200360Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200360Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



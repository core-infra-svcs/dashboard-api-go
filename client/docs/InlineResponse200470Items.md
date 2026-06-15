# InlineResponse200470Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200470Network**](InlineResponse200470Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200470Readings**](InlineResponse200470Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200470Items

`func NewInlineResponse200470Items() *InlineResponse200470Items`

NewInlineResponse200470Items instantiates a new InlineResponse200470Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200470ItemsWithDefaults

`func NewInlineResponse200470ItemsWithDefaults() *InlineResponse200470Items`

NewInlineResponse200470ItemsWithDefaults instantiates a new InlineResponse200470Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200470Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200470Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200470Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200470Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200470Items) GetNetwork() InlineResponse200470Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200470Items) GetNetworkOk() (*InlineResponse200470Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200470Items) SetNetwork(v InlineResponse200470Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200470Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200470Items) GetReadings() []InlineResponse200470Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200470Items) GetReadingsOk() (*[]InlineResponse200470Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200470Items) SetReadings(v []InlineResponse200470Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200470Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



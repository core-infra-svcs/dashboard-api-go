# InlineResponse200458Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Network** | Pointer to [**InlineResponse200458Network**](InlineResponse200458Network.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200458Readings**](InlineResponse200458Readings.md) | Overview history of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200458Items

`func NewInlineResponse200458Items() *InlineResponse200458Items`

NewInlineResponse200458Items instantiates a new InlineResponse200458Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200458ItemsWithDefaults

`func NewInlineResponse200458ItemsWithDefaults() *InlineResponse200458Items`

NewInlineResponse200458ItemsWithDefaults instantiates a new InlineResponse200458Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200458Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200458Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200458Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200458Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200458Items) GetNetwork() InlineResponse200458Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200458Items) GetNetworkOk() (*InlineResponse200458Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200458Items) SetNetwork(v InlineResponse200458Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200458Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200458Items) GetReadings() []InlineResponse200458Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200458Items) GetReadingsOk() (*[]InlineResponse200458Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200458Items) SetReadings(v []InlineResponse200458Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200458Items) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



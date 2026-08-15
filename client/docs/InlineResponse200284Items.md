# InlineResponse200284Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200284Network**](InlineResponse200284Network.md) |  | [optional] 
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse200284Interfaces**](InlineResponse200284Interfaces.md) | The interfaces associated with the given device | [optional] 

## Methods

### NewInlineResponse200284Items

`func NewInlineResponse200284Items() *InlineResponse200284Items`

NewInlineResponse200284Items instantiates a new InlineResponse200284Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200284ItemsWithDefaults

`func NewInlineResponse200284ItemsWithDefaults() *InlineResponse200284Items`

NewInlineResponse200284ItemsWithDefaults instantiates a new InlineResponse200284Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200284Items) GetNetwork() InlineResponse200284Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200284Items) GetNetworkOk() (*InlineResponse200284Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200284Items) SetNetwork(v InlineResponse200284Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200284Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200284Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200284Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200284Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200284Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse200284Items) GetInterfaces() []InlineResponse200284Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200284Items) GetInterfacesOk() (*[]InlineResponse200284Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200284Items) SetInterfaces(v []InlineResponse200284Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse200284Items) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



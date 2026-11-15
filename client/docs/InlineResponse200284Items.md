# InlineResponse200284Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the Catalyst appliance | [optional] 
**Ports** | Pointer to [**[]InlineResponse200284Ports**](InlineResponse200284Ports.md) | Ports associated with the given appliance | [optional] 
**Network** | Pointer to [**InlineResponse200284Network**](InlineResponse200284Network.md) |  | [optional] 

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

### GetPorts

`func (o *InlineResponse200284Items) GetPorts() []InlineResponse200284Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200284Items) GetPortsOk() (*[]InlineResponse200284Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200284Items) SetPorts(v []InlineResponse200284Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200284Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



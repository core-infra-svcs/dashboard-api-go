# InlineResponse200290Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200290Network**](InlineResponse200290Network.md) |  | [optional] 
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse200290Interfaces**](InlineResponse200290Interfaces.md) | The interfaces associated with the given device | [optional] 

## Methods

### NewInlineResponse200290Items

`func NewInlineResponse200290Items() *InlineResponse200290Items`

NewInlineResponse200290Items instantiates a new InlineResponse200290Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200290ItemsWithDefaults

`func NewInlineResponse200290ItemsWithDefaults() *InlineResponse200290Items`

NewInlineResponse200290ItemsWithDefaults instantiates a new InlineResponse200290Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200290Items) GetNetwork() InlineResponse200290Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200290Items) GetNetworkOk() (*InlineResponse200290Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200290Items) SetNetwork(v InlineResponse200290Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200290Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200290Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200290Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200290Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200290Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse200290Items) GetInterfaces() []InlineResponse200290Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200290Items) GetInterfacesOk() (*[]InlineResponse200290Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200290Items) SetInterfaces(v []InlineResponse200290Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse200290Items) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



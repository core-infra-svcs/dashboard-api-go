# InlineResponse200286Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200286Network**](InlineResponse200286Network.md) |  | [optional] 
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse200286Interfaces**](InlineResponse200286Interfaces.md) | The interfaces associated with the given device | [optional] 

## Methods

### NewInlineResponse200286Items

`func NewInlineResponse200286Items() *InlineResponse200286Items`

NewInlineResponse200286Items instantiates a new InlineResponse200286Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200286ItemsWithDefaults

`func NewInlineResponse200286ItemsWithDefaults() *InlineResponse200286Items`

NewInlineResponse200286ItemsWithDefaults instantiates a new InlineResponse200286Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200286Items) GetNetwork() InlineResponse200286Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200286Items) GetNetworkOk() (*InlineResponse200286Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200286Items) SetNetwork(v InlineResponse200286Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200286Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200286Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200286Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200286Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200286Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse200286Items) GetInterfaces() []InlineResponse200286Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse200286Items) GetInterfacesOk() (*[]InlineResponse200286Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse200286Items) SetInterfaces(v []InlineResponse200286Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse200286Items) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



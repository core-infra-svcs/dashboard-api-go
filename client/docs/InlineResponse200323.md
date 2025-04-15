# InlineResponse200323

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the switch. | [optional] 
**Serial** | Pointer to **string** | The serial number of the switch. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the switch. | [optional] 
**Network** | Pointer to [**InlineResponse200323Network**](InlineResponse200323Network.md) |  | [optional] 
**Model** | Pointer to **string** | The model of the switch. | [optional] 
**Ports** | Pointer to [**[]InlineResponse200323Ports**](InlineResponse200323Ports.md) | Ports belonging to the switch | [optional] 

## Methods

### NewInlineResponse200323

`func NewInlineResponse200323() *InlineResponse200323`

NewInlineResponse200323 instantiates a new InlineResponse200323 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200323WithDefaults

`func NewInlineResponse200323WithDefaults() *InlineResponse200323`

NewInlineResponse200323WithDefaults instantiates a new InlineResponse200323 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200323) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200323) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200323) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200323) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200323) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200323) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200323) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200323) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200323) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200323) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200323) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200323) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200323) GetNetwork() InlineResponse200323Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200323) GetNetworkOk() (*InlineResponse200323Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200323) SetNetwork(v InlineResponse200323Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200323) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200323) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200323) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200323) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200323) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200323) GetPorts() []InlineResponse200323Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200323) GetPortsOk() (*[]InlineResponse200323Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200323) SetPorts(v []InlineResponse200323Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200323) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



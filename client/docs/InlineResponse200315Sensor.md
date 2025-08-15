# InlineResponse200315Sensor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number of the sensor. | [optional] 
**Name** | **string** | The name of the sensor. | 
**Mac** | **string** | The MAC address of the sensor. | 

## Methods

### NewInlineResponse200315Sensor

`func NewInlineResponse200315Sensor(name string, mac string, ) *InlineResponse200315Sensor`

NewInlineResponse200315Sensor instantiates a new InlineResponse200315Sensor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200315SensorWithDefaults

`func NewInlineResponse200315SensorWithDefaults() *InlineResponse200315Sensor`

NewInlineResponse200315SensorWithDefaults instantiates a new InlineResponse200315Sensor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200315Sensor) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200315Sensor) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200315Sensor) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200315Sensor) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200315Sensor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200315Sensor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200315Sensor) SetName(v string)`

SetName sets Name field to given value.


### GetMac

`func (o *InlineResponse200315Sensor) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200315Sensor) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200315Sensor) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



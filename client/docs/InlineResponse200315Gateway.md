# InlineResponse200315Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number of the gateway. | [optional] 
**Name** | **string** | The name of the gateway. | 
**Mac** | **string** | The MAC address of the gateway. | 

## Methods

### NewInlineResponse200315Gateway

`func NewInlineResponse200315Gateway(name string, mac string, ) *InlineResponse200315Gateway`

NewInlineResponse200315Gateway instantiates a new InlineResponse200315Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200315GatewayWithDefaults

`func NewInlineResponse200315GatewayWithDefaults() *InlineResponse200315Gateway`

NewInlineResponse200315GatewayWithDefaults instantiates a new InlineResponse200315Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200315Gateway) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200315Gateway) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200315Gateway) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200315Gateway) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200315Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200315Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200315Gateway) SetName(v string)`

SetName sets Name field to given value.


### GetMac

`func (o *InlineResponse200315Gateway) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200315Gateway) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200315Gateway) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



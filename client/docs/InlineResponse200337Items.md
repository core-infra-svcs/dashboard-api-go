# InlineResponse200337Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of device | [optional] 
**BySlot** | Pointer to [**[]InlineResponse200337BySlot**](InlineResponse200337BySlot.md) | Information about the supported SIM slots for the device. Maximum 3 entries. | [optional] 

## Methods

### NewInlineResponse200337Items

`func NewInlineResponse200337Items() *InlineResponse200337Items`

NewInlineResponse200337Items instantiates a new InlineResponse200337Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200337ItemsWithDefaults

`func NewInlineResponse200337ItemsWithDefaults() *InlineResponse200337Items`

NewInlineResponse200337ItemsWithDefaults instantiates a new InlineResponse200337Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200337Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200337Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200337Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200337Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetBySlot

`func (o *InlineResponse200337Items) GetBySlot() []InlineResponse200337BySlot`

GetBySlot returns the BySlot field if non-nil, zero value otherwise.

### GetBySlotOk

`func (o *InlineResponse200337Items) GetBySlotOk() (*[]InlineResponse200337BySlot, bool)`

GetBySlotOk returns a tuple with the BySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySlot

`func (o *InlineResponse200337Items) SetBySlot(v []InlineResponse200337BySlot)`

SetBySlot sets BySlot field to given value.

### HasBySlot

`func (o *InlineResponse200337Items) HasBySlot() bool`

HasBySlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



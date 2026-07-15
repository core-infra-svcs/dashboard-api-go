# InlineResponse200340Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**BySlot** | Pointer to [**[]InlineResponse200340BySlot**](InlineResponse200340BySlot.md) | SIMs associated with the device. Maximum 3 SIM slots per device. | [optional] 

## Methods

### NewInlineResponse200340Items

`func NewInlineResponse200340Items() *InlineResponse200340Items`

NewInlineResponse200340Items instantiates a new InlineResponse200340Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200340ItemsWithDefaults

`func NewInlineResponse200340ItemsWithDefaults() *InlineResponse200340Items`

NewInlineResponse200340ItemsWithDefaults instantiates a new InlineResponse200340Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200340Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200340Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200340Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200340Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetBySlot

`func (o *InlineResponse200340Items) GetBySlot() []InlineResponse200340BySlot`

GetBySlot returns the BySlot field if non-nil, zero value otherwise.

### GetBySlotOk

`func (o *InlineResponse200340Items) GetBySlotOk() (*[]InlineResponse200340BySlot, bool)`

GetBySlotOk returns a tuple with the BySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySlot

`func (o *InlineResponse200340Items) SetBySlot(v []InlineResponse200340BySlot)`

SetBySlot sets BySlot field to given value.

### HasBySlot

`func (o *InlineResponse200340Items) HasBySlot() bool`

HasBySlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



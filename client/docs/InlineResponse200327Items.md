# InlineResponse200327Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**BySlot** | Pointer to [**[]InlineResponse200327BySlot**](InlineResponse200327BySlot.md) | SIMs associated with the device. Maximum 3 SIM slots per device. | [optional] 

## Methods

### NewInlineResponse200327Items

`func NewInlineResponse200327Items() *InlineResponse200327Items`

NewInlineResponse200327Items instantiates a new InlineResponse200327Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327ItemsWithDefaults

`func NewInlineResponse200327ItemsWithDefaults() *InlineResponse200327Items`

NewInlineResponse200327ItemsWithDefaults instantiates a new InlineResponse200327Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200327Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200327Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200327Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200327Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetBySlot

`func (o *InlineResponse200327Items) GetBySlot() []InlineResponse200327BySlot`

GetBySlot returns the BySlot field if non-nil, zero value otherwise.

### GetBySlotOk

`func (o *InlineResponse200327Items) GetBySlotOk() (*[]InlineResponse200327BySlot, bool)`

GetBySlotOk returns a tuple with the BySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySlot

`func (o *InlineResponse200327Items) SetBySlot(v []InlineResponse200327BySlot)`

SetBySlot sets BySlot field to given value.

### HasBySlot

`func (o *InlineResponse200327Items) HasBySlot() bool`

HasBySlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



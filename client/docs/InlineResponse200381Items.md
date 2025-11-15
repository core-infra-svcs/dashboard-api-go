# InlineResponse200381Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Changes** | Pointer to [**[]InlineResponse200381Changes**](InlineResponse200381Changes.md) | Connectivity information of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200381Items

`func NewInlineResponse200381Items() *InlineResponse200381Items`

NewInlineResponse200381Items instantiates a new InlineResponse200381Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200381ItemsWithDefaults

`func NewInlineResponse200381ItemsWithDefaults() *InlineResponse200381Items`

NewInlineResponse200381ItemsWithDefaults instantiates a new InlineResponse200381Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200381Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200381Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200381Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200381Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200381Items) GetChanges() []InlineResponse200381Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200381Items) GetChangesOk() (*[]InlineResponse200381Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200381Items) SetChanges(v []InlineResponse200381Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200381Items) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



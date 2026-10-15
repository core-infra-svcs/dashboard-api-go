# InlineResponse200524Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Changes** | Pointer to [**[]InlineResponse200524Changes**](InlineResponse200524Changes.md) | Connectivity information of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200524Items

`func NewInlineResponse200524Items() *InlineResponse200524Items`

NewInlineResponse200524Items instantiates a new InlineResponse200524Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200524ItemsWithDefaults

`func NewInlineResponse200524ItemsWithDefaults() *InlineResponse200524Items`

NewInlineResponse200524ItemsWithDefaults instantiates a new InlineResponse200524Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200524Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200524Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200524Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200524Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200524Items) GetChanges() []InlineResponse200524Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200524Items) GetChangesOk() (*[]InlineResponse200524Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200524Items) SetChanges(v []InlineResponse200524Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200524Items) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



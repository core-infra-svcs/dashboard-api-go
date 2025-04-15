# InlineResponse200351Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Changes** | Pointer to [**[]InlineResponse200351Changes**](InlineResponse200351Changes.md) | Connectivity information of a wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200351Items

`func NewInlineResponse200351Items() *InlineResponse200351Items`

NewInlineResponse200351Items instantiates a new InlineResponse200351Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200351ItemsWithDefaults

`func NewInlineResponse200351ItemsWithDefaults() *InlineResponse200351Items`

NewInlineResponse200351ItemsWithDefaults instantiates a new InlineResponse200351Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200351Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200351Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200351Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200351Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200351Items) GetChanges() []InlineResponse200351Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200351Items) GetChangesOk() (*[]InlineResponse200351Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200351Items) SetChanges(v []InlineResponse200351Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200351Items) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse2017Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Count** | Pointer to **int32** | Number of pings to send. [1..5], default 5 | [optional] 

## Methods

### NewInlineResponse2017Request

`func NewInlineResponse2017Request() *InlineResponse2017Request`

NewInlineResponse2017Request instantiates a new InlineResponse2017Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2017RequestWithDefaults

`func NewInlineResponse2017RequestWithDefaults() *InlineResponse2017Request`

NewInlineResponse2017RequestWithDefaults instantiates a new InlineResponse2017Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse2017Request) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse2017Request) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse2017Request) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse2017Request) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse2017Request) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse2017Request) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse2017Request) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse2017Request) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



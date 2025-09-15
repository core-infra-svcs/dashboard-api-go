# InlineResponse2016Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Target** | Pointer to **string** | IP address or FQDN to ping | [optional] 
**Count** | Pointer to **int32** | Number of pings to send. [1..5], default 5 | [optional] 

## Methods

### NewInlineResponse2016Request

`func NewInlineResponse2016Request() *InlineResponse2016Request`

NewInlineResponse2016Request instantiates a new InlineResponse2016Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2016RequestWithDefaults

`func NewInlineResponse2016RequestWithDefaults() *InlineResponse2016Request`

NewInlineResponse2016RequestWithDefaults instantiates a new InlineResponse2016Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse2016Request) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse2016Request) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse2016Request) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse2016Request) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTarget

`func (o *InlineResponse2016Request) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineResponse2016Request) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineResponse2016Request) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InlineResponse2016Request) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse2016Request) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse2016Request) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse2016Request) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InlineResponse2016Request) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



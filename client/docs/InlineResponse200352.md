# InlineResponse200352

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSerial** | Pointer to **string** | Serial number of the source switch (must be on a network not bound to a template) | [optional] 
**TargetSerials** | Pointer to **[]string** | Array of serial numbers of one or more target switches (must be on a network not bound to a template) | [optional] 

## Methods

### NewInlineResponse200352

`func NewInlineResponse200352() *InlineResponse200352`

NewInlineResponse200352 instantiates a new InlineResponse200352 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200352WithDefaults

`func NewInlineResponse200352WithDefaults() *InlineResponse200352`

NewInlineResponse200352WithDefaults instantiates a new InlineResponse200352 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSerial

`func (o *InlineResponse200352) GetSourceSerial() string`

GetSourceSerial returns the SourceSerial field if non-nil, zero value otherwise.

### GetSourceSerialOk

`func (o *InlineResponse200352) GetSourceSerialOk() (*string, bool)`

GetSourceSerialOk returns a tuple with the SourceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSerial

`func (o *InlineResponse200352) SetSourceSerial(v string)`

SetSourceSerial sets SourceSerial field to given value.

### HasSourceSerial

`func (o *InlineResponse200352) HasSourceSerial() bool`

HasSourceSerial returns a boolean if a field has been set.

### GetTargetSerials

`func (o *InlineResponse200352) GetTargetSerials() []string`

GetTargetSerials returns the TargetSerials field if non-nil, zero value otherwise.

### GetTargetSerialsOk

`func (o *InlineResponse200352) GetTargetSerialsOk() (*[]string, bool)`

GetTargetSerialsOk returns a tuple with the TargetSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSerials

`func (o *InlineResponse200352) SetTargetSerials(v []string)`

SetTargetSerials sets TargetSerials field to given value.

### HasTargetSerials

`func (o *InlineResponse200352) HasTargetSerials() bool`

HasTargetSerials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



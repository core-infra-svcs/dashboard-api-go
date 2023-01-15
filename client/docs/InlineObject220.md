# InlineObject220

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSerial** | **string** | Serial number of the source switch (must be on a network not bound to a template) | 
**TargetSerials** | **[]string** | Array of serial numbers of one or more target switches (must be on a network not bound to a template) | 

## Methods

### NewInlineObject220

`func NewInlineObject220(sourceSerial string, targetSerials []string, ) *InlineObject220`

NewInlineObject220 instantiates a new InlineObject220 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject220WithDefaults

`func NewInlineObject220WithDefaults() *InlineObject220`

NewInlineObject220WithDefaults instantiates a new InlineObject220 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSerial

`func (o *InlineObject220) GetSourceSerial() string`

GetSourceSerial returns the SourceSerial field if non-nil, zero value otherwise.

### GetSourceSerialOk

`func (o *InlineObject220) GetSourceSerialOk() (*string, bool)`

GetSourceSerialOk returns a tuple with the SourceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSerial

`func (o *InlineObject220) SetSourceSerial(v string)`

SetSourceSerial sets SourceSerial field to given value.


### GetTargetSerials

`func (o *InlineObject220) GetTargetSerials() []string`

GetTargetSerials returns the TargetSerials field if non-nil, zero value otherwise.

### GetTargetSerialsOk

`func (o *InlineObject220) GetTargetSerialsOk() (*[]string, bool)`

GetTargetSerialsOk returns a tuple with the TargetSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSerials

`func (o *InlineObject220) SetTargetSerials(v []string)`

SetTargetSerials sets TargetSerials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


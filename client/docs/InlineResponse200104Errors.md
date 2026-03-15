# InlineResponse200104Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | The serial of the device | 
**Errors** | **[]string** | The errors for the device | 

## Methods

### NewInlineResponse200104Errors

`func NewInlineResponse200104Errors(serial string, errors []string, ) *InlineResponse200104Errors`

NewInlineResponse200104Errors instantiates a new InlineResponse200104Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104ErrorsWithDefaults

`func NewInlineResponse200104ErrorsWithDefaults() *InlineResponse200104Errors`

NewInlineResponse200104ErrorsWithDefaults instantiates a new InlineResponse200104Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200104Errors) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200104Errors) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200104Errors) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetErrors

`func (o *InlineResponse200104Errors) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200104Errors) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200104Errors) SetErrors(v []string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



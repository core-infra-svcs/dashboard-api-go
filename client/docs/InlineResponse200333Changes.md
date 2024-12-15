# InlineResponse200333Changes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | The timestamp of current status of the interface | [optional] 
**Status** | Pointer to **string** | The status of the interface | [optional] 
**Warnings** | Pointer to **[]string** | All warnings present on the port | [optional] 
**Errors** | Pointer to **[]string** | All errors present on the port | [optional] 

## Methods

### NewInlineResponse200333Changes

`func NewInlineResponse200333Changes() *InlineResponse200333Changes`

NewInlineResponse200333Changes instantiates a new InlineResponse200333Changes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200333ChangesWithDefaults

`func NewInlineResponse200333ChangesWithDefaults() *InlineResponse200333Changes`

NewInlineResponse200333ChangesWithDefaults instantiates a new InlineResponse200333Changes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200333Changes) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200333Changes) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200333Changes) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200333Changes) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200333Changes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200333Changes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200333Changes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200333Changes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarnings

`func (o *InlineResponse200333Changes) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *InlineResponse200333Changes) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *InlineResponse200333Changes) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *InlineResponse200333Changes) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200333Changes) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200333Changes) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200333Changes) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200333Changes) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



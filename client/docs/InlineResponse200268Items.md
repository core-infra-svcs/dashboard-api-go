# InlineResponse200268Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipeline** | Pointer to [**InlineResponse200268Pipeline**](InlineResponse200268Pipeline.md) |  | [optional] 
**JobId** | Pointer to **string** | Job ID | [optional] 
**Operation** | Pointer to [**InlineResponse200268Operation**](InlineResponse200268Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Job status | [optional] 
**Errors** | Pointer to [**[]InlineResponse200268Errors**](InlineResponse200268Errors.md) | Job errors | [optional] 
**Details** | Pointer to [**[]InlineResponse200268Details**](InlineResponse200268Details.md) | Job details (type-specific) | [optional] 

## Methods

### NewInlineResponse200268Items

`func NewInlineResponse200268Items() *InlineResponse200268Items`

NewInlineResponse200268Items instantiates a new InlineResponse200268Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200268ItemsWithDefaults

`func NewInlineResponse200268ItemsWithDefaults() *InlineResponse200268Items`

NewInlineResponse200268ItemsWithDefaults instantiates a new InlineResponse200268Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipeline

`func (o *InlineResponse200268Items) GetPipeline() InlineResponse200268Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *InlineResponse200268Items) GetPipelineOk() (*InlineResponse200268Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *InlineResponse200268Items) SetPipeline(v InlineResponse200268Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *InlineResponse200268Items) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetJobId

`func (o *InlineResponse200268Items) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse200268Items) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse200268Items) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse200268Items) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200268Items) GetOperation() InlineResponse200268Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200268Items) GetOperationOk() (*InlineResponse200268Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200268Items) SetOperation(v InlineResponse200268Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200268Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200268Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200268Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200268Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200268Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200268Items) GetErrors() []InlineResponse200268Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200268Items) GetErrorsOk() (*[]InlineResponse200268Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200268Items) SetErrors(v []InlineResponse200268Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200268Items) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200268Items) GetDetails() []InlineResponse200268Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200268Items) GetDetailsOk() (*[]InlineResponse200268Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200268Items) SetDetails(v []InlineResponse200268Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200268Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



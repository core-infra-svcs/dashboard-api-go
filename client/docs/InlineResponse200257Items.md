# InlineResponse200257Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipeline** | Pointer to [**InlineResponse200257Pipeline**](InlineResponse200257Pipeline.md) |  | [optional] 
**JobId** | Pointer to **string** | Job ID | [optional] 
**Operation** | Pointer to [**InlineResponse200257Operation**](InlineResponse200257Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Job status | [optional] 
**Errors** | Pointer to [**[]InlineResponse200257Errors**](InlineResponse200257Errors.md) | Job errors | [optional] 
**Details** | Pointer to [**[]InlineResponse200257Details**](InlineResponse200257Details.md) | Job details (type-specific) | [optional] 

## Methods

### NewInlineResponse200257Items

`func NewInlineResponse200257Items() *InlineResponse200257Items`

NewInlineResponse200257Items instantiates a new InlineResponse200257Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200257ItemsWithDefaults

`func NewInlineResponse200257ItemsWithDefaults() *InlineResponse200257Items`

NewInlineResponse200257ItemsWithDefaults instantiates a new InlineResponse200257Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipeline

`func (o *InlineResponse200257Items) GetPipeline() InlineResponse200257Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *InlineResponse200257Items) GetPipelineOk() (*InlineResponse200257Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *InlineResponse200257Items) SetPipeline(v InlineResponse200257Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *InlineResponse200257Items) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetJobId

`func (o *InlineResponse200257Items) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse200257Items) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse200257Items) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse200257Items) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200257Items) GetOperation() InlineResponse200257Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200257Items) GetOperationOk() (*InlineResponse200257Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200257Items) SetOperation(v InlineResponse200257Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200257Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200257Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200257Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200257Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200257Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200257Items) GetErrors() []InlineResponse200257Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200257Items) GetErrorsOk() (*[]InlineResponse200257Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200257Items) SetErrors(v []InlineResponse200257Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200257Items) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200257Items) GetDetails() []InlineResponse200257Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200257Items) GetDetailsOk() (*[]InlineResponse200257Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200257Items) SetDetails(v []InlineResponse200257Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200257Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



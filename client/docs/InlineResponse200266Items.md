# InlineResponse200266Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipeline** | Pointer to [**InlineResponse200266Pipeline**](InlineResponse200266Pipeline.md) |  | [optional] 
**JobId** | Pointer to **string** | Job ID | [optional] 
**Operation** | Pointer to [**InlineResponse200266Operation**](InlineResponse200266Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Job status | [optional] 
**Errors** | Pointer to [**[]InlineResponse200266Errors**](InlineResponse200266Errors.md) | Job errors | [optional] 
**Details** | Pointer to [**[]InlineResponse200266Details**](InlineResponse200266Details.md) | Job details (type-specific) | [optional] 

## Methods

### NewInlineResponse200266Items

`func NewInlineResponse200266Items() *InlineResponse200266Items`

NewInlineResponse200266Items instantiates a new InlineResponse200266Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200266ItemsWithDefaults

`func NewInlineResponse200266ItemsWithDefaults() *InlineResponse200266Items`

NewInlineResponse200266ItemsWithDefaults instantiates a new InlineResponse200266Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipeline

`func (o *InlineResponse200266Items) GetPipeline() InlineResponse200266Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *InlineResponse200266Items) GetPipelineOk() (*InlineResponse200266Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *InlineResponse200266Items) SetPipeline(v InlineResponse200266Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *InlineResponse200266Items) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetJobId

`func (o *InlineResponse200266Items) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse200266Items) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse200266Items) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse200266Items) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200266Items) GetOperation() InlineResponse200266Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200266Items) GetOperationOk() (*InlineResponse200266Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200266Items) SetOperation(v InlineResponse200266Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200266Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200266Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200266Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200266Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200266Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200266Items) GetErrors() []InlineResponse200266Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200266Items) GetErrorsOk() (*[]InlineResponse200266Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200266Items) SetErrors(v []InlineResponse200266Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200266Items) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200266Items) GetDetails() []InlineResponse200266Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200266Items) GetDetailsOk() (*[]InlineResponse200266Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200266Items) SetDetails(v []InlineResponse200266Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200266Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200276Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipeline** | Pointer to [**InlineResponse200276Pipeline**](InlineResponse200276Pipeline.md) |  | [optional] 
**JobId** | Pointer to **string** | Job ID | [optional] 
**Operation** | Pointer to [**InlineResponse200276Operation**](InlineResponse200276Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Job status | [optional] 
**Errors** | Pointer to [**[]InlineResponse200276Errors**](InlineResponse200276Errors.md) | Job errors | [optional] 
**Details** | Pointer to [**[]InlineResponse200276Details**](InlineResponse200276Details.md) | Job details (type-specific) | [optional] 

## Methods

### NewInlineResponse200276Items

`func NewInlineResponse200276Items() *InlineResponse200276Items`

NewInlineResponse200276Items instantiates a new InlineResponse200276Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200276ItemsWithDefaults

`func NewInlineResponse200276ItemsWithDefaults() *InlineResponse200276Items`

NewInlineResponse200276ItemsWithDefaults instantiates a new InlineResponse200276Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipeline

`func (o *InlineResponse200276Items) GetPipeline() InlineResponse200276Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *InlineResponse200276Items) GetPipelineOk() (*InlineResponse200276Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *InlineResponse200276Items) SetPipeline(v InlineResponse200276Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *InlineResponse200276Items) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetJobId

`func (o *InlineResponse200276Items) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse200276Items) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse200276Items) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse200276Items) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200276Items) GetOperation() InlineResponse200276Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200276Items) GetOperationOk() (*InlineResponse200276Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200276Items) SetOperation(v InlineResponse200276Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200276Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200276Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200276Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200276Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200276Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200276Items) GetErrors() []InlineResponse200276Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200276Items) GetErrorsOk() (*[]InlineResponse200276Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200276Items) SetErrors(v []InlineResponse200276Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200276Items) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200276Items) GetDetails() []InlineResponse200276Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200276Items) GetDetailsOk() (*[]InlineResponse200276Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200276Items) SetDetails(v []InlineResponse200276Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200276Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



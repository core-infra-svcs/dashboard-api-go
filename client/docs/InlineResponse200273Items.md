# InlineResponse200273Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipeline** | Pointer to [**InlineResponse200273Pipeline**](InlineResponse200273Pipeline.md) |  | [optional] 
**JobId** | Pointer to **string** | Job ID | [optional] 
**Operation** | Pointer to [**InlineResponse200273Operation**](InlineResponse200273Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Job status | [optional] 
**Errors** | Pointer to [**[]InlineResponse200273Errors**](InlineResponse200273Errors.md) | Job errors | [optional] 
**Details** | Pointer to [**[]InlineResponse200273Details**](InlineResponse200273Details.md) | Job details (type-specific) | [optional] 

## Methods

### NewInlineResponse200273Items

`func NewInlineResponse200273Items() *InlineResponse200273Items`

NewInlineResponse200273Items instantiates a new InlineResponse200273Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200273ItemsWithDefaults

`func NewInlineResponse200273ItemsWithDefaults() *InlineResponse200273Items`

NewInlineResponse200273ItemsWithDefaults instantiates a new InlineResponse200273Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipeline

`func (o *InlineResponse200273Items) GetPipeline() InlineResponse200273Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *InlineResponse200273Items) GetPipelineOk() (*InlineResponse200273Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *InlineResponse200273Items) SetPipeline(v InlineResponse200273Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *InlineResponse200273Items) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetJobId

`func (o *InlineResponse200273Items) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *InlineResponse200273Items) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *InlineResponse200273Items) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *InlineResponse200273Items) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200273Items) GetOperation() InlineResponse200273Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200273Items) GetOperationOk() (*InlineResponse200273Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200273Items) SetOperation(v InlineResponse200273Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200273Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200273Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200273Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200273Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200273Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200273Items) GetErrors() []InlineResponse200273Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200273Items) GetErrorsOk() (*[]InlineResponse200273Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200273Items) SetErrors(v []InlineResponse200273Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200273Items) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200273Items) GetDetails() []InlineResponse200273Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200273Items) GetDetailsOk() (*[]InlineResponse200273Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200273Items) SetDetails(v []InlineResponse200273Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200273Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



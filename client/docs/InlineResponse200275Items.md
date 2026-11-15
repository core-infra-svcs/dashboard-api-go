# InlineResponse200275Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineId** | Pointer to **string** | Pipeline ID | [optional] 
**Operation** | Pointer to [**InlineResponse200275Operation**](InlineResponse200275Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Pipeline status. &#x60;pending&#x60; pipelines have not started, &#x60;active&#x60; pipelines have started but not finished, &#x60;success&#x60; pipelines completed successfully, and &#x60;error&#x60; pipelines failed. | [optional] 

## Methods

### NewInlineResponse200275Items

`func NewInlineResponse200275Items() *InlineResponse200275Items`

NewInlineResponse200275Items instantiates a new InlineResponse200275Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200275ItemsWithDefaults

`func NewInlineResponse200275ItemsWithDefaults() *InlineResponse200275Items`

NewInlineResponse200275ItemsWithDefaults instantiates a new InlineResponse200275Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineId

`func (o *InlineResponse200275Items) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *InlineResponse200275Items) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *InlineResponse200275Items) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *InlineResponse200275Items) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200275Items) GetOperation() InlineResponse200275Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200275Items) GetOperationOk() (*InlineResponse200275Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200275Items) SetOperation(v InlineResponse200275Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200275Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200275Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200275Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200275Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200275Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



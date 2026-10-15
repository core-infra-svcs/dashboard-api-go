# InlineResponse200272Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineId** | Pointer to **string** | Pipeline ID | [optional] 
**Operation** | Pointer to [**InlineResponse200272Operation**](InlineResponse200272Operation.md) |  | [optional] 
**Status** | Pointer to **string** | Pipeline status. &#x60;pending&#x60; pipelines have not started, &#x60;active&#x60; pipelines have started but not finished, &#x60;success&#x60; pipelines completed successfully, and &#x60;error&#x60; pipelines failed. | [optional] 

## Methods

### NewInlineResponse200272Items

`func NewInlineResponse200272Items() *InlineResponse200272Items`

NewInlineResponse200272Items instantiates a new InlineResponse200272Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200272ItemsWithDefaults

`func NewInlineResponse200272ItemsWithDefaults() *InlineResponse200272Items`

NewInlineResponse200272ItemsWithDefaults instantiates a new InlineResponse200272Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineId

`func (o *InlineResponse200272Items) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *InlineResponse200272Items) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *InlineResponse200272Items) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *InlineResponse200272Items) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse200272Items) GetOperation() InlineResponse200272Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse200272Items) GetOperationOk() (*InlineResponse200272Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse200272Items) SetOperation(v InlineResponse200272Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse200272Items) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200272Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200272Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200272Items) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200272Items) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



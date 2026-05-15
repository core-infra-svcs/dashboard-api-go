# InlineResponse200127

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse200127Products**](InlineResponse200127Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse200127Stages**](InlineResponse200127Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200126Reasons**](InlineResponse200126Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse200127

`func NewInlineResponse200127() *InlineResponse200127`

NewInlineResponse200127 instantiates a new InlineResponse200127 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200127WithDefaults

`func NewInlineResponse200127WithDefaults() *InlineResponse200127`

NewInlineResponse200127WithDefaults instantiates a new InlineResponse200127 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse200127) GetProducts() InlineResponse200127Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse200127) GetProductsOk() (*InlineResponse200127Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse200127) SetProducts(v InlineResponse200127Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse200127) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse200127) GetStages() []InlineResponse200127Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse200127) GetStagesOk() (*[]InlineResponse200127Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse200127) SetStages(v []InlineResponse200127Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse200127) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200127) GetReasons() []InlineResponse200126Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200127) GetReasonsOk() (*[]InlineResponse200126Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200127) SetReasons(v []InlineResponse200126Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200127) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



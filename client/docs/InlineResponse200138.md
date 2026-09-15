# InlineResponse200138

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse200138Products**](InlineResponse200138Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse200138Stages**](InlineResponse200138Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200137Reasons**](InlineResponse200137Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse200138

`func NewInlineResponse200138() *InlineResponse200138`

NewInlineResponse200138 instantiates a new InlineResponse200138 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200138WithDefaults

`func NewInlineResponse200138WithDefaults() *InlineResponse200138`

NewInlineResponse200138WithDefaults instantiates a new InlineResponse200138 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse200138) GetProducts() InlineResponse200138Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse200138) GetProductsOk() (*InlineResponse200138Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse200138) SetProducts(v InlineResponse200138Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse200138) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse200138) GetStages() []InlineResponse200138Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse200138) GetStagesOk() (*[]InlineResponse200138Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse200138) SetStages(v []InlineResponse200138Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse200138) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200138) GetReasons() []InlineResponse200137Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200138) GetReasonsOk() (*[]InlineResponse200137Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200138) SetReasons(v []InlineResponse200137Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200138) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



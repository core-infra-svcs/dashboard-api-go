# InlineResponse20037

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20037Products**](InlineResponse20037Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20037Stages**](InlineResponse20037Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20036Reasons**](InlineResponse20036Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20037

`func NewInlineResponse20037() *InlineResponse20037`

NewInlineResponse20037 instantiates a new InlineResponse20037 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20037WithDefaults

`func NewInlineResponse20037WithDefaults() *InlineResponse20037`

NewInlineResponse20037WithDefaults instantiates a new InlineResponse20037 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20037) GetProducts() InlineResponse20037Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20037) GetProductsOk() (*InlineResponse20037Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20037) SetProducts(v InlineResponse20037Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20037) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20037) GetStages() []InlineResponse20037Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20037) GetStagesOk() (*[]InlineResponse20037Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20037) SetStages(v []InlineResponse20037Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20037) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20037) GetReasons() []InlineResponse20036Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20037) GetReasonsOk() (*[]InlineResponse20036Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20037) SetReasons(v []InlineResponse20036Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20037) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



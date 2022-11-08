# InlineResponse20019

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20019Products**](InlineResponse20019Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20019Stages**](InlineResponse20019Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20018Reasons**](InlineResponse20018Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20019

`func NewInlineResponse20019() *InlineResponse20019`

NewInlineResponse20019 instantiates a new InlineResponse20019 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20019WithDefaults

`func NewInlineResponse20019WithDefaults() *InlineResponse20019`

NewInlineResponse20019WithDefaults instantiates a new InlineResponse20019 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20019) GetProducts() InlineResponse20019Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20019) GetProductsOk() (*InlineResponse20019Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20019) SetProducts(v InlineResponse20019Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20019) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20019) GetStages() []InlineResponse20019Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20019) GetStagesOk() (*[]InlineResponse20019Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20019) SetStages(v []InlineResponse20019Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20019) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20019) GetReasons() []InlineResponse20018Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20019) GetReasonsOk() (*[]InlineResponse20018Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20019) SetReasons(v []InlineResponse20018Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20019) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



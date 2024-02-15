# InlineResponse20044

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20044Products**](InlineResponse20044Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20044Stages**](InlineResponse20044Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20043Reasons**](InlineResponse20043Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20044

`func NewInlineResponse20044() *InlineResponse20044`

NewInlineResponse20044 instantiates a new InlineResponse20044 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20044WithDefaults

`func NewInlineResponse20044WithDefaults() *InlineResponse20044`

NewInlineResponse20044WithDefaults instantiates a new InlineResponse20044 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20044) GetProducts() InlineResponse20044Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20044) GetProductsOk() (*InlineResponse20044Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20044) SetProducts(v InlineResponse20044Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20044) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20044) GetStages() []InlineResponse20044Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20044) GetStagesOk() (*[]InlineResponse20044Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20044) SetStages(v []InlineResponse20044Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20044) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20044) GetReasons() []InlineResponse20043Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20044) GetReasonsOk() (*[]InlineResponse20043Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20044) SetReasons(v []InlineResponse20043Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20044) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



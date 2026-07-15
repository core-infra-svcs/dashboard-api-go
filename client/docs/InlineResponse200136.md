# InlineResponse200136

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse200136Products**](InlineResponse200136Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse200136Stages**](InlineResponse200136Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200135Reasons**](InlineResponse200135Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse200136

`func NewInlineResponse200136() *InlineResponse200136`

NewInlineResponse200136 instantiates a new InlineResponse200136 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200136WithDefaults

`func NewInlineResponse200136WithDefaults() *InlineResponse200136`

NewInlineResponse200136WithDefaults instantiates a new InlineResponse200136 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse200136) GetProducts() InlineResponse200136Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse200136) GetProductsOk() (*InlineResponse200136Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse200136) SetProducts(v InlineResponse200136Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse200136) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse200136) GetStages() []InlineResponse200136Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse200136) GetStagesOk() (*[]InlineResponse200136Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse200136) SetStages(v []InlineResponse200136Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse200136) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200136) GetReasons() []InlineResponse200135Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200136) GetReasonsOk() (*[]InlineResponse200135Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200136) SetReasons(v []InlineResponse200135Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200136) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



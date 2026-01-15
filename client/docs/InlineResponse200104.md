# InlineResponse200104

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse200104Products**](InlineResponse200104Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse200104Stages**](InlineResponse200104Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200103Reasons**](InlineResponse200103Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse200104

`func NewInlineResponse200104() *InlineResponse200104`

NewInlineResponse200104 instantiates a new InlineResponse200104 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104WithDefaults

`func NewInlineResponse200104WithDefaults() *InlineResponse200104`

NewInlineResponse200104WithDefaults instantiates a new InlineResponse200104 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse200104) GetProducts() InlineResponse200104Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse200104) GetProductsOk() (*InlineResponse200104Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse200104) SetProducts(v InlineResponse200104Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse200104) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse200104) GetStages() []InlineResponse200104Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse200104) GetStagesOk() (*[]InlineResponse200104Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse200104) SetStages(v []InlineResponse200104Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse200104) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200104) GetReasons() []InlineResponse200103Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200104) GetReasonsOk() (*[]InlineResponse200103Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200104) SetReasons(v []InlineResponse200103Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200104) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200141

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse200141Products**](InlineResponse200141Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse200141Stages**](InlineResponse200141Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200140Reasons**](InlineResponse200140Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse200141

`func NewInlineResponse200141() *InlineResponse200141`

NewInlineResponse200141 instantiates a new InlineResponse200141 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200141WithDefaults

`func NewInlineResponse200141WithDefaults() *InlineResponse200141`

NewInlineResponse200141WithDefaults instantiates a new InlineResponse200141 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse200141) GetProducts() InlineResponse200141Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse200141) GetProductsOk() (*InlineResponse200141Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse200141) SetProducts(v InlineResponse200141Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse200141) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse200141) GetStages() []InlineResponse200141Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse200141) GetStagesOk() (*[]InlineResponse200141Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse200141) SetStages(v []InlineResponse200141Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse200141) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200141) GetReasons() []InlineResponse200140Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200141) GetReasonsOk() (*[]InlineResponse200140Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200141) SetReasons(v []InlineResponse200140Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200141) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



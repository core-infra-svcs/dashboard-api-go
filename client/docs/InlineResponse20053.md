# InlineResponse20053

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20053Products**](InlineResponse20053Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20053Stages**](InlineResponse20053Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20052Reasons**](InlineResponse20052Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20053

`func NewInlineResponse20053() *InlineResponse20053`

NewInlineResponse20053 instantiates a new InlineResponse20053 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20053WithDefaults

`func NewInlineResponse20053WithDefaults() *InlineResponse20053`

NewInlineResponse20053WithDefaults instantiates a new InlineResponse20053 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20053) GetProducts() InlineResponse20053Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20053) GetProductsOk() (*InlineResponse20053Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20053) SetProducts(v InlineResponse20053Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20053) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20053) GetStages() []InlineResponse20053Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20053) GetStagesOk() (*[]InlineResponse20053Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20053) SetStages(v []InlineResponse20053Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20053) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20053) GetReasons() []InlineResponse20052Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20053) GetReasonsOk() (*[]InlineResponse20052Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20053) SetReasons(v []InlineResponse20052Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20053) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse20038

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20038Products**](InlineResponse20038Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20038Stages**](InlineResponse20038Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20037Reasons**](InlineResponse20037Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20038

`func NewInlineResponse20038() *InlineResponse20038`

NewInlineResponse20038 instantiates a new InlineResponse20038 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20038WithDefaults

`func NewInlineResponse20038WithDefaults() *InlineResponse20038`

NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20038) GetProducts() InlineResponse20038Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20038) GetProductsOk() (*InlineResponse20038Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20038) SetProducts(v InlineResponse20038Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20038) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20038) GetStages() []InlineResponse20038Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20038) GetStagesOk() (*[]InlineResponse20038Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20038) SetStages(v []InlineResponse20038Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20038) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20038) GetReasons() []InlineResponse20037Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20038) GetReasonsOk() (*[]InlineResponse20037Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20038) SetReasons(v []InlineResponse20037Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20038) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



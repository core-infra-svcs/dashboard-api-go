# InlineResponse20082

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | Pointer to [**[]InlineResponse20082Mappings**](InlineResponse20082Mappings.md) | An array of DSCP to CoS mappings. An empty array will reset the mappings to default. | [optional] 

## Methods

### NewInlineResponse20082

`func NewInlineResponse20082() *InlineResponse20082`

NewInlineResponse20082 instantiates a new InlineResponse20082 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20082WithDefaults

`func NewInlineResponse20082WithDefaults() *InlineResponse20082`

NewInlineResponse20082WithDefaults instantiates a new InlineResponse20082 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *InlineResponse20082) GetMappings() []InlineResponse20082Mappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *InlineResponse20082) GetMappingsOk() (*[]InlineResponse20082Mappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *InlineResponse20082) SetMappings(v []InlineResponse20082Mappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *InlineResponse20082) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



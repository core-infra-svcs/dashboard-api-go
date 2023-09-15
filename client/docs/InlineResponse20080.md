# InlineResponse20080

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | Pointer to [**[]InlineResponse20080Mappings**](InlineResponse20080Mappings.md) | An array of DSCP to CoS mappings. An empty array will reset the mappings to default. | [optional] 

## Methods

### NewInlineResponse20080

`func NewInlineResponse20080() *InlineResponse20080`

NewInlineResponse20080 instantiates a new InlineResponse20080 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20080WithDefaults

`func NewInlineResponse20080WithDefaults() *InlineResponse20080`

NewInlineResponse20080WithDefaults instantiates a new InlineResponse20080 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *InlineResponse20080) GetMappings() []InlineResponse20080Mappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *InlineResponse20080) GetMappingsOk() (*[]InlineResponse20080Mappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *InlineResponse20080) SetMappings(v []InlineResponse20080Mappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *InlineResponse20080) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200160

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | Pointer to [**[]InlineResponse200160Mappings**](InlineResponse200160Mappings.md) | An array of DSCP to CoS mappings. An empty array will reset the mappings to default. | [optional] 

## Methods

### NewInlineResponse200160

`func NewInlineResponse200160() *InlineResponse200160`

NewInlineResponse200160 instantiates a new InlineResponse200160 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200160WithDefaults

`func NewInlineResponse200160WithDefaults() *InlineResponse200160`

NewInlineResponse200160WithDefaults instantiates a new InlineResponse200160 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *InlineResponse200160) GetMappings() []InlineResponse200160Mappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *InlineResponse200160) GetMappingsOk() (*[]InlineResponse200160Mappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *InlineResponse200160) SetMappings(v []InlineResponse200160Mappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *InlineResponse200160) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



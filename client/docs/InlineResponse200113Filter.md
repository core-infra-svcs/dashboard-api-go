# InlineResponse200113Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Out** | Pointer to **[]string** | Filters routes advertised to an eBGP neighbor. Each entry must be an IPv4 CIDR string (e.g., &#39;10.0.0.0/8&#39; or &#39;10.0.0.5/32&#39;). | [optional] 

## Methods

### NewInlineResponse200113Filter

`func NewInlineResponse200113Filter() *InlineResponse200113Filter`

NewInlineResponse200113Filter instantiates a new InlineResponse200113Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200113FilterWithDefaults

`func NewInlineResponse200113FilterWithDefaults() *InlineResponse200113Filter`

NewInlineResponse200113FilterWithDefaults instantiates a new InlineResponse200113Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOut

`func (o *InlineResponse200113Filter) GetOut() []string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *InlineResponse200113Filter) GetOutOk() (*[]string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *InlineResponse200113Filter) SetOut(v []string)`

SetOut sets Out field to given value.

### HasOut

`func (o *InlineResponse200113Filter) HasOut() bool`

HasOut returns a boolean if a field has been set.

### SetOutNil

`func (o *InlineResponse200113Filter) SetOutNil(b bool)`

 SetOutNil sets the value for Out to be an explicit nil

### UnsetOut
`func (o *InlineResponse200113Filter) UnsetOut()`

UnsetOut ensures that no value is present for Out, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



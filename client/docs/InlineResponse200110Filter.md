# InlineResponse200110Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Out** | Pointer to **[]string** | Filters routes advertised to an eBGP neighbor. Each entry must be an IPv4 CIDR string (e.g., &#39;10.0.0.0/8&#39; or &#39;10.0.0.5/32&#39;). | [optional] 

## Methods

### NewInlineResponse200110Filter

`func NewInlineResponse200110Filter() *InlineResponse200110Filter`

NewInlineResponse200110Filter instantiates a new InlineResponse200110Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110FilterWithDefaults

`func NewInlineResponse200110FilterWithDefaults() *InlineResponse200110Filter`

NewInlineResponse200110FilterWithDefaults instantiates a new InlineResponse200110Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOut

`func (o *InlineResponse200110Filter) GetOut() []string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *InlineResponse200110Filter) GetOutOk() (*[]string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *InlineResponse200110Filter) SetOut(v []string)`

SetOut sets Out field to given value.

### HasOut

`func (o *InlineResponse200110Filter) HasOut() bool`

HasOut returns a boolean if a field has been set.

### SetOutNil

`func (o *InlineResponse200110Filter) SetOutNil(b bool)`

 SetOutNil sets the value for Out to be an explicit nil

### UnsetOut
`func (o *InlineResponse200110Filter) UnsetOut()`

UnsetOut ensures that no value is present for Out, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



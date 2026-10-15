# InlineResponse200311EbgpNeighborFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Out** | Pointer to **[]string** | Filters routes advertised to an eBGP neighbor. Each entry must be an IPv4 CIDR string. Supported only for MX 26.2 and above | [optional] 

## Methods

### NewInlineResponse200311EbgpNeighborFilter

`func NewInlineResponse200311EbgpNeighborFilter() *InlineResponse200311EbgpNeighborFilter`

NewInlineResponse200311EbgpNeighborFilter instantiates a new InlineResponse200311EbgpNeighborFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311EbgpNeighborFilterWithDefaults

`func NewInlineResponse200311EbgpNeighborFilterWithDefaults() *InlineResponse200311EbgpNeighborFilter`

NewInlineResponse200311EbgpNeighborFilterWithDefaults instantiates a new InlineResponse200311EbgpNeighborFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOut

`func (o *InlineResponse200311EbgpNeighborFilter) GetOut() []string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *InlineResponse200311EbgpNeighborFilter) GetOutOk() (*[]string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *InlineResponse200311EbgpNeighborFilter) SetOut(v []string)`

SetOut sets Out field to given value.

### HasOut

`func (o *InlineResponse200311EbgpNeighborFilter) HasOut() bool`

HasOut returns a boolean if a field has been set.

### SetOutNil

`func (o *InlineResponse200311EbgpNeighborFilter) SetOutNil(b bool)`

 SetOutNil sets the value for Out to be an explicit nil

### UnsetOut
`func (o *InlineResponse200311EbgpNeighborFilter) UnsetOut()`

UnsetOut ensures that no value is present for Out, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



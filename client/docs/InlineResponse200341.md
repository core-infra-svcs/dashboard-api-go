# InlineResponse200341

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp of the start of the interval. | [optional] 
**Draw** | Pointer to **float32** | The PoE power draw in watts for all switch ports in the organization for the given interval. | [optional] 

## Methods

### NewInlineResponse200341

`func NewInlineResponse200341() *InlineResponse200341`

NewInlineResponse200341 instantiates a new InlineResponse200341 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200341WithDefaults

`func NewInlineResponse200341WithDefaults() *InlineResponse200341`

NewInlineResponse200341WithDefaults instantiates a new InlineResponse200341 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200341) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200341) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200341) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200341) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDraw

`func (o *InlineResponse200341) GetDraw() float32`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *InlineResponse200341) GetDrawOk() (*float32, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraw

`func (o *InlineResponse200341) SetDraw(v float32)`

SetDraw sets Draw field to given value.

### HasDraw

`func (o *InlineResponse200341) HasDraw() bool`

HasDraw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



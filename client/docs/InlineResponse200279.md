# InlineResponse200279

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp of the start of the interval. | [optional] 
**Draw** | Pointer to **float32** | The PoE power draw in watts for all switch ports in the organization for the given interval. | [optional] 

## Methods

### NewInlineResponse200279

`func NewInlineResponse200279() *InlineResponse200279`

NewInlineResponse200279 instantiates a new InlineResponse200279 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200279WithDefaults

`func NewInlineResponse200279WithDefaults() *InlineResponse200279`

NewInlineResponse200279WithDefaults instantiates a new InlineResponse200279 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200279) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200279) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200279) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200279) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDraw

`func (o *InlineResponse200279) GetDraw() float32`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *InlineResponse200279) GetDrawOk() (*float32, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraw

`func (o *InlineResponse200279) SetDraw(v float32)`

SetDraw sets Draw field to given value.

### HasDraw

`func (o *InlineResponse200279) HasDraw() bool`

HasDraw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



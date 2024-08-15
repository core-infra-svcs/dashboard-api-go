# InlineResponse200219Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200219Totals**](InlineResponse200219Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200219ByAlertType**](InlineResponse200219ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200219Items

`func NewInlineResponse200219Items(segmentStart time.Time, totals InlineResponse200219Totals, byAlertType []InlineResponse200219ByAlertType, ) *InlineResponse200219Items`

NewInlineResponse200219Items instantiates a new InlineResponse200219Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200219ItemsWithDefaults

`func NewInlineResponse200219ItemsWithDefaults() *InlineResponse200219Items`

NewInlineResponse200219ItemsWithDefaults instantiates a new InlineResponse200219Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200219Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200219Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200219Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200219Items) GetTotals() InlineResponse200219Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200219Items) GetTotalsOk() (*InlineResponse200219Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200219Items) SetTotals(v InlineResponse200219Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200219Items) GetByAlertType() []InlineResponse200219ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200219Items) GetByAlertTypeOk() (*[]InlineResponse200219ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200219Items) SetByAlertType(v []InlineResponse200219ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



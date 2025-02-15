# InlineResponse200236Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200236Totals**](InlineResponse200236Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200236ByAlertType**](InlineResponse200236ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200236Items

`func NewInlineResponse200236Items(segmentStart time.Time, totals InlineResponse200236Totals, byAlertType []InlineResponse200236ByAlertType, ) *InlineResponse200236Items`

NewInlineResponse200236Items instantiates a new InlineResponse200236Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200236ItemsWithDefaults

`func NewInlineResponse200236ItemsWithDefaults() *InlineResponse200236Items`

NewInlineResponse200236ItemsWithDefaults instantiates a new InlineResponse200236Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200236Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200236Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200236Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200236Items) GetTotals() InlineResponse200236Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200236Items) GetTotalsOk() (*InlineResponse200236Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200236Items) SetTotals(v InlineResponse200236Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200236Items) GetByAlertType() []InlineResponse200236ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200236Items) GetByAlertTypeOk() (*[]InlineResponse200236ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200236Items) SetByAlertType(v []InlineResponse200236ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200298Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**InlineResponse200298Totals**](InlineResponse200298Totals.md) |  | 
**ByAlertType** | [**[]InlineResponse200298ByAlertType**](InlineResponse200298ByAlertType.md) | Totals by Type | 

## Methods

### NewInlineResponse200298Items

`func NewInlineResponse200298Items(segmentStart time.Time, totals InlineResponse200298Totals, byAlertType []InlineResponse200298ByAlertType, ) *InlineResponse200298Items`

NewInlineResponse200298Items instantiates a new InlineResponse200298Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200298ItemsWithDefaults

`func NewInlineResponse200298ItemsWithDefaults() *InlineResponse200298Items`

NewInlineResponse200298ItemsWithDefaults instantiates a new InlineResponse200298Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *InlineResponse200298Items) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *InlineResponse200298Items) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *InlineResponse200298Items) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *InlineResponse200298Items) GetTotals() InlineResponse200298Totals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InlineResponse200298Items) GetTotalsOk() (*InlineResponse200298Totals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InlineResponse200298Items) SetTotals(v InlineResponse200298Totals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *InlineResponse200298Items) GetByAlertType() []InlineResponse200298ByAlertType`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *InlineResponse200298Items) GetByAlertTypeOk() (*[]InlineResponse200298ByAlertType, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *InlineResponse200298Items) SetByAlertType(v []InlineResponse200298ByAlertType)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



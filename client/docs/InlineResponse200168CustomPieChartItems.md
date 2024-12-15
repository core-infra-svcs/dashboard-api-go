# InlineResponse200168CustomPieChartItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom pie chart item. | 
**Type** | **string** |     The signature type for the custom pie chart item. Can be one of &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39;.  | 
**Value** | **string** |     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details).  | 

## Methods

### NewInlineResponse200168CustomPieChartItems

`func NewInlineResponse200168CustomPieChartItems(name string, type_ string, value string, ) *InlineResponse200168CustomPieChartItems`

NewInlineResponse200168CustomPieChartItems instantiates a new InlineResponse200168CustomPieChartItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200168CustomPieChartItemsWithDefaults

`func NewInlineResponse200168CustomPieChartItemsWithDefaults() *InlineResponse200168CustomPieChartItems`

NewInlineResponse200168CustomPieChartItemsWithDefaults instantiates a new InlineResponse200168CustomPieChartItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200168CustomPieChartItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200168CustomPieChartItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200168CustomPieChartItems) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InlineResponse200168CustomPieChartItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200168CustomPieChartItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200168CustomPieChartItems) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *InlineResponse200168CustomPieChartItems) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineResponse200168CustomPieChartItems) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineResponse200168CustomPieChartItems) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



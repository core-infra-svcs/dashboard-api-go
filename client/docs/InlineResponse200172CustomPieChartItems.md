# InlineResponse200172CustomPieChartItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom pie chart item. | 
**Type** | **string** |     The signature type for the custom pie chart item. Can be one of &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39;.  | 
**Value** | **string** |     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details).  | 

## Methods

### NewInlineResponse200172CustomPieChartItems

`func NewInlineResponse200172CustomPieChartItems(name string, type_ string, value string, ) *InlineResponse200172CustomPieChartItems`

NewInlineResponse200172CustomPieChartItems instantiates a new InlineResponse200172CustomPieChartItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200172CustomPieChartItemsWithDefaults

`func NewInlineResponse200172CustomPieChartItemsWithDefaults() *InlineResponse200172CustomPieChartItems`

NewInlineResponse200172CustomPieChartItemsWithDefaults instantiates a new InlineResponse200172CustomPieChartItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200172CustomPieChartItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200172CustomPieChartItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200172CustomPieChartItems) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *InlineResponse200172CustomPieChartItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200172CustomPieChartItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200172CustomPieChartItems) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *InlineResponse200172CustomPieChartItems) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineResponse200172CustomPieChartItems) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineResponse200172CustomPieChartItems) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



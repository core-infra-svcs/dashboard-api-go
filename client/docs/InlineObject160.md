# InlineObject160

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |     The traffic analysis mode for the network. Can be one of &#39;disabled&#39; (do not collect traffic types),     &#39;basic&#39; (collect generic traffic categories), or &#39;detailed&#39; (collect destination hostnames).  | [optional] 
**CustomPieChartItems** | Pointer to [**[]InlineResponse200162CustomPieChartItems**](InlineResponse200162CustomPieChartItems.md) | The list of items that make up the custom pie chart for traffic reporting. | [optional] 

## Methods

### NewInlineObject160

`func NewInlineObject160() *InlineObject160`

NewInlineObject160 instantiates a new InlineObject160 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject160WithDefaults

`func NewInlineObject160WithDefaults() *InlineObject160`

NewInlineObject160WithDefaults instantiates a new InlineObject160 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject160) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject160) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject160) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject160) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomPieChartItems

`func (o *InlineObject160) GetCustomPieChartItems() []InlineResponse200162CustomPieChartItems`

GetCustomPieChartItems returns the CustomPieChartItems field if non-nil, zero value otherwise.

### GetCustomPieChartItemsOk

`func (o *InlineObject160) GetCustomPieChartItemsOk() (*[]InlineResponse200162CustomPieChartItems, bool)`

GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPieChartItems

`func (o *InlineObject160) SetCustomPieChartItems(v []InlineResponse200162CustomPieChartItems)`

SetCustomPieChartItems sets CustomPieChartItems field to given value.

### HasCustomPieChartItems

`func (o *InlineObject160) HasCustomPieChartItems() bool`

HasCustomPieChartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



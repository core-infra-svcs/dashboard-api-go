# InlineResponse200175

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |     The traffic analysis mode for the network. Can be one of &#39;disabled&#39; (do not collect traffic types),     &#39;basic&#39; (collect generic traffic categories), or &#39;detailed&#39; (collect destination hostnames).  | [optional] 
**CustomPieChartItems** | Pointer to [**[]InlineResponse200175CustomPieChartItems**](InlineResponse200175CustomPieChartItems.md) | The list of items that make up the custom pie chart for traffic reporting. | [optional] 

## Methods

### NewInlineResponse200175

`func NewInlineResponse200175() *InlineResponse200175`

NewInlineResponse200175 instantiates a new InlineResponse200175 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200175WithDefaults

`func NewInlineResponse200175WithDefaults() *InlineResponse200175`

NewInlineResponse200175WithDefaults instantiates a new InlineResponse200175 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200175) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200175) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200175) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200175) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomPieChartItems

`func (o *InlineResponse200175) GetCustomPieChartItems() []InlineResponse200175CustomPieChartItems`

GetCustomPieChartItems returns the CustomPieChartItems field if non-nil, zero value otherwise.

### GetCustomPieChartItemsOk

`func (o *InlineResponse200175) GetCustomPieChartItemsOk() (*[]InlineResponse200175CustomPieChartItems, bool)`

GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPieChartItems

`func (o *InlineResponse200175) SetCustomPieChartItems(v []InlineResponse200175CustomPieChartItems)`

SetCustomPieChartItems sets CustomPieChartItems field to given value.

### HasCustomPieChartItems

`func (o *InlineResponse200175) HasCustomPieChartItems() bool`

HasCustomPieChartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



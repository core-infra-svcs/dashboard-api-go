# InlineObject170

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |     The traffic analysis mode for the network. Can be one of &#39;disabled&#39; (do not collect traffic types),     &#39;basic&#39; (collect generic traffic categories), or &#39;detailed&#39; (collect destination hostnames).  | [optional] 
**CustomPieChartItems** | Pointer to [**[]InlineResponse200177CustomPieChartItems**](InlineResponse200177CustomPieChartItems.md) | The list of items that make up the custom pie chart for traffic reporting. | [optional] 

## Methods

### NewInlineObject170

`func NewInlineObject170() *InlineObject170`

NewInlineObject170 instantiates a new InlineObject170 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject170WithDefaults

`func NewInlineObject170WithDefaults() *InlineObject170`

NewInlineObject170WithDefaults instantiates a new InlineObject170 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject170) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject170) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject170) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject170) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomPieChartItems

`func (o *InlineObject170) GetCustomPieChartItems() []InlineResponse200177CustomPieChartItems`

GetCustomPieChartItems returns the CustomPieChartItems field if non-nil, zero value otherwise.

### GetCustomPieChartItemsOk

`func (o *InlineObject170) GetCustomPieChartItemsOk() (*[]InlineResponse200177CustomPieChartItems, bool)`

GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPieChartItems

`func (o *InlineObject170) SetCustomPieChartItems(v []InlineResponse200177CustomPieChartItems)`

SetCustomPieChartItems sets CustomPieChartItems field to given value.

### HasCustomPieChartItems

`func (o *InlineObject170) HasCustomPieChartItems() bool`

HasCustomPieChartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



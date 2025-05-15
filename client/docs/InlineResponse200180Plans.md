# InlineResponse200180Plans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the pricing plan to update. | [optional] 
**Price** | Pointer to **float32** | The price of the billing plan. | [optional] 
**BandwidthLimits** | Pointer to [**InlineResponse200180BandwidthLimits**](InlineResponse200180BandwidthLimits.md) |  | [optional] 
**TimeLimit** | Pointer to **string** | The time limit of the pricing plan in minutes. | [optional] 

## Methods

### NewInlineResponse200180Plans

`func NewInlineResponse200180Plans() *InlineResponse200180Plans`

NewInlineResponse200180Plans instantiates a new InlineResponse200180Plans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200180PlansWithDefaults

`func NewInlineResponse200180PlansWithDefaults() *InlineResponse200180Plans`

NewInlineResponse200180PlansWithDefaults instantiates a new InlineResponse200180Plans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200180Plans) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200180Plans) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200180Plans) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200180Plans) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *InlineResponse200180Plans) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse200180Plans) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse200180Plans) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InlineResponse200180Plans) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *InlineResponse200180Plans) GetBandwidthLimits() InlineResponse200180BandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *InlineResponse200180Plans) GetBandwidthLimitsOk() (*InlineResponse200180BandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *InlineResponse200180Plans) SetBandwidthLimits(v InlineResponse200180BandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *InlineResponse200180Plans) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.

### GetTimeLimit

`func (o *InlineResponse200180Plans) GetTimeLimit() string`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *InlineResponse200180Plans) GetTimeLimitOk() (*string, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *InlineResponse200180Plans) SetTimeLimit(v string)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *InlineResponse200180Plans) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



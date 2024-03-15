# InlineResponse200113Plans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the pricing plan to update. | [optional] 
**Price** | Pointer to **float32** | The price of the billing plan. | [optional] 
**BandwidthLimits** | Pointer to [**InlineResponse200113BandwidthLimits**](InlineResponse200113BandwidthLimits.md) |  | [optional] 
**TimeLimit** | Pointer to **string** | The time limit of the pricing plan in minutes. | [optional] 

## Methods

### NewInlineResponse200113Plans

`func NewInlineResponse200113Plans() *InlineResponse200113Plans`

NewInlineResponse200113Plans instantiates a new InlineResponse200113Plans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200113PlansWithDefaults

`func NewInlineResponse200113PlansWithDefaults() *InlineResponse200113Plans`

NewInlineResponse200113PlansWithDefaults instantiates a new InlineResponse200113Plans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200113Plans) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200113Plans) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200113Plans) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200113Plans) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrice

`func (o *InlineResponse200113Plans) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InlineResponse200113Plans) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InlineResponse200113Plans) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InlineResponse200113Plans) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *InlineResponse200113Plans) GetBandwidthLimits() InlineResponse200113BandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *InlineResponse200113Plans) GetBandwidthLimitsOk() (*InlineResponse200113BandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *InlineResponse200113Plans) SetBandwidthLimits(v InlineResponse200113BandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *InlineResponse200113Plans) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.

### GetTimeLimit

`func (o *InlineResponse200113Plans) GetTimeLimit() string`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *InlineResponse200113Plans) GetTimeLimitOk() (*string, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *InlineResponse200113Plans) SetTimeLimit(v string)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *InlineResponse200113Plans) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



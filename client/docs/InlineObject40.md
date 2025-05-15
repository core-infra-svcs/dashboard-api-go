# InlineObject40

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20050DefaultDestinations**](InlineResponse20050DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20050Alerts**](InlineResponse20050Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20050Muting**](InlineResponse20050Muting.md) |  | [optional] 

## Methods

### NewInlineObject40

`func NewInlineObject40() *InlineObject40`

NewInlineObject40 instantiates a new InlineObject40 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject40WithDefaults

`func NewInlineObject40WithDefaults() *InlineObject40`

NewInlineObject40WithDefaults instantiates a new InlineObject40 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject40) GetDefaultDestinations() InlineResponse20050DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject40) GetDefaultDestinationsOk() (*InlineResponse20050DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject40) SetDefaultDestinations(v InlineResponse20050DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject40) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject40) GetAlerts() []InlineResponse20050Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject40) GetAlertsOk() (*[]InlineResponse20050Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject40) SetAlerts(v []InlineResponse20050Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject40) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject40) GetMuting() InlineResponse20050Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject40) GetMutingOk() (*InlineResponse20050Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject40) SetMuting(v InlineResponse20050Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject40) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



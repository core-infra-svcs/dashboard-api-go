# InlineObject41

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20058DefaultDestinations**](InlineResponse20058DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20058Alerts**](InlineResponse20058Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20058Muting**](InlineResponse20058Muting.md) |  | [optional] 

## Methods

### NewInlineObject41

`func NewInlineObject41() *InlineObject41`

NewInlineObject41 instantiates a new InlineObject41 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject41WithDefaults

`func NewInlineObject41WithDefaults() *InlineObject41`

NewInlineObject41WithDefaults instantiates a new InlineObject41 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject41) GetDefaultDestinations() InlineResponse20058DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject41) GetDefaultDestinationsOk() (*InlineResponse20058DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject41) SetDefaultDestinations(v InlineResponse20058DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject41) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject41) GetAlerts() []InlineResponse20058Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject41) GetAlertsOk() (*[]InlineResponse20058Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject41) SetAlerts(v []InlineResponse20058Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject41) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject41) GetMuting() InlineResponse20058Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject41) GetMutingOk() (*InlineResponse20058Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject41) SetMuting(v InlineResponse20058Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject41) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



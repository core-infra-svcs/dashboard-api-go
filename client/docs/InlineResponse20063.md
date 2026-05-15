# InlineResponse20063

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**InlineResponse20063DefaultDestinations**](InlineResponse20063DefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]InlineResponse20063Alerts**](InlineResponse20063Alerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**InlineResponse20063Muting**](InlineResponse20063Muting.md) |  | [optional] 

## Methods

### NewInlineResponse20063

`func NewInlineResponse20063() *InlineResponse20063`

NewInlineResponse20063 instantiates a new InlineResponse20063 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20063WithDefaults

`func NewInlineResponse20063WithDefaults() *InlineResponse20063`

NewInlineResponse20063WithDefaults instantiates a new InlineResponse20063 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineResponse20063) GetDefaultDestinations() InlineResponse20063DefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineResponse20063) GetDefaultDestinationsOk() (*InlineResponse20063DefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineResponse20063) SetDefaultDestinations(v InlineResponse20063DefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineResponse20063) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineResponse20063) GetAlerts() []InlineResponse20063Alerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineResponse20063) GetAlertsOk() (*[]InlineResponse20063Alerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineResponse20063) SetAlerts(v []InlineResponse20063Alerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineResponse20063) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineResponse20063) GetMuting() InlineResponse20063Muting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineResponse20063) GetMutingOk() (*InlineResponse20063Muting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineResponse20063) SetMuting(v InlineResponse20063Muting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineResponse20063) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



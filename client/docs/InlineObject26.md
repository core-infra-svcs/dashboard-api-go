# InlineObject26

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**NetworksNetworkIdAlertsSettingsDefaultDestinations**](NetworksNetworkIdAlertsSettingsDefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]NetworksNetworkIdAlertsSettingsAlerts**](NetworksNetworkIdAlertsSettingsAlerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 

## Methods

### NewInlineObject26

`func NewInlineObject26() *InlineObject26`

NewInlineObject26 instantiates a new InlineObject26 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject26WithDefaults

`func NewInlineObject26WithDefaults() *InlineObject26`

NewInlineObject26WithDefaults instantiates a new InlineObject26 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject26) GetDefaultDestinations() NetworksNetworkIdAlertsSettingsDefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject26) GetDefaultDestinationsOk() (*NetworksNetworkIdAlertsSettingsDefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject26) SetDefaultDestinations(v NetworksNetworkIdAlertsSettingsDefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject26) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject26) GetAlerts() []NetworksNetworkIdAlertsSettingsAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject26) GetAlertsOk() (*[]NetworksNetworkIdAlertsSettingsAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject26) SetAlerts(v []NetworksNetworkIdAlertsSettingsAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject26) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



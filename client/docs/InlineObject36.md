# InlineObject36

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**NetworksNetworkIdAlertsSettingsDefaultDestinations**](NetworksNetworkIdAlertsSettingsDefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]NetworksNetworkIdAlertsSettingsAlerts**](NetworksNetworkIdAlertsSettingsAlerts.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**NetworksNetworkIdAlertsSettingsMuting**](NetworksNetworkIdAlertsSettingsMuting.md) |  | [optional] 

## Methods

### NewInlineObject36

`func NewInlineObject36() *InlineObject36`

NewInlineObject36 instantiates a new InlineObject36 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject36WithDefaults

`func NewInlineObject36WithDefaults() *InlineObject36`

NewInlineObject36WithDefaults instantiates a new InlineObject36 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *InlineObject36) GetDefaultDestinations() NetworksNetworkIdAlertsSettingsDefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *InlineObject36) GetDefaultDestinationsOk() (*NetworksNetworkIdAlertsSettingsDefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *InlineObject36) SetDefaultDestinations(v NetworksNetworkIdAlertsSettingsDefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *InlineObject36) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *InlineObject36) GetAlerts() []NetworksNetworkIdAlertsSettingsAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *InlineObject36) GetAlertsOk() (*[]NetworksNetworkIdAlertsSettingsAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *InlineObject36) SetAlerts(v []NetworksNetworkIdAlertsSettingsAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *InlineObject36) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *InlineObject36) GetMuting() NetworksNetworkIdAlertsSettingsMuting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *InlineObject36) GetMutingOk() (*NetworksNetworkIdAlertsSettingsMuting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *InlineObject36) SetMuting(v NetworksNetworkIdAlertsSettingsMuting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *InlineObject36) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OrganizationsOrganizationIdAlertsProfilesAlertCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The total duration in seconds that the threshold should be crossed before alerting | [optional] 
**Window** | Pointer to **int32** | The look back period in seconds for sensing the alert | [optional] 
**BitRateBps** | Pointer to **int32** | The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts. | [optional] 
**Interface** | Pointer to **string** | The uplink observed for the alert | [optional] 

## Methods

### NewOrganizationsOrganizationIdAlertsProfilesAlertCondition

`func NewOrganizationsOrganizationIdAlertsProfilesAlertCondition() *OrganizationsOrganizationIdAlertsProfilesAlertCondition`

NewOrganizationsOrganizationIdAlertsProfilesAlertCondition instantiates a new OrganizationsOrganizationIdAlertsProfilesAlertCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAlertsProfilesAlertConditionWithDefaults

`func NewOrganizationsOrganizationIdAlertsProfilesAlertConditionWithDefaults() *OrganizationsOrganizationIdAlertsProfilesAlertCondition`

NewOrganizationsOrganizationIdAlertsProfilesAlertConditionWithDefaults instantiates a new OrganizationsOrganizationIdAlertsProfilesAlertCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetWindow

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetWindow() int32`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetWindowOk() (*int32, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) SetWindow(v int32)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetBitRateBps

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetBitRateBps() int32`

GetBitRateBps returns the BitRateBps field if non-nil, zero value otherwise.

### GetBitRateBpsOk

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetBitRateBpsOk() (*int32, bool)`

GetBitRateBpsOk returns a tuple with the BitRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitRateBps

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) SetBitRateBps(v int32)`

SetBitRateBps sets BitRateBps field to given value.

### HasBitRateBps

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) HasBitRateBps() bool`

HasBitRateBps returns a boolean if a field has been set.

### GetInterface

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *OrganizationsOrganizationIdAlertsProfilesAlertCondition) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



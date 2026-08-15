# OrganizationsOrganizationIdDevicesCellularDataProfilesRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | **string** | The logical SIM identifier that the rule is applied to. One of: &#39;sim1&#39;, &#39;sim2&#39;, or &#39;esim&#39;. Devices included in this group must have an active card/profile with the defined SIM slot in order to use this rule. | 
**Uplink** | [**OrganizationsOrganizationIdDevicesCellularDataProfilesUplink**](OrganizationsOrganizationIdDevicesCellularDataProfilesUplink.md) |  | 
**Cap** | [**OrganizationsOrganizationIdDevicesCellularDataProfilesCap**](OrganizationsOrganizationIdDevicesCellularDataProfilesCap.md) |  | 
**Actions** | Pointer to [**[]InlineResponse200344Actions**](InlineResponse200344Actions.md) | The defined actions that will occur when the profile rule is triggered. No more than two actions may be defined for a rule. | [optional] 

## Methods

### NewOrganizationsOrganizationIdDevicesCellularDataProfilesRules

`func NewOrganizationsOrganizationIdDevicesCellularDataProfilesRules(slot string, uplink OrganizationsOrganizationIdDevicesCellularDataProfilesUplink, cap OrganizationsOrganizationIdDevicesCellularDataProfilesCap, ) *OrganizationsOrganizationIdDevicesCellularDataProfilesRules`

NewOrganizationsOrganizationIdDevicesCellularDataProfilesRules instantiates a new OrganizationsOrganizationIdDevicesCellularDataProfilesRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdDevicesCellularDataProfilesRulesWithDefaults

`func NewOrganizationsOrganizationIdDevicesCellularDataProfilesRulesWithDefaults() *OrganizationsOrganizationIdDevicesCellularDataProfilesRules`

NewOrganizationsOrganizationIdDevicesCellularDataProfilesRulesWithDefaults instantiates a new OrganizationsOrganizationIdDevicesCellularDataProfilesRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) SetSlot(v string)`

SetSlot sets Slot field to given value.


### GetUplink

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetUplink() OrganizationsOrganizationIdDevicesCellularDataProfilesUplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetUplinkOk() (*OrganizationsOrganizationIdDevicesCellularDataProfilesUplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) SetUplink(v OrganizationsOrganizationIdDevicesCellularDataProfilesUplink)`

SetUplink sets Uplink field to given value.


### GetCap

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetCap() OrganizationsOrganizationIdDevicesCellularDataProfilesCap`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetCapOk() (*OrganizationsOrganizationIdDevicesCellularDataProfilesCap, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) SetCap(v OrganizationsOrganizationIdDevicesCellularDataProfilesCap)`

SetCap sets Cap field to given value.


### GetActions

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetActions() []InlineResponse200344Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) GetActionsOk() (*[]InlineResponse200344Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) SetActions(v []InlineResponse200344Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *OrganizationsOrganizationIdDevicesCellularDataProfilesRules) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OrganizationsOrganizationIdAssuranceAlertsScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]OrganizationsOrganizationIdAssuranceAlertsScopeDevices**](OrganizationsOrganizationIdAssuranceAlertsScopeDevices.md) | Description of affected devices | [optional] 
**Applications** | Pointer to **[]map[string]interface{}** | Applications affected by the alert | [optional] 
**Peers** | Pointer to **[]map[string]interface{}** | Peers affected by the alert | [optional] 

## Methods

### NewOrganizationsOrganizationIdAssuranceAlertsScope

`func NewOrganizationsOrganizationIdAssuranceAlertsScope() *OrganizationsOrganizationIdAssuranceAlertsScope`

NewOrganizationsOrganizationIdAssuranceAlertsScope instantiates a new OrganizationsOrganizationIdAssuranceAlertsScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdAssuranceAlertsScopeWithDefaults

`func NewOrganizationsOrganizationIdAssuranceAlertsScopeWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsScope`

NewOrganizationsOrganizationIdAssuranceAlertsScopeWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetDevices() []OrganizationsOrganizationIdAssuranceAlertsScopeDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetDevicesOk() (*[]OrganizationsOrganizationIdAssuranceAlertsScopeDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) SetDevices(v []OrganizationsOrganizationIdAssuranceAlertsScopeDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetApplications

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetApplications() []map[string]interface{}`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetApplicationsOk() (*[]map[string]interface{}, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) SetApplications(v []map[string]interface{})`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPeers

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetPeers() []map[string]interface{}`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) GetPeersOk() (*[]map[string]interface{}, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) SetPeers(v []map[string]interface{})`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *OrganizationsOrganizationIdAssuranceAlertsScope) HasPeers() bool`

HasPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



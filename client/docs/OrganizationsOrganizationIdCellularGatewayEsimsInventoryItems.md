# OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice**](OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether eSIM is currently active SIM on Device | [optional] 
**Eid** | Pointer to **string** | eSIM EID | [optional] 
**LastUpdatedAt** | Pointer to **string** | Last update of eSIM | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork**](OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork.md) |  | [optional] 
**Profiles** | Pointer to [**[]OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles**](OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles.md) | eSIM Profile Information | [optional] 

## Methods

### NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems

`func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems`

NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItems instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItemsWithDefaults

`func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItemsWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems`

NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryItemsWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetDevice() OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetDeviceOk() (*OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetDevice(v OrganizationsOrganizationIdCellularGatewayEsimsInventoryDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetActive

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEid

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetNetwork() OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetNetworkOk() (*OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetNetwork(v OrganizationsOrganizationIdCellularGatewayEsimsInventoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProfiles

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetProfiles() []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) GetProfilesOk() (*[]OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) SetProfiles(v []OrganizationsOrganizationIdCellularGatewayEsimsInventoryProfiles)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



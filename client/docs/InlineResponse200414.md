# InlineResponse200414

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network identifier | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**Url** | Pointer to **string** | Network clients list URL | [optional] 
**Tags** | Pointer to **[]string** | Network tags | [optional] 
**Group** | Pointer to [**OrganizationsOrganizationIdSummaryTopNetworksByStatusGroup**](OrganizationsOrganizationIdSummaryTopNetworksByStatusGroup.md) |  | [optional] 
**Clients** | Pointer to [**OrganizationsOrganizationIdSummaryTopNetworksByStatusClients**](OrganizationsOrganizationIdSummaryTopNetworksByStatusClients.md) |  | [optional] 
**Statuses** | Pointer to [**OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses**](OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses.md) |  | [optional] 
**Devices** | Pointer to [**OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices**](OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices.md) |  | [optional] 
**ProductTypes** | Pointer to **[]string** | Product types in network | [optional] 

## Methods

### NewInlineResponse200414

`func NewInlineResponse200414() *InlineResponse200414`

NewInlineResponse200414 instantiates a new InlineResponse200414 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200414WithDefaults

`func NewInlineResponse200414WithDefaults() *InlineResponse200414`

NewInlineResponse200414WithDefaults instantiates a new InlineResponse200414 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200414) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200414) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200414) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200414) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200414) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200414) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200414) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200414) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200414) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200414) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200414) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200414) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200414) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200414) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200414) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200414) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGroup

`func (o *InlineResponse200414) GetGroup() OrganizationsOrganizationIdSummaryTopNetworksByStatusGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InlineResponse200414) GetGroupOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InlineResponse200414) SetGroup(v OrganizationsOrganizationIdSummaryTopNetworksByStatusGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InlineResponse200414) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200414) GetClients() OrganizationsOrganizationIdSummaryTopNetworksByStatusClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200414) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200414) SetClients(v OrganizationsOrganizationIdSummaryTopNetworksByStatusClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200414) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetStatuses

`func (o *InlineResponse200414) GetStatuses() OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *InlineResponse200414) GetStatusesOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *InlineResponse200414) SetStatuses(v OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *InlineResponse200414) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200414) GetDevices() OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200414) GetDevicesOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200414) SetDevices(v OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200414) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200414) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200414) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200414) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200414) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



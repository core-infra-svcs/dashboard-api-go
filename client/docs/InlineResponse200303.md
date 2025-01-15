# InlineResponse200303

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the SSID | [optional] 
**Usage** | Pointer to [**OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage**](OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage.md) |  | [optional] 
**Clients** | Pointer to [**OrganizationsOrganizationIdSummaryTopSsidsByUsageClients**](OrganizationsOrganizationIdSummaryTopSsidsByUsageClients.md) |  | [optional] 

## Methods

### NewInlineResponse200303

`func NewInlineResponse200303() *InlineResponse200303`

NewInlineResponse200303 instantiates a new InlineResponse200303 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200303WithDefaults

`func NewInlineResponse200303WithDefaults() *InlineResponse200303`

NewInlineResponse200303WithDefaults instantiates a new InlineResponse200303 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200303) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200303) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200303) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200303) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsage

`func (o *InlineResponse200303) GetUsage() OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineResponse200303) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineResponse200303) SetUsage(v OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineResponse200303) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetClients

`func (o *InlineResponse200303) GetClients() OrganizationsOrganizationIdSummaryTopSsidsByUsageClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse200303) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopSsidsByUsageClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse200303) SetClients(v OrganizationsOrganizationIdSummaryTopSsidsByUsageClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse200303) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



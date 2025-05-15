# InlineResponse200230

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network identifier | [optional] 
**Serial** | Pointer to **string** | The uplink serial | [optional] 
**Model** | Pointer to **string** | The uplink model | [optional] 
**LastReportedAt** | Pointer to **time.Time** | Last reported time for the device | [optional] 
**HighAvailability** | Pointer to [**OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability**](OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability.md) |  | [optional] 
**Uplinks** | Pointer to [**[]OrganizationsOrganizationIdApplianceUplinkStatusesUplinks**](OrganizationsOrganizationIdApplianceUplinkStatusesUplinks.md) | Uplinks | [optional] 

## Methods

### NewInlineResponse200230

`func NewInlineResponse200230() *InlineResponse200230`

NewInlineResponse200230 instantiates a new InlineResponse200230 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200230WithDefaults

`func NewInlineResponse200230WithDefaults() *InlineResponse200230`

NewInlineResponse200230WithDefaults instantiates a new InlineResponse200230 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200230) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200230) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200230) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200230) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200230) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200230) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200230) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200230) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200230) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200230) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200230) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200230) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *InlineResponse200230) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200230) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200230) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *InlineResponse200230) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetHighAvailability

`func (o *InlineResponse200230) GetHighAvailability() OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *InlineResponse200230) GetHighAvailabilityOk() (*OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *InlineResponse200230) SetHighAvailability(v OrganizationsOrganizationIdApplianceUplinkStatusesHighAvailability)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *InlineResponse200230) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.

### GetUplinks

`func (o *InlineResponse200230) GetUplinks() []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *InlineResponse200230) GetUplinksOk() (*[]OrganizationsOrganizationIdApplianceUplinkStatusesUplinks, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *InlineResponse200230) SetUplinks(v []OrganizationsOrganizationIdApplianceUplinkStatusesUplinks)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *InlineResponse200230) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200446

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the readings. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNetwork**](OrganizationsOrganizationIdSensorReadingsHistoryNetwork.md) |  | [optional] 
**Readings** | Pointer to [**[]OrganizationsOrganizationIdSensorReadingsLatestReadings**](OrganizationsOrganizationIdSensorReadingsLatestReadings.md) | Array of latest readings from the sensor. Each object represents a single reading for a single metric. | [optional] 

## Methods

### NewInlineResponse200446

`func NewInlineResponse200446() *InlineResponse200446`

NewInlineResponse200446 instantiates a new InlineResponse200446 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200446WithDefaults

`func NewInlineResponse200446WithDefaults() *InlineResponse200446`

NewInlineResponse200446WithDefaults instantiates a new InlineResponse200446 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200446) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200446) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200446) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200446) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200446) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200446) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200446) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200446) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200446) GetReadings() []OrganizationsOrganizationIdSensorReadingsLatestReadings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200446) GetReadingsOk() (*[]OrganizationsOrganizationIdSensorReadingsLatestReadings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200446) SetReadings(v []OrganizationsOrganizationIdSensorReadingsLatestReadings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200446) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



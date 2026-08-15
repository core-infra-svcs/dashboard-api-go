# InlineResponse200425

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number of the sensor that took the readings. | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdSensorReadingsHistoryNetwork**](OrganizationsOrganizationIdSensorReadingsHistoryNetwork.md) |  | [optional] 
**Readings** | Pointer to [**[]OrganizationsOrganizationIdSensorReadingsLatestReadings**](OrganizationsOrganizationIdSensorReadingsLatestReadings.md) | Array of latest readings from the sensor. Each object represents a single reading for a single metric. | [optional] 

## Methods

### NewInlineResponse200425

`func NewInlineResponse200425() *InlineResponse200425`

NewInlineResponse200425 instantiates a new InlineResponse200425 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200425WithDefaults

`func NewInlineResponse200425WithDefaults() *InlineResponse200425`

NewInlineResponse200425WithDefaults instantiates a new InlineResponse200425 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200425) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200425) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200425) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200425) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200425) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200425) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200425) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200425) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200425) GetReadings() []OrganizationsOrganizationIdSensorReadingsLatestReadings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200425) GetReadingsOk() (*[]OrganizationsOrganizationIdSensorReadingsLatestReadings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200425) SetReadings(v []OrganizationsOrganizationIdSensorReadingsLatestReadings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200425) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200249

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp, in iso8601 format, at which the event happened | [optional] 
**Device** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice**](OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice.md) |  | [optional] 
**Details** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails**](OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails.md) |  | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork**](OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork.md) |  | [optional] 

## Methods

### NewInlineResponse200249

`func NewInlineResponse200249() *InlineResponse200249`

NewInlineResponse200249 instantiates a new InlineResponse200249 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200249WithDefaults

`func NewInlineResponse200249WithDefaults() *InlineResponse200249`

NewInlineResponse200249WithDefaults instantiates a new InlineResponse200249 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200249) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200249) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200249) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200249) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200249) GetDevice() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200249) GetDeviceOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200249) SetDevice(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200249) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200249) GetDetails() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200249) GetDetailsOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200249) SetDetails(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200249) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200249) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200249) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200249) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200249) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



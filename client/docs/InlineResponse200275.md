# InlineResponse200275

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp, in iso8601 format, at which the event happened | [optional] 
**Device** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice**](OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice.md) |  | [optional] 
**Details** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails**](OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails.md) |  | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork**](OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork.md) |  | [optional] 

## Methods

### NewInlineResponse200275

`func NewInlineResponse200275() *InlineResponse200275`

NewInlineResponse200275 instantiates a new InlineResponse200275 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200275WithDefaults

`func NewInlineResponse200275WithDefaults() *InlineResponse200275`

NewInlineResponse200275WithDefaults instantiates a new InlineResponse200275 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200275) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200275) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200275) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200275) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetDevice

`func (o *InlineResponse200275) GetDevice() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200275) GetDeviceOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200275) SetDevice(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200275) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200275) GetDetails() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200275) GetDetailsOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200275) SetDetails(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200275) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200275) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200275) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200275) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesChangeHistoryNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200275) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



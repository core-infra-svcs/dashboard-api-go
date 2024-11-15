# OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Ts** | Pointer to **time.Time** | Failover time | [optional] 
**Reason** | Pointer to **string** | Failover reason | [optional] 
**Failed** | Pointer to [**OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed**](OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed.md) |  | [optional] 
**Active** | Pointer to [**OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive**](OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems

`func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems`

NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItemsWithDefaults

`func NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItemsWithDefaults() *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems`

NewOrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItemsWithDefaults instantiates a new OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTs

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetReason

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetFailed

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetFailed() OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetFailedOk() (*OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetFailed(v OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryFailed)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetActive

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetActive() OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) GetActiveOk() (*OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) SetActive(v OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryActive)`

SetActive sets Active field to given value.

### HasActive

`func (o *OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



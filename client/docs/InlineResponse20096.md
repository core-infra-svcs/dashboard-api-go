# InlineResponse20096

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | License ID | [optional] 
**LicenseType** | Pointer to **string** | License type | [optional] 
**LicenseKey** | Pointer to **string** | License key | [optional] 
**OrderNumber** | Pointer to **string** | Order number | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device the license is assigned to | [optional] 
**NetworkId** | Pointer to **string** | ID of the network the license is assigned to | [optional] 
**State** | Pointer to **string** | The state of the license. All queued licenses have a status of &#x60;recentlyQueued&#x60;. | [optional] 
**SeatCount** | Pointer to **int32** | The number of seats of the license. Only applicable to SM licenses. | [optional] 
**TotalDurationInDays** | Pointer to **int32** | The duration of the license plus all permanently queued licenses associated with it | [optional] 
**DurationInDays** | Pointer to **int32** | The duration of the individual license | [optional] 
**PermanentlyQueuedLicenses** | Pointer to [**[]OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses**](OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses.md) | DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial&#x3D; to retrieved queued licenses for a given device. | [optional] 
**ClaimDate** | Pointer to **string** | The date the license was claimed into the organization | [optional] 
**ActivationDate** | Pointer to **string** | The date the license started burning | [optional] 
**ExpirationDate** | Pointer to **string** | The date the license will expire | [optional] 
**HeadLicenseId** | Pointer to **string** | The id of the head license this license is queued behind. If there is no head license, it returns nil. | [optional] 

## Methods

### NewInlineResponse20096

`func NewInlineResponse20096() *InlineResponse20096`

NewInlineResponse20096 instantiates a new InlineResponse20096 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20096WithDefaults

`func NewInlineResponse20096WithDefaults() *InlineResponse20096`

NewInlineResponse20096WithDefaults instantiates a new InlineResponse20096 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20096) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20096) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20096) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20096) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseType

`func (o *InlineResponse20096) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *InlineResponse20096) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *InlineResponse20096) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *InlineResponse20096) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicenseKey

`func (o *InlineResponse20096) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *InlineResponse20096) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *InlineResponse20096) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *InlineResponse20096) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetOrderNumber

`func (o *InlineResponse20096) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InlineResponse20096) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InlineResponse20096) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *InlineResponse20096) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse20096) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse20096) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse20096) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse20096) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse20096) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20096) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20096) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20096) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetState

`func (o *InlineResponse20096) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InlineResponse20096) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InlineResponse20096) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InlineResponse20096) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSeatCount

`func (o *InlineResponse20096) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *InlineResponse20096) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *InlineResponse20096) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.

### HasSeatCount

`func (o *InlineResponse20096) HasSeatCount() bool`

HasSeatCount returns a boolean if a field has been set.

### GetTotalDurationInDays

`func (o *InlineResponse20096) GetTotalDurationInDays() int32`

GetTotalDurationInDays returns the TotalDurationInDays field if non-nil, zero value otherwise.

### GetTotalDurationInDaysOk

`func (o *InlineResponse20096) GetTotalDurationInDaysOk() (*int32, bool)`

GetTotalDurationInDaysOk returns a tuple with the TotalDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationInDays

`func (o *InlineResponse20096) SetTotalDurationInDays(v int32)`

SetTotalDurationInDays sets TotalDurationInDays field to given value.

### HasTotalDurationInDays

`func (o *InlineResponse20096) HasTotalDurationInDays() bool`

HasTotalDurationInDays returns a boolean if a field has been set.

### GetDurationInDays

`func (o *InlineResponse20096) GetDurationInDays() int32`

GetDurationInDays returns the DurationInDays field if non-nil, zero value otherwise.

### GetDurationInDaysOk

`func (o *InlineResponse20096) GetDurationInDaysOk() (*int32, bool)`

GetDurationInDaysOk returns a tuple with the DurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInDays

`func (o *InlineResponse20096) SetDurationInDays(v int32)`

SetDurationInDays sets DurationInDays field to given value.

### HasDurationInDays

`func (o *InlineResponse20096) HasDurationInDays() bool`

HasDurationInDays returns a boolean if a field has been set.

### GetPermanentlyQueuedLicenses

`func (o *InlineResponse20096) GetPermanentlyQueuedLicenses() []OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses`

GetPermanentlyQueuedLicenses returns the PermanentlyQueuedLicenses field if non-nil, zero value otherwise.

### GetPermanentlyQueuedLicensesOk

`func (o *InlineResponse20096) GetPermanentlyQueuedLicensesOk() (*[]OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses, bool)`

GetPermanentlyQueuedLicensesOk returns a tuple with the PermanentlyQueuedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentlyQueuedLicenses

`func (o *InlineResponse20096) SetPermanentlyQueuedLicenses(v []OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses)`

SetPermanentlyQueuedLicenses sets PermanentlyQueuedLicenses field to given value.

### HasPermanentlyQueuedLicenses

`func (o *InlineResponse20096) HasPermanentlyQueuedLicenses() bool`

HasPermanentlyQueuedLicenses returns a boolean if a field has been set.

### GetClaimDate

`func (o *InlineResponse20096) GetClaimDate() string`

GetClaimDate returns the ClaimDate field if non-nil, zero value otherwise.

### GetClaimDateOk

`func (o *InlineResponse20096) GetClaimDateOk() (*string, bool)`

GetClaimDateOk returns a tuple with the ClaimDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimDate

`func (o *InlineResponse20096) SetClaimDate(v string)`

SetClaimDate sets ClaimDate field to given value.

### HasClaimDate

`func (o *InlineResponse20096) HasClaimDate() bool`

HasClaimDate returns a boolean if a field has been set.

### GetActivationDate

`func (o *InlineResponse20096) GetActivationDate() string`

GetActivationDate returns the ActivationDate field if non-nil, zero value otherwise.

### GetActivationDateOk

`func (o *InlineResponse20096) GetActivationDateOk() (*string, bool)`

GetActivationDateOk returns a tuple with the ActivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationDate

`func (o *InlineResponse20096) SetActivationDate(v string)`

SetActivationDate sets ActivationDate field to given value.

### HasActivationDate

`func (o *InlineResponse20096) HasActivationDate() bool`

HasActivationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InlineResponse20096) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InlineResponse20096) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InlineResponse20096) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InlineResponse20096) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetHeadLicenseId

`func (o *InlineResponse20096) GetHeadLicenseId() string`

GetHeadLicenseId returns the HeadLicenseId field if non-nil, zero value otherwise.

### GetHeadLicenseIdOk

`func (o *InlineResponse20096) GetHeadLicenseIdOk() (*string, bool)`

GetHeadLicenseIdOk returns a tuple with the HeadLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadLicenseId

`func (o *InlineResponse20096) SetHeadLicenseId(v string)`

SetHeadLicenseId sets HeadLicenseId field to given value.

### HasHeadLicenseId

`func (o *InlineResponse20096) HasHeadLicenseId() bool`

HasHeadLicenseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



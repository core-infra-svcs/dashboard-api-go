# InlineResponse200371Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Wireless LAN controller cloud ID | [optional] 
**Mode** | Pointer to **string** | Wireless LAN controller redundancy SSO (stateful switchover) | [optional] 
**Enabled** | Pointer to **bool** | Wireless LAN controller redundancy enablement | [optional] 
**Failover** | Pointer to [**InlineResponse200371Failover**](InlineResponse200371Failover.md) |  | [optional] 
**MobilityMac** | Pointer to **string** | Wireless LAN controller redundancy mobility mac  | [optional] 

## Methods

### NewInlineResponse200371Items

`func NewInlineResponse200371Items() *InlineResponse200371Items`

NewInlineResponse200371Items instantiates a new InlineResponse200371Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200371ItemsWithDefaults

`func NewInlineResponse200371ItemsWithDefaults() *InlineResponse200371Items`

NewInlineResponse200371ItemsWithDefaults instantiates a new InlineResponse200371Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200371Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200371Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200371Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200371Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200371Items) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200371Items) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200371Items) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200371Items) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200371Items) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200371Items) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200371Items) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200371Items) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailover

`func (o *InlineResponse200371Items) GetFailover() InlineResponse200371Failover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineResponse200371Items) GetFailoverOk() (*InlineResponse200371Failover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineResponse200371Items) SetFailover(v InlineResponse200371Failover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineResponse200371Items) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetMobilityMac

`func (o *InlineResponse200371Items) GetMobilityMac() string`

GetMobilityMac returns the MobilityMac field if non-nil, zero value otherwise.

### GetMobilityMacOk

`func (o *InlineResponse200371Items) GetMobilityMacOk() (*string, bool)`

GetMobilityMacOk returns a tuple with the MobilityMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilityMac

`func (o *InlineResponse200371Items) SetMobilityMac(v string)`

SetMobilityMac sets MobilityMac field to given value.

### HasMobilityMac

`func (o *InlineResponse200371Items) HasMobilityMac() bool`

HasMobilityMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



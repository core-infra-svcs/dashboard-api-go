# AdministeredLicensingSubscriptionSubscriptionsCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seats** | Pointer to [**AdministeredLicensingSubscriptionSubscriptionsCountsSeats**](AdministeredLicensingSubscriptionSubscriptionsCountsSeats.md) |  | [optional] 
**Networks** | Pointer to **int32** | Number of networks bound to this subscription | [optional] 

## Methods

### NewAdministeredLicensingSubscriptionSubscriptionsCounts

`func NewAdministeredLicensingSubscriptionSubscriptionsCounts() *AdministeredLicensingSubscriptionSubscriptionsCounts`

NewAdministeredLicensingSubscriptionSubscriptionsCounts instantiates a new AdministeredLicensingSubscriptionSubscriptionsCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministeredLicensingSubscriptionSubscriptionsCountsWithDefaults

`func NewAdministeredLicensingSubscriptionSubscriptionsCountsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsCounts`

NewAdministeredLicensingSubscriptionSubscriptionsCountsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeats

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetSeats() AdministeredLicensingSubscriptionSubscriptionsCountsSeats`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetSeatsOk() (*AdministeredLicensingSubscriptionSubscriptionsCountsSeats, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) SetSeats(v AdministeredLicensingSubscriptionSubscriptionsCountsSeats)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### GetNetworks

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetNetworks() int32`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetNetworksOk() (*int32, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) SetNetworks(v int32)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



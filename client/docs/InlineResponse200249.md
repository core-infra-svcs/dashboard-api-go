# InlineResponse200249

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not Hotspot 2.0 for this SSID is enabled | [optional] 
**Operator** | Pointer to [**InlineResponse200249Operator**](InlineResponse200249Operator.md) |  | [optional] 
**Venue** | Pointer to [**InlineResponse200249Venue**](InlineResponse200249Venue.md) |  | [optional] 
**NetworkAccessType** | Pointer to **string** | The network type of this SSID | [optional] 
**Domains** | Pointer to **[]string** | An array of domain names | [optional] 
**RoamConsortOis** | Pointer to **[]string** | An array of roaming consortium OIs (hexadecimal number 3-5 octets in length) | [optional] 
**MccMncs** | Pointer to [**[]InlineResponse200249MccMncs**](InlineResponse200249MccMncs.md) | An array of MCC/MNC pairs | [optional] 
**NaiRealms** | Pointer to [**[]InlineResponse200249NaiRealms**](InlineResponse200249NaiRealms.md) | An array of NAI realms | [optional] 

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

### GetEnabled

`func (o *InlineResponse200249) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200249) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200249) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200249) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOperator

`func (o *InlineResponse200249) GetOperator() InlineResponse200249Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InlineResponse200249) GetOperatorOk() (*InlineResponse200249Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InlineResponse200249) SetOperator(v InlineResponse200249Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *InlineResponse200249) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetVenue

`func (o *InlineResponse200249) GetVenue() InlineResponse200249Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *InlineResponse200249) GetVenueOk() (*InlineResponse200249Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *InlineResponse200249) SetVenue(v InlineResponse200249Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *InlineResponse200249) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetNetworkAccessType

`func (o *InlineResponse200249) GetNetworkAccessType() string`

GetNetworkAccessType returns the NetworkAccessType field if non-nil, zero value otherwise.

### GetNetworkAccessTypeOk

`func (o *InlineResponse200249) GetNetworkAccessTypeOk() (*string, bool)`

GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessType

`func (o *InlineResponse200249) SetNetworkAccessType(v string)`

SetNetworkAccessType sets NetworkAccessType field to given value.

### HasNetworkAccessType

`func (o *InlineResponse200249) HasNetworkAccessType() bool`

HasNetworkAccessType returns a boolean if a field has been set.

### GetDomains

`func (o *InlineResponse200249) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *InlineResponse200249) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *InlineResponse200249) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *InlineResponse200249) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetRoamConsortOis

`func (o *InlineResponse200249) GetRoamConsortOis() []string`

GetRoamConsortOis returns the RoamConsortOis field if non-nil, zero value otherwise.

### GetRoamConsortOisOk

`func (o *InlineResponse200249) GetRoamConsortOisOk() (*[]string, bool)`

GetRoamConsortOisOk returns a tuple with the RoamConsortOis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamConsortOis

`func (o *InlineResponse200249) SetRoamConsortOis(v []string)`

SetRoamConsortOis sets RoamConsortOis field to given value.

### HasRoamConsortOis

`func (o *InlineResponse200249) HasRoamConsortOis() bool`

HasRoamConsortOis returns a boolean if a field has been set.

### GetMccMncs

`func (o *InlineResponse200249) GetMccMncs() []InlineResponse200249MccMncs`

GetMccMncs returns the MccMncs field if non-nil, zero value otherwise.

### GetMccMncsOk

`func (o *InlineResponse200249) GetMccMncsOk() (*[]InlineResponse200249MccMncs, bool)`

GetMccMncsOk returns a tuple with the MccMncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccMncs

`func (o *InlineResponse200249) SetMccMncs(v []InlineResponse200249MccMncs)`

SetMccMncs sets MccMncs field to given value.

### HasMccMncs

`func (o *InlineResponse200249) HasMccMncs() bool`

HasMccMncs returns a boolean if a field has been set.

### GetNaiRealms

`func (o *InlineResponse200249) GetNaiRealms() []InlineResponse200249NaiRealms`

GetNaiRealms returns the NaiRealms field if non-nil, zero value otherwise.

### GetNaiRealmsOk

`func (o *InlineResponse200249) GetNaiRealmsOk() (*[]InlineResponse200249NaiRealms, bool)`

GetNaiRealmsOk returns a tuple with the NaiRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiRealms

`func (o *InlineResponse200249) SetNaiRealms(v []InlineResponse200249NaiRealms)`

SetNaiRealms sets NaiRealms field to given value.

### HasNaiRealms

`func (o *InlineResponse200249) HasNaiRealms() bool`

HasNaiRealms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



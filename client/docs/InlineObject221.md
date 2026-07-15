# InlineObject221

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether or not Hotspot 2.0 for this SSID is enabled | [optional] 
**Operator** | Pointer to [**InlineResponse200249Operator**](InlineResponse200249Operator.md) |  | [optional] 
**Venue** | Pointer to [**NetworksNetworkIdWirelessSsidsNumberHotspot20Venue**](NetworksNetworkIdWirelessSsidsNumberHotspot20Venue.md) |  | [optional] 
**NetworkAccessType** | Pointer to **string** | The network type of this SSID (&#39;Private network&#39;, &#39;Private network with guest access&#39;, &#39;Chargeable public network&#39;, &#39;Free public network&#39;, &#39;Personal device network&#39;, &#39;Emergency services only network&#39;, &#39;Test or experimental&#39;, &#39;Wildcard&#39;) | [optional] 
**Domains** | Pointer to **[]string** | An array of domain names | [optional] 
**RoamConsortOis** | Pointer to **[]string** | An array of roaming consortium OIs (hexadecimal number 3-5 octets in length) | [optional] 
**MccMncs** | Pointer to [**[]InlineResponse200249MccMncs**](InlineResponse200249MccMncs.md) | An array of MCC/MNC pairs | [optional] 
**NaiRealms** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms**](NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms.md) | An array of NAI realms | [optional] 

## Methods

### NewInlineObject221

`func NewInlineObject221() *InlineObject221`

NewInlineObject221 instantiates a new InlineObject221 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject221WithDefaults

`func NewInlineObject221WithDefaults() *InlineObject221`

NewInlineObject221WithDefaults instantiates a new InlineObject221 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject221) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject221) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject221) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject221) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOperator

`func (o *InlineObject221) GetOperator() InlineResponse200249Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InlineObject221) GetOperatorOk() (*InlineResponse200249Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InlineObject221) SetOperator(v InlineResponse200249Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *InlineObject221) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetVenue

`func (o *InlineObject221) GetVenue() NetworksNetworkIdWirelessSsidsNumberHotspot20Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *InlineObject221) GetVenueOk() (*NetworksNetworkIdWirelessSsidsNumberHotspot20Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *InlineObject221) SetVenue(v NetworksNetworkIdWirelessSsidsNumberHotspot20Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *InlineObject221) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetNetworkAccessType

`func (o *InlineObject221) GetNetworkAccessType() string`

GetNetworkAccessType returns the NetworkAccessType field if non-nil, zero value otherwise.

### GetNetworkAccessTypeOk

`func (o *InlineObject221) GetNetworkAccessTypeOk() (*string, bool)`

GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessType

`func (o *InlineObject221) SetNetworkAccessType(v string)`

SetNetworkAccessType sets NetworkAccessType field to given value.

### HasNetworkAccessType

`func (o *InlineObject221) HasNetworkAccessType() bool`

HasNetworkAccessType returns a boolean if a field has been set.

### GetDomains

`func (o *InlineObject221) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *InlineObject221) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *InlineObject221) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *InlineObject221) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetRoamConsortOis

`func (o *InlineObject221) GetRoamConsortOis() []string`

GetRoamConsortOis returns the RoamConsortOis field if non-nil, zero value otherwise.

### GetRoamConsortOisOk

`func (o *InlineObject221) GetRoamConsortOisOk() (*[]string, bool)`

GetRoamConsortOisOk returns a tuple with the RoamConsortOis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamConsortOis

`func (o *InlineObject221) SetRoamConsortOis(v []string)`

SetRoamConsortOis sets RoamConsortOis field to given value.

### HasRoamConsortOis

`func (o *InlineObject221) HasRoamConsortOis() bool`

HasRoamConsortOis returns a boolean if a field has been set.

### GetMccMncs

`func (o *InlineObject221) GetMccMncs() []InlineResponse200249MccMncs`

GetMccMncs returns the MccMncs field if non-nil, zero value otherwise.

### GetMccMncsOk

`func (o *InlineObject221) GetMccMncsOk() (*[]InlineResponse200249MccMncs, bool)`

GetMccMncsOk returns a tuple with the MccMncs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMccMncs

`func (o *InlineObject221) SetMccMncs(v []InlineResponse200249MccMncs)`

SetMccMncs sets MccMncs field to given value.

### HasMccMncs

`func (o *InlineObject221) HasMccMncs() bool`

HasMccMncs returns a boolean if a field has been set.

### GetNaiRealms

`func (o *InlineObject221) GetNaiRealms() []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms`

GetNaiRealms returns the NaiRealms field if non-nil, zero value otherwise.

### GetNaiRealmsOk

`func (o *InlineObject221) GetNaiRealmsOk() (*[]NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms, bool)`

GetNaiRealmsOk returns a tuple with the NaiRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiRealms

`func (o *InlineObject221) SetNaiRealms(v []NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms)`

SetNaiRealms sets NaiRealms field to given value.

### HasNaiRealms

`func (o *InlineObject221) HasNaiRealms() bool`

HasNaiRealms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



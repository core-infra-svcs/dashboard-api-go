# InlineObject110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, mDNS gateway is enabled for this SSID and cluster. | [optional] 
**Rules** | Pointer to [**[]NetworksNetworkIdCampusGatewaySsidsNumberMdnsRules**](NetworksNetworkIdCampusGatewaySsidsNumberMdnsRules.md) | List of mDNS forwarding rules. | [optional] 

## Methods

### NewInlineObject110

`func NewInlineObject110() *InlineObject110`

NewInlineObject110 instantiates a new InlineObject110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject110WithDefaults

`func NewInlineObject110WithDefaults() *InlineObject110`

NewInlineObject110WithDefaults instantiates a new InlineObject110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject110) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject110) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject110) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject110) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject110) GetRules() []NetworksNetworkIdCampusGatewaySsidsNumberMdnsRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject110) GetRulesOk() (*[]NetworksNetworkIdCampusGatewaySsidsNumberMdnsRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject110) SetRules(v []NetworksNetworkIdCampusGatewaySsidsNumberMdnsRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject110) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



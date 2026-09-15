# InlineResponse200493SsidAddressing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Addressing mode | [optional] 
**AdultContentFiltering** | Pointer to [**InlineResponse200493SsidAddressingAdultContentFiltering**](InlineResponse200493SsidAddressingAdultContentFiltering.md) |  | [optional] 
**DnsRewrite** | Pointer to [**InlineResponse200493SsidAddressingDnsRewrite**](InlineResponse200493SsidAddressingDnsRewrite.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**InlineResponse200493SsidAddressingBonjourForwarding**](InlineResponse200493SsidAddressingBonjourForwarding.md) |  | [optional] 
**Mdns** | Pointer to [**InlineResponse200493SsidAddressingMdns**](InlineResponse200493SsidAddressingMdns.md) |  | [optional] 
**Radius** | Pointer to [**InlineResponse200493SsidAddressingRadius**](InlineResponse200493SsidAddressingRadius.md) |  | [optional] 

## Methods

### NewInlineResponse200493SsidAddressing

`func NewInlineResponse200493SsidAddressing() *InlineResponse200493SsidAddressing`

NewInlineResponse200493SsidAddressing instantiates a new InlineResponse200493SsidAddressing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200493SsidAddressingWithDefaults

`func NewInlineResponse200493SsidAddressingWithDefaults() *InlineResponse200493SsidAddressing`

NewInlineResponse200493SsidAddressingWithDefaults instantiates a new InlineResponse200493SsidAddressing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200493SsidAddressing) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200493SsidAddressing) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200493SsidAddressing) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200493SsidAddressing) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAdultContentFiltering

`func (o *InlineResponse200493SsidAddressing) GetAdultContentFiltering() InlineResponse200493SsidAddressingAdultContentFiltering`

GetAdultContentFiltering returns the AdultContentFiltering field if non-nil, zero value otherwise.

### GetAdultContentFilteringOk

`func (o *InlineResponse200493SsidAddressing) GetAdultContentFilteringOk() (*InlineResponse200493SsidAddressingAdultContentFiltering, bool)`

GetAdultContentFilteringOk returns a tuple with the AdultContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContentFiltering

`func (o *InlineResponse200493SsidAddressing) SetAdultContentFiltering(v InlineResponse200493SsidAddressingAdultContentFiltering)`

SetAdultContentFiltering sets AdultContentFiltering field to given value.

### HasAdultContentFiltering

`func (o *InlineResponse200493SsidAddressing) HasAdultContentFiltering() bool`

HasAdultContentFiltering returns a boolean if a field has been set.

### GetDnsRewrite

`func (o *InlineResponse200493SsidAddressing) GetDnsRewrite() InlineResponse200493SsidAddressingDnsRewrite`

GetDnsRewrite returns the DnsRewrite field if non-nil, zero value otherwise.

### GetDnsRewriteOk

`func (o *InlineResponse200493SsidAddressing) GetDnsRewriteOk() (*InlineResponse200493SsidAddressingDnsRewrite, bool)`

GetDnsRewriteOk returns a tuple with the DnsRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRewrite

`func (o *InlineResponse200493SsidAddressing) SetDnsRewrite(v InlineResponse200493SsidAddressingDnsRewrite)`

SetDnsRewrite sets DnsRewrite field to given value.

### HasDnsRewrite

`func (o *InlineResponse200493SsidAddressing) HasDnsRewrite() bool`

HasDnsRewrite returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineResponse200493SsidAddressing) GetBonjourForwarding() InlineResponse200493SsidAddressingBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineResponse200493SsidAddressing) GetBonjourForwardingOk() (*InlineResponse200493SsidAddressingBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineResponse200493SsidAddressing) SetBonjourForwarding(v InlineResponse200493SsidAddressingBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineResponse200493SsidAddressing) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.

### GetMdns

`func (o *InlineResponse200493SsidAddressing) GetMdns() InlineResponse200493SsidAddressingMdns`

GetMdns returns the Mdns field if non-nil, zero value otherwise.

### GetMdnsOk

`func (o *InlineResponse200493SsidAddressing) GetMdnsOk() (*InlineResponse200493SsidAddressingMdns, bool)`

GetMdnsOk returns a tuple with the Mdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdns

`func (o *InlineResponse200493SsidAddressing) SetMdns(v InlineResponse200493SsidAddressingMdns)`

SetMdns sets Mdns field to given value.

### HasMdns

`func (o *InlineResponse200493SsidAddressing) HasMdns() bool`

HasMdns returns a boolean if a field has been set.

### GetRadius

`func (o *InlineResponse200493SsidAddressing) GetRadius() InlineResponse200493SsidAddressingRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineResponse200493SsidAddressing) GetRadiusOk() (*InlineResponse200493SsidAddressingRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineResponse200493SsidAddressing) SetRadius(v InlineResponse200493SsidAddressingRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineResponse200493SsidAddressing) HasRadius() bool`

HasRadius returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



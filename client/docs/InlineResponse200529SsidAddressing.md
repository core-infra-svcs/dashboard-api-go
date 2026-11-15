# InlineResponse200529SsidAddressing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Addressing mode | [optional] 
**AdultContentFiltering** | Pointer to [**InlineResponse200529SsidAddressingAdultContentFiltering**](InlineResponse200529SsidAddressingAdultContentFiltering.md) |  | [optional] 
**DnsRewrite** | Pointer to [**InlineResponse200529SsidAddressingDnsRewrite**](InlineResponse200529SsidAddressingDnsRewrite.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**InlineResponse200529SsidAddressingBonjourForwarding**](InlineResponse200529SsidAddressingBonjourForwarding.md) |  | [optional] 
**Mdns** | Pointer to [**InlineResponse200529SsidAddressingMdns**](InlineResponse200529SsidAddressingMdns.md) |  | [optional] 
**Radius** | Pointer to [**InlineResponse200529SsidAddressingRadius**](InlineResponse200529SsidAddressingRadius.md) |  | [optional] 

## Methods

### NewInlineResponse200529SsidAddressing

`func NewInlineResponse200529SsidAddressing() *InlineResponse200529SsidAddressing`

NewInlineResponse200529SsidAddressing instantiates a new InlineResponse200529SsidAddressing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200529SsidAddressingWithDefaults

`func NewInlineResponse200529SsidAddressingWithDefaults() *InlineResponse200529SsidAddressing`

NewInlineResponse200529SsidAddressingWithDefaults instantiates a new InlineResponse200529SsidAddressing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200529SsidAddressing) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200529SsidAddressing) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200529SsidAddressing) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200529SsidAddressing) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAdultContentFiltering

`func (o *InlineResponse200529SsidAddressing) GetAdultContentFiltering() InlineResponse200529SsidAddressingAdultContentFiltering`

GetAdultContentFiltering returns the AdultContentFiltering field if non-nil, zero value otherwise.

### GetAdultContentFilteringOk

`func (o *InlineResponse200529SsidAddressing) GetAdultContentFilteringOk() (*InlineResponse200529SsidAddressingAdultContentFiltering, bool)`

GetAdultContentFilteringOk returns a tuple with the AdultContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContentFiltering

`func (o *InlineResponse200529SsidAddressing) SetAdultContentFiltering(v InlineResponse200529SsidAddressingAdultContentFiltering)`

SetAdultContentFiltering sets AdultContentFiltering field to given value.

### HasAdultContentFiltering

`func (o *InlineResponse200529SsidAddressing) HasAdultContentFiltering() bool`

HasAdultContentFiltering returns a boolean if a field has been set.

### GetDnsRewrite

`func (o *InlineResponse200529SsidAddressing) GetDnsRewrite() InlineResponse200529SsidAddressingDnsRewrite`

GetDnsRewrite returns the DnsRewrite field if non-nil, zero value otherwise.

### GetDnsRewriteOk

`func (o *InlineResponse200529SsidAddressing) GetDnsRewriteOk() (*InlineResponse200529SsidAddressingDnsRewrite, bool)`

GetDnsRewriteOk returns a tuple with the DnsRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRewrite

`func (o *InlineResponse200529SsidAddressing) SetDnsRewrite(v InlineResponse200529SsidAddressingDnsRewrite)`

SetDnsRewrite sets DnsRewrite field to given value.

### HasDnsRewrite

`func (o *InlineResponse200529SsidAddressing) HasDnsRewrite() bool`

HasDnsRewrite returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineResponse200529SsidAddressing) GetBonjourForwarding() InlineResponse200529SsidAddressingBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineResponse200529SsidAddressing) GetBonjourForwardingOk() (*InlineResponse200529SsidAddressingBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineResponse200529SsidAddressing) SetBonjourForwarding(v InlineResponse200529SsidAddressingBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineResponse200529SsidAddressing) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.

### GetMdns

`func (o *InlineResponse200529SsidAddressing) GetMdns() InlineResponse200529SsidAddressingMdns`

GetMdns returns the Mdns field if non-nil, zero value otherwise.

### GetMdnsOk

`func (o *InlineResponse200529SsidAddressing) GetMdnsOk() (*InlineResponse200529SsidAddressingMdns, bool)`

GetMdnsOk returns a tuple with the Mdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdns

`func (o *InlineResponse200529SsidAddressing) SetMdns(v InlineResponse200529SsidAddressingMdns)`

SetMdns sets Mdns field to given value.

### HasMdns

`func (o *InlineResponse200529SsidAddressing) HasMdns() bool`

HasMdns returns a boolean if a field has been set.

### GetRadius

`func (o *InlineResponse200529SsidAddressing) GetRadius() InlineResponse200529SsidAddressingRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineResponse200529SsidAddressing) GetRadiusOk() (*InlineResponse200529SsidAddressingRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineResponse200529SsidAddressing) SetRadius(v InlineResponse200529SsidAddressingRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineResponse200529SsidAddressing) HasRadius() bool`

HasRadius returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



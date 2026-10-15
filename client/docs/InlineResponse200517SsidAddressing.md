# InlineResponse200517SsidAddressing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Addressing mode | [optional] 
**AdultContentFiltering** | Pointer to [**InlineResponse200517SsidAddressingAdultContentFiltering**](InlineResponse200517SsidAddressingAdultContentFiltering.md) |  | [optional] 
**DnsRewrite** | Pointer to [**InlineResponse200517SsidAddressingDnsRewrite**](InlineResponse200517SsidAddressingDnsRewrite.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**InlineResponse200517SsidAddressingBonjourForwarding**](InlineResponse200517SsidAddressingBonjourForwarding.md) |  | [optional] 
**Mdns** | Pointer to [**InlineResponse200517SsidAddressingMdns**](InlineResponse200517SsidAddressingMdns.md) |  | [optional] 
**Radius** | Pointer to [**InlineResponse200517SsidAddressingRadius**](InlineResponse200517SsidAddressingRadius.md) |  | [optional] 

## Methods

### NewInlineResponse200517SsidAddressing

`func NewInlineResponse200517SsidAddressing() *InlineResponse200517SsidAddressing`

NewInlineResponse200517SsidAddressing instantiates a new InlineResponse200517SsidAddressing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200517SsidAddressingWithDefaults

`func NewInlineResponse200517SsidAddressingWithDefaults() *InlineResponse200517SsidAddressing`

NewInlineResponse200517SsidAddressingWithDefaults instantiates a new InlineResponse200517SsidAddressing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200517SsidAddressing) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200517SsidAddressing) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200517SsidAddressing) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200517SsidAddressing) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAdultContentFiltering

`func (o *InlineResponse200517SsidAddressing) GetAdultContentFiltering() InlineResponse200517SsidAddressingAdultContentFiltering`

GetAdultContentFiltering returns the AdultContentFiltering field if non-nil, zero value otherwise.

### GetAdultContentFilteringOk

`func (o *InlineResponse200517SsidAddressing) GetAdultContentFilteringOk() (*InlineResponse200517SsidAddressingAdultContentFiltering, bool)`

GetAdultContentFilteringOk returns a tuple with the AdultContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContentFiltering

`func (o *InlineResponse200517SsidAddressing) SetAdultContentFiltering(v InlineResponse200517SsidAddressingAdultContentFiltering)`

SetAdultContentFiltering sets AdultContentFiltering field to given value.

### HasAdultContentFiltering

`func (o *InlineResponse200517SsidAddressing) HasAdultContentFiltering() bool`

HasAdultContentFiltering returns a boolean if a field has been set.

### GetDnsRewrite

`func (o *InlineResponse200517SsidAddressing) GetDnsRewrite() InlineResponse200517SsidAddressingDnsRewrite`

GetDnsRewrite returns the DnsRewrite field if non-nil, zero value otherwise.

### GetDnsRewriteOk

`func (o *InlineResponse200517SsidAddressing) GetDnsRewriteOk() (*InlineResponse200517SsidAddressingDnsRewrite, bool)`

GetDnsRewriteOk returns a tuple with the DnsRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRewrite

`func (o *InlineResponse200517SsidAddressing) SetDnsRewrite(v InlineResponse200517SsidAddressingDnsRewrite)`

SetDnsRewrite sets DnsRewrite field to given value.

### HasDnsRewrite

`func (o *InlineResponse200517SsidAddressing) HasDnsRewrite() bool`

HasDnsRewrite returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineResponse200517SsidAddressing) GetBonjourForwarding() InlineResponse200517SsidAddressingBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineResponse200517SsidAddressing) GetBonjourForwardingOk() (*InlineResponse200517SsidAddressingBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineResponse200517SsidAddressing) SetBonjourForwarding(v InlineResponse200517SsidAddressingBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineResponse200517SsidAddressing) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.

### GetMdns

`func (o *InlineResponse200517SsidAddressing) GetMdns() InlineResponse200517SsidAddressingMdns`

GetMdns returns the Mdns field if non-nil, zero value otherwise.

### GetMdnsOk

`func (o *InlineResponse200517SsidAddressing) GetMdnsOk() (*InlineResponse200517SsidAddressingMdns, bool)`

GetMdnsOk returns a tuple with the Mdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdns

`func (o *InlineResponse200517SsidAddressing) SetMdns(v InlineResponse200517SsidAddressingMdns)`

SetMdns sets Mdns field to given value.

### HasMdns

`func (o *InlineResponse200517SsidAddressing) HasMdns() bool`

HasMdns returns a boolean if a field has been set.

### GetRadius

`func (o *InlineResponse200517SsidAddressing) GetRadius() InlineResponse200517SsidAddressingRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineResponse200517SsidAddressing) GetRadiusOk() (*InlineResponse200517SsidAddressingRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineResponse200517SsidAddressing) SetRadius(v InlineResponse200517SsidAddressingRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineResponse200517SsidAddressing) HasRadius() bool`

HasRadius returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



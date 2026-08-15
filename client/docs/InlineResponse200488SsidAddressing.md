# InlineResponse200488SsidAddressing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Addressing mode | [optional] 
**AdultContentFiltering** | Pointer to [**InlineResponse200488SsidAddressingAdultContentFiltering**](InlineResponse200488SsidAddressingAdultContentFiltering.md) |  | [optional] 
**DnsRewrite** | Pointer to [**InlineResponse200488SsidAddressingDnsRewrite**](InlineResponse200488SsidAddressingDnsRewrite.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**InlineResponse200488SsidAddressingBonjourForwarding**](InlineResponse200488SsidAddressingBonjourForwarding.md) |  | [optional] 
**Mdns** | Pointer to [**InlineResponse200488SsidAddressingMdns**](InlineResponse200488SsidAddressingMdns.md) |  | [optional] 
**Radius** | Pointer to [**InlineResponse200488SsidAddressingRadius**](InlineResponse200488SsidAddressingRadius.md) |  | [optional] 

## Methods

### NewInlineResponse200488SsidAddressing

`func NewInlineResponse200488SsidAddressing() *InlineResponse200488SsidAddressing`

NewInlineResponse200488SsidAddressing instantiates a new InlineResponse200488SsidAddressing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200488SsidAddressingWithDefaults

`func NewInlineResponse200488SsidAddressingWithDefaults() *InlineResponse200488SsidAddressing`

NewInlineResponse200488SsidAddressingWithDefaults instantiates a new InlineResponse200488SsidAddressing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200488SsidAddressing) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200488SsidAddressing) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200488SsidAddressing) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200488SsidAddressing) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAdultContentFiltering

`func (o *InlineResponse200488SsidAddressing) GetAdultContentFiltering() InlineResponse200488SsidAddressingAdultContentFiltering`

GetAdultContentFiltering returns the AdultContentFiltering field if non-nil, zero value otherwise.

### GetAdultContentFilteringOk

`func (o *InlineResponse200488SsidAddressing) GetAdultContentFilteringOk() (*InlineResponse200488SsidAddressingAdultContentFiltering, bool)`

GetAdultContentFilteringOk returns a tuple with the AdultContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContentFiltering

`func (o *InlineResponse200488SsidAddressing) SetAdultContentFiltering(v InlineResponse200488SsidAddressingAdultContentFiltering)`

SetAdultContentFiltering sets AdultContentFiltering field to given value.

### HasAdultContentFiltering

`func (o *InlineResponse200488SsidAddressing) HasAdultContentFiltering() bool`

HasAdultContentFiltering returns a boolean if a field has been set.

### GetDnsRewrite

`func (o *InlineResponse200488SsidAddressing) GetDnsRewrite() InlineResponse200488SsidAddressingDnsRewrite`

GetDnsRewrite returns the DnsRewrite field if non-nil, zero value otherwise.

### GetDnsRewriteOk

`func (o *InlineResponse200488SsidAddressing) GetDnsRewriteOk() (*InlineResponse200488SsidAddressingDnsRewrite, bool)`

GetDnsRewriteOk returns a tuple with the DnsRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRewrite

`func (o *InlineResponse200488SsidAddressing) SetDnsRewrite(v InlineResponse200488SsidAddressingDnsRewrite)`

SetDnsRewrite sets DnsRewrite field to given value.

### HasDnsRewrite

`func (o *InlineResponse200488SsidAddressing) HasDnsRewrite() bool`

HasDnsRewrite returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *InlineResponse200488SsidAddressing) GetBonjourForwarding() InlineResponse200488SsidAddressingBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *InlineResponse200488SsidAddressing) GetBonjourForwardingOk() (*InlineResponse200488SsidAddressingBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *InlineResponse200488SsidAddressing) SetBonjourForwarding(v InlineResponse200488SsidAddressingBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *InlineResponse200488SsidAddressing) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.

### GetMdns

`func (o *InlineResponse200488SsidAddressing) GetMdns() InlineResponse200488SsidAddressingMdns`

GetMdns returns the Mdns field if non-nil, zero value otherwise.

### GetMdnsOk

`func (o *InlineResponse200488SsidAddressing) GetMdnsOk() (*InlineResponse200488SsidAddressingMdns, bool)`

GetMdnsOk returns a tuple with the Mdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdns

`func (o *InlineResponse200488SsidAddressing) SetMdns(v InlineResponse200488SsidAddressingMdns)`

SetMdns sets Mdns field to given value.

### HasMdns

`func (o *InlineResponse200488SsidAddressing) HasMdns() bool`

HasMdns returns a boolean if a field has been set.

### GetRadius

`func (o *InlineResponse200488SsidAddressing) GetRadius() InlineResponse200488SsidAddressingRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InlineResponse200488SsidAddressing) GetRadiusOk() (*InlineResponse200488SsidAddressingRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InlineResponse200488SsidAddressing) SetRadius(v InlineResponse200488SsidAddressingRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *InlineResponse200488SsidAddressing) HasRadius() bool`

HasRadius returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



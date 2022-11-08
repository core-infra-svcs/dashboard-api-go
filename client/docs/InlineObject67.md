# InlineObject67

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpLeaseTime** | Pointer to **string** | DHCP Lease time for all MG of the network. Possible values are &#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;. | [optional] 
**DnsNameservers** | Pointer to **string** | DNS name servers mode for all MG of the network. Possible values are: &#39;upstream_dns&#39;, &#39;google_dns&#39;, &#39;opendns&#39;, &#39;custom&#39;. | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | list of fixed IPs representing the the DNS Name servers when the mode is &#39;custom&#39; | [optional] 

## Methods

### NewInlineObject67

`func NewInlineObject67() *InlineObject67`

NewInlineObject67 instantiates a new InlineObject67 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject67WithDefaults

`func NewInlineObject67WithDefaults() *InlineObject67`

NewInlineObject67WithDefaults instantiates a new InlineObject67 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpLeaseTime

`func (o *InlineObject67) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineObject67) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineObject67) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineObject67) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *InlineObject67) GetDnsNameservers() string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *InlineObject67) GetDnsNameserversOk() (*string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *InlineObject67) SetDnsNameservers(v string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *InlineObject67) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *InlineObject67) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *InlineObject67) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *InlineObject67) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *InlineObject67) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse20082

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpLeaseTime** | Pointer to **string** | DHCP Lease time for all MG in the network. | [optional] 
**DnsNameservers** | Pointer to **string** | DNS name servers mode for all MG in the network. | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | List of fixed IPs representing the the DNS Name servers when the mode is &#39;custom&#39;. | [optional] 

## Methods

### NewInlineResponse20082

`func NewInlineResponse20082() *InlineResponse20082`

NewInlineResponse20082 instantiates a new InlineResponse20082 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20082WithDefaults

`func NewInlineResponse20082WithDefaults() *InlineResponse20082`

NewInlineResponse20082WithDefaults instantiates a new InlineResponse20082 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpLeaseTime

`func (o *InlineResponse20082) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *InlineResponse20082) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *InlineResponse20082) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *InlineResponse20082) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *InlineResponse20082) GetDnsNameservers() string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *InlineResponse20082) GetDnsNameserversOk() (*string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *InlineResponse20082) SetDnsNameservers(v string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *InlineResponse20082) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *InlineResponse20082) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *InlineResponse20082) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *InlineResponse20082) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *InlineResponse20082) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse200111

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network identifier | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**ByUplink** | Pointer to [**[]OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink**](OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink.md) | Uplink usage | [optional] 

## Methods

### NewInlineResponse200111

`func NewInlineResponse200111() *InlineResponse200111`

NewInlineResponse200111 instantiates a new InlineResponse200111 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111WithDefaults

`func NewInlineResponse200111WithDefaults() *InlineResponse200111`

NewInlineResponse200111WithDefaults instantiates a new InlineResponse200111 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200111) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200111) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200111) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200111) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200111) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200111) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200111) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200111) HasName() bool`

HasName returns a boolean if a field has been set.

### GetByUplink

`func (o *InlineResponse200111) GetByUplink() []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink`

GetByUplink returns the ByUplink field if non-nil, zero value otherwise.

### GetByUplinkOk

`func (o *InlineResponse200111) GetByUplinkOk() (*[]OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink, bool)`

GetByUplinkOk returns a tuple with the ByUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByUplink

`func (o *InlineResponse200111) SetByUplink(v []OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink)`

SetByUplink sets ByUplink field to given value.

### HasByUplink

`func (o *InlineResponse200111) HasByUplink() bool`

HasByUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



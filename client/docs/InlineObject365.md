# InlineObject365

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the SSID profile | 
**Precedence** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesPrecedence**](OrganizationsOrganizationIdWirelessSsidsProfilesPrecedence.md) |  | [optional] 
**Ssid** | [**OrganizationsOrganizationIdWirelessSsidsProfilesSsid**](OrganizationsOrganizationIdWirelessSsidsProfilesSsid.md) |  | 

## Methods

### NewInlineObject365

`func NewInlineObject365(name string, ssid OrganizationsOrganizationIdWirelessSsidsProfilesSsid, ) *InlineObject365`

NewInlineObject365 instantiates a new InlineObject365 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject365WithDefaults

`func NewInlineObject365WithDefaults() *InlineObject365`

NewInlineObject365WithDefaults instantiates a new InlineObject365 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject365) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject365) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject365) SetName(v string)`

SetName sets Name field to given value.


### GetPrecedence

`func (o *InlineObject365) GetPrecedence() OrganizationsOrganizationIdWirelessSsidsProfilesPrecedence`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *InlineObject365) GetPrecedenceOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesPrecedence, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *InlineObject365) SetPrecedence(v OrganizationsOrganizationIdWirelessSsidsProfilesPrecedence)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *InlineObject365) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetSsid

`func (o *InlineObject365) GetSsid() OrganizationsOrganizationIdWirelessSsidsProfilesSsid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineObject365) GetSsidOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineObject365) SetSsid(v OrganizationsOrganizationIdWirelessSsidsProfilesSsid)`

SetSsid sets Ssid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineResponse20066

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The device serial | [optional] 
**Network** | Pointer to [**InlineResponse20066Network**](InlineResponse20066Network.md) |  | [optional] 
**RfProfile** | Pointer to [**InlineResponse20066RfProfile**](InlineResponse20066RfProfile.md) |  | [optional] 
**Radios** | Pointer to [**[]InlineResponse20066Radios**](InlineResponse20066Radios.md) | Individual radios on the device | [optional] 

## Methods

### NewInlineResponse20066

`func NewInlineResponse20066() *InlineResponse20066`

NewInlineResponse20066 instantiates a new InlineResponse20066 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20066WithDefaults

`func NewInlineResponse20066WithDefaults() *InlineResponse20066`

NewInlineResponse20066WithDefaults instantiates a new InlineResponse20066 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20066) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20066) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20066) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20066) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20066) GetNetwork() InlineResponse20066Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20066) GetNetworkOk() (*InlineResponse20066Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20066) SetNetwork(v InlineResponse20066Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20066) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRfProfile

`func (o *InlineResponse20066) GetRfProfile() InlineResponse20066RfProfile`

GetRfProfile returns the RfProfile field if non-nil, zero value otherwise.

### GetRfProfileOk

`func (o *InlineResponse20066) GetRfProfileOk() (*InlineResponse20066RfProfile, bool)`

GetRfProfileOk returns a tuple with the RfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfProfile

`func (o *InlineResponse20066) SetRfProfile(v InlineResponse20066RfProfile)`

SetRfProfile sets RfProfile field to given value.

### HasRfProfile

`func (o *InlineResponse20066) HasRfProfile() bool`

HasRfProfile returns a boolean if a field has been set.

### GetRadios

`func (o *InlineResponse20066) GetRadios() []InlineResponse20066Radios`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *InlineResponse20066) GetRadiosOk() (*[]InlineResponse20066Radios, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *InlineResponse20066) SetRadios(v []InlineResponse20066Radios)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *InlineResponse20066) HasRadios() bool`

HasRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



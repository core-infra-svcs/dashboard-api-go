# InlineResponse20064

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The device serial | [optional] 
**Network** | Pointer to [**InlineResponse20064Network**](InlineResponse20064Network.md) |  | [optional] 
**RfProfile** | Pointer to [**InlineResponse20064RfProfile**](InlineResponse20064RfProfile.md) |  | [optional] 
**Radios** | Pointer to [**[]InlineResponse20064Radios**](InlineResponse20064Radios.md) | Individual radios on the device | [optional] 

## Methods

### NewInlineResponse20064

`func NewInlineResponse20064() *InlineResponse20064`

NewInlineResponse20064 instantiates a new InlineResponse20064 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20064WithDefaults

`func NewInlineResponse20064WithDefaults() *InlineResponse20064`

NewInlineResponse20064WithDefaults instantiates a new InlineResponse20064 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse20064) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20064) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20064) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20064) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20064) GetNetwork() InlineResponse20064Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20064) GetNetworkOk() (*InlineResponse20064Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20064) SetNetwork(v InlineResponse20064Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20064) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRfProfile

`func (o *InlineResponse20064) GetRfProfile() InlineResponse20064RfProfile`

GetRfProfile returns the RfProfile field if non-nil, zero value otherwise.

### GetRfProfileOk

`func (o *InlineResponse20064) GetRfProfileOk() (*InlineResponse20064RfProfile, bool)`

GetRfProfileOk returns a tuple with the RfProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfProfile

`func (o *InlineResponse20064) SetRfProfile(v InlineResponse20064RfProfile)`

SetRfProfile sets RfProfile field to given value.

### HasRfProfile

`func (o *InlineResponse20064) HasRfProfile() bool`

HasRfProfile returns a boolean if a field has been set.

### GetRadios

`func (o *InlineResponse20064) GetRadios() []InlineResponse20064Radios`

GetRadios returns the Radios field if non-nil, zero value otherwise.

### GetRadiosOk

`func (o *InlineResponse20064) GetRadiosOk() (*[]InlineResponse20064Radios, bool)`

GetRadiosOk returns a tuple with the Radios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadios

`func (o *InlineResponse20064) SetRadios(v []InlineResponse20064Radios)`

SetRadios sets Radios field to given value.

### HasRadios

`func (o *InlineResponse20064) HasRadios() bool`

HasRadios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



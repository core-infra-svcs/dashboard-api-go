# InlineResponse200324

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the client | [optional] 
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the client | [optional] 
**Records** | Pointer to [**[]InlineResponse200324Records**](InlineResponse200324Records.md) | The clients that appear on any networks within an organization | [optional] 

## Methods

### NewInlineResponse200324

`func NewInlineResponse200324() *InlineResponse200324`

NewInlineResponse200324 instantiates a new InlineResponse200324 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200324WithDefaults

`func NewInlineResponse200324WithDefaults() *InlineResponse200324`

NewInlineResponse200324WithDefaults instantiates a new InlineResponse200324 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InlineResponse200324) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineResponse200324) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineResponse200324) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineResponse200324) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200324) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200324) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200324) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200324) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineResponse200324) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineResponse200324) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineResponse200324) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineResponse200324) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetRecords

`func (o *InlineResponse200324) GetRecords() []InlineResponse200324Records`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *InlineResponse200324) GetRecordsOk() (*[]InlineResponse200324Records, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *InlineResponse200324) SetRecords(v []InlineResponse200324Records)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *InlineResponse200324) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



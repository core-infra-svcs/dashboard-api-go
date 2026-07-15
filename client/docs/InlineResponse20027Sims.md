# InlineResponse20027Sims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. eSIM slots use the raw slot value for the device, such as &#39;sim2&#39; or &#39;sim3&#39;. | [optional] 
**Iccid** | Pointer to **string** | Integrated Circuit Card Identification Number | [optional] 
**Imsi** | Pointer to **string** | International Mobile Subscriber Identity | [optional] 
**Msisdn** | Pointer to **string** | Mobile Station Integrated Services Digital Network | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using &#39;simOrdering&#39;. | [optional] [default to false]
**Status** | Pointer to **string** | Status of the SIM card. | [optional] 
**Apns** | Pointer to [**[]InlineResponse20027Apns**](InlineResponse20027Apns.md) | APN configurations. If empty, the default APN will be used. | [optional] 

## Methods

### NewInlineResponse20027Sims

`func NewInlineResponse20027Sims() *InlineResponse20027Sims`

NewInlineResponse20027Sims instantiates a new InlineResponse20027Sims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027SimsWithDefaults

`func NewInlineResponse20027SimsWithDefaults() *InlineResponse20027Sims`

NewInlineResponse20027SimsWithDefaults instantiates a new InlineResponse20027Sims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse20027Sims) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse20027Sims) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse20027Sims) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse20027Sims) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIccid

`func (o *InlineResponse20027Sims) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *InlineResponse20027Sims) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *InlineResponse20027Sims) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *InlineResponse20027Sims) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetImsi

`func (o *InlineResponse20027Sims) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *InlineResponse20027Sims) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *InlineResponse20027Sims) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *InlineResponse20027Sims) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetMsisdn

`func (o *InlineResponse20027Sims) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *InlineResponse20027Sims) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *InlineResponse20027Sims) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *InlineResponse20027Sims) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetIsPrimary

`func (o *InlineResponse20027Sims) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *InlineResponse20027Sims) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *InlineResponse20027Sims) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *InlineResponse20027Sims) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20027Sims) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20027Sims) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20027Sims) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20027Sims) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApns

`func (o *InlineResponse20027Sims) GetApns() []InlineResponse20027Apns`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *InlineResponse20027Sims) GetApnsOk() (*[]InlineResponse20027Apns, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *InlineResponse20027Sims) SetApns(v []InlineResponse20027Apns)`

SetApns sets Apns field to given value.

### HasApns

`func (o *InlineResponse20027Sims) HasApns() bool`

HasApns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



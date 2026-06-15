# InlineResponse200330Sims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | Identifier for the logical SIM slot. Physical slots are reported as sim1 or sim2. Embedded SIM is reported as esim. Devices will only report SIM slots that are supported by the device. For devices with an eSIM, the esim entry will indicate the status of the embedded SIM if it has an active profile on it. | [optional] 
**Type** | Pointer to **string** | Type of SIM (sim or esim) | [optional] 
**Active** | Pointer to **bool** | Whether a SIM is installed in the slot or the eSIM has an active Cellular Data Management Profile. | [optional] 

## Methods

### NewInlineResponse200330Sims

`func NewInlineResponse200330Sims() *InlineResponse200330Sims`

NewInlineResponse200330Sims instantiates a new InlineResponse200330Sims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200330SimsWithDefaults

`func NewInlineResponse200330SimsWithDefaults() *InlineResponse200330Sims`

NewInlineResponse200330SimsWithDefaults instantiates a new InlineResponse200330Sims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse200330Sims) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200330Sims) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200330Sims) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200330Sims) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200330Sims) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200330Sims) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200330Sims) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200330Sims) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200330Sims) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200330Sims) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200330Sims) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200330Sims) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



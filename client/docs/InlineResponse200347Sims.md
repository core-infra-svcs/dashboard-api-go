# InlineResponse200347Sims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | Identifier for the logical SIM slot. Physical slots are reported as sim1 or sim2. Embedded SIM is reported as esim. Devices will only report SIM slots that are supported by the device. For devices with an eSIM, the esim entry will indicate the status of the embedded SIM if it has an active profile on it. | [optional] 
**Type** | Pointer to **string** | Type of SIM (sim or esim) | [optional] 
**Active** | Pointer to **bool** | Whether a SIM is installed in the slot or the eSIM has an active Cellular Data Management Profile. | [optional] 

## Methods

### NewInlineResponse200347Sims

`func NewInlineResponse200347Sims() *InlineResponse200347Sims`

NewInlineResponse200347Sims instantiates a new InlineResponse200347Sims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200347SimsWithDefaults

`func NewInlineResponse200347SimsWithDefaults() *InlineResponse200347Sims`

NewInlineResponse200347SimsWithDefaults instantiates a new InlineResponse200347Sims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse200347Sims) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200347Sims) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200347Sims) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200347Sims) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200347Sims) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200347Sims) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200347Sims) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200347Sims) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200347Sims) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200347Sims) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200347Sims) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200347Sims) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



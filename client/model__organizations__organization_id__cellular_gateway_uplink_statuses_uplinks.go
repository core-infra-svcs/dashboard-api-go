/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks struct for OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks
type OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks struct {
	// Uplink interface
	Interface *string `json:"interface,omitempty"`
	// Uplink status
	Status *string `json:"status,omitempty"`
	// Uplink IP
	Ip *string `json:"ip,omitempty"`
	// Network Provider
	Provider *string `json:"provider,omitempty"`
	// Public IP
	PublicIp *string `json:"publicIp,omitempty"`
	// Uplink model
	Model *string `json:"model,omitempty"`
	SignalStat *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat `json:"signalStat,omitempty"`
	// Mobile Country Code
	Mcc *string `json:"mcc,omitempty"`
	// Mobile Network Code
	Mnc *string `json:"mnc,omitempty"`
	Roaming *OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming `json:"roaming,omitempty"`
	// Connection Type
	ConnectionType *string `json:"connectionType,omitempty"`
	// Access Point Name
	Apn *string `json:"apn,omitempty"`
	// Gateway IP
	Gateway *string `json:"gateway,omitempty"`
	// Primary DNS IP
	Dns1 *string `json:"dns1,omitempty"`
	// Secondary DNS IP
	Dns2 *string `json:"dns2,omitempty"`
	// Signal Type
	SignalType *string `json:"signalType,omitempty"`
	// Maximum Transmission Unit
	Mtu *int32 `json:"mtu,omitempty"`
	// Integrated Circuit Card Identification Number
	Iccid *string `json:"iccid,omitempty"`
	// International Mobile Subscriber Identity
	Imsi *string `json:"imsi,omitempty"`
	// Mobile Station Integrated Services Digital Network
	Msisdn *string `json:"msisdn,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks instantiates a new OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks {
	this := OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinksWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinksWithDefaults() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks {
	this := OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetInterface(v string) {
	o.Interface = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetStatus(v string) {
	o.Status = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetIp(v string) {
	o.Ip = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetProvider(v string) {
	o.Provider = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetModel(v string) {
	o.Model = &v
}

// GetSignalStat returns the SignalStat field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetSignalStat() OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat {
	if o == nil || isNil(o.SignalStat) {
		var ret OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat
		return ret
	}
	return *o.SignalStat
}

// GetSignalStatOk returns a tuple with the SignalStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetSignalStatOk() (*OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat, bool) {
	if o == nil || isNil(o.SignalStat) {
    return nil, false
	}
	return o.SignalStat, true
}

// HasSignalStat returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasSignalStat() bool {
	if o != nil && !isNil(o.SignalStat) {
		return true
	}

	return false
}

// SetSignalStat gets a reference to the given OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat and assigns it to the SignalStat field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetSignalStat(v OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) {
	o.SignalStat = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMcc() string {
	if o == nil || isNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMccOk() (*string, bool) {
	if o == nil || isNil(o.Mcc) {
    return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasMcc() bool {
	if o != nil && !isNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMnc() string {
	if o == nil || isNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMncOk() (*string, bool) {
	if o == nil || isNil(o.Mnc) {
    return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasMnc() bool {
	if o != nil && !isNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetMnc(v string) {
	o.Mnc = &v
}

// GetRoaming returns the Roaming field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetRoaming() OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming {
	if o == nil || isNil(o.Roaming) {
		var ret OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming
		return ret
	}
	return *o.Roaming
}

// GetRoamingOk returns a tuple with the Roaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetRoamingOk() (*OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming, bool) {
	if o == nil || isNil(o.Roaming) {
    return nil, false
	}
	return o.Roaming, true
}

// HasRoaming returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasRoaming() bool {
	if o != nil && !isNil(o.Roaming) {
		return true
	}

	return false
}

// SetRoaming gets a reference to the given OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming and assigns it to the Roaming field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetRoaming(v OrganizationsOrganizationIdCellularGatewayUplinkStatusesRoaming) {
	o.Roaming = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetConnectionType() string {
	if o == nil || isNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetConnectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionType) {
    return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasConnectionType() bool {
	if o != nil && !isNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetApn returns the Apn field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetApn() string {
	if o == nil || isNil(o.Apn) {
		var ret string
		return ret
	}
	return *o.Apn
}

// GetApnOk returns a tuple with the Apn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetApnOk() (*string, bool) {
	if o == nil || isNil(o.Apn) {
    return nil, false
	}
	return o.Apn, true
}

// HasApn returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasApn() bool {
	if o != nil && !isNil(o.Apn) {
		return true
	}

	return false
}

// SetApn gets a reference to the given string and assigns it to the Apn field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetApn(v string) {
	o.Apn = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetGateway(v string) {
	o.Gateway = &v
}

// GetDns1 returns the Dns1 field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetDns1() string {
	if o == nil || isNil(o.Dns1) {
		var ret string
		return ret
	}
	return *o.Dns1
}

// GetDns1Ok returns a tuple with the Dns1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetDns1Ok() (*string, bool) {
	if o == nil || isNil(o.Dns1) {
    return nil, false
	}
	return o.Dns1, true
}

// HasDns1 returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasDns1() bool {
	if o != nil && !isNil(o.Dns1) {
		return true
	}

	return false
}

// SetDns1 gets a reference to the given string and assigns it to the Dns1 field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetDns1(v string) {
	o.Dns1 = &v
}

// GetDns2 returns the Dns2 field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetDns2() string {
	if o == nil || isNil(o.Dns2) {
		var ret string
		return ret
	}
	return *o.Dns2
}

// GetDns2Ok returns a tuple with the Dns2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetDns2Ok() (*string, bool) {
	if o == nil || isNil(o.Dns2) {
    return nil, false
	}
	return o.Dns2, true
}

// HasDns2 returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasDns2() bool {
	if o != nil && !isNil(o.Dns2) {
		return true
	}

	return false
}

// SetDns2 gets a reference to the given string and assigns it to the Dns2 field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetDns2(v string) {
	o.Dns2 = &v
}

// GetSignalType returns the SignalType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetSignalType() string {
	if o == nil || isNil(o.SignalType) {
		var ret string
		return ret
	}
	return *o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetSignalTypeOk() (*string, bool) {
	if o == nil || isNil(o.SignalType) {
    return nil, false
	}
	return o.SignalType, true
}

// HasSignalType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasSignalType() bool {
	if o != nil && !isNil(o.SignalType) {
		return true
	}

	return false
}

// SetSignalType gets a reference to the given string and assigns it to the SignalType field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetSignalType(v string) {
	o.SignalType = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMtu() int32 {
	if o == nil || isNil(o.Mtu) {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMtuOk() (*int32, bool) {
	if o == nil || isNil(o.Mtu) {
    return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasMtu() bool {
	if o != nil && !isNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetMtu(v int32) {
	o.Mtu = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetIccid() string {
	if o == nil || isNil(o.Iccid) {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetIccidOk() (*string, bool) {
	if o == nil || isNil(o.Iccid) {
    return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasIccid() bool {
	if o != nil && !isNil(o.Iccid) {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetIccid(v string) {
	o.Iccid = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetImsi() string {
	if o == nil || isNil(o.Imsi) {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetImsiOk() (*string, bool) {
	if o == nil || isNil(o.Imsi) {
    return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasImsi() bool {
	if o != nil && !isNil(o.Imsi) {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetImsi(v string) {
	o.Imsi = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMsisdn() string {
	if o == nil || isNil(o.Msisdn) {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) GetMsisdnOk() (*string, bool) {
	if o == nil || isNil(o.Msisdn) {
    return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) HasMsisdn() bool {
	if o != nil && !isNil(o.Msisdn) {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) SetMsisdn(v string) {
	o.Msisdn = &v
}

func (o OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.SignalStat) {
		toSerialize["signalStat"] = o.SignalStat
	}
	if !isNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !isNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !isNil(o.Roaming) {
		toSerialize["roaming"] = o.Roaming
	}
	if !isNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if !isNil(o.Apn) {
		toSerialize["apn"] = o.Apn
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.Dns1) {
		toSerialize["dns1"] = o.Dns1
	}
	if !isNil(o.Dns2) {
		toSerialize["dns2"] = o.Dns2
	}
	if !isNil(o.SignalType) {
		toSerialize["signalType"] = o.SignalType
	}
	if !isNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !isNil(o.Iccid) {
		toSerialize["iccid"] = o.Iccid
	}
	if !isNil(o.Imsi) {
		toSerialize["imsi"] = o.Imsi
	}
	if !isNil(o.Msisdn) {
		toSerialize["msisdn"] = o.Msisdn
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks struct {
	value *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) Get() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) Set(val *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks(val *OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks {
	return &NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



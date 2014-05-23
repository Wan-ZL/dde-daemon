// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingIp4ConfigKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_IP4_CONFIG_METHOD:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		t = ktypeWrapperIpv4Addresses
	case NM_SETTING_IP4_CONFIG_DNS:
		t = ktypeWrapperIpv4Dns
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_ROUTES:
		t = ktypeWrapperIpv4Routes
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		t = ktypeString
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		t = ktypeBoolean
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		t = ktypeBoolean
	}
	return
}

// Check is key in current setting field
func isKeyInSettingIp4Config(key string) bool {
	switch key {
	case NM_SETTING_IP4_CONFIG_METHOD:
		return true
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		return true
	case NM_SETTING_IP4_CONFIG_DNS:
		return true
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		return true
	case NM_SETTING_IP4_CONFIG_ROUTES:
		return true
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		return true
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		return true
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		return true
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		return true
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		return true
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		return true
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		return true
	}
	return false
}

// Get key's default value
func getSettingIp4ConfigDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_IP4_CONFIG_METHOD:
		value = ""
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		value = nil
	case NM_SETTING_IP4_CONFIG_DNS:
		value = nil
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		value = ""
	case NM_SETTING_IP4_CONFIG_ROUTES:
		value = nil
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		value = false
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		value = false
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		value = ""
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		value = true
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		value = ""
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		value = false
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		value = true
	}
	return
}

// Get JSON value generally
func generalGetSettingIp4ConfigKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingIp4ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP4_CONFIG_METHOD:
		value = getSettingIp4ConfigMethodJSON(data)
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		value = getSettingIp4ConfigAddressesJSON(data)
	case NM_SETTING_IP4_CONFIG_DNS:
		value = getSettingIp4ConfigDnsJSON(data)
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		value = getSettingIp4ConfigDnsSearchJSON(data)
	case NM_SETTING_IP4_CONFIG_ROUTES:
		value = getSettingIp4ConfigRoutesJSON(data)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		value = getSettingIp4ConfigIgnoreAutoRoutesJSON(data)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		value = getSettingIp4ConfigIgnoreAutoDnsJSON(data)
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		value = getSettingIp4ConfigDhcpClientIdJSON(data)
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		value = getSettingIp4ConfigDhcpSendHostnameJSON(data)
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		value = getSettingIp4ConfigDhcpHostnameJSON(data)
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		value = getSettingIp4ConfigNeverDefaultJSON(data)
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		value = getSettingIp4ConfigMayFailJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingIp4ConfigKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingIp4ConfigKeyJSON: invalide key", key)
	case NM_SETTING_IP4_CONFIG_METHOD:
		err = logicSetSettingIp4ConfigMethodJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_ADDRESSES:
		err = setSettingIp4ConfigAddressesJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DNS:
		err = setSettingIp4ConfigDnsJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DNS_SEARCH:
		err = setSettingIp4ConfigDnsSearchJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_ROUTES:
		err = setSettingIp4ConfigRoutesJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES:
		err = setSettingIp4ConfigIgnoreAutoRoutesJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS:
		err = setSettingIp4ConfigIgnoreAutoDnsJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID:
		err = setSettingIp4ConfigDhcpClientIdJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME:
		err = setSettingIp4ConfigDhcpSendHostnameJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME:
		err = setSettingIp4ConfigDhcpHostnameJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_NEVER_DEFAULT:
		err = setSettingIp4ConfigNeverDefaultJSON(data, valueJSON)
	case NM_SETTING_IP4_CONFIG_MAY_FAIL:
		err = setSettingIp4ConfigMayFailJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingIp4ConfigMethodExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD)
}
func isSettingIp4ConfigAddressesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES)
}
func isSettingIp4ConfigDnsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS)
}
func isSettingIp4ConfigDnsSearchExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH)
}
func isSettingIp4ConfigRoutesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES)
}
func isSettingIp4ConfigIgnoreAutoRoutesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES)
}
func isSettingIp4ConfigIgnoreAutoDnsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS)
}
func isSettingIp4ConfigDhcpClientIdExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID)
}
func isSettingIp4ConfigDhcpSendHostnameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME)
}
func isSettingIp4ConfigDhcpHostnameExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME)
}
func isSettingIp4ConfigNeverDefaultExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT)
}
func isSettingIp4ConfigMayFailExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL)
}

// Ensure field and key exists and not empty
func ensureFieldSettingIp4ConfigExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_IP4_CONFIG_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_IP4_CONFIG_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_IP4_CONFIG_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_IP4_CONFIG_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_IP4_CONFIG_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_IP4_CONFIG_SETTING_NAME))
	}
}
func ensureSettingIp4ConfigMethodNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigMethodExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigMethod(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigAddressesNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigAddressesExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigAddresses(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigDnsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigDnsExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDns(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigDnsSearchNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigDnsSearchExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDnsSearch(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigRoutesNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigRoutes(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigIgnoreAutoRoutesNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigIgnoreAutoRoutesExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigIgnoreAutoDnsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigIgnoreAutoDnsExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigDhcpClientIdNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigDhcpClientIdExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDhcpClientId(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigDhcpSendHostnameNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigDhcpSendHostnameExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigDhcpHostnameNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigDhcpHostnameExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingIp4ConfigDhcpHostname(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingIp4ConfigNeverDefaultNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigNeverDefaultExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingIp4ConfigMayFailNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingIp4ConfigMayFailExists(data) {
		rememberError(errs, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingIp4ConfigMethod(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingIp4ConfigMethod: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigAddresses(data connectionData) (value [][]uint32) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES)
	value, ok := ivalue.([][]uint32)
	if !ok {
		logger.Warningf("getSettingIp4ConfigAddresses: value type is invalid, should be [][]uint32, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigDns(data connectionData) (value []uint32) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS)
	value, ok := ivalue.([]uint32)
	if !ok {
		logger.Warningf("getSettingIp4ConfigDns: value type is invalid, should be []uint32, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigDnsSearch(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingIp4ConfigDnsSearch: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigRoutes(data connectionData) (value [][]uint32) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES)
	value, ok := ivalue.([][]uint32)
	if !ok {
		logger.Warningf("getSettingIp4ConfigRoutes: value type is invalid, should be [][]uint32, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigIgnoreAutoRoutes(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Warningf("getSettingIp4ConfigIgnoreAutoRoutes: value type is invalid, should be bool, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigIgnoreAutoDns(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Warningf("getSettingIp4ConfigIgnoreAutoDns: value type is invalid, should be bool, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigDhcpClientId(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingIp4ConfigDhcpClientId: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigDhcpSendHostname(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Warningf("getSettingIp4ConfigDhcpSendHostname: value type is invalid, should be bool, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigDhcpHostname(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingIp4ConfigDhcpHostname: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigNeverDefault(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Warningf("getSettingIp4ConfigNeverDefault: value type is invalid, should be bool, %v", ivalue)
	}
	return
}
func getSettingIp4ConfigMayFail(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Warningf("getSettingIp4ConfigMayFail: value type is invalid, should be bool, %v", ivalue)
	}
	return
}

// Setter
func setSettingIp4ConfigMethod(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, value)
}
func setSettingIp4ConfigAddresses(data connectionData, value [][]uint32) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, value)
}
func setSettingIp4ConfigDns(data connectionData, value []uint32) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, value)
}
func setSettingIp4ConfigDnsSearch(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, value)
}
func setSettingIp4ConfigRoutes(data connectionData, value [][]uint32) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, value)
}
func setSettingIp4ConfigIgnoreAutoRoutes(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, value)
}
func setSettingIp4ConfigIgnoreAutoDns(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, value)
}
func setSettingIp4ConfigDhcpClientId(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, value)
}
func setSettingIp4ConfigDhcpSendHostname(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, value)
}
func setSettingIp4ConfigDhcpHostname(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, value)
}
func setSettingIp4ConfigNeverDefault(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, value)
}
func setSettingIp4ConfigMayFail(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, value)
}

// JSON Getter
func getSettingIp4ConfigMethodJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_METHOD))
	return
}
func getSettingIp4ConfigAddressesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ADDRESSES))
	return
}
func getSettingIp4ConfigDnsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS))
	return
}
func getSettingIp4ConfigDnsSearchJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS_SEARCH))
	return
}
func getSettingIp4ConfigRoutesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ROUTES))
	return
}
func getSettingIp4ConfigIgnoreAutoRoutesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES))
	return
}
func getSettingIp4ConfigIgnoreAutoDnsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS))
	return
}
func getSettingIp4ConfigDhcpClientIdJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID))
	return
}
func getSettingIp4ConfigDhcpSendHostnameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME))
	return
}
func getSettingIp4ConfigDhcpHostnameJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME))
	return
}
func getSettingIp4ConfigNeverDefaultJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_NEVER_DEFAULT))
	return
}
func getSettingIp4ConfigMayFailJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_MAY_FAIL))
	return
}

// JSON Setter
func setSettingIp4ConfigMethodJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_METHOD))
}
func setSettingIp4ConfigAddressesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ADDRESSES))
}
func setSettingIp4ConfigDnsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS))
}
func setSettingIp4ConfigDnsSearchJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DNS_SEARCH))
}
func setSettingIp4ConfigRoutesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_ROUTES))
}
func setSettingIp4ConfigIgnoreAutoRoutesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES))
}
func setSettingIp4ConfigIgnoreAutoDnsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS))
}
func setSettingIp4ConfigDhcpClientIdJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID))
}
func setSettingIp4ConfigDhcpSendHostnameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME))
}
func setSettingIp4ConfigDhcpHostnameJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME))
}
func setSettingIp4ConfigNeverDefaultJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_NEVER_DEFAULT))
}
func setSettingIp4ConfigMayFailJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL, valueJSON, getSettingIp4ConfigKeyType(NM_SETTING_IP4_CONFIG_MAY_FAIL))
}

// Logic JSON Setter
func logicSetSettingIp4ConfigMethodJSON(data connectionData, valueJSON string) (err error) {
	err = setSettingIp4ConfigMethodJSON(data, valueJSON)
	if err != nil {
		return
	}
	if isSettingIp4ConfigMethodExists(data) {
		value := getSettingIp4ConfigMethod(data)
		err = logicSetSettingIp4ConfigMethod(data, value)
	}
	return
}

// Remover
func removeSettingIp4ConfigMethod(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_METHOD)
}
func removeSettingIp4ConfigAddresses(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES)
}
func removeSettingIp4ConfigDns(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS)
}
func removeSettingIp4ConfigDnsSearch(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS_SEARCH)
}
func removeSettingIp4ConfigRoutes(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES)
}
func removeSettingIp4ConfigIgnoreAutoRoutes(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_ROUTES)
}
func removeSettingIp4ConfigIgnoreAutoDns(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_IGNORE_AUTO_DNS)
}
func removeSettingIp4ConfigDhcpClientId(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_CLIENT_ID)
}
func removeSettingIp4ConfigDhcpSendHostname(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_SEND_HOSTNAME)
}
func removeSettingIp4ConfigDhcpHostname(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DHCP_HOSTNAME)
}
func removeSettingIp4ConfigNeverDefault(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_NEVER_DEFAULT)
}
func removeSettingIp4ConfigMayFail(data connectionData) {
	removeSettingKey(data, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_MAY_FAIL)
}

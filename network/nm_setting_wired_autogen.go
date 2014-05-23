// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingWiredKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRED_PORT:
		t = ktypeString
	case NM_SETTING_WIRED_SPEED:
		t = ktypeUint32
	case NM_SETTING_WIRED_DUPLEX:
		t = ktypeString
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		t = ktypeBoolean
	case NM_SETTING_WIRED_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		t = ktypeArrayString
	case NM_SETTING_WIRED_MTU:
		t = ktypeUint32
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		t = ktypeArrayString
	case NM_SETTING_WIRED_S390_NETTYPE:
		t = ktypeString
	case NM_SETTING_WIRED_S390_OPTIONS:
		t = ktypeDictStringString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingWired(key string) bool {
	switch key {
	case NM_SETTING_WIRED_PORT:
		return true
	case NM_SETTING_WIRED_SPEED:
		return true
	case NM_SETTING_WIRED_DUPLEX:
		return true
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		return true
	case NM_SETTING_WIRED_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		return true
	case NM_SETTING_WIRED_MTU:
		return true
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		return true
	case NM_SETTING_WIRED_S390_NETTYPE:
		return true
	case NM_SETTING_WIRED_S390_OPTIONS:
		return true
	}
	return false
}

// Get key's default value
func getSettingWiredDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_WIRED_PORT:
		value = ""
	case NM_SETTING_WIRED_SPEED:
		value = 0
	case NM_SETTING_WIRED_DUPLEX:
		value = ""
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		value = true
	case NM_SETTING_WIRED_MAC_ADDRESS:
		value = ""
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		value = ""
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		value = nil
	case NM_SETTING_WIRED_MTU:
		value = 0
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		value = nil
	case NM_SETTING_WIRED_S390_NETTYPE:
		value = ""
	case NM_SETTING_WIRED_S390_OPTIONS:
		value = nil
	}
	return
}

// Get JSON value generally
func generalGetSettingWiredKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingWiredKeyJSON: invalide key", key)
	case NM_SETTING_WIRED_PORT:
		value = getSettingWiredPortJSON(data)
	case NM_SETTING_WIRED_SPEED:
		value = getSettingWiredSpeedJSON(data)
	case NM_SETTING_WIRED_DUPLEX:
		value = getSettingWiredDuplexJSON(data)
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		value = getSettingWiredAutoNegotiateJSON(data)
	case NM_SETTING_WIRED_MAC_ADDRESS:
		value = getSettingWiredMacAddressJSON(data)
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		value = getSettingWiredClonedMacAddressJSON(data)
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		value = getSettingWiredMacAddressBlacklistJSON(data)
	case NM_SETTING_WIRED_MTU:
		value = getSettingWiredMtuJSON(data)
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		value = getSettingWiredS390SubchannelsJSON(data)
	case NM_SETTING_WIRED_S390_NETTYPE:
		value = getSettingWiredS390NettypeJSON(data)
	case NM_SETTING_WIRED_S390_OPTIONS:
		value = getSettingWiredS390OptionsJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingWiredKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingWiredKeyJSON: invalide key", key)
	case NM_SETTING_WIRED_PORT:
		err = setSettingWiredPortJSON(data, valueJSON)
	case NM_SETTING_WIRED_SPEED:
		err = setSettingWiredSpeedJSON(data, valueJSON)
	case NM_SETTING_WIRED_DUPLEX:
		err = setSettingWiredDuplexJSON(data, valueJSON)
	case NM_SETTING_WIRED_AUTO_NEGOTIATE:
		err = setSettingWiredAutoNegotiateJSON(data, valueJSON)
	case NM_SETTING_WIRED_MAC_ADDRESS:
		err = setSettingWiredMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRED_CLONED_MAC_ADDRESS:
		err = setSettingWiredClonedMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST:
		err = setSettingWiredMacAddressBlacklistJSON(data, valueJSON)
	case NM_SETTING_WIRED_MTU:
		err = setSettingWiredMtuJSON(data, valueJSON)
	case NM_SETTING_WIRED_S390_SUBCHANNELS:
		err = setSettingWiredS390SubchannelsJSON(data, valueJSON)
	case NM_SETTING_WIRED_S390_NETTYPE:
		err = setSettingWiredS390NettypeJSON(data, valueJSON)
	case NM_SETTING_WIRED_S390_OPTIONS:
		err = setSettingWiredS390OptionsJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingWiredPortExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT)
}
func isSettingWiredSpeedExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED)
}
func isSettingWiredDuplexExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX)
}
func isSettingWiredAutoNegotiateExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE)
}
func isSettingWiredMacAddressExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS)
}
func isSettingWiredClonedMacAddressExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS)
}
func isSettingWiredMacAddressBlacklistExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST)
}
func isSettingWiredMtuExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU)
}
func isSettingWiredS390SubchannelsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS)
}
func isSettingWiredS390NettypeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE)
}
func isSettingWiredS390OptionsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS)
}

// Ensure field and key exists and not empty
func ensureFieldSettingWiredExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_WIRED_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_WIRED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_WIRED_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_WIRED_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_WIRED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_WIRED_SETTING_NAME))
	}
}
func ensureSettingWiredPortNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredPortExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredPort(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredSpeedNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredSpeedExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWiredDuplexNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredDuplexExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredDuplex(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredAutoNegotiateNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredAutoNegotiateExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWiredMacAddressNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredClonedMacAddressNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredClonedMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredClonedMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredMacAddressBlacklistNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredMacAddressBlacklistExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredMacAddressBlacklist(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredMtuNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredMtuExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWiredS390SubchannelsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredS390SubchannelsExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredS390Subchannels(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredS390NettypeNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredS390NettypeExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredS390Nettype(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWiredS390OptionsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWiredS390OptionsExists(data) {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWiredS390Options(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingWiredPort(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingWiredPort: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingWiredSpeed(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED)
	value, ok := ivalue.(uint32)
	if !ok {
		logger.Warningf("getSettingWiredSpeed: value type is invalid, should be uint32, %v", ivalue)
	}
	return
}
func getSettingWiredDuplex(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingWiredDuplex: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingWiredAutoNegotiate(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE)
	value, ok := ivalue.(bool)
	if !ok {
		logger.Warningf("getSettingWiredAutoNegotiate: value type is invalid, should be bool, %v", ivalue)
	}
	return
}
func getSettingWiredMacAddress(data connectionData) (value []byte) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS)
	value, ok := ivalue.([]byte)
	if !ok {
		logger.Warningf("getSettingWiredMacAddress: value type is invalid, should be []byte, %v", ivalue)
	}
	return
}
func getSettingWiredClonedMacAddress(data connectionData) (value []byte) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS)
	value, ok := ivalue.([]byte)
	if !ok {
		logger.Warningf("getSettingWiredClonedMacAddress: value type is invalid, should be []byte, %v", ivalue)
	}
	return
}
func getSettingWiredMacAddressBlacklist(data connectionData) (value []string) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST)
	value, ok := ivalue.([]string)
	if !ok {
		logger.Warningf("getSettingWiredMacAddressBlacklist: value type is invalid, should be []string, %v", ivalue)
	}
	return
}
func getSettingWiredMtu(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU)
	value, ok := ivalue.(uint32)
	if !ok {
		logger.Warningf("getSettingWiredMtu: value type is invalid, should be uint32, %v", ivalue)
	}
	return
}
func getSettingWiredS390Subchannels(data connectionData) (value []string) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS)
	value, ok := ivalue.([]string)
	if !ok {
		logger.Warningf("getSettingWiredS390Subchannels: value type is invalid, should be []string, %v", ivalue)
	}
	return
}
func getSettingWiredS390Nettype(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE)
	value, ok := ivalue.(string)
	if !ok {
		logger.Warningf("getSettingWiredS390Nettype: value type is invalid, should be string, %v", ivalue)
	}
	return
}
func getSettingWiredS390Options(data connectionData) (value map[string]string) {
	ivalue := getSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS)
	value, ok := ivalue.(map[string]string)
	if !ok {
		logger.Warningf("getSettingWiredS390Options: value type is invalid, should be map[string]string, %v", ivalue)
	}
	return
}

// Setter
func setSettingWiredPort(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, value)
}
func setSettingWiredSpeed(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, value)
}
func setSettingWiredDuplex(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, value)
}
func setSettingWiredAutoNegotiate(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, value)
}
func setSettingWiredMacAddress(data connectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, value)
}
func setSettingWiredClonedMacAddress(data connectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, value)
}
func setSettingWiredMacAddressBlacklist(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, value)
}
func setSettingWiredMtu(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, value)
}
func setSettingWiredS390Subchannels(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, value)
}
func setSettingWiredS390Nettype(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, value)
}
func setSettingWiredS390Options(data connectionData, value map[string]string) {
	setSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, value)
}

// JSON Getter
func getSettingWiredPortJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, getSettingWiredKeyType(NM_SETTING_WIRED_PORT))
	return
}
func getSettingWiredSpeedJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, getSettingWiredKeyType(NM_SETTING_WIRED_SPEED))
	return
}
func getSettingWiredDuplexJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, getSettingWiredKeyType(NM_SETTING_WIRED_DUPLEX))
	return
}
func getSettingWiredAutoNegotiateJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, getSettingWiredKeyType(NM_SETTING_WIRED_AUTO_NEGOTIATE))
	return
}
func getSettingWiredMacAddressJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS))
	return
}
func getSettingWiredClonedMacAddressJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, getSettingWiredKeyType(NM_SETTING_WIRED_CLONED_MAC_ADDRESS))
	return
}
func getSettingWiredMacAddressBlacklistJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST))
	return
}
func getSettingWiredMtuJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, getSettingWiredKeyType(NM_SETTING_WIRED_MTU))
	return
}
func getSettingWiredS390SubchannelsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, getSettingWiredKeyType(NM_SETTING_WIRED_S390_SUBCHANNELS))
	return
}
func getSettingWiredS390NettypeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, getSettingWiredKeyType(NM_SETTING_WIRED_S390_NETTYPE))
	return
}
func getSettingWiredS390OptionsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, getSettingWiredKeyType(NM_SETTING_WIRED_S390_OPTIONS))
	return
}

// JSON Setter
func setSettingWiredPortJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_PORT))
}
func setSettingWiredSpeedJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_SPEED))
}
func setSettingWiredDuplexJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_DUPLEX))
}
func setSettingWiredAutoNegotiateJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_AUTO_NEGOTIATE))
}
func setSettingWiredMacAddressJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS))
}
func setSettingWiredClonedMacAddressJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_CLONED_MAC_ADDRESS))
}
func setSettingWiredMacAddressBlacklistJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST))
}
func setSettingWiredMtuJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_MTU))
}
func setSettingWiredS390SubchannelsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_S390_SUBCHANNELS))
}
func setSettingWiredS390NettypeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_S390_NETTYPE))
}
func setSettingWiredS390OptionsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS, valueJSON, getSettingWiredKeyType(NM_SETTING_WIRED_S390_OPTIONS))
}

// Logic JSON Setter

// Remover
func removeSettingWiredPort(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_PORT)
}
func removeSettingWiredSpeed(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_SPEED)
}
func removeSettingWiredDuplex(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_DUPLEX)
}
func removeSettingWiredAutoNegotiate(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_AUTO_NEGOTIATE)
}
func removeSettingWiredMacAddress(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS)
}
func removeSettingWiredClonedMacAddress(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_CLONED_MAC_ADDRESS)
}
func removeSettingWiredMacAddressBlacklist(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MAC_ADDRESS_BLACKLIST)
}
func removeSettingWiredMtu(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU)
}
func removeSettingWiredS390Subchannels(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_SUBCHANNELS)
}
func removeSettingWiredS390Nettype(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_NETTYPE)
}
func removeSettingWiredS390Options(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_S390_OPTIONS)
}

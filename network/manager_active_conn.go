/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package network

import (
	"pkg.deepin.io/dde/daemon/network/nm"
	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
	. "pkg.deepin.io/lib/gettext"
)

type activeConnection struct {
	path dbus.ObjectPath
	typ  string

	Devices []dbus.ObjectPath
	Id      string
	Uuid    string
	State   uint32
	Vpn     bool
}

type activeConnectionInfo struct {
	IsPrimaryConnection bool
	Device              dbus.ObjectPath
	SettingPath         dbus.ObjectPath
	ConnectionType      string
	ConnectionName      string
	ConnectionUuid      string
	MobileNetworkType   string
	Security            string
	DeviceType          string
	DeviceInterface     string
	HwAddress           string
	Speed               string
	Ip4                 ip4ConnectionInfo
	Ip6                 ip6ConnectionInfo
	Hotspot             hotspotConnectionInfo
}
type ip4ConnectionInfo struct {
	Address  string
	Mask     string
	Gateways []string
	Dnses    []string
}
type ip6ConnectionInfo struct {
	Address  string
	Prefix   string
	Gateways []string
	Dnses    []string
}
type hotspotConnectionInfo struct {
	Ssid string
	Band string
}

func (m *Manager) initActiveConnectionManage() {
	m.initActiveConnections()

	// custom dbus watcher to catch all signals about active
	// connection, including vpn connection
	senderNm := "org.freedesktop.NetworkManager"
	pathNm := "/org/freedesktop/NetworkManager"
	interfaceNm := senderNm
	interfaceActive := "org.freedesktop.NetworkManager.Connection.Active"
	interfaceVpn := "org.freedesktop.NetworkManager.VPN.Connection"
	memberProperties := "PropertiesChanged"
	memberStateChanged := "StateChanged"
	memberVpnState := "VpnStateChanged"
	m.dbusWatcher.watch("type=signal," + "path=" + pathNm + ",interface=" + interfaceNm + ",member=" + memberProperties)
	m.dbusWatcher.watch("type=signal,sender=" + senderNm + ",interface=" + interfaceActive + ",member=" + memberStateChanged)
	m.dbusWatcher.watch("type=signal,sender=" + senderNm + ",interface=" + interfaceVpn + ",member=" + memberVpnState)

	// update active connection properties
	m.dbusWatcher.connect(func(s *dbus.Signal) {
		if s.Name == interfaceNm+"."+memberProperties && len(s.Body) == 1 {
			var props = make(map[string]dbus.Variant)
			props, _ = s.Body[0].(map[string]dbus.Variant)
			for prop, value := range props {
				if prop == "ActiveConnections" {
					// rebuild ActiveConnections prop
					m.rebuildActiveConnections(value.Value().([]dbus.ObjectPath))
				}
			}
			return
		}

		if s.Name == interfaceActive+"."+memberStateChanged && len(s.Body) >= 2 {
			state, _ := s.Body[0].(uint32)
			m.doUpdateActiveConnection(s.Path, state)
		}
	})

	// handle notification for vpn connections
	m.dbusWatcher.connect(func(s *dbus.Signal) {
		if s.Name == interfaceVpn+"."+memberVpnState && len(s.Body) >= 2 {
			state, _ := s.Body[0].(uint32)
			reason, _ := s.Body[1].(uint32)
			m.doHandleVpnNotification(s.Path, state, reason)
		}
	})
}

func (m *Manager) initActiveConnections() {
	m.activeConnectionsLock.Lock()
	defer m.activeConnectionsLock.Unlock()
	m.activeConnections = make(map[dbus.ObjectPath]*activeConnection)
	for _, path := range nmGetActiveConnections() {
		m.activeConnections[path] = m.newActiveConnection(path)
	}
	m.updatePropActiveConnections()
}

func (m *Manager) doHandleVpnNotification(apath dbus.ObjectPath, state, reason uint32) {
	m.activeConnectionsLock.Lock()
	defer m.activeConnectionsLock.Unlock()

	// get the corresponding active connection
	aconn, ok := m.activeConnections[apath]
	if !ok {
		return
	}

	// notification for vpn
	if isVpnConnectionStateActivated(state) {
		// FIXME: looks like a NetworkManger issue, when user
		// disconnect a connectiong vpn, the vpn state will changed to
		// nm.NM_VPN_CONNECTION_STATE_ACTIVATED first
		if reason != nm.NM_VPN_CONNECTION_STATE_REASON_USER_DISCONNECTED {
			notifyVpnConnected(aconn.Id)
		}
	} else if isVpnConnectionStateDeactivate(state) {
		notifyVpnDisconnected(aconn.Id)
	} else if isVpnConnectionStateFailed(state) {
		notifyVpnFailed(aconn.Id, reason)
	}

	if isVpnConnectionStateInActivating(state) {
		m.switchHandler.doEnableVpn(true)
	} else {
		// if vpn authentication dialog pop-up, notify to kill them
		connPath, _ := nmGetConnectionByUuid(aconn.Uuid)
		// TODO:
		//m.agent.cancelVpnAuthDialog(connPath)
		_ = connPath

		delete(m.activeConnections, apath)
	}
}
func (m *Manager) rebuildActiveConnections(apaths []dbus.ObjectPath) {
	m.clearActiveConnections()

	m.activeConnectionsLock.Lock()
	defer m.activeConnectionsLock.Unlock()
	for _, apath := range apaths {
		aconn := m.newActiveConnection(apath)
		m.activeConnections[apath] = aconn
	}

	m.updatePropActiveConnections()
}
func (m *Manager) doUpdateActiveConnection(apath dbus.ObjectPath, state uint32) {
	m.activeConnectionsLock.Lock()
	defer m.activeConnectionsLock.Unlock()

	aconn, ok := m.activeConnections[apath]
	if !ok {
		return
	}

	switch state {
	case nm.NM_ACTIVE_CONNECTION_STATE_ACTIVATING,
		nm.NM_ACTIVE_CONNECTION_STATE_ACTIVATED:
		// re-get all the active date especially vpn state for the
		// new connection
		aconn = m.newActiveConnection(apath)
		logger.Infof("add active connection %#v", aconn)
		m.activeConnections[apath] = aconn
	case nm.NM_ACTIVE_CONNECTION_STATE_DEACTIVATING,
		nm.NM_ACTIVE_CONNECTION_STATE_DEACTIVATED:
		// nothing to do
		logger.Infof("remove active connection %#v", aconn)
		// vpn's active connection will be removed after giving a
		// notification
	}

	m.updatePropActiveConnections()
}

func (m *Manager) newActiveConnection(path dbus.ObjectPath) (aconn *activeConnection) {
	aconn = &activeConnection{path: path}
	nmAConn, err := nmNewActiveConnection(path)
	if err != nil {
		return
	}

	aconn.State, _ = nmAConn.State().Get(0)
	aconn.Devices, _ = nmAConn.Devices().Get(0)
	aconn.typ, _ = nmAConn.Type().Get(0)
	aconn.Uuid, _ = nmAConn.Uuid().Get(0)
	aconn.Vpn, _ = nmAConn.Vpn().Get(0)
	if cpath, err := nmGetConnectionByUuid(aconn.Uuid); err == nil {
		aconn.Id = nmGetConnectionId(cpath)
	}

	return
}

func (m *Manager) clearActiveConnections() {
	m.activeConnectionsLock.Lock()
	defer m.activeConnectionsLock.Unlock()
	m.activeConnections = make(map[dbus.ObjectPath]*activeConnection)
	m.updatePropActiveConnections()
}

func (m *Manager) GetActiveConnectionInfo() (acinfosJSON string, busErr *dbus.Error) {
	var acinfos []activeConnectionInfo
	// get activated devices' connection information
	for _, devPath := range nmGetDevices() {
		if isDeviceStateActivated(nmGetDeviceState(devPath)) {
			if info, err := m.doGetActiveConnectionInfo(nmGetDeviceActiveConnection(devPath), devPath); err == nil {
				acinfos = append(acinfos, info)
			}
		}
	}
	// get activated vpn connection information
	for _, apath := range nmGetVpnActiveConnections() {
		if nmAConn, err := nmNewActiveConnection(apath); err == nil {
			if devs, _ := nmAConn.Devices().Get(0); len(devs) > 0 {
				devPath := devs[0]
				if info, err := m.doGetActiveConnectionInfo(apath, devPath); err == nil {
					acinfos = append(acinfos, info)
				}
			}
		}
	}
	acinfosJSON, err := marshalJSON(acinfos)
	busErr = dbusutil.ToError(err)
	return
}

func (m *Manager) doGetActiveConnectionInfo(apath, devPath dbus.ObjectPath) (acinfo activeConnectionInfo, err error) {
	var connType, connName, mobileNetworkType, security, devType, devIfc, hwAddress, speed string
	var ip4Address, ip4Mask string
	var ip4Gateways, ip4Dnses []string
	var ip6Address, ip6Prefix string
	var ip6Gateways, ip6Dnses []string
	var ip4Info ip4ConnectionInfo
	var ip6Info ip6ConnectionInfo
	var hotspotInfo hotspotConnectionInfo

	// active connection
	nmAConn, err := nmNewActiveConnection(apath)
	if err != nil {
		return
	}

	nmAConnConnection, _ := nmAConn.Connection().Get(0)
	nmConn, err := nmNewSettingsConnection(nmAConnConnection)
	if err != nil {
		return
	}

	// device
	nmDev, err := nmNewDevice(devPath)
	if err != nil {
		return
	}

	deviceType, _ := nmDev.DeviceType().Get(0)
	devType = getCustomDeviceType(deviceType)
	devIfc, _ = nmDev.Interface().Get(0)
	if devType == deviceModem {
		devUdi, _ := nmDev.Udi().Get(0)
		mobileNetworkType = mmGetModemMobileNetworkType(dbus.ObjectPath(devUdi))
	}

	// connection data
	hwAddress, err = nmGeneralGetDeviceHwAddr(devPath, false)
	if err != nil {
		hwAddress = ""
	}
	speed = nmGeneralGetDeviceSpeed(devPath)

	cdata, err := nmConn.GetSettings(0)
	if err != nil {
		return
	}
	connName = getSettingConnectionId(cdata)
	connType = getCustomConnectionType(cdata)
	if connType == connectionWirelessHotspot {
		hotspotInfo.Ssid = decodeSsid(getSettingWirelessSsid(cdata))
		band := getSettingWirelessBand(cdata)
		switch band {
		case "a":
			hotspotInfo.Band = Tr("A (5 GHz)")
		case "bg":
			hotspotInfo.Band = Tr("BG (2.4 GHz)")
		default:
			hotspotInfo.Band = Tr("Automatic")
		}
	}

	// security
	use8021xSecurity := false
	switch getSettingConnectionType(cdata) {
	case nm.NM_SETTING_WIRED_SETTING_NAME:
		if getSettingVk8021xEnable(cdata) {
			use8021xSecurity = true
		} else {
			security = Tr("None")
		}
	case nm.NM_SETTING_WIRELESS_SETTING_NAME:
		switch getSettingVkWirelessSecurityKeyMgmt(cdata) {
		case "none":
			security = Tr("None")
		case "wep":
			security = Tr("WEP 40/128-bit Key")
		case "wpa-psk":
			security = Tr("WPA/WPA2 Personal")
		case "wpa-eap":
			use8021xSecurity = true
		}
	}
	if use8021xSecurity {
		switch getSettingVk8021xEap(cdata) {
		case "tls":
			security = "EAP/" + Tr("TLS")
		case "md5":
			security = "EAP/" + Tr("MD5")
		case "leap":
			security = "EAP/" + Tr("LEAP")
		case "fast":
			security = "EAP/" + Tr("FAST")
		case "ttls":
			security = "EAP/" + Tr("Tunneled TLS")
		case "peap":
			security = "EAP/" + Tr("Protected EAP")
		}
	}

	// ipv4
	if ip4Path, _ := nmAConn.Ip4Config().Get(0); isNmObjectPathValid(ip4Path) {
		ip4Address, ip4Mask, ip4Gateways, ip4Dnses = nmGetIp4ConfigInfo(ip4Path)
	}
	ip4Info = ip4ConnectionInfo{
		Address:  ip4Address,
		Mask:     ip4Mask,
		Gateways: ip4Gateways,
		Dnses:    ip4Dnses,
	}

	// ipv6
	if ip6Path, _ := nmAConn.Ip6Config().Get(0); isNmObjectPathValid(ip6Path) {
		ip6Address, ip6Prefix, ip6Gateways, ip6Dnses = nmGetIp6ConfigInfo(ip6Path)
	}
	ip6Info = ip6ConnectionInfo{
		Address:  ip6Address,
		Prefix:   ip6Prefix,
		Gateways: ip6Gateways,
		Dnses:    ip6Dnses,
	}

	nmAConnUuid, _ := nmAConn.Uuid().Get(0)
	acinfo = activeConnectionInfo{
		IsPrimaryConnection: nmGetPrimaryConnection() == apath,
		Device:              devPath,
		SettingPath:         nmConn.Path_(),
		ConnectionType:      connType,
		ConnectionName:      connName,
		ConnectionUuid:      nmAConnUuid,
		MobileNetworkType:   mobileNetworkType,
		Security:            security,
		DeviceType:          devType,
		DeviceInterface:     devIfc,
		HwAddress:           hwAddress,
		Speed:               speed,
		Ip4:                 ip4Info,
		Ip6:                 ip6Info,
		Hotspot:             hotspotInfo,
	}
	return
}

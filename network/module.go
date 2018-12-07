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
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/dde/daemon/network/proxychains"
	"pkg.deepin.io/lib/log"
	libnotify "pkg.deepin.io/lib/notify"
)

var (
	logger  = log.NewLogger("daemon/network")
	manager *Manager
)

func init() {
	loader.Register(newModule(logger))
	proxychains.SetLogger(logger)
}

type Module struct {
	*loader.ModuleBase
}

func newModule(logger *log.Logger) *Module {
	module := new(Module)
	module.ModuleBase = loader.NewModuleBase("network", module, logger)
	return module
}

func (d *Module) GetDependencies() []string {
	return []string{}
}

func (d *Module) Start() error {
	libnotify.Init("dde-session-daemon")
	if manager != nil {
		return nil
	}

	initSlices() // initialize slice code
	service := loader.GetService()

	initSysSignalLoop()
	manager = NewManager(service)
	manager.init()

	managerServerObj, err := service.NewServerObject(dbusPath, manager)
	if err != nil {
		return err
	}

	managerServerObj.SetWriteCallback(manager, "NetworkingEnabled", manager.networkingEnabledWriteCb)
	managerServerObj.SetWriteCallback(manager, "VpnEnabled", manager.vpnEnabledWriteCb)

	err = managerServerObj.Export()
	if err != nil {
		logger.Error("failed to export manager:", err)
		manager = nil
		return err
	}

	manager.proxyChainsManager = proxychains.NewManager(service)
	err = service.Export(proxychains.DBusPath, manager.proxyChainsManager)
	if err != nil {
		logger.Warning("failed to export proxyChainsManager:", err)
		manager.proxyChainsManager = nil
		return err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		return err
	}

	go func() {
		initDBusDaemon()
		watchNetworkManagerRestart(manager)
	}()
	return nil
}

func (d *Module) Stop() error {
	if manager == nil {
		return nil
	}

	service := loader.GetService()

	err := service.ReleaseName(dbusServiceName)
	if err != nil {
		logger.Warning(err)
	}

	manager.destroy()
	destroyDBusDaemon()
	sysSigLoop.Stop()
	service.StopExport(manager)

	if manager.proxyChainsManager != nil {
		service.StopExport(manager.proxyChainsManager)
		manager.proxyChainsManager = nil
	}

	manager = nil
	return nil
}

package main

import (
        "dbus/com/deepin/api/setdatetime"
        "dlib"
        "dlib/dbus"
        "dlib/dbus/property"
        "dlib/gio-2.0"
        "dlib/logger"
        "github.com/howeyc/fsnotify"
)

const (
        _DATE_TIME_DEST = "com.deepin.daemon.DateAndTime"
        _DATE_TIME_PATH = "/com/deepin/daemon/DateAndTime"
        _DATA_TIME_IFC  = "com.deepin.daemon.DateAndTime"

        _DATE_TIME_SCHEMA = "com.deepin.dde.datetime"
        _TIME_ZONE_FILE   = "/etc/timezone"
)

var (
        busConn      *dbus.Conn
        dateSettings = gio.NewSettings(_DATE_TIME_SCHEMA)

        setDate     *setdatetime.SetDateTime
        zoneWatcher *fsnotify.Watcher
)

type Manager struct {
        AutoSetTime      *property.GSettingsBoolProperty `access:"readwrite"`
        Use24HourDisplay *property.GSettingsBoolProperty `access:"readwrite"`
        CurrentTimeZone  string
}

func (op *Manager) SetDate(d string) (bool, error) {
        ret, err := setDate.SetCurrentDate(d)
        if err != nil {
                logger.Printf("Set Date - '%s' Failed: %s\n",
                        d, err)
                return false, err
        }
        return ret, nil
}

func (op *Manager) SetTime(t string) (bool, error) {
        ret, err := setDate.SetCurrentTime(t)
        if err != nil {
                logger.Printf("Set Time - '%s' Failed: %s\n",
                        t, err)
                return false, err
        }
        return ret, nil
}

func (op *Manager) TimezoneCityList() []string {
        return getZoneCityList()
}

func (op *Manager) SetTimeZone(zone string) bool {
        //tz := convertCityToZone(zone)
        tz := getKeyByValue(zone)
        //logger.Println("tz:", tz)
        _, err := setDate.SetTimezone(tz)
        if err != nil {
                logger.Printf("Set TimeZone - '%s' Failed: %s\n",
                        zone, err)
                return false
        }
        op.setPropName("CurrentTimeZone")
        return true
}

func (op *Manager) SetAutoSetTime(auto bool) (bool, error) {
        var (
                ret     bool
                err     error
        )

        if auto {
                ret, err = setDate.SetNtpUsing(true)
        } else {
                ret, err = setDate.SetNtpUsing(false)
        }

        if err != nil {
                logger.Printf("Set NTP - %d Failed: %s\n",
                        auto, err)
                return false, err
        }
        return ret, nil
}

func (op *Manager) SyncNtpTime() bool {
        ret, _ := setDate.SyncNtpTime()
        return ret
}

func NewDateAndTime() *Manager {
        m := &Manager{}

        m.AutoSetTime = property.NewGSettingsBoolProperty(
                m, "AutoSetTime",
                dateSettings, "is-auto-set")
        m.Use24HourDisplay = property.NewGSettingsBoolProperty(
                m, "Use24HourDisplay",
                dateSettings, "is-24hour")

        m.setPropName("CurrentTimeZone")
        m.listenSettings()
        m.listenZone()

        return m
}

func Init() {
        var err error

        setDate, err = setdatetime.NewSetDateTime("/com/deepin/api/SetDateTime")
        if err != nil {
                logger.Println("New SetDateTime Failed:", err)
                panic(err)
        }

        zoneWatcher, err = fsnotify.NewWatcher()
        if err != nil {
                logger.Println("New FS Watcher Failed:", err)
                panic(err)
        }
}

func main() {
        defer func() {
                if err := recover(); err != nil {
                        logger.Println("recover error:", err)
                }
        }()

        Init()
        defer zoneWatcher.Close()

        date := NewDateAndTime()
        err := dbus.InstallOnSession(date)
        if err != nil {
                logger.Println("Install Session DBus Failed:", err)
                panic(err)
        }

        if date.AutoSetTime.Get() {
                go setDate.SetNtpUsing(true)
        }
        dbus.DealWithUnhandledMessage()
        dlib.StartLoop()
}

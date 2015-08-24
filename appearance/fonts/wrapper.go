package fonts

// #cgo pkg-config: fontconfig
// #include <stdlib.h>
// #include "font_list.h"
import "C"

import (
	"os"
	"strings"
	"unsafe"
)

const (
	defaultLang      = "en"
	defaultLangDelim = "|"
	defaultNameDelim = ","
	spaceTypeMono    = "100"
)

// family ex: 'sans', 'serif', 'monospace'
// cRet: `SourceCodePro-Medium.otf: "Source Code Pro" "Medium"`
func fcFontMatch(family string) string {
	cFamily := C.CString(family)
	defer C.free(unsafe.Pointer(cFamily))
	cRet := C.font_match(cFamily)
	ret := C.GoString(cRet)

	if len(ret) == 0 {
		return ""
	}

	tmp := strings.Split(ret, ":")
	if len(tmp) != 2 {
		return ""
	}

	return strings.Split(tmp[1], "\"")[1]
}

func fcInfosToFonts() Fonts {
	var num = C.int(0)
	list := C.get_font_info_list(&num)
	if num < 1 {
		return nil
	}
	defer C.font_info_list_free(list, num)

	listPtr := uintptr(unsafe.Pointer(list))
	itemLen := unsafe.Sizeof(*list)

	var infos Fonts
	for i := C.int(0); i < num; i++ {
		cItem := (*C.FcInfo)(unsafe.Pointer(
			listPtr + uintptr(i)*itemLen))

		infos = append(infos, fcInfoToFont(cItem))
	}
	return infos
}

func fcInfoToFont(cInfo *C.FcInfo) *Font {
	names := strings.Split(C.GoString(cInfo.fullname), defaultNameDelim)
	nameLang := strings.Split(C.GoString(cInfo.fullnamelang),
		defaultNameDelim)
	families := strings.Split(C.GoString(cInfo.family), defaultNameDelim)
	familyLang := strings.Split(C.GoString(cInfo.familylang),
		defaultNameDelim)

	var info = Font{
		Id: getItemByIndex(indexList(defaultLang,
			nameLang), names),
		Name: getItemByIndex(indexList(getCurLang(),
			nameLang), names),
		Family: getItemByIndex(indexList(defaultLang,
			familyLang), families),
		FamilyName: getItemByIndex(indexList(getCurLang(),
			familyLang), families),
		File: C.GoString(cInfo.filename),
		Styles: strings.Split(C.GoString(cInfo.style),
			defaultNameDelim),
		Lang: strings.Split(C.GoString(cInfo.lang),
			defaultLangDelim),
	}
	info.Monospace = isMonospace(info.Name,
		C.GoString(cInfo.spacing))
	info.Deletable = isDeletable(info.File)

	return &info
}

func isMonospace(name, spacing string) bool {
	if spacing == spaceTypeMono ||
		strings.Contains(strings.ToLower(name), "mono") {
		return true
	}

	return false
}

func isDeletable(file string) bool {
	if strings.Contains(file, os.Getenv("HOME")) {
		return true
	}
	return false
}

func getItemByIndex(idx int, list []string) string {
	if len(list) == 0 {
		return ""
	}

	if idx < 0 || len(list) <= idx {
		return list[0]
	}

	return list[idx]
}

func indexList(item string, list []string) int {
	for i, v := range list {
		if v == item {
			return i
		}
	}
	return -1
}

func getCurLang() string {
	lang := os.Getenv("LANGUAGE")
	if len(lang) == 0 {
		return defaultLang
	}

	switch lang {
	case "zh_CN":
		lang = "zh-cn"
	case "zh_TW", "zh_HK":
		lang = "zh-tw"
	}

	return lang
}
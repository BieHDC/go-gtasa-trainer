package subhook

//#define SUBHOOK_STATIC
//#include "subhook_c/subhook.c"
import "C"
import (
	"unsafe"

	"golang.org/x/sys/windows"
)

//fixme:
// this could get a native port, maybe even with c2go

// Wrappers for the API

func SubhookNew[Detour any](original uintptr, detour Detour) C.subhook_t {
	functionaddress := *(*uintptr)(unsafe.Pointer(&detour)) //highly sus
	return C.subhook_new(unsafe.Pointer(original), unsafe.Pointer(functionaddress), C.SUBHOOK_TRAMPOLINE)
}

func SubhookInstall(sh C.subhook_t) int {
	return int(C.subhook_install(sh))
}

func SubhookRemove(sh C.subhook_t) int {
	return int(C.subhook_remove(sh))
}

func SubhookFree(sh C.subhook_t) {
	C.subhook_free(sh)
}

// returns a function pointer you can call
func SubhookGetTrampoline(sh C.subhook_t) unsafe.Pointer {
	return C.subhook_get_trampoline(sh)
}

func SubhookTypeonly() C.subhook_t {
	var zero C.subhook_t
	return zero
}

// The other method
// Returns the previous value at address
func SwapVtablePtr[Detour any](address uintptr, value Detour) uintptr {
	functionaddress := *(*uintptr)(unsafe.Pointer(&value)) //highly sus

	const PTR_SIZE = 4
	var oldprotect uint32
	err := windows.VirtualProtect(address, PTR_SIZE, windows.PAGE_EXECUTE_READWRITE, &oldprotect)
	if err != nil {
		return 0 //failed
	}

	oldaddressptr := (*uint32)(unsafe.Pointer(&address))
	oldaddress := uintptr(*oldaddressptr)
	*oldaddressptr = uint32(functionaddress)

	var throwaway uint32
	err = windows.VirtualProtect(address, PTR_SIZE, oldprotect, &throwaway)
	if err != nil {
		return 0
	}

	return oldaddress
}

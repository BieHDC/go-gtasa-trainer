package main

/*
// Shared definitions
#include <windows.h>


// Hooked Functions
typedef void (*fpRenderScene)(void);
extern fpRenderScene ogrenderscene;
void RenderScene(void);
void RenderSceneFirst(void);

typedef int (*fpRsEventHandler)(int, void*);
extern fpRsEventHandler ogrseventhandler;
int RsEventHandler(int event, void* param);

typedef LRESULT CALLBACK(*fpMainWndProc)(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam);
extern fpMainWndProc ogmainwndproc;
LRESULT CALLBACK myMainWndProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam);
*/
import "C"

import (
	. "gtasamod/subhook"
)

const renderaddress = uintptr(0x53DF40)

var hookrender = SubhookTypeonly()

func SetRenderFuncFirstTime() {
	hookrender = SubhookNew(renderaddress, C.RenderSceneFirst)
	SubhookInstall(hookrender)
	C.ogrenderscene = (C.fpRenderScene)(SubhookGetTrampoline(hookrender))
}

func SetRenderFuncLooping() {
	SubhookRemove(hookrender)
	SubhookFree(hookrender)
	hookrender = SubhookNew(renderaddress, C.RenderScene)
	SubhookInstall(hookrender)
	C.ogrenderscene = (C.fpRenderScene)(SubhookGetTrampoline(hookrender))
}

func SetupFunctionsHooks() {
	SetRenderFuncFirstTime()

	const menurenderaddress = uintptr(0x619B60)
	hookeventhandler := SubhookNew(menurenderaddress, C.RsEventHandler)
	SubhookInstall(hookeventhandler)
	C.ogrseventhandler = (C.fpRsEventHandler)(SubhookGetTrampoline(hookeventhandler))

	const mainwndprocaddress = uintptr(0x747EB0)
	hookmainwndproc := SubhookNew(mainwndprocaddress, C.myMainWndProc)
	SubhookInstall(hookmainwndproc)
	C.ogmainwndproc = (C.fpMainWndProc)(SubhookGetTrampoline(hookmainwndproc))
}

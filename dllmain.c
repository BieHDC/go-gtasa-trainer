// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// Shared definitions
#include <stdint.h>
#include <stdio.h>
#include <windows.h>
// Using the sometimes-maybe generated "dinput8.h" is too flaky.
#include "_cgo_export.h"

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// Defines
// If enabled it does more verbose printing we only use for debugging,
//	but not for normal use
// #define DEBUG

// For Printing Stuff
#define FIXME(M, ...)                                                          \
  (printf("[FIXME] (%s:%d) " M "\n", __FILE__, __LINE__, ##__VA_ARGS__))
#define INFO(M, ...)                                                           \
  (printf("[INFO] (%s:%d) " M "\n", __FILE__, __LINE__, ##__VA_ARGS__))
#define ERR(M, ...)                                                            \
  (printf("[ERR] (%s:%d) " M "\n", __FILE__, __LINE__, ##__VA_ARGS__))

#ifdef DEBUG
#define SPAM(M, ...)                                                           \
  (printf("[SPAM] (%s:%d) " M "\n", __FILE__, __LINE__, ##__VA_ARGS__))
#else
#define SPAM(...)
#endif

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// Helpers
void cstrappend(char *dst, const char *str1) {
  size_t dst_len = strlen(dst);
  size_t i = 0;

  if (dst) {
    for (i = 0; str1[i] != '\0'; i++)
      (dst)[dst_len + i] = str1[i];

    (dst)[dst_len + i] = '\0';
  }
}

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// Original Function
typedef HRESULT(__stdcall *DIRECTINPUT8CREATE)(HINSTANCE, DWORD, REFIID,
                                               LPVOID *, LPUNKNOWN);
DIRECTINPUT8CREATE fpDirectInput8Create;
__declspec(dllexport) HRESULT
    __stdcall DirectInput8Create(HINSTANCE hinst, DWORD dwVersion,
                                 REFIID riidltf, LPVOID *ppvOut,
                                 LPUNKNOWN punkOuter) {
  if (!fpDirectInput8Create) {
    char syspath[PATH_MAX];
    UINT sysret = GetSystemDirectoryA(syspath, PATH_MAX - 1);
    if ((sysret == 0) || (sysret > PATH_MAX - 1)) {
      // Fail
      ERR("syspath buffer too small, or general fail: got %u but max is %u",
          sysret, PATH_MAX - 1);
      return E_OUTOFMEMORY; //=DIERR_OUTOFMEMORY //Lets at least not crash
    }

    unsigned int lenofstring = strlen(syspath) + strlen("\\dinput8.dll");
    if (lenofstring > PATH_MAX) {
      ERR("syspath and dll name too long: got %u but max is %u", lenofstring,
          PATH_MAX);
      return E_OUTOFMEMORY; //=DIERR_OUTOFMEMORY //Lets at least not crash
    }

    cstrappend(syspath, "\\dinput8.dll");
    HMODULE hMod = LoadLibraryA(syspath);
    if (hMod) {
      fpDirectInput8Create =
          (DIRECTINPUT8CREATE)GetProcAddress(hMod, "DirectInput8Create");
      INFO("fpDirectInput8Create: %p", (void *)fpDirectInput8Create);
    } else {
      MessageBoxA(NULL, "Failed to load original dinput8.dll", "", 0);
      return E_OUTOFMEMORY; //=DIERR_OUTOFMEMORY //Lets at least not crash
    }
  }
  return fpDirectInput8Create(hinst, dwVersion, riidltf, ppvOut, punkOuter);
}

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// DllMain Stuff when we need to initialise things
// DO NOT call Go functions inside here, it will lead to a hang!
// Go init() functions will run after DllMain returns, use this place
// to start off the stuff you need.
BOOL APIENTRY DllMain(HMODULE hModule, DWORD reason, LPVOID lpReserved) {
  switch (reason) {
  case DLL_PROCESS_ATTACH: {
    // If we have a console, we attach to it for printing
    BOOL attached = AttachConsole(ATTACH_PARENT_PROCESS);
    if (attached) {
      freopen("CON", "w", stdout);
      freopen("CON", "r", stdin);
      freopen("CON", "w", stderr);
    }
    DisableThreadLibraryCalls(
        hModule); // Disabled unless we need it at some point
    INFO("DINPUT8.dll Hook Loaded");
  } break;
  case DLL_THREAD_ATTACH:
  case DLL_THREAD_DETACH:
    // ignored
    return TRUE;
  case DLL_PROCESS_DETACH: {
    INFO("DINPUT8.dll Hook Unloaded");
  } break;
  } // switch

  return TRUE;
}
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// Before C can call to Go, Go must be fully started. Go's init function will
//	call this and thats how we know we can now do things
void GoRuntimeHasFullyLoaded(void) {
  SPAM("Go is ready to receive calls");
  GODLLmain();
}

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// Userfunctions
// For rundll32 testing
__declspec(dllexport) void GolangIsGoing(void) {
  INFO("Called GolangIsGoing");
  Ccalls();
  Sleep(15000);
}

// hook for stuff while the game is rendering (not in menu etc)
// 0x53DF40
typedef void (*fpRenderScene)(void);
fpRenderScene ogrenderscene = NULL;
// This is for the first time init
void RenderSceneFirst(void) {
  RenderSceneFirstGo();
  ogrenderscene();
}
// This is for the rest of the lifetime
void RenderScene(void) {
  RenderSceneGo();
  ogrenderscene();
}

// catch game quit event so we can stop the world and get rid off some flaky
// crashes 0x619B60
#define rsQUITAPP 30
typedef int (*fpRsEventHandler)(int, void *);
fpRsEventHandler ogrseventhandler = NULL;
int RsEventHandler(int event, void *param) {
  if (event == rsQUITAPP) {
    GameIsQuitting();
  }
  return ogrseventhandler(event, param);
}

typedef enum {
  Backspace = 1000, // VK_BACK
  Tab,              // VK_TAB
  Enter,            // VK_RETURN
  Shift,            // VK_SHIFT
  ShiftL,           // VK_LSHIFT
  ShiftR,           // VK_RSHIFT
  Ctrl,             // VK_CONTROL
  CtrlL,            // VK_LCONTROL
  CtrlR,            // VK_RCONTROL
  Alt,              // VK_MENU
  AltL,             // VK_LMENU
  AltR,             // VK_RMENU
  Pause,            // VK_PAUSE
  Capslock,         // VK_CAPITAL
  Esc,              // VK_ESCAPE
  PageUp,           // VK_PRIOR
  PageDown,         // VK_NEXT
  End,              // VK_END
  Home,             // VK_HOME
  ArrLeft,          // VK_LEFT
  ArrUp,            // VK_UP
  ArrRight,         // VK_RIGHT
  ArrDown,          // VK_DOWN
  Insert,           // VK_INSERT
  Delete,           // VK_DELETE
  Numpad0,          // VK_NUMPAD0
  Numpad1,          // VK_NUMPAD1
  Numpad2,          // VK_NUMPAD2
  Numpad3,          // VK_NUMPAD3
  Numpad4,          // VK_NUMPAD4
  Numpad5,          // VK_NUMPAD5
  Numpad6,          // VK_NUMPAD6
  Numpad7,          // VK_NUMPAD7
  Numpad8,          // VK_NUMPAD8
  Numpad9,          // VK_NUMPAD9
  F1,               // VK_F1
  F2,               // VK_F2
  F3,               // VK_F3
  F4,               // VK_F4
  F5,               // VK_F5
  F6,               // VK_F6
  F7,               // VK_F7
  F8,               // VK_F8
  F9,               // VK_F9
  F10,              // VK_F10
  F11,              // VK_F11
  F12               // VK_F12
} SpecialKeys;

int validRange(int16_t c) {
  //       <<  -  chars  -  >>     << - special keys - >>
  return (c >= 32 && c <= 255) || (c >= 1000 && c <= 1046);
}

// we are going to hook this, do our default checking, return false if we
// handled the event in our code return ogcall() if its not our business!
const GoUint8 GoFalse = 0;
const GoUint8 GoTrue = 1;
static GoUint8 inputting = GoFalse;

// 0x747EB0
typedef LRESULT CALLBACK (*fpMainWndProc)(HWND hWnd, UINT uMsg, WPARAM wParam,
                                          LPARAM lParam);
fpMainWndProc ogmainwndproc = NULL;

// static int32_t* gGameState = (int32_t*)0xC8D4C0;
LRESULT CALLBACK myMainWndProc(HWND hWnd, UINT uMsg, WPARAM wParamOG,
                               LPARAM lParam) {
  // can skip the intros
  // but can also cause random hangs during loadings
  /*
  switch (*gGameState) {
          case 0:
          case 1:
          case 3: {
                  *gGameState = 4;
          } break;
  }
  */

  WPARAM wParam = wParamOG; // do not clobber the value for ogcall
  int up = 0;
  switch (uMsg) {
  case WM_KEYDOWN:
  case WM_KEYUP:
  case WM_SYSKEYDOWN:
  case WM_SYSKEYUP: {
    // int down = !((lParam >> 31) & 1);
    up = !!((lParam >> 31) & 1);
    // int ctrl = GetKeyState(VK_CONTROL) & (1 << 15);

    switch (wParam) {
    case VK_BACK:
      wParam = Backspace;
      break;
    case VK_TAB:
      wParam = Tab;
      break;
    case VK_RETURN:
      wParam = Enter;
      break;
    case VK_SHIFT:
      wParam = Shift;
      break;
    case VK_LSHIFT:
      wParam = ShiftL;
      break;
    case VK_RSHIFT:
      wParam = ShiftR;
      break;
    case VK_CONTROL:
      wParam = Ctrl;
      break;
    case VK_LCONTROL:
      wParam = CtrlL;
      break;
    case VK_RCONTROL:
      wParam = CtrlR;
      break;
    case VK_MENU:
      wParam = Alt;
      break;
    case VK_LMENU:
      wParam = AltL;
      break;
    case VK_RMENU:
      wParam = AltR;
      break;
    case VK_PAUSE:
      wParam = Pause;
      break;
    case VK_CAPITAL:
      wParam = Capslock;
      break;
    case VK_ESCAPE:
      wParam = Esc;
      break;
    case VK_PRIOR:
      wParam = PageUp;
      break;
    case VK_NEXT:
      wParam = PageDown;
      break;
    case VK_END:
      wParam = End;
      break;
    case VK_HOME:
      wParam = Home;
      break;
    case VK_LEFT:
      wParam = ArrLeft;
      break;
    case VK_UP:
      wParam = ArrUp;
      break;
    case VK_RIGHT:
      wParam = ArrRight;
      break;
    case VK_DOWN:
      wParam = ArrDown;
      break;
    case VK_INSERT:
      wParam = Insert;
      break;
    case VK_DELETE:
      wParam = Delete;
      break;
    case VK_NUMPAD0:
      wParam = Numpad0;
      break;
    case VK_NUMPAD1:
      wParam = Numpad1;
      break;
    case VK_NUMPAD2:
      wParam = Numpad2;
      break;
    case VK_NUMPAD3:
      wParam = Numpad3;
      break;
    case VK_NUMPAD4:
      wParam = Numpad4;
      break;
    case VK_NUMPAD5:
      wParam = Numpad5;
      break;
    case VK_NUMPAD6:
      wParam = Numpad6;
      break;
    case VK_NUMPAD7:
      wParam = Numpad7;
      break;
    case VK_NUMPAD8:
      wParam = Numpad8;
      break;
    case VK_NUMPAD9:
      wParam = Numpad9;
      break;
    case VK_F1:
      wParam = F1;
      break;
    case VK_F2:
      wParam = F2;
      break;
    case VK_F3:
      wParam = F3;
      break;
    case VK_F4:
      wParam = F4;
      break;
    case VK_F5:
      wParam = F5;
      break;
    case VK_F6:
      wParam = F6;
      break;
    case VK_F7:
      wParam = F7;
      break;
    case VK_F8:
      wParam = F8;
      break;
    case VK_F9:
      wParam = F9;
      break;
    case VK_F10:
      wParam = F10;
      break;
    case VK_F11:
      wParam = F11;
      break;
    case VK_F12:
      wParam = F12;
      break;
    // special chars we do not convert we do not care about
    default: {
      goto nope;
    } break;
    } // switch (wParam)
  } break;

  case WM_CHAR: {
    if (wParam < 32) {
      goto nope;
    }
  } break;

  // we got a type of message we do not care about, like window resize and etc
  default: {
    goto notourmessage;
  } break;

  } // switch (uMsg)

  // if we care about it and its not keyup (special keys)
  if (validRange(wParam) && !up) {
    if (TypedChars((GoUint32)wParam, &inputting)) {
      return 0; // if we handled it, return
    }           // otherwise we fall through to the og game code
  }

// not our thing, but a type of message we cared about
// block forwarding if we are inputting
nope:
  // if we are currently capturing input, we block key events
  if (inputting == GoTrue) {
    return 0;
  }

// Was not our thing, forward to the game
notourmessage:
  return ogmainwndproc(hWnd, uMsg, wParamOG, lParam);
}

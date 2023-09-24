@echo off
rem echo "!!! disabled c recompiling, reenable!! >> -a after build"
set CGO_ENABLED=1&& set GOARCH=386&& go build -a -buildmode=c-shared -ldflags="-w -s -H=windowsgui -extldflags \"-Wl,--kill-at,--enable-stdcall-fixup\"" -o dinput8.dll

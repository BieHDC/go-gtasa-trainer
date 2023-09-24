CGO_ENABLED=1 GOARCH=386 GOOS=windows CC=i686-w64-mingw32-gcc go build -a -buildmode=c-shared -ldflags="-w -s -H=windowsgui -extldflags \"-Wl,--kill-at,--enable-stdcall-fixup\"" -o dinput8.dll

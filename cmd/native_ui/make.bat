@echo off
rem go build -tags debug
fyne package -release -os windows
set ANDROID_NDK_HOME=C:\android-ndk-r25c&& fyne package -release -os android
rem Get it at https://github.com/android/ndk

PATH=$PATH:~/go/bin
fyne build -release -os linux -o biehdc.trainer.gtasa
CGO_ENABLED=1 GOARCH=386 GOOS=windows CC=i686-w64-mingw32-gcc fyne package -release -os windows
ANDROID_NDK_HOME=~/code/android-ndk-r25c fyne package -release -os android/arm
# Get it at https://github.com/android/ndk

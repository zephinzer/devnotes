# Setting up Android SDK on Linux

> These instructions were created in the context of setting up my Android device for use with React-Native using **ONLY** command line tools.

## Pre-setup Notes
These steps have worked for an Ubuntu 16.04 installation on a Lenovo ThinkPad T480s.

`uname -a` output:

```
Linux ${USER} 4.13.0-43-generic #48~16.04.1-Ubuntu SMP Thu $MONTH $DATE $HR:$MIN:$SEC UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```

> There's a list of [gotchas](#gotchas) that wasted me lots of time. Check 'em out if you encounter any errors.

## Instructions

1. Go to https://developer.android.com/studio/#downloads and grab the right SDK tools for your system.
2. Assuming you downloaded it to `~/Downloads`, run `sudo mkdir -p /opt/android && mv ~/Downloads/sdk-tools-linux-******.zip /opt/android`. This moves your downloaded `.zip` to the correct directory.
3. Modify the directory settings to your user so that `sdkmanager` (in step 9) can write to the `$ANDROID_HOME/licenses` directory. Do this with: `sudo chown $USER:$USER -R /opt/android`. 
4. Go to the `/opt/android` directory and run `unzip ./sdk-tools-linux-*******.zip`. A `./tools` directory should appear.
5. Add `/opt/android/tools:/opt/android/tools/bin` into your `$PATH` variable.
6. Add `export ANDROID_HOME=/opt/android` as an environment variable.
7. Install the Java JDK if you don't already have it (check with `java -version`) with `sudo apt update && sudo apt install defaultjdk`.
8. Add `export JAVA_HOME= /usr/lib/jvm/java-${JAVA_MAJOR_VERSION}-openjdk-amd64` where `${JAVA_MAJOR_VERSION}` is 8 or 9 (at point of publishing - July 2018).
9. Accept the SDK licenses by running `sdkmanager --licenses` (they'll try to make you accept all, if you're wary of terms of uses, make sure you only accept the main SDK license)

## Preparing Android Device
1. Go to your **System Settings > System > About phone**
2. Repeatedly tap the **Build number** field to enable developer mode
3. Go to the new **Developer options** item in **System Settings > System**
4. Scroll to USB debugging and turn it on.

## Running Android App on Device
1. Follow the instructions at https://facebook.github.io/react-native/docs/running-on-device. It's also replicated below just incase the site gets removed
2. Run `lsusb` and get the line corresponding to your phone. Mine was `Bus 003 Device 002: ID 18d1:xxxx Google Inc. ` (Google pixel). Note the 18d1. That's the identifier for your phone. We'll use it in the next step
3. Run `echo SUBSYSTEM=="usb", ATTR{idVendor}=="18d1", MODE="0666", GROUP="plugdev" | sudo tee /etc/udev/rules.d/51-android-usb.rules`.
4. Check it works by running `adb devices -l`. If you see a device ID followed by `no permissions ...`, follow the next few steps, or else **you're in**.
5. If this doesn't work, shut down the ADB server: `sudo adb kill-server`
6. **Unplug your device** and confirm that USB debugging is on
7. Turn the ADB server back on: `sudo adb start-server`
8. Plug your device in. You should see a prompt on your device asking for permissions, tap on the **Always allow** checkbox and hit **OK**.
9. Done.


## Gotchas

### `sdkmanager --licenses` cannot accept licenses
If you run `sdkmanager --licenses` and the licenses are still not accepted, it's:

1. `$ANDROID_HOME` is not set
2. `sdkmanager` has no permissions to `$ANDROID_HOME/licenses`
3. `$ANDROID_HOME` is set to your tools directory (the stupid mistake I made)

This fails silently and wasted me lots of time.

### `.android/repositories.cfg` not found
The error message sounds like it was your duty to make sure it was there. It'll be there if you installed Android Studio, but if you installed the command line tools only, it won't. To solve this, simply run:

```bash
touch ~/.android/repositories.cfg
```
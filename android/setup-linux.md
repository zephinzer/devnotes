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
3. Go to the `/opt/android` directory and run `unzip ./sdk-tools-linux-*******.zip`. A `./tools` directory should appear.
4. Add `/opt/android/tools:/opt/android/tools/bin` into your `$PATH` variable.
5. Add `export ANDROID_HOME=/opt/android` as an environment variable.
6. Install the Java JDK if you don't already have it (check with `java -version`) with `sudo apt update && sudo apt install defaultjdk`.
7. Add `export JAVA_HOME= /usr/lib/jvm/java-${JAVA_MAJOR_VERSION}-openjdk-amd64` where `${JAVA_MAJOR_VERSION}` is 8 or 9 (at point of publishing - July 2018).
8. Accept the SDK licenses by running `sdkmanager --licenses` (they'll try to make you accept all, if you're wary of terms of uses, make sure you only accept the main SDK license)

## Gotchas

### `sdkmanager --licenses` cannot accept licenses
If you run `sdkmanager --licenses` and the licenses are still not accepted, it's:

1. `$ANDROID_HOME` is not set
2. `sdkmanager` has no permissions to `$ANDROID_HOME/licenses`
3. `$ANDROID_HOME` is set to your tools directory (the stupid mistake I made)

This fails silently and wasted me lots of time.
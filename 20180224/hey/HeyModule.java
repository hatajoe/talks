package com.hey;

import hey.Hey;
...

public class HeyModule extends ReactContextBaseJavaModule {
    ...

    @ReactMethod
    public void connect(String url, String origin, final String eventName) {
        Hey.connect(url, origin, new hey.OnReceive() {
            @Override
            public void receive(String msg) {
                reactContext
                        .getJSModule(DeviceEventManagerModule.RCTDeviceEventEmitter.class)
                        .emit(eventName, msg);
            }
        });
    }
}

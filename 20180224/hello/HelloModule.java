package com.app;

import hello.Hello;
import ... 

public class HelloModule extends ReactContextBaseJavaModule {
    public HelloModule(ReactApplicationContext reactContext) {
        super(reactContext);
    }

    @Override
    public String getName() {
        return "Core";
    }

    @ReactMethod
    public void Hello() {
        Hello.hello();
    }
}

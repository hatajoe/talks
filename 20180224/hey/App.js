import React, { Component } from 'react';
import {
	...
	DeviceEventEmitter
} from 'react-native';

...

export default class App extends Component<Props> {
	constructor(props) {
		...

		DeviceEventEmitter.addListener("onReceive", function(e: Event) {
			...
			that.setState({received: e});
  	});

		...
	}
	...
});

import * as models from './models';

export interface go {
  "main": {
    "App": {
		Greet(arg1:string):Promise<string>
		GreetAsyncViaChannel():Promise<any>
		GreetAsyncViaEvent():Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}

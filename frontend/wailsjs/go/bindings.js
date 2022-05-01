// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
const go = {
  "main": {
    "App": {
      /**
       * Greet
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "Greet": (arg1) => {
        return window.go.main.App.Greet(arg1);
      },
      /**
       * GreetAsyncViaChannel
       * @returns {Promise<any>}  - Go Type: <-chan string
       */
      "GreetAsyncViaChannel": () => {
        return window.go.main.App.GreetAsyncViaChannel();
      },
      /**
       * GreetAsyncViaEvent
       * @returns {Promise<void>} 
       */
      "GreetAsyncViaEvent": () => {
        return window.go.main.App.GreetAsyncViaEvent();
      },
    },
  },

};
export default go;

const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      this.baseUrl = config.baseUrl;
      
    },
    baseUrl: "https://localhost:8443",
    testingType: "e2e", 
    specPattern: '**/*.spec.js',   }, 
    env: {
      "USERNAME": "username",
      "PASSWORD": "password"
    },
   
});

const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      this.baseUrl = config.baseUrl;
      
    },
  }, 
    env: {
      "USERNAME": "admin",
      "PASSWORD": "Ho9h3GNZAL5jdJvg"
    }
});

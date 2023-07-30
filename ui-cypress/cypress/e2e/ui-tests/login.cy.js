describe('Rancher Web UI Tests', () => {
  beforeEach(() => {
  
    cy.visit('https://localhost:8443/dashboard/auth/login')
    cy.get('[class="text-center login-welcome"]').should('be.visible');  })

  it('Check if main webpage opens up', () => {
    cy.get('[class="text-center login-welcome"]').should('contain', 'Welcome to Rancher');
  })

  it('Check if main webpage title is correct', () => {
    cy.title().should('eq', 'Rancher');
  })

  it('Login to rancher', () => {
    const username = Cypress.env('USERNAME'); 
    const password = Cypress.env('PASSWORD'); 
    cy.get('#username').type(username)
    cy.get('#password').type(password)
    cy.get('#submit').click()

  })

 })

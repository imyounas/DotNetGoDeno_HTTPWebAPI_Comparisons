interface Todo {
  id: number;
  title: string;
  dueBy: string; 
 isComplete:boolean;
}
  const sampleTodos: Todo[] = [
    { id: 1, title: "Set up a new Deno 2 project", dueBy: new Date().toISOString().split('T')[0] , isComplete: false},
    { id: 2, title: "Install necessary NuGet packages", dueBy: new Date(Date.now() + 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 3, title: "Create solution file", dueBy: new Date(Date.now() + 2 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 4, title: "Add project references for database", dueBy: new Date(Date.now() + 3 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 5, title: "Configure logging services", dueBy: new Date(Date.now() + 4 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 6, title: "Set up dependency injection", dueBy: new Date(Date.now() + 5 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 7, title: "Implement basic user authentication", dueBy: new Date(Date.now() + 6 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 8, title: "Add validation for input fields", dueBy: new Date(Date.now() + 7 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 9, title: "Create unit tests for validation", dueBy: new Date(Date.now() + 8 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 10, title: "Add middleware for error handling", dueBy: new Date(Date.now() + 9 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 11, title: "Design the database schema", dueBy: new Date(Date.now() + 10 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 12, title: "Create the database context", dueBy: new Date(Date.now() + 11 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 13, title: "Implement repository pattern", dueBy: new Date(Date.now() + 12 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 14, title: "Add logging to repository methods", dueBy: new Date(Date.now() + 13 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 15, title: "Create models for API response", dueBy: new Date(Date.now() + 14 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 16, title: "Set up Swagger for API documentation", dueBy: new Date(Date.now() + 15 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 17, title: "Create sample controller", dueBy: new Date(Date.now() + 16 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 18, title: "Add actions for CRUD operations in controller", dueBy: new Date(Date.now() + 17 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 19, title: "Implement data seeding for testing", dueBy: new Date(Date.now() + 18 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 20, title: "Create custom exception class", dueBy: new Date(Date.now() + 19 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 21, title: "Set up JWT authentication", dueBy: new Date(Date.now() + 20 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 22, title: "Test JWT authentication middleware", dueBy: new Date(Date.now() + 21 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 23, title: "Add user roles and permissions", dueBy: new Date(Date.now() + 22 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 24, title: "Create the user registration API", dueBy: new Date(Date.now() + 23 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 25, title: "Implement password hashing", dueBy: new Date(Date.now() + 24 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 26, title: "Create user login API", dueBy: new Date(Date.now() + 25 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 27, title: "Test login functionality", dueBy: new Date(Date.now() + 26 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 28, title: "Create an API endpoint for fetching user profile", dueBy: new Date(Date.now() + 27 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 29, title: "Set up auto-mapping between DTOs and models", dueBy: new Date(Date.now() + 28 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 30, title: "Add logging for user actions", dueBy: new Date(Date.now() + 29 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 31, title: "Implement pagination for list views", dueBy: new Date(Date.now() + 30 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 32, title: "Create model for user preferences", dueBy: new Date(Date.now() + 31 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 33, title: "Add validation to user preference model", dueBy: new Date(Date.now() + 32 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 34, title: "Create API endpoint for updating user preferences", dueBy: new Date(Date.now() + 33 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 35, title: "Test updating user preferences endpoint", dueBy: new Date(Date.now() + 34 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 36, title: "Add caching to frequently accessed endpoints", dueBy: new Date(Date.now() + 35 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 37, title: "Set up caching policies", dueBy: new Date(Date.now() + 36 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 38, title: "Create custom middleware for rate limiting", dueBy: new Date(Date.now() + 37 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 39, title: "Write integration tests for API endpoints", dueBy: new Date(Date.now() + 38 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 40, title: "Implement email sending functionality", dueBy: new Date(Date.now() + 39 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 41, title: "Create templates for email notifications", dueBy: new Date(Date.now() + 40 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 42, title: "Integrate with an SMTP server", dueBy: new Date(Date.now() + 41 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 43, title: "Set up environment variables for email credentials", dueBy: new Date(Date.now() + 42 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 44, title: "Create password reset API", dueBy: new Date(Date.now() + 43 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 45, title: "Test password reset functionality", dueBy: new Date(Date.now() + 44 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 46, title: "Implement 2FA for user authentication", dueBy: new Date(Date.now() + 45 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 47, title: "Set up Azure DevOps pipeline for CI/CD", dueBy: new Date(Date.now() + 46 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 48, title: "Configure build and release pipelines", dueBy: new Date(Date.now() + 47 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 49, title: "Add deployment scripts for staging environment", dueBy: new Date(Date.now() + 48 * 86400000).toISOString().split('T')[0] , isComplete: false},
    { id: 50, title: "Configure database migration in CI/CD pipeline", dueBy: new Date(Date.now() + 49 * 86400000).toISOString().split('T')[0] , isComplete: false}
  ];

  export { sampleTodos };  
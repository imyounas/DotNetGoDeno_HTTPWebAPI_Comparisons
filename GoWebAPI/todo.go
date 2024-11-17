package main

import "time"

type Todo struct {
	Id         int        `json:"id"`
	Title      *string    `json:"title,omitempty"`
	DueBy      *time.Time `json:"dueBy,omitempty"`
	IsComplete bool       `json:"isComplete"`
}

func GetToDos() []Todo {

	now := time.Now()
	sampleTodos := []Todo{
		{Id: 1, Title: strPtr("Set up a new Go 1.23 project"), DueBy: datePtr(now)},
		{Id: 2, Title: strPtr("Install necessary NuGet packages"), DueBy: datePtr(now.AddDate(0, 0, 1))},
		{Id: 3, Title: strPtr("Create solution file"), DueBy: datePtr(now.AddDate(0, 0, 2))},
		{Id: 4, Title: strPtr("Add project references for database"), DueBy: datePtr(now.AddDate(0, 0, 3))},
		{Id: 5, Title: strPtr("Configure logging services"), DueBy: datePtr(now.AddDate(0, 0, 4))},
		{Id: 6, Title: strPtr("Set up dependency injection"), DueBy: datePtr(now.AddDate(0, 0, 5))},
		{Id: 7, Title: strPtr("Implement basic user authentication"), DueBy: datePtr(now.AddDate(0, 0, 6))},
		{Id: 8, Title: strPtr("Add validation for input fields"), DueBy: datePtr(now.AddDate(0, 0, 7))},
		{Id: 9, Title: strPtr("Create unit tests for validation"), DueBy: datePtr(now.AddDate(0, 0, 8))},
		{Id: 10, Title: strPtr("Add middleware for error handling"), DueBy: datePtr(now.AddDate(0, 0, 9))},
		{Id: 11, Title: strPtr("Design the database schema"), DueBy: datePtr(now.AddDate(0, 0, 10))},
		{Id: 12, Title: strPtr("Create the database context"), DueBy: datePtr(now.AddDate(0, 0, 11))},
		{Id: 13, Title: strPtr("Implement repository pattern"), DueBy: datePtr(now.AddDate(0, 0, 12))},
		{Id: 14, Title: strPtr("Add logging to repository methods"), DueBy: datePtr(now.AddDate(0, 0, 13))},
		{Id: 15, Title: strPtr("Create models for API response"), DueBy: datePtr(now.AddDate(0, 0, 14))},
		{Id: 16, Title: strPtr("Set up Swagger for API documentation"), DueBy: datePtr(now.AddDate(0, 0, 15))},
		{Id: 17, Title: strPtr("Create sample controller"), DueBy: datePtr(now.AddDate(0, 0, 16))},
		{Id: 18, Title: strPtr("Add actions for CRUD operations in controller"), DueBy: datePtr(now.AddDate(0, 0, 17))},
		{Id: 19, Title: strPtr("Implement data seeding for testing"), DueBy: datePtr(now.AddDate(0, 0, 18))},
		{Id: 20, Title: strPtr("Create custom exception class"), DueBy: datePtr(now.AddDate(0, 0, 19))},
		{Id: 21, Title: strPtr("Set up JWT authentication"), DueBy: datePtr(now.AddDate(0, 0, 20))},
		{Id: 22, Title: strPtr("Test JWT authentication middleware"), DueBy: datePtr(now.AddDate(0, 0, 21))},
		{Id: 23, Title: strPtr("Add user roles and permissions"), DueBy: datePtr(now.AddDate(0, 0, 22))},
		{Id: 24, Title: strPtr("Create the user registration API"), DueBy: datePtr(now.AddDate(0, 0, 23))},
		{Id: 25, Title: strPtr("Implement password hashing"), DueBy: datePtr(now.AddDate(0, 0, 24))},
		{Id: 26, Title: strPtr("Create user login API"), DueBy: datePtr(now.AddDate(0, 0, 25))},
		{Id: 27, Title: strPtr("Test login functionality"), DueBy: datePtr(now.AddDate(0, 0, 26))},
		{Id: 28, Title: strPtr("Create an API endpoint for fetching user profile"), DueBy: datePtr(now.AddDate(0, 0, 27))},
		{Id: 29, Title: strPtr("Set up auto-mapping between DTOs and models"), DueBy: datePtr(now.AddDate(0, 0, 28))},
		{Id: 30, Title: strPtr("Add logging for user actions"), DueBy: datePtr(now.AddDate(0, 0, 29))},
		{Id: 31, Title: strPtr("Implement pagination for list views"), DueBy: datePtr(now.AddDate(0, 0, 30))},
		{Id: 32, Title: strPtr("Create model for user preferences"), DueBy: datePtr(now.AddDate(0, 0, 31))},
		{Id: 33, Title: strPtr("Add validation to user preference model"), DueBy: datePtr(now.AddDate(0, 0, 32))},
		{Id: 34, Title: strPtr("Create API endpoint for updating user preferences"), DueBy: datePtr(now.AddDate(0, 0, 33))},
		{Id: 35, Title: strPtr("Test updating user preferences endpoint"), DueBy: datePtr(now.AddDate(0, 0, 34))},
		{Id: 36, Title: strPtr("Add caching to frequently accessed endpoints"), DueBy: datePtr(now.AddDate(0, 0, 35))},
		{Id: 37, Title: strPtr("Set up caching policies"), DueBy: datePtr(now.AddDate(0, 0, 36))},
		{Id: 38, Title: strPtr("Create custom middleware for rate limiting"), DueBy: datePtr(now.AddDate(0, 0, 37))},
		{Id: 39, Title: strPtr("Write integration tests for API endpoints"), DueBy: datePtr(now.AddDate(0, 0, 38))},
		{Id: 40, Title: strPtr("Implement email sending functionality"), DueBy: datePtr(now.AddDate(0, 0, 39))},
		{Id: 41, Title: strPtr("Create templates for email notifications"), DueBy: datePtr(now.AddDate(0, 0, 40))},
		{Id: 42, Title: strPtr("Integrate with an SMTP server"), DueBy: datePtr(now.AddDate(0, 0, 41))},
		{Id: 43, Title: strPtr("Set up environment variables for email credentials"), DueBy: datePtr(now.AddDate(0, 0, 42))},
		{Id: 44, Title: strPtr("Create password reset API"), DueBy: datePtr(now.AddDate(0, 0, 43))},
		{Id: 45, Title: strPtr("Test password reset functionality"), DueBy: datePtr(now.AddDate(0, 0, 44))},
		{Id: 46, Title: strPtr("Implement 2FA for user authentication"), DueBy: datePtr(now.AddDate(0, 0, 45))},
		{Id: 47, Title: strPtr("Set up Azure DevOps pipeline for CI/CD"), DueBy: datePtr(now.AddDate(0, 0, 46))},
		{Id: 48, Title: strPtr("Configure build and release pipelines"), DueBy: datePtr(now.AddDate(0, 0, 47))},
		{Id: 49, Title: strPtr("Add deployment scripts for staging environment"), DueBy: datePtr(now.AddDate(0, 0, 48))},
		{Id: 50, Title: strPtr("Configure database migration in CI/CD pipeline"), DueBy: datePtr(now.AddDate(0, 0, 49))},
	}

	return sampleTodos
}

func strPtr(s string) *string {
	return &s
}
func datePtr(t time.Time) *time.Time {
	return &t
}

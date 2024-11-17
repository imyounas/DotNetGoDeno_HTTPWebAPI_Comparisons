using System.Text.Json.Serialization;

namespace DotNet9AOTWebAPI
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateSlimBuilder(args);

            builder.Services.ConfigureHttpJsonOptions(options =>
            {
                options.SerializerOptions.TypeInfoResolverChain.Insert(0, AppJsonSerializerContext.Default);
            });

            var app = builder.Build();

            var sampleTodos = GetTods();

            var todosApi = app.MapGroup("/todos");
            todosApi.MapGet("/", () => sampleTodos);
            todosApi.MapGet("/{id}", (int id) =>
                sampleTodos.FirstOrDefault(a => a.Id == id) is { } todo
                    ? Results.Ok(todo)
                    : Results.NotFound());

            app.Run();
        }

        public static Todo[] GetTods()
        {
            var sampleTodos = new Todo[] {
            new(1, "Set up a new .NET 9 project", DateOnly.FromDateTime(DateTime.Now)),
            new(2, "Install necessary NuGet packages", DateOnly.FromDateTime(DateTime.Now.AddDays(1))),
            new(3, "Create solution file", DateOnly.FromDateTime(DateTime.Now.AddDays(2))),
            new(4, "Add project references for database", DateOnly.FromDateTime(DateTime.Now.AddDays(3))),
            new(5, "Configure logging services", DateOnly.FromDateTime(DateTime.Now.AddDays(4))),
            new(6, "Set up dependency injection", DateOnly.FromDateTime(DateTime.Now.AddDays(5))),
            new(7, "Implement basic user authentication", DateOnly.FromDateTime(DateTime.Now.AddDays(6))),
            new(8, "Add validation for input fields", DateOnly.FromDateTime(DateTime.Now.AddDays(7))),
            new(9, "Create unit tests for validation", DateOnly.FromDateTime(DateTime.Now.AddDays(8))),
            new(10, "Add middleware for error handling", DateOnly.FromDateTime(DateTime.Now.AddDays(9))),
            new(11, "Design the database schema", DateOnly.FromDateTime(DateTime.Now.AddDays(10))),
            new(12, "Create the database context", DateOnly.FromDateTime(DateTime.Now.AddDays(11))),
            new(13, "Implement repository pattern", DateOnly.FromDateTime(DateTime.Now.AddDays(12))),
            new(14, "Add logging to repository methods", DateOnly.FromDateTime(DateTime.Now.AddDays(13))),
            new(15, "Create models for API response", DateOnly.FromDateTime(DateTime.Now.AddDays(14))),
            new(16, "Set up Swagger for API documentation", DateOnly.FromDateTime(DateTime.Now.AddDays(15))),
            new(17, "Create sample controller", DateOnly.FromDateTime(DateTime.Now.AddDays(16))),
            new(18, "Add actions for CRUD operations in controller", DateOnly.FromDateTime(DateTime.Now.AddDays(17))),
            new(19, "Implement data seeding for testing", DateOnly.FromDateTime(DateTime.Now.AddDays(18))),
            new(20, "Create custom exception class", DateOnly.FromDateTime(DateTime.Now.AddDays(19))),
            new(21, "Set up JWT authentication", DateOnly.FromDateTime(DateTime.Now.AddDays(20))),
            new(22, "Test JWT authentication middleware", DateOnly.FromDateTime(DateTime.Now.AddDays(21))),
            new(23, "Add user roles and permissions", DateOnly.FromDateTime(DateTime.Now.AddDays(22))),
            new(24, "Create the user registration API", DateOnly.FromDateTime(DateTime.Now.AddDays(23))),
            new(25, "Implement password hashing", DateOnly.FromDateTime(DateTime.Now.AddDays(24))),
            new(26, "Create user login API", DateOnly.FromDateTime(DateTime.Now.AddDays(25))),
            new(27, "Test login functionality", DateOnly.FromDateTime(DateTime.Now.AddDays(26))),
            new(28, "Create an API endpoint for fetching user profile", DateOnly.FromDateTime(DateTime.Now.AddDays(27))),
            new(29, "Set up auto-mapping between DTOs and models", DateOnly.FromDateTime(DateTime.Now.AddDays(28))),
            new(30, "Add logging for user actions", DateOnly.FromDateTime(DateTime.Now.AddDays(29))),
            new(31, "Implement pagination for list views", DateOnly.FromDateTime(DateTime.Now.AddDays(30))),
            new(32, "Create model for user preferences", DateOnly.FromDateTime(DateTime.Now.AddDays(31))),
            new(33, "Add validation to user preference model", DateOnly.FromDateTime(DateTime.Now.AddDays(32))),
            new(34, "Create API endpoint for updating user preferences", DateOnly.FromDateTime(DateTime.Now.AddDays(33))),
            new(35, "Test updating user preferences endpoint", DateOnly.FromDateTime(DateTime.Now.AddDays(34))),
            new(36, "Add caching to frequently accessed endpoints", DateOnly.FromDateTime(DateTime.Now.AddDays(35))),
            new(37, "Set up caching policies", DateOnly.FromDateTime(DateTime.Now.AddDays(36))),
            new(38, "Create custom middleware for rate limiting", DateOnly.FromDateTime(DateTime.Now.AddDays(37))),
            new(39, "Write integration tests for API endpoints", DateOnly.FromDateTime(DateTime.Now.AddDays(38))),
            new(40, "Implement email sending functionality", DateOnly.FromDateTime(DateTime.Now.AddDays(39))),
            new(41, "Create templates for email notifications", DateOnly.FromDateTime(DateTime.Now.AddDays(40))),
            new(42, "Integrate with an SMTP server", DateOnly.FromDateTime(DateTime.Now.AddDays(41))),
            new(43, "Set up environment variables for email credentials", DateOnly.FromDateTime(DateTime.Now.AddDays(42))),
            new(44, "Create password reset API", DateOnly.FromDateTime(DateTime.Now.AddDays(43))),
            new(45, "Test password reset functionality", DateOnly.FromDateTime(DateTime.Now.AddDays(44))),
            new(46, "Implement 2FA for user authentication", DateOnly.FromDateTime(DateTime.Now.AddDays(45))),
            new(47, "Set up Azure DevOps pipeline for CI/CD", DateOnly.FromDateTime(DateTime.Now.AddDays(46))),
            new(48, "Configure build and release pipelines", DateOnly.FromDateTime(DateTime.Now.AddDays(47))),
            new(49, "Add deployment scripts for staging environment", DateOnly.FromDateTime(DateTime.Now.AddDays(48))),
            new(50, "Configure database migration in CI/CD pipeline", DateOnly.FromDateTime(DateTime.Now.AddDays(49)))
            };


            return sampleTodos;
        }
    }



    public record Todo(int Id, string? Title, DateOnly? DueBy = null, bool IsComplete = false);

    [JsonSerializable(typeof(Todo[]))]
    internal partial class AppJsonSerializerContext : JsonSerializerContext
    {

    }
}

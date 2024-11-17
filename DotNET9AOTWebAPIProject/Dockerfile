# Stage 1: Base image for building the application
FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build

# Install clang and dependencies required for AOT compilation
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    clang \
    gcc \
    g++ \
    zlib1g-dev \
    libcurl4-openssl-dev \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /src

# Copy the solution and project files into the container
COPY ["DotNet9AOTWebAPI/DotNet9AOTWebAPI.csproj", "DotNet9AOTWebAPI/"]

# Restore the dependencies for the project
RUN dotnet restore "./DotNet9AOTWebAPI/DotNet9AOTWebAPI.csproj"

# Copy all other source files into the container
COPY . .

# Set the working directory to the project folder
WORKDIR "/src/DotNet9AOTWebAPI"

# Build the project in Release configuration
RUN dotnet build "./DotNet9AOTWebAPI.csproj" -c Release -o /app/build

# Stage 2: Publish the application with AOT enabled
FROM build AS publish

# Publish the application with AOT compilation enabled
RUN dotnet publish "./DotNet9AOTWebAPI.csproj" -c Release -o /app/publish /p:PublishAot=true /p:UseAppHost=true

# Stage 3: Create the final runtime image
FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS final

# Set the working directory inside the final container
WORKDIR /app

# Expose port 8080 for the Web API
EXPOSE 8080

# Copy the published output from the publish stage
COPY --from=publish /app/publish .

# Set the entry point to the AOT-compiled application
ENTRYPOINT ["./DotNet9AOTWebAPI"]

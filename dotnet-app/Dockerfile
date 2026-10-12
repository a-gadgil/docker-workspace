# ---------------------------------------------------
# Stage 1: Build & Publish
# ---------------------------------------------------
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src

# Copy project file and restore dependencies
COPY FreshApp.csproj ./
RUN dotnet restore

# Copy remaining source code and publish
COPY . .
RUN dotnet publish FreshApp.csproj -c Release -o /app/publish /p:UseAppHost=false

# ---------------------------------------------------
# Stage 2: Runtime Environment
# ---------------------------------------------------
FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS final
WORKDIR /app

# Run as non-root security user
USER $APP_UID

# Copy published artifacts from build stage
COPY --from=build /app/publish .

# Expose port and set entrypoint
EXPOSE 8080
ENV ASPNETCORE_URLS=http://+:8080

ENTRYPOINT ["dotnet", "FreshApp.dll"]
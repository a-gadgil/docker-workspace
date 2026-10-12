var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Hello from the Fresh .NET 8 DevOps Lab!");

app.Run();
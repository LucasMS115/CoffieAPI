### Initial Instructions:

I would like to design the back-end of the project as first step. 

Before writing any code, we will first define the architecture and project structure. Do NOT generate any code yet.

## Project Context

We are creating a RESTful API focused on **manage cofee recipes**. Our target audience are cofee hobbists.

The application will use Golang for all the back-end services.

The goal is to build a clean, maintainable, production-like architecture that follows modern best practices for Go applications.

This project will prioritize:

- clear separation of concerns
- testability
- extensibility

I do not have experience with go. 
My main languages nowdays are Javascript, Java and Clojure.
I need you to include the rationale behind some decisions of frameworks and main code usage, ideally, create paralels with the languages I mentioned.

IMPORTANT: this is an MVP, lets use the minimum information to make the core features work in our api.

## Architecture Guidelines

Follow these architectural principles:

### 1. **Clean Monolitic Architecture**
Use naming conventios and file organization commonly used for Golang services for the project structre.

If possible, organize the project by feature (feature-based architecture), we might break it into hexagonal microsservices later.


---

### 2. **Layer separation**

Keep a clean separation among database, business rules and api.
Endpoints must never expose domain entities directly.

You can use something like:
- incoming requests schemas
- outgoing responses schemas
- domain schemas
- database entities schemas
- adapters from one to another

--- 


### 3. **Thin controllers**

In Java terms, controllers should only:
- receive HTTP requests
- validate inputs (we need a validation layer)
- delegate to services
- return responses
Business logic belongs only in the service layer.

### 4. **Testing strategy**
We will implement two types of tests:
- **Unit tests**
    - focus on service logic
    - use mocking for dependencies
    - write the code in a given, when, then format. 
    - Use the principles of equivalence class testing to generate test cases
    - Use the FIRST principles
- **Integration tests**
    - test full HTTP flow with database mocked
    - validate API behavior (request -> response)

### 5. **Clean naming**
Use descriptive names for directories, files and functions.
Avoid generic names like:
- Utils
- Helpers
- Misc
And abreviations like (use full descriptive names):
ctx -> requestContext
req -> registerRequest, createRecipeRequest, createRatingRequest
db -> databaseConnection
mux -> serveMux
s -> service ou postgres...Store
u, c, r -> user, coffee, recipe, rating
dsn -> dataSourceName
addr -> address


### 8. Error handling 
Our RESTful apis need to have correct HTTP error codes
Error responses need to be informative, with error code, meaningful message, and related fields.

### 9. Database
We should have clean entities and there relations. We can have a normalized database to start.

### 10. Idiomatic Go Naming

Always use idiomatic Go terms for naming — avoid DDD/jargon from other ecosystems. Examples:
- **`Store`** instead of `Repository` (e.g., `RecipeStore`, `UserStore`)
- **`DB`** / **`Database`** for data access (e.g., `recipeDB`)
- **`Handler`** instead of `Controller`
- Prefer short, clear names: `pkg.Store` over `pkg.UserRepositoryInterface`
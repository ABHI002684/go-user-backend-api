
---

# 📄 `Reasoning.md` 


# Reasoning & Learning Approach

## Background

Before this assignment, my backend development experience was primarily in **Node.js**.  
I had not worked with **Go** for backend development prior to this task.

The goal was to implement the required functionality while quickly adapting to Go’s syntax, tooling, and backend design patterns.

---

## Approach

I approached this project by mapping familiar Node.js concepts to Go equivalents:

- Express.js → Fiber
- Controllers → Handlers
- Services → Service layer
- ORM / Query builders → SQLC
- dotenv → godotenv
- Middleware patterns → Fiber middleware

This helped me focus on system design while learning Go incrementally.

---

## Architecture & Design

The application follows a **layered architecture**:

- **Handler Layer**  
  Handles HTTP requests, input validation, and responses.

- **Service Layer**  
  Contains business logic such as parsing the DOB string and calculating the user’s age.

- **Repository Layer**  
  Handles database interactions using SQLC-generated queries.

This separation of concerns improves readability, maintainability, and testability.

---

## Age Calculation

The user’s age is **not stored** in the database.  
Instead, it is calculated dynamically using Go’s `time` package when user data is retrieved.

This avoids data inconsistency and ensures the age is always accurate.

---

## SQLC & Type Safety

SQLC was used to generate type-safe Go code from SQL queries.  
Compared to typical Node.js ORMs, SQLC enforces correctness at compile time, which helped me better understand Go’s strong type system.

---

## Configuration Management

Application configuration is handled using environment variables loaded from a `.env` file.  
This follows the same best practices I use in Node.js applications and avoids hardcoding sensitive information.

---

## Trade-offs & Limitations

Due to time constraints and the learning curve:
- Authentication and authorization were not implemented
- Pagination and filtering were omitted
- Focus was placed on correctness, clarity, and clean architecture

---

## Future Improvements

With more time, I would:
- Add unit tests for the service layer
- Implement pagination for list endpoints
- Add Docker support
- Improve error handling and observability

---

## Conclusion

This project demonstrates my ability to quickly transition from **Node.js to Go**, apply core backend principles, and build a clean, maintainable API while learning a new language and ecosystem.

---

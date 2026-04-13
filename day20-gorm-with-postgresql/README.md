# Day 20: GORM with PostgreSQL

This directory demonstrates how to use **GORM**, the most popular Object-Relational Mapper (ORM) for Go, to simplify database interactions.

## 📝 Key Concepts Covered

- **GORM Foundations**:
  - `gorm.Open`: Establishing a connection with the PostgreSQL driver.
  - **Auto-Migration**: Using `DB.AutoMigrate(&User{})` to automatically create or update database tables from Go structs.
- **ORM CRUD Operations**:
  - **Create**: `DB.Create(&user)` to insert records.
  - **Read**: `DB.Find(&users)` for all records, and `DB.First(&user, id)` for a single record.
  - **Update**: `DB.Save(&user)` to update an entire record.
  - **Delete**: `DB.Delete(&user)` to remove a record.
- **Struct Tags for GORM**: Using tags like `gorm:"primaryKey;autoIncrement"` and `gorm:"uniqueIndex"` to define database constraints.
- **Gin Integration**: Combining Gin's routing with GORM's database logic for a fully database-driven REST API.

## 📂 Files

- [main.go](main.go): A comprehensive example of a Gin API backed by GORM and PostgreSQL.

## 🚀 How to Run

1.  Configure your `.env` file with `DATABASE_URL`.
2.  Run the server:
    ```bash
    go run day20-gorm-with-postgresql/main.go
    ```
3.  The application will automatically create the `users` table upon startup!

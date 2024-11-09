# Secret-Note-API-SPA-Marwan-Radwan

Create a web application that allows users to securely share self-destructing secret notes.

## Functionalities:

### 1. Note Creation:

- Users can create a note with text content
- Generate a unique, secure URL for each note
- Set an expiration time or number of views before destruction

### 2. Note Retrieval:

- Users can view a note using the unique URL
- After viewing or upon expiration, the note is permanently deleted

### 3. Security Features:

- Use secure random generation for URLs (should be able to browse /note/1, /note/2, etc. The URLs shouldn't be predictable)

### 4. Backend:

- Swagger docs for the API
- Develop a restful API using Go
- Use a database (e.g., sqlite or postgresql) for storing the notes

### 5. Frontend:

- Create a simple, responsive UI using Vue

# How to Run

```golang
docker compose up
```

when you finnish

```golang
docker compose up
```

<p>
Frontend will run port 8080
</p>
Backend will run on port 3000

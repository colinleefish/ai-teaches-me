# GMDB - Go Movie Database 🎬

A Django-inspired movie database API built with Gin + GORM + PostgreSQL.

## 🎯 Features
- Movies, actors, and awards management
- Many-to-many relationships between entities
- RESTful API endpoints
- Database migrations and seeding

## 🗄️ Database
```
postgres://gmdb_user:gmdb_pass@localhost:5432/gmdb_db
```

## 📁 Project Structure

```
gmdb/
├── main.go                 # Application entry point
├── config/
│   └── database.go         # Database connection & config
├── models/
│   ├── movie.go            # Movie model with associations
│   ├── actor.go            # Actor model with associations  
│   ├── award.go            # Award model (movies/actors)
│   └── associations.go     # M2N relationship definitions
├── handlers/
│   ├── movie_handlers.go   # Movie CRUD endpoints
│   ├── actor_handlers.go   # Actor CRUD endpoints
│   └── award_handlers.go   # Award CRUD endpoints
├── routes/
│   └── routes.go           # Route definitions & middleware
├── services/
│   ├── movie_service.go    # Business logic for movies
│   ├── actor_service.go    # Business logic for actors
│   └── award_service.go    # Business logic for awards
├── utils/
│   ├── response.go         # Standard API responses
│   └── validation.go       # Custom validation helpers
└── migrations/
    └── seed.go             # Sample data for testing
```

## 🔗 Entity Relationships

### Models
- **Movie**: id, title, director, year, genre, rating, description
- **Actor**: id, name, birth_date, nationality, biography
- **Award**: id, name, category, year, description

### Many-to-Many Relationships
- **MovieActors**: movies ↔ actors (with role field)
- **MovieAwards**: movies ↔ awards  
- **ActorAwards**: actors ↔ awards

## 🚀 API Endpoints

### Movies
- `GET /api/v1/movies` - List all movies (with actors, awards)
- `POST /api/v1/movies` - Create movie
- `GET /api/v1/movies/:id` - Get movie details
- `PUT /api/v1/movies/:id` - Update movie
- `DELETE /api/v1/movies/:id` - Delete movie
- `POST /api/v1/movies/:id/actors` - Add actor to movie
- `DELETE /api/v1/movies/:id/actors/:actor_id` - Remove actor from movie

### Actors  
- `GET /api/v1/actors` - List all actors (with movies, awards)
- `POST /api/v1/actors` - Create actor
- `GET /api/v1/actors/:id` - Get actor details
- `PUT /api/v1/actors/:id` - Update actor
- `DELETE /api/v1/actors/:id` - Delete actor

### Awards
- `GET /api/v1/awards` - List all awards
- `POST /api/v1/awards` - Create award
- `GET /api/v1/awards/:id` - Get award details
- `PUT /api/v1/awards/:id` - Update award
- `DELETE /api/v1/awards/:id` - Delete award

## 📋 Implementation Checklist

### Phase 1: Foundation
- [ ] `config/database.go` - Database connection
- [ ] `models/movie.go` - Movie model
- [ ] `models/actor.go` - Actor model  
- [ ] `models/award.go` - Award model
- [ ] `models/associations.go` - M2N relationships

### Phase 2: Basic CRUD
- [ ] `handlers/movie_handlers.go` - Movie endpoints
- [ ] `handlers/actor_handlers.go` - Actor endpoints
- [ ] `handlers/award_handlers.go` - Award endpoints
- [ ] `routes/routes.go` - Route setup
- [ ] `main.go` - Wire everything together

### Phase 3: Business Logic
- [ ] `services/movie_service.go` - Movie business logic
- [ ] `services/actor_service.go` - Actor business logic
- [ ] `services/award_service.go` - Award business logic
- [ ] `utils/response.go` - Standardized responses
- [ ] `utils/validation.go` - Custom validations

### Phase 4: Advanced Features
- [ ] Association endpoints (add/remove actors from movies)
- [ ] Query filtering and sorting
- [ ] `migrations/seed.go` - Sample data
- [ ] Error handling middleware
- [ ] Request logging middleware

## 🧪 Testing Commands

```bash
# Start server
go run .

# Test endpoints
curl http://localhost:8080/api/v1/movies
curl -X POST http://localhost:8080/api/v1/movies -H "Content-Type: application/json" -d '{...}'
```


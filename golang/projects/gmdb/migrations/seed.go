package migrations

import (
	"log"
	"time"

	"gmdb/models"
	"gmdb/utils"

	"gorm.io/gorm"
)

// SeedData populates the database with sample data
func SeedData(db *gorm.DB) error {
	log.Println("Seeding database with sample data...")

	// Check if data already exists
	var actorCount int64
	if err := db.Model(&models.Actor{}).Count(&actorCount).Error; err != nil {
		return err
	}

	if actorCount > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Sample birth dates
	birthDate1, _ := time.Parse("2006-01-02", "1956-07-09")
	birthDate2, _ := time.Parse("2006-01-02", "1974-04-04")
	birthDate3, _ := time.Parse("2006-01-02", "1967-04-07")
	// Create actors
	actors := []models.Actor{
		{
			ID:        utils.NewUUIDv7(),
			Name:      "Tom Hanks",
			BirthDate: &birthDate1,
			Biography: "Thomas Jeffrey Hanks is an American actor and filmmaker known for his comedic and dramatic roles.",
		},
		{
			ID:        utils.NewUUIDv7(),
			Name:      "Heath Ledger",
			BirthDate: &birthDate2,
			Biography: "Heath Andrew Ledger was an Australian actor, photographer, and music video director.",
		},
		{
			ID:        utils.NewUUIDv7(),
			Name:      "Russell Crowe",
			BirthDate: &birthDate3,
			Biography: "Russell Ira Crowe is an actor, director and musician known for his fiery temper and intense performances.",
		},
	}

	// Create movies
	movies := []models.Movie{
		{
			ID:          utils.NewUUIDv7(),
			Title:       "Forrest Gump",
			Year:        1994,
			Director:    "Robert Zemeckis",
			Genre:       "Drama",
			Description: "The presidencies of Kennedy and Johnson, the events of Vietnam, Watergate and other historical events unfold through the perspective of an Alabama man.",
			Rating:      8.8,
		},
		{
			ID:          utils.NewUUIDv7(),
			Title:       "The Dark Knight",
			Year:        2008,
			Director:    "Christopher Nolan",
			Genre:       "Action",
			Description: "When the menace known as the Joker wreaks havoc on Gotham, Batman must accept one of the greatest psychological and physical tests.",
			Rating:      9.0,
		},
		{
			ID:          utils.NewUUIDv7(),
			Title:       "Gladiator",
			Year:        2000,
			Director:    "Ridley Scott",
			Genre:       "Action",
			Description: "A former Roman General sets out to exact vengeance against the corrupt emperor who murdered his family.",
			Rating:      8.5,
		},
	}

	// Create awards
	awards := []models.Award{
		{
			ID:          utils.NewUUIDv7(),
			Name:        "Academy Award for Best Actor",
			Category:    "Best Actor",
			Year:        1995,
			ActorID:     &actors[0].ID, // Tom Hanks
			MovieID:     &movies[0].ID, // Forrest Gump
			Description: "Won for outstanding performance as Forrest Gump",
		},
		{
			ID:          utils.NewUUIDv7(),
			Name:        "Academy Award for Best Supporting Actor",
			Category:    "Best Supporting Actor",
			Year:        2009,
			ActorID:     &actors[1].ID, // Heath Ledger
			MovieID:     &movies[1].ID, // The Dark Knight
			Description: "Posthumously won for his iconic portrayal of the Joker",
		},
		{
			ID:          utils.NewUUIDv7(),
			Name:        "Academy Award for Best Actor",
			Category:    "Best Actor",
			Year:        2001,
			ActorID:     &actors[2].ID, // Russell Crowe
			MovieID:     &movies[2].ID, // Gladiator
			Description: "Won for his powerful performance as Maximus",
		},
	}

	// Insert data
	if err := db.Create(&actors).Error; err != nil {
		return err
	}
	log.Printf("Created %d actors", len(actors))

	if err := db.Create(&movies).Error; err != nil {
		return err
	}
	log.Printf("Created %d movies", len(movies))

	if err := db.Create(&awards).Error; err != nil {
		return err
	}
	log.Printf("Created %d awards", len(awards))

	// Create many-to-many relationships (actor-movie associations)
	// Tom Hanks in Forrest Gump
	if err := db.Model(&actors[0]).Association("Movies").Append(&movies[0]); err != nil {
		return err
	}

	// Heath Ledger in The Dark Knight
	if err := db.Model(&actors[1]).Association("Movies").Append(&movies[1]); err != nil {
		return err
	}

	// Russell Crowe in Gladiator
	if err := db.Model(&actors[2]).Association("Movies").Append(&movies[2]); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully!")
	return nil
}

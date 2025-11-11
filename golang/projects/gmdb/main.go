package main

import (
	"fmt"
	"log"
	"os"

	"gmdb/config"
	"gmdb/handlers"
	"gmdb/migrations"
	"gmdb/routes"
	"gmdb/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "gmdb",
	Short: "GMDB - Go Movie Database API",
	Long:  "A REST API for managing movies, actors, and awards built with Go and Gin.",
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		if err := config.LoadConfig(configFile); err != nil {
			log.Fatal("Failed to load config:", err)
		}

		// Connect to database
		config.ConnectDB()

		// Initialize services
		db := config.GetDB()
		actorService := services.NewActorService(db)

		// Initialize handlers with services
		handlers.InitActorHandlers(actorService)

		// Set Gin mode based on environment
		if config.GlobalConfig.App.Environment == "production" {
			gin.SetMode(gin.ReleaseMode)
		}

		// Initialize Gin router
		r := gin.Default()

		// Setup routes
		routes.SetupRoutes(r)

		// Start server
		serverAddr := fmt.Sprintf("%s:%d",
			config.GlobalConfig.Server.Host,
			config.GlobalConfig.Server.Port,
		)

		log.Printf("Starting %s v%s server on %s",
			config.GlobalConfig.App.Name,
			config.GlobalConfig.App.Version,
			serverAddr,
		)

		if err := r.Run(serverAddr); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with sample data",
	Long:  "Populates the database with sample actors, movies, and awards for testing purposes.",
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		if err := config.LoadConfig(configFile); err != nil {
			log.Fatal("Failed to load config:", err)
		}

		// Connect to database (this also runs migrations)
		config.ConnectDB()

		// Seed the database
		if err := migrations.SeedData(config.GetDB()); err != nil {
			log.Fatal("Failed to seed database:", err)
		}

		fmt.Println("Database seeded successfully! 🎉")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "gmdb.yaml", "config file path")

	// Add subcommands
	rootCmd.AddCommand(seedCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

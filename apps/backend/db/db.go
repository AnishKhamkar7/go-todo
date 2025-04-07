package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AnishKhamkar7/todo-api/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbConfig() {
	err := godotenv.Load("../../.env") // ✅ Adjust the path based on where your .env file is
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println(databaseURL)

	if databaseURL == "" {
		log.Fatalf("❌ DATABASE_URL environment variable not set")
	}

	client, err := ent.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to open database: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Database migration successful!")

}

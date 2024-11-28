package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	// Connect to MySQL
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_api?parseTime=true&loc=Local")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test DB Connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Set up Gin router
	r := gin.Default()

	// Define routes
	r.GET("/current-time", getCurrentTime)
	r.GET("/logged-times", getLoggedTimes)

	// Start the server
	r.Run(":8080") // Default port 8080
}

func getCurrentTime(c *gin.Context) {
	// Get current time in Toronto timezone
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Printf("Error loading timezone: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load timezone"})
		return
	}
	currentTime := time.Now().In(loc)

	// Log time to database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log time"})
		return
	}

	// Respond with JSON
	c.JSON(http.StatusOK, TimeResponse{CurrentTime: currentTime.Format(time.RFC3339)})
}

func getLoggedTimes(c *gin.Context) {
	rows, err := db.Query("SELECT timestamp FROM time_log")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve logged times"})
		return
	}
	defer rows.Close()

	var times []string
	for rows.Next() {
		var timestamp time.Time
		err = rows.Scan(&timestamp)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		times = append(times, timestamp.Format(time.RFC3339))
	}

	c.JSON(http.StatusOK, gin.H{"logged_times": times})
}

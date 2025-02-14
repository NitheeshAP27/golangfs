package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "bank_app_db"
var mongoCollectionBank string = "bank_details"

// Database variables
var mongoclient *mongo.Client
var bankCollection *mongo.Collection

// Model BankDetails for Collection "bank_details"
type BankDetails struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HolderName string             `json:"holder_name" bson:"holder_name"`
	PhoneNo   string             `json:"phone_no" bson:"phone_no"`
	AccountType string            `json:"account_type" bson:"account_type"`
}

// Connect to MongoDB
func connectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoclient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	bankCollection = mongoclient.Database(mongoDbName).Collection(mongoCollectionBank)
	fmt.Println("Connected to MongoDB!")
}

// POST /bank
func createBankDetails(c *gin.Context) {
	var bankDetails BankDetails

	// Bind JSON body to bankDetails
	if err := c.BindJSON(&bankDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert bank details into MongoDB
	result, err := bankCollection.InsertOne(ctx, bankDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bank details"})
		return
	}

	// Extract the inserted ID
	bankId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse inserted ID"})
		return
	}
	bankDetails.ID = bankId

	// Return created bank details
	c.JSON(http.StatusCreated, gin.H{
		"message": "Bank details created successfully",
		"bank":    bankDetails,
	})
}

// GET /bank
func readAllBankDetails(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := bankCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bank details"})
		return
	}
	defer cursor.Close(ctx)

	// Ensure bank details is an empty slice, not nil
	bankDetailsList := []BankDetails{}
	if err := cursor.All(ctx, &bankDetailsList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse bank details"})
		return
	}

	c.JSON(http.StatusOK, bankDetailsList)
}

// GET /bank/:id
func readBankDetailsById(c *gin.Context) {
	id := c.Param("id")

	// Convert string ID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var bankDetails BankDetails
	err = bankCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&bankDetails)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank details not found"})
		return
	}

	c.JSON(http.StatusOK, bankDetails)
}

// PUT /bank/:id
func updateBankDetails(c *gin.Context) {
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedBankDetails BankDetails
	if err := c.BindJSON(&updatedBankDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := bankCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updatedBankDetails})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bank details"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank details not found"})
		return
	}

	// Return updated bank details
	c.JSON(http.StatusOK, gin.H{
		"message": "Bank details updated successfully",
		"bank":    updatedBankDetails,
	})
}

// DELETE /bank/:id
func deleteBankDetails(c *gin.Context) {
	id := c.Param("id")
	// Convert string ID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := bankCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bank details"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank details not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank details deleted successfully"})
}

func main() {
	// Connect to MongoDB
	connectDB()

	// Set up Gin router
	r := gin.Default()
	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Routes
	r.POST("/bank", createBankDetails)
	r.GET("/bank", readAllBankDetails)
	r.GET("/bank/:id", readBankDetailsById)
	r.PUT("/bank/:id", updateBankDetails)
	r.DELETE("/bank/:id", deleteBankDetails)

	// Start server
	r.Run(":8080")
}

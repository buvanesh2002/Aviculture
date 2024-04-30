package config

import (
	"context"
	"log"

	//"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var LoginCollection *mongo.Collection
var AddFlockCollection *mongo.Collection
var AddEntryCollection *mongo.Collection
var AddRemainderCollection *mongo.Collection
var AddFlockEntriesCollection *mongo.Collection
var AddShopCollection *mongo.Collection
var AddCartCollection *mongo.Collection
var AddCustomerCollection *mongo.Collection
var OrderCollection *mongo.Collection

func init() {
	// mongourl := viper.GetString("mongourl")
	mongourl := "mongodb+srv://rohith:rohith@cluster0.cgwpnv8.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	log.Println("================", mongourl)
	Ctx := context.Background()
	defer Ctx.Done()
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(mongourl))
	if err != nil {
		log.Println("Error connecting to", err)
	} else {
		log.Println("mongo connected successfully")
	}
	if err := client.Ping(Ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB server:", err)
	}
	log.Println("Connected to MongoDB successfully")
	//db := viper.GetString("db")

	// LoginCollection = client.Database(viper.GetString("db")).Collection("users")
	// AddFlockCollection = client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	// AddEntryCollection = client.Database(db).Collection(viper.GetString("AddEntry"))
	// AddRemainderCollection = client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))
	// AddFlockEntriesCollection = client.Database(viper.GetString("db")).Collection(viper.GetString("AddflockEntries"))
	// AddShopCollection = client.Database(viper.GetString("db")).Collection(viper.GetString("AddShop"))
	// AddCartCollection = client.Database(viper.GetString("db")).Collection(viper.GetString("AddCart"))
	LoginCollection = client.Database("aviculture").Collection("users")
	AddFlockCollection = client.Database("aviculture").Collection("Flock")
	AddEntryCollection = client.Database("aviculture").Collection("Entries")
	AddRemainderCollection = client.Database("aviculture").Collection("Reminder")
	//AddFlockEntriesCollection = client.Database("aviculture").Collection(viper.GetString("AddflockEntries"))
	AddShopCollection = client.Database("aviculture").Collection("Shop")
	AddCartCollection = client.Database("aviculture").Collection("Cart")
	AddCustomerCollection = client.Database("aviculture").Collection("Customer")
	OrderCollection = client.Database("aviculture").Collection("orders")
}

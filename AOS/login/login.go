package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"log"
	"run/config"
	"run/dto"

	"time"

	//"github.com/aws/aws-sdk-go/aws/client"
	// "github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func Login(value dto.Logindata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  login service +++++++++++++++++++++++++")

	log.Println(value)
	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())

	// collection := Client.Database(viper.GetString("db")).Collection("users")

	filter := bson.M{"email": value.Email, "password": value.Password}

	err := config.LoginCollection.FindOne(context.Background(), filter).Decode(&value)
	log.Println(err)
	if err != nil {
		return "Invalid Credentials", err
	} else {
		res := AgeCalculator()
		log.Println("++++++++++++++++res++++++++++", res)
		return "Login successful", nil
	}
}

func AddFlock(value dto.Flockdata) (string, error) {

	log.Println("++++++++++++++++++++++++++++  AddFlock service +++++++++++++++++++++++++")
	randomID, err := generateRandomID(10)
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}
	value.ID = randomID

	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())

	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))

	data := bson.M{
		"_id":       value.ID,
		"flockName": value.FlockName,
		"breedName": value.BreedName,
		"startDate": value.StartDate,
		"startAge":  value.StartAge,
		//	"age":          value.Age,
		"openingBirds": value.NoBirds,
		"shedNumber":   value.ShedNumber,
		"active":       value.Active,
		"createdAt":    value.CreatedAt,
		"updatedAt":    value.UpdatedAt,
		"image":        value.Image,
	}

	_, err = config.AddFlockCollection.InsertOne(context.Background(), data)
	if err != nil {
		log.Println("Error inserting document:", err)
		return "", err
	}

	return "Flock added successfully", nil
}

func ListFlock() *[]dto.Flockdata {
	AgeCalculator()
	log.Println("===============List Credentials===============")
	info := []dto.Flockdata{}
	// Client := config.GetConfig()
	// db := viper.GetString("db")
	// log.Println("db name : ", db)
	// // this is also add flock collection
	// collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	cur, err := config.AddFlockCollection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
	}
	// defer Client.Disconnect(context.Background())
	defer cur.Close(context.Background())

	err = cur.All(context.Background(), &info)
	if err != nil {
		log.Println("Error while fetching documents:", err)
		return nil
	}
	return &info
}

// func ListFlockbyid(id string) (dto.Flockdata, error) {
// 	log.Println("===============ListFlockbyid Credentials===============")

// 	var info dto.Flockdata
// 	id = "759c91c579"
// 	Client := config.GetConfig()
// 	db := viper.GetString("db")

// 	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
// 	result, _ := collection.CountDocuments(context.Background(), &info)
// 	// log.Println("Collection list error : ", result.Err())
// 	// return dto.Flockdata{}, result.Err()
// 	log.Println("buvanesh-----", result)

// 	// if err := result.Decode(&info); err != nil {
// 	//     log.Println("Error while decoding document:", err)
// 	//     return dto.Flockdata{}, err
// 	// }

// 	return info, nil

// }

func ListFlockbyid(id string) (dto.Flockdata, error) {
	log.Println("===============ListFlockbyid Credentials===============")

	var info dto.Flockdata
	// Client := config.GetConfig()
	// db := viper.GetString("db")

	// collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	filter := bson.M{"_id": id}
	result := config.AddFlockCollection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		log.Println("Error while fetching document:", result.Err())
		return dto.Flockdata{}, result.Err()
	}

	if err := result.Decode(&info); err != nil {
		log.Println("Error while decoding document:", err)
		return dto.Flockdata{}, err
	}

	return info, nil
}
func DailyEntry(value dto.DailyEntry) (string, error) {
	log.Println("..............", value)

	log.Println("============= Daily Entry =================")
	ctx := context.TODO()
	defer ctx.Done()
	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())
	// db := viper.GetString("db")
	// collection := Client.Database(db).Collection(viper.GetString("AddEntry"))
	_, err := config.AddEntryCollection.InsertOne(ctx, value)
	if err != nil {
		log.Println("Error Inserting Document:", err)
		return " ", err
	}
	UpdateEntry(value)
	UpdateFlockEntries(value)
	log.Println(value)
	return "Entry added", nil
}
func UpdateFlock(value dto.Flockdata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  UpdateFlock service +++++++++++++++++++++++++")

	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	filter := bson.M{"_id": value.ID}
	log.Println("active=", value.Active)
	log.Println("ID=", value.ID)
	update := bson.M{
		"$set": bson.M{
			"flockName":    value.FlockName,
			"breedName":    value.BreedName,
			"startDate":    value.StartDate,
			"startAge":     value.StartAge,
			"active":       value.Active,
			"openingBirds": value.NoBirds,
			"shedNumber":   value.ShedNumber,
			"updatedAt":    time.Now().String(),
		},
	}

	_, err := config.AddFlockCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating document:", err)
		return "", err
	}

	return "Flock updated successfully", nil
}

func generateRandomID(length int) (string, error) {
	byteSize := length / 2
	if length%2 != 0 {
		byteSize++
	}

	bytes := make([]byte, byteSize)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	id := hex.EncodeToString(bytes)

	id = id[:length]

	return id, nil
}

func AgeCalculator() string {
	log.Println("===============Age Update called===============")
	var (
		flockarray  []dto.Flockdata
		createdDate time.Time
		duration    time.Duration
		days        int
		filter      bson.M
		update      bson.M
	)

	// Client := config.GetConfig()
	// db := viper.GetString("db")
	// collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	Date := time.Now()
	ctx := context.TODO()
	search := bson.M{}
	cursor, err := config.AddFlockCollection.Find(ctx, search)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var flock dto.Flockdata
		if err := cursor.Decode(&flock); err != nil {
			log.Println("Error decoding document:", err)
			continue
		}
		flockarray = append(flockarray, flock)
	}
	for _, flock := range flockarray {
		createdDate, _ = time.Parse("2006-01-02", flock.StartDate)
		duration = Date.Sub(createdDate)
		days = (flock.StartAge)
		days = days + int(duration.Hours()/24)
		filter = bson.M{"_id": flock.ID}
		update = bson.M{
			"$set": bson.M{
				"age": int32(days),
			},
		}
		_, err := config.AddFlockCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println("Error updating document:", err)
		}
	}
	return "success"

}

func AddReminder(value dto.Reminder) (string, error) {
	log.Println("++++++++++++++++++++++++++++  AddReminder service +++++++++++++++++++++++++")
	randomID, err := generateRandomID(5)
	if err != nil {
		log.Println("Error:", err)
		return "", err
	}
	value.ReminderId = randomID
	value.Status = "true"
	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())

	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))

	_, err1 := config.AddRemainderCollection.InsertOne(context.Background(), value)
	if err1 != nil {
		log.Println("Error inserting document:", err1)
		return "", err1
	}
	return "Reminder added successfully", nil
}

func ShowReminders(value dto.Reminder) *[]dto.Reminder {
	log.Println("===============List Credentials===============")
	DeleteReminder()
	info := []dto.Reminder{}
	// Client := config.GetConfig()
	// db := viper.GetString("db")
	// log.Println("db name : ", db)
	// collection := Client.Database(db).Collection(viper.GetString("AddReminder"))
	cur, err := config.AddRemainderCollection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
	}
	// defer Client.Disconnect(context.Background())
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &info)
	if err != nil {
		log.Println("Error while fetching documents:", err)
		return nil
	}
	return &info
}
func UpdateEntry(value dto.DailyEntry) {
	// client := config.GetConfig()
	// defer client.Disconnect(context.Background())
	// collection := client.Database(viper.GetString("db")).Collection(viper.GetString("AddFlock"))
	filter := bson.M{"_id": value.ID}
	update := bson.M{
		"$push": bson.M{"entry": value},
	}
	_, err := config.AddFlockCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("error while updating:", err)
	}
	log.Println("successfully updated")
}

func DeleteReminder() string {

	log.Println("------------Reminder Deletion callled ")
	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))

	currentDate := time.Now().Format("2006-01-02")

	// Define filter to find completed reminders
	filter := bson.M{
		"afterdate": bson.M{"$lt": currentDate}, // Reminders with date less than current date are completed
	}

	// Perform deletion operation
	_, err := config.AddRemainderCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return "failed to delete"
	}
	return "deleted  successfully"
}

func UpdateFlockEntries(value dto.DailyEntry) string {
	log.Println("------------Update FlockEntries-------------")
	// Client := config.GetConfig()
	// defer Client.Disconnect(context.TODO())
	//collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddflockEntries"))
	var flock dto.Flockdata
	var flockEntries dto.ListEntry
	flockEntries.EntryDate = value.Date
	flockEntries.Mortality = value.Mortality
	flockEntries.BirdsSold = value.BirdsSold

	// FlockCollection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))

	flockfilter := bson.M{
		"_id": value.ID,
	}

	err := config.AddFlockCollection.FindOne(context.TODO(), flockfilter).Decode(&flock)
	flockEntries.Age = flock.Age

	log.Println("flockvalues-------------------------", flock)
	//log.Println("err in updateflockentry :", err)
	if err == nil && len(flock.ListEntry) > 0 {
		log.Println("inside if ----------", flock.ListEntry)
		lastEntry := (flock.ListEntry)[len(flock.ListEntry)-1]
		flockEntries.OpeningBirds = lastEntry.ClosingBirds
		flockEntries.CumMortality = lastEntry.CumMortality + flockEntries.Mortality
		flockEntries.EggsPerDay = value.Eggs + (value.Trays * 30)
		flockEntries.EggProducion = lastEntry.EggProducion + flockEntries.EggsPerDay
		flockEntries.Feed = value.Feed
		flockEntries.CumFPE = lastEntry.CumFPE + flockEntries.FeedPerEgg
		flockEntries.ClosingBirds = flockEntries.OpeningBirds - (flockEntries.Mortality + flockEntries.BirdsSold)
		flockEntries.FeedPerBird = (flockEntries.Feed * 1000) / float32(flockEntries.ClosingBirds)
		flockEntries.MortalityPer = float32(flockEntries.CumMortality / flock.NoBirds)
		flockEntries.ProductionPer = float32(flockEntries.EggProducion) / float32(flockEntries.ClosingBirds)
		if flockEntries.EggProducion != 0 {
			flockEntries.FeedPerEgg = float32(flockEntries.Feed*1000) / float32(flockEntries.EggProducion)
			flockEntries.CumFPE = (lastEntry.CumFPE + flockEntries.FeedPerBird) / float32(flockEntries.EggProducion)
		}
		flockEntries.TotalFeed = lastEntry.TotalFeed + flockEntries.Feed

	} else {

		log.Println("inside else=========")
		flockEntries.OpeningBirds = flock.NoBirds
		flockEntries.CumMortality = flockEntries.Mortality
		flockEntries.EggsPerDay = value.Eggs + (value.Trays * 30)
		flockEntries.EggProducion = flockEntries.EggsPerDay
		flockEntries.Feed = value.Feed
		flockEntries.ClosingBirds = flockEntries.OpeningBirds - (flockEntries.Mortality + flockEntries.BirdsSold)
		flockEntries.FeedPerBird = float32(flockEntries.Feed*1000) / float32(flockEntries.ClosingBirds)
		flockEntries.MortalityPer = float32(flockEntries.CumMortality / flock.NoBirds)
		flockEntries.ProductionPer = float32(flockEntries.EggProducion) / float32(flockEntries.ClosingBirds)
		if flockEntries.EggProducion != 0 {
			flockEntries.FeedPerEgg = float32(flockEntries.Feed*1000) / float32(flockEntries.EggProducion)
			flockEntries.CumFPE = flockEntries.FeedPerBird / float32(flockEntries.EggProducion)
		}
		flockEntries.TotalFeed = flockEntries.Feed
	}

	filter := bson.M{
		"_id": value.ID,
	}
	query := bson.M{
		"$push": bson.M{
			"listentry": flockEntries,
		},
	}
	_, err = config.AddFlockCollection.UpdateOne(context.Background(), filter, query)
	if err != nil {
		log.Println("error inserting", err)
		return "updation in listEntry failed"
	}
	return "updation in listEntry successfully"
}

func ListFlockEntry() []dto.Flockdata {
	log.Println("----------------Lsit Flock Entry----------------")
	//var listarray []dto.ListEntry
	var flock []dto.Flockdata
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	// log.Println("----connected to DB------------------------")
	cur, err := config.AddFlockCollection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("error finding:", err)
		log.Println(err)
	}
	//defer Client.Disconnect(context.Background())
	defer cur.Close(context.TODO())
	log.Println("fsdfsdfSdfffef")
	// var EntryLength int
	_ = cur.All(context.TODO(), &flock)

	log.Println(flock)
	return flock
}

func ListParticularFlock(Id string) ([]dto.ListEntry, error) {
	log.Println("========= list Particular Flock =================")
	log.Println("==========id", Id)
	var entry []dto.ListEntry
	var flock dto.Flockdata
	flock.ID = Id
	filter := bson.M{"_id": flock.ID}
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	err := config.AddFlockCollection.FindOne(context.TODO(), filter).Decode(&flock)
	if err != nil {
		log.Println("error fetching:", err)

	}
	entry = append(entry, flock.ListEntry...)
	return entry, err
}
func ListReminder() []dto.Reminder {
	log.Println("----------------Lsit Flock Entry----------------")

	var remainder []dto.Reminder
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))
	// log.Println("----connected to DB------------------------")
	cur, err := config.AddFlockCollection.Find(context.Background(), bson.M{"status": "true"})
	if err != nil {
		log.Println("error finding:", err)
		log.Println(err)
	}
	// defer Client.Disconnect(context.Background())
	defer cur.Close(context.TODO())

	_ = cur.All(context.TODO(), &remainder)

	log.Println(remainder)
	return remainder
}

func ShopList() *[]dto.ListShop {
	var flock []dto.Flockdata
	var shopdata []dto.ListShop
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	// shopCollection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddShop"))
	cur, err := config.AddFlockCollection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
		return nil
	}
	//defer Client.Disconnect(context.Background())
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &flock)
	if err != nil {
		log.Println("Error while fetching documents:", err)
		return nil
	}

	for _, f := range flock {
		//log.Println(f)
		lastentry := len(f.ListEntry)
		if lastentry > 0 {

			log.Println("breed:", f.BreedName)

			filter := bson.M{"breedName": f.BreedName}
			var existingShop dto.ListShop
			err := config.AddShopCollection.FindOne(context.Background(), filter).Decode(&existingShop)
			if err == nil {

				log.Println("Document with BreedName", f.BreedName, "already exists, updating values")
				update := bson.M{
					"$set": bson.M{
						"noBirds":   f.ListEntry[lastentry-1].ClosingBirds,
						"noEgg":     f.ListEntry[lastentry-1].EggProducion,
						"birdprice": 50,
						"eggprice":  5,
					},
				}
				log.Println("in update", f.BreedName)
				_, err := config.AddShopCollection.UpdateOne(context.Background(), filter, update)
				if err != nil {
					log.Println("Error updating document:", err)
					continue
				}
			} else if err == mongo.ErrNoDocuments {
				randomID, err := generateRandomID(10)
				if err != nil {
					log.Println("Error:", err)
					return nil
				}
				ID := randomID
				log.Println("in insertion", f.BreedName)
				_, err = config.AddShopCollection.InsertOne(context.Background(), dto.ListShop{
					ID:        ID,
					BreedName: f.BreedName,
					Nobirds:   f.ListEntry[lastentry-1].ClosingBirds,
					NoEgg:     f.ListEntry[lastentry-1].EggProducion,
					Birdprice: 50,
					EggPrice:  5,
					Image:     f.Image,
				})
				if err != nil {
					log.Println("Error while inserting document:", err)
				}
			} else {
				log.Println("Error while finding document:", err)
				continue
			}
		}
	}

	shopCur, err := config.AddShopCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error while fetching shop documents:", err)
		return nil
	}
	defer shopCur.Close(context.Background())
	err = shopCur.All(context.Background(), &shopdata)
	if err != nil {
		log.Println("Error while fetching shop documents:", err)
		return nil
	}

	return &shopdata
}

func fetchShopDataFromDB(client *mongo.Client, id string) (*dto.ListShop, error) {
	var shopData dto.ListShop
	// collection := client.Database("Login").Collection("Shop")
	filter := bson.M{"_id": id}
	err := config.AddShopCollection.FindOne(context.Background(), filter).Decode(&shopData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No document found with ID %s\n", id)
			return nil, nil
		}
		log.Println("Error fetching shop data:", err)
		return nil, err
	}
	return &shopData, nil
}

func ShopListWithIDs(id string) []dto.ListShop {
	var shopList []dto.ListShop
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddCart"))
	var shop dto.ListShop
	log.Println(id)
	id = removeFirstLastChar(id)
	log.Println(id)
	filter := bson.M{"id": id}
	err := config.AddCartCollection.FindOne(context.Background(), filter).Decode(&shop)
	if err != nil {

		if err == mongo.ErrNoDocuments {

			err = config.AddShopCollection.FindOne(context.Background(), bson.M{"id": id}).Decode(&shop)
			log.Println(shop)
			if err != nil {
				log.Println(err)
				if err == mongo.ErrNoDocuments {
					log.Printf("No document found with ID %s\n", id)
					return nil
				}
				log.Println("Error fetching shop Data:", err)
				return nil
			} else {
				log.Println("In Else")
				shop.BirdQuantity = 0
				shop.EggQuantity = 0
				shop.TotalAmount = 0
				iy, err := config.AddCartCollection.InsertOne(context.Background(), shop)
				log.Println(iy)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			log.Println("Error fetching shop data:", err)
		}

	}

	query := bson.M{}
	cur, err := config.AddCartCollection.Find(context.Background(), query)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var shop dto.ListShop
		err := cur.Decode(&shop)
		if err != nil {
			log.Fatal("Error decoding document:", err)
		}
		shopList = append(shopList, shop)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("Error:", err)
	}

	return shopList
}

func removeFirstLastChar(input string) string {
	if len(input) < 2 {
		// If the string has less than 2 characters, there's nothing to remove
		return input
	}
	// Return the substring starting from the second character up to the second-to-last character
	return input[1 : len(input)-1]
}

func RemoveFromGlobalArray(id string) string {
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddCart"))
	filter := bson.M{"id": id}
	result, err := config.AddCartCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Error removing product from cart: %v\n", err)
		return "Error removing product from cart"
	}

	if result.DeletedCount == 0 {
		log.Printf("No product found with ID %s in the cart\n", id)
		return "No product found"
	}

	log.Printf("Product with ID %s successfully removed from the cart\n", id)
	return "Product has been successfully removed from the cart"

}

func ListCart() []dto.ListShop {
	var shopList []dto.ListShop
	query := bson.M{}
	// Client := config.GetConfig()
	// collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddCart"))
	cur, err := config.AddCartCollection.Find(context.Background(), query)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var shop dto.ListShop
		err := cur.Decode(&shop)
		if err != nil {
			log.Fatal("Error decoding document:", err)
		}

		shopList = append(shopList, shop)
	}
	log.Println(shopList[0].BirdQuantity)
	log.Println(shopList)

	if err := cur.Err(); err != nil {
		log.Fatal("Error:", err)
	}

	return shopList
}

func UpdateEggQuantity(id string, eggquantity int, totalamount int) string {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"eggquantity": eggquantity, "totalamount": totalamount}}
	options := options.Update().SetUpsert(false)

	// Perform the update operation
	result, err := config.AddCartCollection.UpdateOne(context.Background(), filter, update, options)
	if err != nil {
		log.Println(err)
		return "Egg Quantity is Not Updated"
	} else {
		log.Println(result)
	}
	return "Egg Quantity is Updated Successfully"
}

func UpdateBirdQuantity(id string, birdquantity int, totalamount int) string {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"birdquantity": birdquantity, "totalamount": totalamount}}
	options := options.Update().SetUpsert(false)

	// Perform the update operation
	result, err := config.AddCartCollection.UpdateOne(context.Background(), filter, update, options)
	if err != nil {
		log.Println(err)
		return "Bird Quantity is Not Updated"
	} else {
		log.Println(result)
	}
	return "Bird Quantity is Updated Successfully"
}

func AddCustomer(customer dto.CustomerReg) (string, error) {
	log.Println("++++++++++++++++++++++++++++  AddCustomer service +++++++++++++++++++++++++")
	ctx := context.Background()
	var existingCustomer dto.CustomerReg
	err := config.AddCustomerCollection.FindOne(ctx, bson.M{"email": customer.Email}).Decode(&existingCustomer)
	if err == nil {
		return "", errors.New("email already exists")
	} else if err != mongo.ErrNoDocuments {
		log.Println("Error checking for existing email:", err)
		return "", err
	}

	// Insert customer into MongoDB
	_, err = config.AddCustomerCollection.InsertOne(ctx, customer)
	if err != nil {
		log.Println("Error inserting document:", err)
		return "", err
	}

	return "Customer added successfully", nil
}

func CustomerLogin(value dto.Logindata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  login service +++++++++++++++++++++++++")

	log.Println(value)
	// Client := config.GetConfig()
	// defer Client.Disconnect(context.Background())

	// collection := Client.Database(viper.GetString("db")).Collection("users")

	filter := bson.M{"email": value.Email, "password": value.Password}

	err := config.AddCustomerCollection.FindOne(context.Background(), filter).Decode(&value)
	log.Println(err)
	if err != nil {
		return "Invalid Credentials", err
	} else {
		return "Login successful", nil
	}
}

func AddAdmin(admin dto.AdminReg) (string, error) {
	log.Println("++++++++++++++++++++++++++++  AddCustomer service +++++++++++++++++++++++++")
	ctx := context.Background()
	var existingCustomer dto.CustomerReg
	err := config.LoginCollection.FindOne(ctx, bson.M{"email": admin.Email}).Decode(&existingCustomer)
	if err == nil {
		return "", errors.New("email already exists")
	} else if err != mongo.ErrNoDocuments {
		log.Println("Error checking for existing email:", err)
		return "", err
	}

	// Insert customer into MongoDB
	_, err = config.LoginCollection.InsertOne(ctx, admin)
	if err != nil {
		log.Println("Error inserting document:", err)
		return "", err
	}

	return "Admin added successfully", nil
}

func PlaceOrder() (string, error) {
	log.Println("-----------------------------IN placeorder----------------")
	ctx := context.Background()
	var totalprice int

	query := bson.M{}
	cur, err := config.AddCartCollection.Find(context.Background(), query)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var shop dto.ListShop
		err := cur.Decode(&shop)
		if err != nil {
			log.Fatal("Error decoding document:", err)
		}
		var flock dto.Flockdata
		queryfind := bson.M{"breedName": shop.BreedName}
		err = config.AddFlockCollection.FindOne(ctx, queryfind).Decode(&flock)
		lastentry := len(flock.ListEntry)
		flock.ListEntry[lastentry-1].ClosingBirds -= shop.BirdQuantity
		flock.ListEntry[lastentry-1].EggProducion -= shop.EggQuantity
		update := bson.M{"$set": bson.M{
			"listentry": flock.ListEntry,
		},
		}
		totalprice += shop.TotalAmount
		_, err = config.AddFlockCollection.UpdateOne(ctx, queryfind, update)
		if err!=nil {
			fmt.Println("err updtaing",err)
		}

	}
	_ = config.AddCartCollection.Drop(ctx)

	go SendOrderConformation("buvaneshwaran.ee20@bitsathy.ac.in", strconv.Itoa(totalprice), strconv.Itoa(totalprice+50),time.Now().Format("12 Jan 2024"),"djhsgdj","5")
	return "sfd", err
}

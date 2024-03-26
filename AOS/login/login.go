package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"log"
	"run/config"
	"run/dto"

	"time"

	//"github.com/aws/aws-sdk-go/aws/client"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func Login(value dto.Logindata) (string, error) {
	log.Println("++++++++++++++++++++++++++++  login service +++++++++++++++++++++++++")

	log.Println(value)
	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())

	collection := Client.Database(viper.GetString("db")).Collection("users")

	filter := bson.M{"email": value.Email, "password": value.Password}

	err := collection.FindOne(context.Background(), filter).Decode(&value)
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

	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())

	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))

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
	}

	_, err = collection.InsertOne(context.Background(), data)
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
	Client := config.GetConfig()
	db := viper.GetString("db")
	log.Println("db name : ", db)

	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	cur, err := collection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
	}
	defer Client.Disconnect(context.Background())
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
	Client := config.GetConfig()
	db := viper.GetString("db")

	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)
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
	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())
	db := viper.GetString("db")
	collection := Client.Database(db).Collection(viper.GetString("AddEntry"))
	_, err := collection.InsertOne(ctx, value)
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

	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())
	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
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

	_, err := collection.UpdateOne(context.Background(), filter, update)
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

	Client := config.GetConfig()
	db := viper.GetString("db")
	collection := Client.Database(db).Collection(viper.GetString("Addflock"))
	Date := time.Now()
	ctx := context.TODO()
	search := bson.M{}
	cursor, err := collection.Find(ctx, search)
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
		_, err := collection.UpdateOne(context.Background(), filter, update)
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
	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())

	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))

	_, err1 := collection.InsertOne(context.Background(), value)
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
	Client := config.GetConfig()
	db := viper.GetString("db")
	log.Println("db name : ", db)
	collection := Client.Database(db).Collection(viper.GetString("AddReminder"))
	cur, err := collection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
	}
	defer Client.Disconnect(context.Background())
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &info)
	if err != nil {
		log.Println("Error while fetching documents:", err)
		return nil
	}
	return &info
}
func UpdateEntry(value dto.DailyEntry) {
	client := config.GetConfig()
	defer client.Disconnect(context.Background())
	collection := client.Database(viper.GetString("db")).Collection(viper.GetString("AddFlock"))
	filter := bson.M{"_id": value.ID}
	update := bson.M{
		"$push": bson.M{"entry": value},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("error while updating:", err)
	}
	log.Println("successfully updated")
}

func DeleteReminder() string {

	log.Println("------------Reminder Deletion callled ")
	Client := config.GetConfig()
	defer Client.Disconnect(context.Background())
	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))

	currentDate := time.Now().Format("2006-01-02")

	// Define filter to find completed reminders
	filter := bson.M{
		"afterdate": bson.M{"$lt": currentDate}, // Reminders with date less than current date are completed
	}

	// Perform deletion operation
	_, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return "failed to delete"
	}
	return "deleted  successfully"
}

func UpdateFlockEntries(value dto.DailyEntry) string {
	log.Println("------------Update FlockEntries-------------")
	Client := config.GetConfig()
	defer Client.Disconnect(context.TODO())
	//collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddflockEntries"))
	var flock dto.Flockdata
	var flockEntries dto.ListEntry
	flockEntries.EntryDate = value.Date
	flockEntries.Mortality = value.Mortality
	flockEntries.BirdsSold = value.BirdsSold

	FlockCollection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))

	flockfilter := bson.M{
		"_id": value.ID,
	}

	err := FlockCollection.FindOne(context.TODO(), flockfilter).Decode(&flock)
	flockEntries.Age = flock.Age

	log.Println("flockvalues-------------------------", flock)
	log.Println("err in updateflockentry :", err)
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
		flockEntries.ProductionPer =float32 (flockEntries.EggProducion) / float32(flockEntries.ClosingBirds)
		if flockEntries.EggProducion != 0 {
			flockEntries.FeedPerEgg = float32(flockEntries.Feed * 1000) / float32(flockEntries.EggProducion)
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
		flockEntries.FeedPerBird = float32(flockEntries.Feed * 1000) / float32(flockEntries.ClosingBirds)
		flockEntries.MortalityPer = float32(flockEntries.CumMortality / flock.NoBirds)
		flockEntries.ProductionPer = float32(flockEntries.EggProducion) /float32( flockEntries.ClosingBirds)
		if flockEntries.EggProducion != 0 {
			flockEntries.FeedPerEgg = float32(flockEntries.Feed * 1000) / float32(flockEntries.EggProducion)
			flockEntries.CumFPE = flockEntries.FeedPerBird /float32( flockEntries.EggProducion)
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
	_, err = FlockCollection.UpdateOne(context.Background(), filter, query)
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
	Client := config.GetConfig()
	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	log.Println("----connected to DB------------------------")
	cur, err := collection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("error finding:", err)
		log.Println(err)
	}
	defer Client.Disconnect(context.Background())
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
	Client := config.GetConfig()
	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	err := collection.FindOne(context.TODO(), filter).Decode(&flock)
	if err != nil {
		log.Println("error fetching:", err)

	}
	entry = append(entry, flock.ListEntry...)
	return entry, err
}
func ListReminder() []dto.Reminder {
	log.Println("----------------Lsit Flock Entry----------------")

	var remainder []dto.Reminder
	Client := config.GetConfig()
	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddReminder"))
	log.Println("----connected to DB------------------------")
	cur, err := collection.Find(context.Background(), bson.M{"status": "true"})
	if err != nil {
		log.Println("error finding:", err)
		log.Println(err)
	}
	defer Client.Disconnect(context.Background())
	defer cur.Close(context.TODO())

	_ = cur.All(context.TODO(), &remainder)

	log.Println(remainder)
	return remainder
}

func ShopList() *[]dto.ListShop {
	var flock []dto.Flockdata
	var shopdata []dto.ListShop
	Client := config.GetConfig()
	collection := Client.Database(viper.GetString("db")).Collection(viper.GetString("Addflock"))
	shopCollection := Client.Database(viper.GetString("db")).Collection(viper.GetString("AddShop"))
	cur, err := collection.Find(context.Background(), bson.M{"active": "true"})
	if err != nil {
		log.Println("Collection list error : ", err)
		return nil
	}
	defer Client.Disconnect(context.Background())
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &flock)
	if err != nil {
		log.Println("Error while fetching documents:", err)
		return nil
	}

	for _, f := range flock {
		lastentry := len(f.ListEntry)
		if lastentry > 0 {

			log.Println("breed:", f.BreedName)

			filter := bson.M{"breedName": f.BreedName}
			var existingShop dto.ListShop
			err := shopCollection.FindOne(context.Background(), filter).Decode(&existingShop)
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
				_, err := shopCollection.UpdateOne(context.Background(), filter, update)
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

				_, err = shopCollection.InsertOne(context.Background(), dto.ListShop{
					ID:        ID,
					BreedName: f.BreedName,
					Nobirds:   f.ListEntry[lastentry-1].ClosingBirds,
					NoEgg:     f.ListEntry[lastentry-1].EggProducion,
					Birdprice: 50,
					EggPrice:  5,
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

	shopCur, err := shopCollection.Find(context.Background(), bson.M{})
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
	collection := client.Database("Login").Collection("Shop")
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&shopData)
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
	return shopList
}

func RemoveFromGlobalArray(id string) string {

	return ""
}

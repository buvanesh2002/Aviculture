package dto

type Logindata struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Flockdata struct {
	Image      string       `json:"image" bson:"image"`
	ID         string       `json:"id,omitempty" bson:"_id,omitempty"`
	FlockName  string       `json:"flockName,omitempty" bson:"flockName,omitempty"`
	BreedName  string       `json:"breedName,omitempty" bson:"breedName,omitempty"`
	StartDate  string       `json:"startDate,omitempty" bson:"startDate,omitempty"`
	StartAge   int          `json:"startAge,omitempty" bson:"startAge,omitempty"`
	Age        int          `json:"age,omitempty" bson:"age"`
	NoBirds    int          `json:"openingBirds,omitempty" bson:"openingBirds,omitempty"`
	ShedNumber string       `json:"shedNumber,omitempty" bson:"shedNumber,omitempty"`
	Active     string       `json:"active,omitempty" bson:"active,omitempty"`
	CreatedAt  string       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  string       `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Entry      []DailyEntry `json:"entry,omitempty" bson:"entry,omitempty"`
	ListEntry  []ListEntry  `json:"listentry,omitempty" bson:"listentry,omitempty"`
}

type DailyEntry struct {
	ID        string  `json:"id" bson:"id"`
	Date      string  `json:"date,omitempty" bson:"date,omitempty"`
	Mortality int     `json:"mortality,omitempty" bson:"mortality,omitempty"`
	Eggs      int     `json:"extraeggs,omitempty" bson:"extraeggs,omitempty"`
	Feed      float32 `json:"feed,omitempty" bson:"feed,omitempty"`
	BirdsSold int     `json:"birdssold,omitempty" bson:"birdssold,omitempty"`
	CountErr  int     `json:"counterr,omitempty" bson:"counterr,omitempty"`
	Remarks   string  `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Trays     int     `json:"trays,omitempty" bson:"trays,omitempty"`
}

type Reminder struct {
	ReminderId string `json:"reminderId,omitempty" bson:"reminderId,omitempty"`
	Name       string `json:"remindername" bson:remindername"`
	BeforeDate string `json:"beforedate" bson:"beforedate"`
	AfterDate  string `json:"afterdate" bson:"afterdate"`
	Date       string `json:"reminderdate" bson:"reminderdate"`
	Remarks    string `json:"remarks" bson:"remarks"`
	Status     string `json:"status" bson:"status"`
}
type ListEntry struct {
	EntryDate     string  `json:"entrydate" bson:"entrydate"`
	Age           int     `json:"age" bson:"age"`
	OpeningBirds  int     `json:"openingbirds" bson:"openingbirds"`
	Mortality     int     `json:"mortality" bson:"mortality"`
	BirdsSold     int     `json:"birdssold" bson:"birdssold"`
	ClosingBirds  int     `json:"closingbirds" bson:"closingbirds"`
	CumMortality  int     `json:"cummortality" bson:"cummortality"`
	MortalityPer  float32 `json:"mortalitypercent" bson:"mortalitypercent"`
	EggsPerDay    int     `json:"eggsperDay" bson:"eggsperDay"`
	EggProducion  int     `json:"eggproducion" bson:"eggproducion"`
	ProductionPer float32 `json:"productionpercent" bson:"productionpercent"`
	// HHP string  `json:"curHHP" bson:"curHHP"`
	// CumHHP string `json:"cumHHP" bson:"cumHHP"`
	Feed        float32 `json:"feed" bson:"feed"`
	FeedPerBird float32 `json:"feedperBird" bson:"feedperBird"`
	FeedPerEgg  float32 `json:"feedperEgg" bson:"feedperEgg"`
	CumFPE      float32 `json:"cumFPE" bson:"cumFPE"`
	TotalFeed   float32 `json:"totalFeed" bson:"totalFeed"`
}

type ListShop struct {
	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	Image        string `json:"image" bson:"image"`
	BreedName    string `json:"breedName,omitempty" bson:"breedName"`
	Nobirds      int    `json:"noBirds,omitempty" bson:"noBirds"`
	NoEgg        int    `json:"noEgg,omitempty" bson:"noEgg"`
	Birdprice    int    `json:"birdprice,omitempty"  bson:"birdprice"`
	EggPrice     int    `json:"eggprice,omitempty"  bson:"eggprice"`
	EggQuantity  int    `json:"eggquantity"  bson:"eggquantity"`
	BirdQuantity int    `json:"birdquantity"  bson:"birdquantity" `
	TotalAmount int     `json:"totalamount"  bson:"totalamount" `
}


type CustomerReg struct {
    Name     string `json:"name" bson:"name"`
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
    Phone    string `json:"phone" bson:"phone"`
    Address  string `json:"address" bson:"address"`
    Pincode  string `json:"pincode" bson:"pincode"`
}

type AdminReg struct {
    Name     string `json:"name" bson:"name"`
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
    Phone    string `json:"phone" bson:"phone"`
    Address  string `json:"address" bson:"address"`
    Pincode  string `json:"pincode" bson:"pincode"`
}


type Order struct {
	Address    string `json:"address" bson:"address"`
    Companyname    string `json:"companyname" bson:"companyname"`
    Country string `json:"country" bson:"country"`
    EmailAddress    string `json:"emailaddress" bson:"emailaddress"`
    FirstName  string `json:"firstname" bson:"firstname"`
    LastName  string `json:"lastname" bson:"lastname"`
	OrderNotes  string `json:"ordernotes" bson:"ordernotes"`
	Phone  string `json:"phone" bson:"phone"`
	PostalCode  string `json:"postalcode" bson:"postalcode"`
	State  string `json:"state" bson:"state"`
    BreedName    string `json:"breedname,omitempty" bson:"breedname"`
	EggQuantity  int    `json:"eggquantity"  bson:"eggquantity"`
	BirdQuantity int    `json:"birdquantity"  bson:"birdquantity" `
	TotalAmount int     `json:"totalamount"  bson:"totalamount" `

}
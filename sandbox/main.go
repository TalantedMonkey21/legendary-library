package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	Content string `json:"content" gorm:"not null"`
}

type Router struct {
	Db *gorm.DB
}

type DbConfig struct {
	Host string
	User string
	Password string
	Dbname string
	Port string
	Sslmode string
}


func WriteJSONError(w http.ResponseWriter, code int, e string) {
	WriteJSONResponse(w, code, map[string]string{"error":e})
}

func WriteJSONResponse(w http.ResponseWriter, code int, s any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(s)
}

func (rt *Router) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"Hello World"})
}

func (rt *Router) CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteJSONError(w, http.StatusMethodNotAllowed, "Incorrect method")
		return
	}
	var n Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if len(n.Content) < 3 {
		WriteJSONError(w, http.StatusBadRequest, "Write more!!!")
		return
	}
	if err := rt.Db.Create(&n).Error; err != nil {
		WriteJSONError(w, http.StatusInternalServerError, "Cannot create note")
		return
	}
	WriteJSONResponse(w, http.StatusCreated, n)
}

func getEnv(key, defValue string) string {
	if key == "" {
		fmt.Printf("Not found %v, use default\n", key)
		return defValue
	}
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Not found %v value, use default\n", key)
		return defValue
	}
	return value
}
func main() {
	db := DbConfig{
		getEnv("POSTGRES_HOST", "localhost"),
		getEnv("POSTGRES_USER", "admin"),
		getEnv("POSTGRES_PASSWORD", "supersecret"),
		getEnv("POSTGRES_DB", "lectures"),
		getEnv("POSTGRES_PORT", "5432"),
		getEnv("Sslmode", "disable"),
	}
	Dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", db.Host, db.User, db.Password, db.Dbname, db.Port, db.Sslmode)
	fmt.Println(Dsn)
	connect, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
		os.Exit(1)
	}
	err = connect.AutoMigrate(&Note{})
	if err != nil {
		fmt.Println("Cannot migrate")
		os.Exit(1)
	}
	fmt.Println("Migrate complete!")
	r := &Router{Db:connect}
	router := http.NewServeMux()
	router.HandleFunc("/health", r.Health)
	router.HandleFunc("/create", r.CreateNote)
	fmt.Println("Поднимаю сервер")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
}


// func deletionTwo (m map[int]string) map[int]string {
// 	for k := range m {
// 		if k % 2 == 0 {
// 			delete(m, k)
// 		}
// 	}
// 	return m
// }

// func main() {
// 	mapa := map[int]string{
// 		1: "a",
// 		2: "b",
// 		3: "c",
// 		4: "d",
// 		5: "e",
// 		6: "f",
// 		7: "g",
// 		8: "h",
// 		9: "i",
// 		10: "j",
// 	}
// 	fmt.Println(deletionTwo(mapa))
// }

// func InvertMap(m map[string]string) map[string][]string {
// 	invertedMap := make(map[string][]string)
// 	for k, v := range m {
// 		invertedMap[v] = append(invertedMap[v], k)
// 	}
// 	return invertedMap
// }


// func main() {
// 	mapa := map[string]string{
// 		"apple": "apple@apple.com",
// 		"banana": "banana@banana.com",
// 		"cherry": "cherry@cherry.com",
// 		"paeach": "apple@apple.com",
// 	}
// 	fmt.Println(InvertMap(mapa))
// }
// func mergeMaps(a, b map[string]int) map[string]int {
// 	MapAB := make(map[string]int)
// 	maps.Copy(MapAB, a)
// 	for k, v := range b {
// 		MapAB[k] += v
// 	}
// 	return MapAB
// }

// func main() {
// 	mapA := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
// 	mapB := map[string]int{"banana": 20, "date": 4, "elderberry": 5}

// 	fmt.Println(mergeMaps(mapA, mapB))
// }
// var Students map[int]string


// func GetName(id int) (string, bool) {
// 	name, ok := Students[id]
// 	return name, ok
// } 


// func main() {
// 	Students = make(map[int]string)
// 	Students[1] = "Alice"
// 	Students[2] = "Bob"
// 	Students[3] = "Charlie"
// 	ids := []int{1, 2, 5}

// 	for _, id := range ids {
// 		fmt.Printf("Check ID %d\n", id)
// 		if name, ok := GetName(id); ok {
// 			fmt.Printf("%s, found in list\n", name)
// 		} else {
// 			fmt.Printf("Not found in list\n")
// 		}
// 	}



// }

// func main() {
// 	fmt.Println(uniqueElements([]int{10, 20, 10, 30, 20, 20}))

// }


// func uniqueElements(arr []int) ([]int, int, int) {
// 	counts := make(map[int]int)
// 	for _, num := range arr {
// 		counts[num]++
// 	}
// 	maxCount := 0
// 	maxNum := 0

// 	for num, count := range counts {
// 		if count > maxCount {
// 			maxNum = num
// 			maxCount = count
// 		}
// 	}

// 	unique := make(map[int]bool)
// 	uniqueList := []int{}
// 	for _, num := range arr {
// 		unique[num] = true
// 	}

// 	for num := range unique {
// 		uniqueList = append(uniqueList, num)
// 	}
// 	return uniqueList, maxCount, maxNum
// }

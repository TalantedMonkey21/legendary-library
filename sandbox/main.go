package main

import (
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	
}
//Redis Kafka Sql Algoritms Sobes

func (h Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	
}

func main() {

}

// func main() {
// 	nums := []int{4, 5, 6, 7, 0, 1, 2}
// 	target := 0
// 	for i := range nums {
// 		if nums[i] == target {
// 			fmt.Println(i)
// 		}
// 	}
// }



// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/TalantedMonkey21/GoLectures/internal/response"
// )

// type Middleware func(http.Handler) (http.Handler)

// func loggingMiddleware(next http.Handler) (http.Handler) {
// 	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
// 		start := time.Now()
// 		next.ServeHTTP(w, r)
// 		log.Printf("%s, %s, %s", r.Method, r.URL.Path, time.Since(start))
// 	})
// }

// func goodbuyMiddleware(next http.Handler) (http.Handler) {
// 	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 		log.Printf("Goodbuy!")
// 	})
// }

// func setMethod(next http.Handler) (http.Handler) {
// 	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
// 		r.Method = "POST"
// 		next.ServeHTTP(w, r)

// 	})
// }

// func Chain(next http.Handler, m ...Middleware) (http.Handler) {
// 	wrapped := next
// 	for _, f := range m {
// 		wrapped = f(wrapped)
// 	}
// 	return wrapped
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	response.WriteJSONResponse(w, http.StatusOK, map[string]string{"message":"HelloWorld"})
// }

// func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
// 	response.WriteJSONResponse(w, http.StatusOK, map[string]string{"message":"Goodbye"})
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/hello", http.HandlerFunc(helloHandler))
// 	mux.Handle("/goodbuy", http.HandlerFunc(goodbyeHandler))
// 	newMux := Chain(mux, setMethod, loggingMiddleware, goodbuyMiddleware)
// 	http.ListenAndServe(":8080", newMux)
// }


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

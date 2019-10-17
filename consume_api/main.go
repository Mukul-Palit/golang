package main
import(
	"fmt"
	"io/ioutil"
    	"log"
    	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
type Response struct {
    Records []Record `json:"records"`
}
// A struct to map the data fetched from the API
type Record struct {
	ID string `json:"id"`
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Station string `json:"station"`
	LastUpdate string `json:"last_update"`
	Pollutant_id string `json:"pollutant_id"`
	Pollutant_min string `json:"pollutant_min"`
	Pollutant_max string `json:"pollutant_max"`
	Pollutant_avg string `json:"pollutant_avg"`
	Pollutant_unit string `json:"pollutant_unit"`
}

var records Response
func main (){
	for i := 510; i<=700; i=i+10 {
		result := fmt.Sprintf("https://api.data.gov.in/resource/3b01bcb8-0b14-4abf-b6f2-c1bfd384ba69?api-key=579b464db66ec23bdd000001cdd3946e44ce4aad7209ff7b23ac571b&offset=%d&limit=10&format=json",i)
		response, err := http.Get(result)

		if err != nil {
			fmt.Print(err.Error())
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &records)
		//time delay between consecutive fetch request		
		time.Sleep(700 * time.Millisecond)
	
	db, err := sql.Open("mysql", "root:root987@tcp(127.0.0.1:3306)/gov")   
    if err != nil {
        panic(err.Error())
	}

	defer db.Close() 
	var query string =  "Insert into pollution(RefId, Country, State, City, Station, LastUpdate, Pollution_ID, Pollution_min, Pollution_max, Pollution_avg, Pollution_unit) values (?,?,?,?,?,?,?,?,?,?,?)"
	insert, err := db.Prepare(query)
	if err != nil {
        panic(err.Error())
	}
	for i := 0; i < len(records.Records); i++ {
		insert.Exec(
			records.Records[i].ID,
			records.Records[i].Country, 
			records.Records[i].State, 
			records.Records[i].City, 
			records.Records[i].Station, 
			records.Records[i].LastUpdate, 
			records.Records[i].Pollutant_id, 
			records.Records[i].Pollutant_min, 
			records.Records[i].Pollutant_max, 
			records.Records[i].Pollutant_avg, 
			records.Records[i].Pollutant_unit,
		)
	}
	defer insert.Close()

}
}

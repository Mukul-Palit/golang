package main
import(
	"fmt"
	"io/ioutil"
    	"log"
    	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
    Records []Records `json:"records"`
}

// A struct to map our Pokemon's Species which includes it's name
type Records struct {
	ID string `json:"id"`
	Country string `"json:country"`
	State string `"json:state"`
	City string `"json:city"`
	Station string `"json:station"`
	LastUpdate string `"json:last_update"`
	Pollutant_id string `"json:pollutant_id"`
	Pollutant_min string `"json:pollutant_min"`
	Pollutant_max string `"json:pollutant_max"`
	Pollutant_avg string `"json:pollutant_avg"`
	Pollutant_unit string `"json:pollutant_unit"`
}

var responseObject Response

func main (){
	for i := 0; i<=1000; i=i+10 {
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
		//fmt.Println(string(responseData))
		json.Unmarshal(responseData, &responseObject)
	}
	//fmt.Println(responseObject.Name)
	//fmt.Println(len(responseObject.Pokemon))
	
	db, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(IP_ADDRESS OR LOCALHOST:3306)/DATABASE_NAME")   
    if err != nil {
        panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Prepare("Insert into TABLE_NAME values (?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
        panic(err.Error())
	}
	for i := 0; i < len(responseObject.Records); i++ {
		fmt.Println(responseObject.Records[i].ID)
		fmt.Println(responseObject.Records[i].Country)
		fmt.Println(responseObject.Records[i].State)
		//insert.Exec(responseObject.Records[i].ID,)
	}
	defer insert.Close()

}

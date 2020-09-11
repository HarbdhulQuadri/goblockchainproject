// forms.go
package main

import (
    "html/template"
    "net/http"
    "github.com/blockcypher/gobcy"
   "log"
   "fmt"
)

type Response struct{
	//geting started with our variables
	Address string 
	TotalReceived int 
	TotalSent int 
	Balance int 
	UnconfirmedBalance int 
    FinalBalance int 
}
const key = "4d20cbd6a2c74d2f936b8d43502c4572"


	
	
	
	type hotdog int
	
	func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		Address := req.FormValue("address")
    	btc := gobcy.API{key, "btc", "main"}
		addr, err := btc.GetAddrBal(Address, nil)
		 
		data := Response{
			Address :addr.Address, 
			TotalReceived :addr.TotalReceived,
			TotalSent :addr.TotalSent ,
			Balance :addr.Balance ,
			UnconfirmedBalance :addr.UnconfirmedBalance,
			FinalBalance :addr.FinalBalance, 
	
			} 
			fmt.Println("Address =",data.Address,"balance",data.Balance,)

		
       
		tpl.ExecuteTemplate(w, "index.gohtml", data)
	}
	
	
	var tpl *template.Template
	
	func init() {
		tpl = template.Must(template.ParseFiles("index.gohtml"))
	}
	
	func main() {
		var d hotdog
		http.ListenAndServe(":8880", d)
	}
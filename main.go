// forms.go
package main

import (
    "fmt"
    "html/template"
    "net/http"
    "github.com/blockcypher/gobcy"
   // "encoding/json"
)

type Response struct{
	//geting started with our variables
	address string 
	wallet string 
	hdWallet string 
	totalReceived string 
	totalSent string 
	balance string 
	unconfirmedBalance string 
    finalBalance string 
}
const key = "4d20cbd6a2c74d2f936b8d43502c4572"



func main() {
    tmpl := template.Must(template.ParseFiles("form.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }


        address:=   r.FormValue("address")

        // do something with details
        btc := gobcy.API{key, "btc", "main"}
        Addr, err := btc.GetAddrBal(address, nil)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(Addr.Balance)
        fmt.Printf("%+v\n", Addr)


          //var data Response
    
          // unmarshall
         // json.Unmarshal(addr, &data)
        




        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8000", nil)
}
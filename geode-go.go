package main

import "github.com/mross1080/gemfire-golang"
import "encoding/json"
import "fmt"


func main() {
    conn := gemfireGolang.Api{"http://192.168.5.94", "8080"}
    customers := gemfireGolang.Region{conn,"customers"}

    user  := struct{Name string
                    Age int
                    Id string}{
                    "Marty McFly",
                    22,
                    "110"}
    u, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
    }

    //customers.Put(user.Id, u)
    customers.Get(user.Id)

}

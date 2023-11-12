package main

import (
    "net/http"
    "github.com/gin-gonic/gin"

	"fmt"
	"os"
	"io/ioutil"
    "log"
    "net/url"
)

// getAlbums responds with the list of all albums as JSON.
func getCoins(c *gin.Context) {

    client := &http.Client{}
    req, err := http.NewRequest("GET","https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
    if err != nil {
      log.Print(err)
      os.Exit(1)
    }
  
    q := url.Values{}
    q.Add("start", "1")
    q.Add("limit", "5000")
    q.Add("convert", "USD")
  
    req.Header.Set("Accepts", "application/json")
    req.Header.Add("X-CMC_PRO_API_KEY", "8d52310a-2bd0-4c31-893f-70e9c92d0bc4")
    req.URL.RawQuery = q.Encode()
  
  
    resp, err := client.Do(req);
    if err != nil {
      fmt.Println("Error sending request to server")
      os.Exit(1)
    }
    fmt.Println(resp.Status);
    respBody, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println(string(respBody));

    c.IndentedJSON(http.StatusOK, string(respBody))
}
const serverPort = 443

func main() {
    router := gin.Default()
    router.GET("/coins", getCoins)
    router.Run("0.0.0.0:8080")

    
}
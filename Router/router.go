package Router

import (
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
)

func RouterInit(router *gin.Engine) {

  router.Use(cors.Default())
  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "It works On 8081")
  })
  //router.GET("/FPList", FPList.GetList)


  b, err := ioutil.ReadFile("cartest.json")
  if err != nil {
    fmt.Print(err)
  }
  str := string(b)
  router.GET("/data",func(c *gin.Context) {
    c.String(http.StatusOK,"%s",str)
  })

  a, err := ioutil.ReadFile("testnumber.json")
  if err != nil {
    fmt.Print(err)
  }
  str_all := string(a)
  // fmt.Println(str)
  router.GET("/alldata",func(c *gin.Context) {
    c.String(http.StatusOK,"%s",str_all)
  })

  c, err := ioutil.ReadFile("re.json")
  if err != nil {
    fmt.Print(err)
  }
  str_word_cloud := string(c)
  // fmt.Println(str)
  router.GET("/wordcloudData",func(c *gin.Context) {
    c.String(http.StatusOK,"%s",str_word_cloud)
  })

  d, err := ioutil.ReadFile("sales.json")
  if err != nil {
    fmt.Print(err)
  }
  str_sale_number := string(d)
  // fmt.Println(str)
  router.GET("/saleNumber",func(c *gin.Context) {
    c.String(http.StatusOK,"%s",str_sale_number)
  })



}

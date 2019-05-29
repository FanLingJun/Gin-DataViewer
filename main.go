package main

import (
  "awesomeProject/Router"

  //_ "awesomeProject/FPList"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  Router.RouterInit(router)
  //FPList.ReadFile()
  router.Run(":8081")
}


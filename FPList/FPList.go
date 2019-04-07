package FPList

import (
  _ "encoding/json"
  _ "fmt"
  "github.com/gin-gonic/gin"
  _ "os"
)

type CarInfos struct {
  name string
  style string
  price string
  speed string
  accelerate_speed string
  maintenance_period string
  weight string
  displacement string
  housepower string
  power string
}

func ReadFile() {

  /*filePtr, err := os.Open("cartest.json")
  if err != nil {
    fmt.Println("Open file failed [Err:%s]",err.Error())
    return
  }
  defer filePtr.Close()

  var carinfos []CarInfos

  //创建json解码器
  decoder := json.NewDecoder(filePtr)
  err = decoder.Decode(&carinfos)
  if err != nil {
    fmt.Println("Decoder failed",err.Error())
  } else {
    fmt.Println("Decoder success")
    fmt.Println(carinfos)
  }*/

  /*b, err := ioutil.ReadFile("cartest.json")
  if err != nil {
    fmt.Print(err)
  }
  fmt.Println(b)
  str := string(b)
  fmt.Println(str)*/



}



type itemList struct {
  ID     int
  Sort   int
  Name   string
  Status string
  Forums []forumsList
}

type forumsList struct {
  ID        int
  Fgroup    int
  Sort      int
  Name      string
  ShowName  string
  Msg       string
  Interval  int
  CreatedAt string
  UpdateAt  string
  Status    string
}

func GetList(c *gin.Context) {
  flist := []forumsList{
    forumsList{
      ID:        3,
      Fgroup:    4,
      Sort:      1,
      Name:      "每日好货3",
      ShowName:  "",
      Msg:       "msg",
      Interval:  1,
      CreatedAt: "2018-07-20 15:49:28",
      Status:    "n",
    },
  }
  rtList := itemList{
    ID:     1,
    Sort:   1,
    Name:   "每日好货D",
    Status: "n",
  }
  rtList.Forums = flist
  c.JSON(0, rtList)
}

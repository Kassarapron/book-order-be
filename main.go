package main


import (

  "github.com/kassarapron/newproject/controller"

  "github.com/kassarapron/newproject/entity"

  "github.com/gin-gonic/gin"

)


func main() {

  entity.SetupDatabase()

  
  r := gin.Default()
  r.Use(CORSMiddleware())

  // Admin Routes
  r.GET("/Admin"        , controller.ListAdmin)
  r.GET("/Admin/:id"    , controller.GetAdmin)
  r.POST("/Admin"       , controller.CreateAdmin)
  r.PATCH("/Admin"      , controller.UpdateAdmin)
  r.DELETE("/Admin/:id" , controller.DeleteAdmin)
 // Book Routes
 r.GET("/Book"        , controller.ListBook)
 r.GET("/Book/:id"    , controller.GetBook)
 r.POST("/Book"       , controller.CreateBook)
 r.PATCH("/Book"      , controller.UpdateBook)
 r.DELETE("/Book/:id" , controller.DeleteBook)
 // BookType Routes
 r.GET("/BookType"        , controller.ListBookType)
 r.GET("/BookType/:id"    , controller.GetBookType)
 r.POST("/BookType"       , controller.CreateBookType)
 r.PATCH("/BookType"      , controller.UpdateBookType)
 r.DELETE("/BookType/:id" , controller.DeleteBookType)
 // Company Routes
 r.GET("/Company"        , controller.ListCompany)
 r.GET("/Company/:id"    , controller.GetCompany)
 r.POST("/Company"       , controller.CreateCompany)
 r.PATCH("/Company"      , controller.UpdateCompany)
 r.DELETE("/Company/:id" , controller.DeleteCompany)

// Run the server
  r.Run()

}

func CORSMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
    }
    c.Next()
  }
}
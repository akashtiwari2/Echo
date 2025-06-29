package db

import (
  "fmt"
		"os"
		"log"
		"github.com/joho/godotenv"
		"gorm.io/gorm"
		"gorm.io/driver/postgres"
)

func NewDB() *gorm.DB{

	  err:= godotenv.Load()
			if err!=nil{
				log.Fatalln("Env could not be loaded")
			}
   
			url:=fmt.Sprintf("postgres://%s:%s:@%s:%s/%s",os.Getenv("POSTGRES_USER"),os.Getenv("POSTGRES_PW"),
			os.Getenv("POSTGRES_HOST"),os.Getenv("POSTGRES_PORT"),os.Getenv("POSTGRES_DB"))
		
			db,err := gorm.Open(postgres.Open(url), & gorm.Config{})

			if err!=nil{
				 log.Fatalln(err)
			}

			fmt.Println("Connected")



		
return db
	
}

func CloseDB(db *gorm.DB){
	 sdb,_ :=db.DB()
		err:=sdb.Close()
		if err!=nil{
			log.Fatalln(err)
		}
}
package initSetup

import (
	"log"
	"os"
)

func initEnvSetup() {
	os.Setenv("mongoUrl", MONGO_URL)
	os.Setenv("PORT", PORT)
	log.Println("env variable setup done")
}

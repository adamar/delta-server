
package controllers

import (
        "os"
)

func RunFrontend() {

        deltaServer := Router()
        frontendPort := os.Getenv("FRONTEND_PORT")
        if frontendPort == "" {
                frontendPort = "3300"
        }
        deltaServer.Run(":" + frontendPort)

}

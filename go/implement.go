package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const ApiVersion = "1.0.1"

func getIndex(c *gin.Context) {
	c.JSON(http.StatusOK, Beacon{
		ApiVersion: ApiVersion,
	})
}

func getQuery(c *gin.Context) {
	req := BeaconAlleleRequest{}

	err := c.ShouldBind(&req)
	// if err == nil {
	//	if req.AssemblyId == "" {
	//		err = errors.New("Missing assembly id")
	//	} else if req.ReferenceBases == "" {
	//		err = errors.New("missing reference bases")
	//	} else if req.AlternateBases == "" {
	//		err = errors.New("missing alternate bases")
	//	}
	// }
	if err != nil {
		c.JSON(400, BeaconAlleleResponse{
			AlleleRequest: req,
			ApiVersion:    ApiVersion,
			Error: BeaconError{
				ErrorCode:    400,
				ErrorMessage: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, BeaconAlleleResponse{
		ApiVersion:    ApiVersion,
		AlleleRequest: req,
	})
}

func postQuery(c *gin.Context) {
}

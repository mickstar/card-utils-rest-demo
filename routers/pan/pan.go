package pan

import (
	"github.com/gin-gonic/gin"
	"github.com/mickstar/payment-card-utils-go"
	"github.com/mickstar/payment-card-utils-go/Scheme"
	"net/http"
	"strings"
)

type MaskPanRequestBody struct {
	Pan string `json:"pan"`
}

type MaskPanResponseBody struct {
	MaskedPan string `json:"masked_pan"`
}

// MaskPanResponse @title MaskPanResponse
// @description MaskPanResponse
// @name MaskPanResponse
// @Accept  json
// @Produce  json
// @Success 200 {object} MaskPanResponse
func MaskPan(c *gin.Context) {

	var body MaskPanRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	maskedPan := CardUtils.MaskPan(body.Pan)

	res := MaskPanResponseBody{
		MaskedPan: maskedPan,
	}

	c.JSON(200, res)
}

type GenerateRandomPanResponseBody struct {
	Pan string `json:"pan"`
}

func mapScheme(schemeString string) Scheme.Scheme {
	s := strings.ToUpper(schemeString)
	switch s {
	case "VISA":
		return Scheme.Visa
	case "MASTERCARD":
		return Scheme.MasterCard
	case "AMEX", "AMERICAN EXPRESS":
		return Scheme.AmericanExpress
	case "DINERS", "DINERS CLUB":
		return Scheme.DinersClub
	case "DISCOVER":
		return Scheme.Discover
	case "JCB":
		return Scheme.JCB
	case "UNIONPAY":
		return Scheme.UnionPay
	case "BP CARD":
		return Scheme.BPCard
	default:
		return Scheme.Unknown
	}
}

func GenerateRandomPan(c *gin.Context) {
	schemeQuery := c.Query("scheme")

	if schemeQuery == "" {
		schemeQuery = "visa"
	}

	scheme := mapScheme(schemeQuery)
	if (scheme == Scheme.Unknown) {
		c.JSON(400, gin.H{"error": "invalid scheme"})
		return
	}

	pan := CardUtils.GenerateRandomPanOfScheme(scheme)
	c.JSON(200, GenerateRandomPanResponseBody{
		Pan: pan,
	})
}

type ValidatePanRequestBody struct {
	Pan string `json:"pan"`
}

func ValidatePan(c *gin.Context) {
	var panBody ValidatePanRequestBody
	if err := c.BindJSON(&panBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	valid := CardUtils.ValidityCheck(panBody.Pan)

	if valid {
		c.JSON(200, gin.H{"valid": true})
	} else {
		c.JSON(http.StatusConflict, gin.H{"valid": false})
	}
	return
}
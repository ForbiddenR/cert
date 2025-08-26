package rsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var ps = "-----BEGIN CERTIFICATE REQUEST-----\nMIIC8zCCAdsCAQAwgZYxCzAJBgNVBAYTAkNOMREwDwYDVQQIDAhaaGVKaWFuZzEP\nMA0GA1UEBwwGTmluZ0JvMQ8wDQYDVQQKDAZqb3lzb24xFDASBgNVBAsMC0hRNDAw\nMDAwMDAxMR0wGwYDVQQDDBRocXVhdC5qb3lzb25xdWluLmNvbTEdMBsGCSqGSIb3\nDQEJARYObHBoajA3QDE2My5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\nAoIBAQDjSBTjvmMm9xFlKKS6JjaNKOeRy0daH39GCETqANELgQIJWn+q8arJ3TV9\nIVeojmOlILwOaAuGv43EcPQlMD4pAuC22a9yQ8I547DiAJ+zQKWFW9pP0nlrDAQj\nZ19mFgiV7G0zscG0zpfhsp1hZ0SCgsSHioMLWxic4Eev4qCs5EU8GLySoj4Y+3+l\ntBiSceH8Dt7Ncy4+LDGXS6eYkPImnmdYQxZmtNTXXyWqxNZ08wdUa3EkJNmADXl6\nX83UKtplTIRbPihKtUHClz7LSbvRC0xZwIRrgxHqSHg2nQMKqSHqi9pfI0FgFQ8Z\nZ05ZtCW9apkSLTTQrK17mZmJ+PL9AgMBAAGgFzAVBgkqhkiG9w0BCQcxCAwGMTIz\nNDU2MA0GCSqGSIb3DQEBCwUAA4IBAQAPshj+FGDpAH7QMWWnqJn6r5R2f566PRzx\n2qMpk0cWRlvM+Y2FVPIzIDYGmsws84qjAAXpOOGVagN7UNW7QHKrQGirDcBk5O/y\nbWoBqerGv1Qf8fk8uQGq1ial3JofXOum9h1Inqf6f1SV0TKJKgHSGFauUoQDXA8I\noC3vgk6o35l4En/XbR9iCCXASDpSkHZX0D2oidENcFukKlzqTddBiH/ooZo71JrL\n+eloagtb34pF/uiYRC+F1FTFvvO7WQQ66C8vuBAB5fQhXarhUpScpXInuZHfpG1V\n7sdklh/26UZ17zniKELTF2BX8HTqOq8Tb3Nq2EAxVYb78ENynpjD\n-----END CERTIFICATE REQUEST-----"
// var ps = "-----BEGIN CERTIFICATE REQUEST-----\nMIICgjCCAWoCAQEwPTELMAkGA1UEBhMCQ04xDzANBgNVBAoMBmpveXNvbjEdMBsG\nA1UEAwwUaHF1YXQuam95c29ucXVpbi5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IB\nDwAwggEKAoIBAQCZ3nYbqSc0jfV32HfaYX3YrY8YGyW9Oh8CUYBuN3ZluKGNq8kl\niX+3zGn1ctq5uYYOtn3eWPxCAj4C3/VPpbW5F2iu00tYo7LCO2gBwKkh0QzQ1yXb\n9jLPrsNn3xBvYUKyL+yuMRJ79Q7askX7q3N1YcErgcqXHmGp4QzWbLLgNTqo1Or5\nPkv039nmHkmJAZ3pmMWPq0DxriA8GAzarx29ol1nVJxlGMjgmPwhvKd10lt6KqpO\nzus0zSn1WmsfxQCoMsABBUUYVE9TPkci5n1LaXlkT8FZo3xyD7dWE1tlVCjddR+E\nm+XrwjFsoDhDPu0tbuV+av+GfP8ame72U6JzAgMBAAGgADANBgkqhkiG9w0BAQsF\nAAOCAQEAFFCgpkhQOn7hzBXvQ4Rot1gmmfpypJilQBquibRztJ4+6yDcS374ICDp\n87xV1bmySxiT7lC2dsniyqfAzLxlOHSO8lSPhzt43Ym74BMUFhF7/SULqoRmX3lQ\nQHj94cP+tDSuMeULXNPtNd5+VLOVHoWoufGgFW5LggprAcfLY+NbsOgu1+APDDfg\n9+NwyDlCLzrQ5cCfPukTN2DYM5hZfTAT4mFmP3lLf7W4FqA6Z6998usCBZjdbjIb\nfsaJYi2ZL6Jregp2neu1IL827r+nU2IIGr1afHT8dVo8uykbb15NKpiLPNCkm4ZT\n2e4iHKX5T9Wm//tSssKtF60bAIOa8Q==\n-----END CERTIFICATE REQUEST-----\n"
var ps = "-----BEGIN CERTIFICATE REQUEST-----\nMIICgjCCAWoCAQEwPTELMAkGA1UEBhMCQ04xDzANBgNVBAoMBmpveXNvbjEdMBsG\nA1UEAwwUaHF1YXQuam95c29ucXVpbi5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IB\nDwAwggEKAoIBAQCTOk5bMn/9QMPCdVWjkLjfhUJB5zOVGo+umtN2fyMG7nXPesDN\nhy/8Jor+76AcMP9zj5GmyBcvIA986xBUKBxtPGtfxv6yvEaVTI+dpPUtXDDyreIz\nidKCok2/Ob7q4JN1hfLlz11SF4/uhp4aTp8piCOv50T9AqE3bZA7Kszq5ZWPuXt6\nFBJs+UrYInY4ojxtvavi8uKODpuQUkx+kp180Np33X3zQOufkPTtdJSMcle0Ng3w\naFP+JArOOZ0yLt3m9Lch54wGeH13PQezPdxxR6PWWtKq6OF+IujjOQnbeGqIwAvv\nVALo0vOjkdPFCKfyUAMXiyPV3daQBgGY6rZ9AgMBAAGgADANBgkqhkiG9w0BAQsF\nAAOCAQEAa4/kRqJP89GrajHfa/umZ1t/jE2uvT7QgG/DoFheADxReIpedfJPSBgJ\nm3VTaQfwELaXYK2Ho2TAtsusv1NvlYzcFAXZybsvgJi3vRMeiLEVQiFiQzQzzFIl\nAXQ7amyE9Bsj/IOyyHcooalmJxrQWxVS6CTH/k/lyB0cqAZHeCyW60nVmKUieNpn\n4dnFi+NrK7HqDGAM9G5C4CVKTqdmKeDRe+GwS0JfcMm6TIt8tyl/mYKjBiuzyDfh\nqcTlG+D3Jx6x+AEqUX1GichzWYK+zbt2L3aGrq7j8ewFlY6BcfsyeRyRQCkhis5o\ns5i/3i8OkGNbr6zMU83FqbQ0rBiusw==\n-----END CERTIFICATE REQUEST-----\n"

type CertificateRequest struct {
	Data *Certificate `json:"data"`
}

type Certificate struct {
	Certificate string `json:"certificate"`
}

func TestRsaCertificate(t *testing.T) {
	err := Init("ca.pem", "ca.key")
	// if err != nil {
	// 	panic(err)
	// }
	assert.Nil(t, err)
	err = GenerateTempCertificate2Local("test")
	assert.Nil(t, err)
	// if err != nil {
	// 	panic(err)
	// }
	// t.Log(privateKey.Size())
	// t.Log(data)

	// data, err := GenerateTempCertificateFromCSR(ps)
	// assert.Nil(t, err)
	// cert := CertificateRequest{
	// 	Data: &Certificate{
	// 		Certificate: string(data),
	// 	},
	// }
	// result, err := json.Marshal(&cert)
	// assert.Nil(t, err)
	// fmt.Println(string(result))

	// err = GenerateTempCertificate2Local("local")
	// assert.Nil(t, err)
}

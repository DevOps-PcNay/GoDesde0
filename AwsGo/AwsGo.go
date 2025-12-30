package AwsGo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// Funcion para conectarse a AWS y cargar la configuracion
func InicializoAws() {
	// Para iniciaizar (resetar) los valores de las variables.
	Ctx = context.TODO()

	// Permite conectarse a AWS
	// Se deja fija la region a donde nos conectaremos "us-east-1"(Norte de Virginia)
	//
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("Error al cargar la configuracion de AWS: " + err.Error())
	}

}

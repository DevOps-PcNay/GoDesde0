package Secretmanager

import (
	"encoding/json" // Permite trabajar con la condificacion JSon.
	"fmt"
	"github.com/DevOps-PcNay/Twitter-GoLand/AwsGo"
	"github.com/DevOps-PcNay/Twitter-GoLand/Models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Para obtener el Password desde el modulo Secret Manager de AWS
// "Models.Secret" = Es el modelo que se creo para que contenga los datos del Secret Manager, es de tipo struct.
func GetSecret(secretName string) (Models.Secret, error) {
	// Se utilizara para devolverla con los datos del Secret Manager(Clave)
	var datosSecret Models.Secret

	fmt.Println("> Pido Secreto " + secretName) // Para que lo envie a Cloudwatch

	// Inicializamos Secret Manager.
	// AwsGo.Cfg = Se inicializa en otro archivo, pero se puede utilizar en otras funciones (Paquetes)
	svc := secretsmanager.NewFromConfig(AwsGo.Cfg)

	// Se obtiene el Pawwsord dese el modulo "Secret Manager"
	clave, err := svc.GetSecretValue(AwsGo.Ctx, &secretsmanager.GetSecretValueInput{
		// Las cadenas en AWS las maneja de forma diferente, por esta razon se agrega la funcion "String"
		SecretId: aws.String(secretName), // Obtiene los datos  que estan en el servicio AWS de Secret Manager
	})

	if err != nil {
		fmt.Println(">> Error al obtener el Secret Manager " + err.Error()) // Este mensaje de error es escrito en Cloudwatch
		return datosSecret, err                                             // "datosSecret" lo retornara vacio ya que hubo error.
	}

	// Convertirlo a formato JSON
	// clave,err  = En "clave" se encuentran los  5 datos que se encuentran en "Secret Manger"
	// &datosSecret = Siempre va un puntero no va la referencia directa.
	// "clave.SecretString" = Contiene el password que se definio  [ clave, err ]
	// Puntero = Es la direccion de memoria.
	// Es mas rapido trabajar con punteros.
	// &datosSecret = Es donde se asignara el Password, en este caso una vez que se convierta a formato JSon.
	// "clave" = viene cuando se definio al definir la funcion "svc.GetSecretValue"
	// Es un puntero es una direccion de memoria donde esta guardada la variable "datosSecret"
	// En esta variable "clave" vienen apilados los datos del Secret Manager.
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura de Secret OK " + secretName) // Para que lo envie a Cloudwatch

	return datosSecret, nil
}

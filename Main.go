package Main

// Se agrega este comentario para probar el Push cuando es SSH y fue generado solamente local, para despues fucionar.
import (
	"context"
	"fmt"
	"github.com/DevOps-PcNay/Twitter-GoLand/AwsGo"
	"github.com/DevOps-PcNay/Twitter-GoLand/Models"
	"github.com/DevOps-PcNay/Twitter-GoLand/Secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	fmt.Println("Ejecutando Lambda")

	// Inicializo AWS
	AwsGo.InicializoAws()

	// Aquí puedes agregar la lógica de tu Lambda
	if !ValidoParametros() {

		res = &events.APIGatewayProxyResponse{
			// Esta informacion la envia al Postman, solo una vez ya que cuando se cuando se definan los parametros  ya no accesara.
			StatusCode: 400,
			Body:       "Error En las variables de entorno 'SecretName','BucketName','UrlPrefix' ",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil

	} // if !ValidoParametros() {

	// "SecretName" = Variable de entorno.
	// "SecretName" = Aun no se crea la variable de Entorno, se creara en una archivo y se agregara al proyecto.
	SecretModel, err := Secretmanager.GetSecret(os.Getenv("SecretName"))

	// Como buena practica es validar si hubo error, ya que si se colocan instrucciones antes, se pierde el seguimiento de los errores.

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			// Esta informacion la envia al Postman, solo una vez ya que cuando se cuando se definan los parametros  ya no accesara.
			StatusCode: 400,
			Body:       "Error En la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil

	} // if err != nil {

	// -1 = Que inicie la busqueda desde el inicio de la cadena.
	// Quitar en la url "twitterGo/",. que le quite "twitterGo"

	path := strings.Replace(request.PathParameters["twitterGo"], os.Getenv(("urlPrefix")), "", -1)

	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("path"), path)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("method"), request.HTTPMethod)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("user"), SecretModel.Username)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("password"), SecretModel.Password)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("host"), SecretModel.Host)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("database"), SecretModel.Database)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("jwt"), SecretModel.JWTSign)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("body"), request.Body)
	AwsGo.Ctx = context.WithValue(AwsGo.Ctx, Models.Key("bucketname"), os.Getenv("BucketName"))

} //EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

// Para que valide que la lambda debe recibir 3 parametros
func ValidoParametros() bool {
	// Para leer los parametros
	// Esta funcion "os.LookEnv" retorna dos valores "Cadena", "Bool"(Si tiene valor o no)

	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro // Si no tiene, retorna false.
	}

	_, traeParametro = os.LookupEnv("BucketName") // Traer nombbre del Bucket S3 (donde se almacenan las avatar.)
	if !traeParametro {
		return traeParametro // Si no tiene, retorna false.
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro // Si no tiene, retorna false.
	}

	return traeParametro
}

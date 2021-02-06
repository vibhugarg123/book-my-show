package appcontext

func Init() {
	SetupLogger()
	InitMySqlConnection()
}

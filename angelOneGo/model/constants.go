package model

import "time"

// LogLevels
var (
	LogLevel        = "log-level"
	LogLevelInfo    = "info"
	LogLevelDebug   = "debug"
	LogLevelError   = "error"
	LogLevelWarning = "warn"
)

// PackageLevels
var (
	PgressPackageLevel = "pgress"
	ServerPackageLavel = "server"
	ApiPackageLevel    = "api"
	MainPackageLevel   = "main"
	UtilPackageLevel   = "util"
)

// Functions
var (

	//server
	NewServer = "NewServer"

	//store
	NewStore = "NewStore"

	//actuator methods
	ActuatorsURL        = ""
	GetActuators        = "GetActuators"
	GetActuatorByFilter = "GetActuatorByFilter"
	GetActuator         = "GetActuator"
	CreateActuator      = "CreateActuator"
	UpdateActuator      = "UpdateActuator"
	DeleteActuator      = "DeleteActuator"

	//floor methods
	GetFloors        = "GetFloors"
	GetFloorByFilter = "GetFloorByFilter"
	GetFloor         = "GetFloor"
	CreateFloor      = "CreateFloor"
	UpdateFloor      = "UpdateFloor"
	DeleteFloor      = "DeleteFloor"

	//house methods
	GetHouses        = "GetHouses"
	GetHouseByFilter = "GetHouseByFilter"
	GetHouse         = "GetHouse"
	CreateHouse      = "CreateHouse"
	UpdateHouse      = "UpdateHouse"
	DeleteHouse      = "DeleteHouse"

	//organization methods
	GetOrganizations        = "GetOrganizations"
	GetOrganizationByFilter = "GetOrganizationByFilter"
	GetOrganization         = "GetOrganization"
	CreateOrganization      = "CreateOrganization"
	UpdateOrganization      = "UpdateOrganization"
	DeleteOrganization      = "DeleteOrganization"

	//package methods
	GetPackages        = "GetPackages"
	GetPackageByFilter = "GetPackageByFilter"
	GetPackage         = "GetPackage"
	CreatePackage      = "CreatePackage"
	UpdatePackage      = "UpdatePackage"
	DeletePackage      = "DeletePackage"

	//poinRate methods
	GetPointRates        = "GetPointRates"
	GetPointRateByFilter = "GetPointRateByFilter"
	GetPointRate         = "GetPointRate"
	CreatePointRate      = "CreatePointRate"
	UpdatePointRate      = "UpdatePointRate"
	DeletePointRate      = "DeletePointRate"

	//room methods
	GetRooms        = "GetRooms"
	GetRoomByFilter = "GetRoomByFilter"
	GetRoom         = "GetRoom"
	CreateRoom      = "CreateRoom"
	UpdateRoom      = "UpdateRoom"
	DeleteRoom      = "DeleteRoom"

	//sensor methods
	GetSensors        = "GetSensors"
	GetSensorByFilter = "GetSensorByFilter"
	GetSensor         = "GetSensor"
	CreateSensor      = "CreateSensor"
	UpdateSensor      = "UpdateSensor"
	DeleteSensor      = "DeleteSensor"

	//sensorReading methods
	GetSensorReadings        = "GetSensorReadings"
	GetSensorReadingByFilter = "GetSensorReadingByFilter"
	GetSensorReading         = "GetSensorReading"
	CreateSensorReading      = "CreateSensorReading"
	UpdateSensorReading      = "UpdateSensorReading"
	DeleteSensorReading      = "DeleteSensorReading"

	//user methods
	GetUsers        = "GetUsers"
	GetUser         = "GetUser"
	CreateUser      = "CreateUser"
	UpdateUser      = "UpdateUser"
	DeleteUser      = "DeleteUser"
	SignUp          = "SignUp"
	SignIn          = "SignIn"
	GetUserByFilter = "GetUserByFilter"

	Init = "init"
	Log  = "log"

	AuthMiddleware         = "AuthMiddleware"
	AuthMiddlewareComplete = "AuthMiddlewareComplete"
	SetLimitAndPage        = "setLimitAndPag e"
	SetDateRangeFilter     = "setDateRangeFilter"
)

// user type
var (
	HomeAutomationOwner = "HomeAutomationOwner"
	SuperAdminUser      = "superAdmin"
	AdminUser           = "Admin"
	NormalUser          = "User"
	UserTypes           = []string{"HomeAutomationOwner", "superAdmin", "Admin", "User"}
)

// package type
var (
	MonthelyPackage   = "MonthelyPackage"
	QuaterlyPackage   = "QuaterlyPackage"
	HalfYearlyPackage = "HalfYearlyPackage"
	YearlyPackage     = "YearlyPackage"
)

var (
	TokenExpiration = time.Hour * 24
)

var SecretKey = []byte("homeAutomation-secreat-key")

// General
var (
	Value    = "value"
	Email    = "email"
	Password = "password"
	UserID   = "userID"
	Expire   = "exp"

	Authorization = "X-Token"

	DSN = "host=localhost user=iot password=iot dbname=homeautomation port=5432 sslmode=disable"

	DataPerPage = "limit"
	PageNumber  = "page"
	StartDate   = "start_date"
	EndDate     = "end_date"
	TimeLayout  = "2006-01-02 15:04:05.000 -0700"
)

module github.com/kimulaco/daily-steam-api/cmd/gcf

go 1.16

replace daily-steam.app/GetNewApp => ../../GetNewApp

require (
	daily-steam.app/GetNewApp v0.0.0-00010101000000-000000000000
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.0
)

module main.go

go 1.15

replace example.com/data => ../data

replace example.com/filehandling => ../filehandling

replace example.com/types => ../types

replace example.com/mysql => ../mysql

require (
	example.com/data v0.0.0-00010101000000-000000000000
	example.com/filehandling v0.0.0-00010101000000-000000000000
	example.com/mysql v0.0.0-00010101000000-000000000000
	example.com/types v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.5.0
)

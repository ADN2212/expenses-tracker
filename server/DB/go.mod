module expenses-tracker.com/DB

go 1.23.4

require (
	expenses-tracker.com/IO v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

require github.com/joho/godotenv v1.5.1

replace expenses-tracker.com/IO => ../IO

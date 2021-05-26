package main

func main() {
	a := App{
		StorageType: getEnv("STORAGE_TYPE", "local"),
	}
	a.Run(
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASS", ""),
		getEnv("DB_NAME", "postgres"),
	)

}

package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/qytela/example-project-layout/internal/app/providers/database"
	sb "github.com/qytela/example-project-layout/internal/app/providers/supabase"
	"github.com/supabase-community/supabase-go"
)

func ProvideDB() *sqlx.DB {
	return database.InitializeDB()
}

func ProvideSupabase() *supabase.Client {
	return sb.InitializeSupabase()
}

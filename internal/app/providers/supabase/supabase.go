package supabase

import (
	"os"
	"sync"

	"github.com/qytela/example-project-layout/internal/pkg/logger"
	"github.com/supabase-community/supabase-go"
)

var (
	client *supabase.Client
	once   sync.Once
)

func InitializeSupabase() *supabase.Client {
	once.Do(func() {
		var (
			SUPABASE_API_URL  = os.Getenv("SUPABASE_API_URL")
			SUPABASE_ANON_KEY = os.Getenv("SUPABASE_ANON_KEY")
		)

		conn, err := supabase.NewClient(SUPABASE_API_URL, SUPABASE_ANON_KEY, &supabase.ClientOptions{})
		if err != nil {
			logger.MakeLogEntry(nil).Panic("failed to connect supabase: ", err)
		}

		client = conn

		logger.MakeLogEntry(nil).Info("supabase connected successfully")
	})

	return client
}

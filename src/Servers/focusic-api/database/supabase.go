package database

import (
	"focusic-api/config"
	supa "github.com/nedpals/supabase-go"
)

var SupaBaseClient *supa.Client

func Init() {
	url := config.Config
	key := ""
	SupaBaseClient = supa.CreateClient(url, key)
}

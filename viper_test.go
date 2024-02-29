package golang_viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar viper", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "Roni Purwanto", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestYaml(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar viper", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "Roni Purwanto", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestEnvFile(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config.env ")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Belajar Viper", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "Roni Purwanto", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
}

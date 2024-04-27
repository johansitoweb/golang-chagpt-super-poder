package gpt

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var VarGpt *Gpt

type Gpt struct {
	Api   string
	Model string
	Token int
}

func (t *Gpt) SetApi(configfile string, configFileType string, token int, model string) {
	t.getApiKey(configfile, configFileType)
	if model == "" {
		t.Model = openai.GPT3TextDavinci002
	}
	t.Model = model
	t.Token = token

}

func (t *Gpt) ChatWithTextDavinci003(text string) (string, error) {
	t.Model = openai.GPT3TextDavinci003
	return t.chatGPTsimple(text)
}

func (t *Gpt) ChatWithTextDavinci002(text string) (string, error) {
	t.Model = openai.GPT3TextDavinci002
	return t.chatGPTsimple(text)
}

func (t *Gpt) getApiKey(configfile string, configFileType string) (string, error) {
	viper.SetConfigName(configfile)
	viper.SetConfigType(configFileType)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error al leer el archivo de configuración:", err)
		return "", errors.New("no se encontro las apikey ")
	}

	t.Api = viper.GetString("Api")
	return t.Api, nil
}

func (t *Gpt) chatGPTsimple(text string) (string, error) {
	client := openai.NewClient(t.Api)
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:     t.Model,
			MaxTokens: t.Token,
			Prompt:    text,
		},
	)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return "", err
	}
	return resp.Choices[0].Text, nil
}

package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Formatter fornece funções de formatação
type Formatter struct {
	ID        string
	CreatedAt time.Time
}

// NewFormatter cria uma nova instância do formatador
func NewFormatter() *Formatter {
	return &Formatter{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
	}
}

// FormatNumber formata números com casas decimais
func (f *Formatter) FormatNumber(number float64, decimals int) string {
	format := fmt.Sprintf("%%.%df", decimals)
	return fmt.Sprintf(format, number)
}

// FormatCurrency formata valores como moeda
func (f *Formatter) FormatCurrency(value float64, currency string) string {
	formatted := f.FormatNumber(value, 2)
	return fmt.Sprintf("%s %s", currency, formatted)
}

// FormatOperation formata operações matemáticas
func (f *Formatter) FormatOperation(operation string, a, b, result float64) string {
	return fmt.Sprintf("%.2f %s %.2f = %.2f", a, operation, b, result)
}

// FormatError formata erros de forma amigável
func (f *Formatter) FormatError(err error) string {
	return fmt.Sprintf("❌ Erro: %s", err.Error())
}

// FormatSuccess formata mensagens de sucesso
func (f *Formatter) FormatSuccess(message string) string {
	return fmt.Sprintf("✅ %s", message)
}

// FormatHeader cria cabeçalhos formatados
func (f *Formatter) FormatHeader(title string) string {
	line := strings.Repeat("=", len(title)+4)
	return fmt.Sprintf("\n%s\n  %s\n%s\n", line, title, line)
}

// FormatTime formata timestamp
func (f *Formatter) FormatTime(t time.Time) string {
	return t.Format("15:04:05")
}

// GetInfo retorna informações do formatador
func (f *Formatter) GetInfo() string {
	return fmt.Sprintf("Formatador (ID: %s) criado em %s",
		f.ID, f.FormatTime(f.CreatedAt))
}

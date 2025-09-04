package training

import (
	"fmt"
	"os"
	"strings"
)

type DataProvider interface {
  GetData(source string, separator string) ([]string, error)
}

type FileDataProvider struct {}

type SymbolsDataProvider struct {}

// Слой данных, формируем данные с которыми работаем, если из файла
func (f *FileDataProvider) GetData(filePath string, separator string) ([]string , error) {
  fileData, err := os.ReadFile(filePath)
    if err != nil {
	    return make([]string, 0), fmt.Errorf("ошибка чтения файла: %w", err)
    }
  fileSliceData := strings.Split(string(fileData), separator) 
  return fileSliceData, nil
}

func (f *SymbolsDataProvider) GetData(text string, separator string) ([]string, error) {
     return strings.Split(text, separator), nil
}

func CountingSymbols(provider DataProvider, source string, separator string) (int,error) {
  data, err := provider.GetData(source, separator)
  if err != nil {
    return 0, fmt.Errorf("произошла ошибка: %w", err)
  }
  count := CountSymbols(data)
  return count, nil
}

type WordCounter struct {
  provider DataProvider
}

func (w *WordCounter) Count(source string, separator string) (int, error) {
    data, err := w.provider.GetData(source, separator)
    if err != nil {
      return 0, err
    }
    return CountSymbols(data), nil
}


// Конструктор, аналог
func NewWordCounter(provider DataProvider) *WordCounter {
    return &WordCounter{
      provider: provider,
    }
}

func Main() {
  wordCounter := NewWordCounter(&FileDataProvider{})
  count, err := wordCounter.Count("hello.txt", " ")
  if err != nil {
	  fmt.Println(err)
  }
  fmt.Println(count)
}


// Функция бизнес логики для подсчета символов
func CountSymbols(data []string) int {
  count := 0
  for _, symbol := range data {
    if symbol != "" {
	    count += 1
    }
  }
  return count
}

// Теперь ты реализовал:
// Dependency Injection - провайдер внедряется через конструктор
// Инкапсуляцию - внешний код не знает о провайдере
// Гибкость - можно легко подменить провайдер
package training

import (
	"fmt"
	"os"
	"strings"
)

func Main() {
  count, err := CountingSymbols("", "", false)
  if err != nil {
	fmt.Println(err)
  }
  fmt.Println(count)
}


func CountingSymbols(targetEntity string, separator string, isFile bool) (int,error) {
  var targetData []string
  var err error
  if (isFile) {
	targetData, err = GetReadedFileData(targetEntity, separator)
	if err != nil {
		return 0, err
	}
  } else {
	targetData = strings.Split(targetEntity, separator) 
  }
  count := CountWords(targetData)
  return count, nil
}


// Слой данных, формируем данные с которыми работаем, если из файла
func GetReadedFileData(filePath string, separatorPayload string) ([]string, error) { 
  fileData, err := os.ReadFile(filePath)
  if err != nil {
	return make([]string, 0), fmt.Errorf("ошибка чтения файла: %w", err)
  }
  fileSliceData := strings.Split(string(fileData), separatorPayload) 
  return fileSliceData, nil
}
// Функция бизнес логики для подсчета символов
func CountWords(data []string) int {
  count := 0
  for i := 0; i < len(data) ; i++ {
	count += 1
  }
  return count
}
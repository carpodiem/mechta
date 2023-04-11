package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	numApartments    = 6
	gasPricePerCubic = 6.87689
	dataFile         = "year2023.csv"
	initialReading   = 449207
)

func main() {
	apartmentNames := map[int]string{
		1: "1А",
		2: "1Б",
		3: "1В",
		4: "1Г",
		5: "1Д",
		6: "1Е",
	}

	// 1. Запрашиваем текущее показание счетчика и остаток на балансе
	var currentReading, pmonth int64
	var balance float64

	fmt.Print("Введите текущее показание счетчика: ")
	_, _ = fmt.Scanln(&currentReading)

	fmt.Print("Введите остаток на балансе за предыдущий месяц: ")
	_, _ = fmt.Scanln(&balance)

	fmt.Print("Введите месяц за который вносятся показания (от 1 до 12): ")
	_, _ = fmt.Scanln(&pmonth)

	// 2. Получаем предыдущее показание из файла
	prevReading, err := readPreviousReading(dataFile)
	if err != nil {
		fmt.Println("Ошибка при чтении предыдущего показания:", err)
		return
	}

	// 3. Вычисляем потребление и стоимость потребленного газа
	consumption := float64(currentReading - prevReading)
	totalCost := consumption*gasPricePerCubic + balance

	// 4. Вычисляем среднюю стоимость для каждой квартиры
	avgCost := totalCost / float64(numApartments)

	// 5. Записываем среднюю стоимость для 4-й квартиры
	apartmentCosts := make([]float64, numApartments)
	apartmentCosts[5] = avgCost
	roundedCost := math.Round(avgCost)

	// 6. Вычисляем стоимость для остальных квартир с учетом ротации
	for i := 0; i < len(apartmentCosts)-1; i++ {
		rotation := float64((int(pmonth)+i)%5 - 2)
		apartmentCosts[i] = roundedCost - rotation
	}

	// 7. Сохраняем данные в файл
	err = saveDataToFile(dataFile, currentReading, apartmentCosts)
	if err != nil {
		fmt.Println("Ошибка при сохранении данных:", err)
		return
	}

	// 8. Выводим результат для каждой квартиры
	for i, cost := range apartmentCosts {
		apartmentName := apartmentNames[i+1]
		fmt.Printf("Квартира %s - сумма к оплате: %.2f\n", apartmentName, cost)
	}
}

func readPreviousReading(filename string) (int64, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return initialReading, nil
	} else if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	latestRecord := records[len(records)-1]
	prevReading, err := strconv.ParseInt(latestRecord[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return prevReading, nil
}

func saveDataToFile(filename string, currentReading int64, apartmentCosts []float64) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{strconv.FormatInt(currentReading, 10)}
	for _, cost := range apartmentCosts {
		record = append(record, strconv.FormatFloat(cost, 'f', 2, 64))
	}

	return writer.Write(record)
}

package ipgen

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomIPv4() string {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
	)
}

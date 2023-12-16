package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// УСЛОВИЕ:
// Реализовать конкурентную запись данных в map

// можно запустить с флагом race, чтобы проверить, не возникает ли состояния гонки
// 3 СПОСОБА ->

// СПОСОБ 1:
//......................................................................................................................

// MyDataModel1 оборачивает map, чтобы обеспечить доступ на чтение\запись через методы структуры, которые гарантируют использование примитивов синхронизации
type MyDataModel1 struct {
	data map[string]int
	// можно просто встроить sync.Mutex без назначения ему имени mutex, но это сделано для понятного изложения
	mutex sync.Mutex // Механизм блокировки, который позволяет только одной горутине одновременно выполнять критическую секцию кода, защищенную мьютексом (чтение запись).
}

// NewMyDataModel1 Требуется конструктор тк map не инициализируется пустой, а имеет значение nil при простом объявлении объекта структуры
func NewMyDataModel1() *MyDataModel1 {
	return &MyDataModel1{data: make(map[string]int)}
}

// ReadUseMutex читает данные из map (горутины делают это по очереди, даже если вызвали в одно время)
// тк метод работает с mutex, то ресивер должен быть по указателю. тк состояние mutex будет изменяться
func (m *MyDataModel1) ReadUseMutex(key string) (int, bool) {
	// Если Lock не используется, то продолжает работу. Иначе (значит что другая горутина сейчас читает данные) блокируется на этом месте пока иная горутина не вызовет Unlock
	m.mutex.Lock()
	// Сообщаем что операция чтения завершена
	defer m.mutex.Unlock()
	// ok говорит, существует ли такой ключ в map, Если не использовать его, при отсутствии ключа будет возвращаться 0 (для int), а это может быть валидным значением в нашей модели данных.
	value, ok := m.data[key]
	return value, ok
}

// WriteUseMutex записывает данные в map (горутины делают это по очереди, даже если вызвали в одно время)
func (m *MyDataModel1) WriteUseMutex(key string, value int) {
	// Если Lock не используется, то продолжает работу. Иначе (значит что другая горутина сейчас записывает данные) блокируется на этом месте пока иная горутина не вызовет Unlock
	m.mutex.Lock()
	// Сообщаем что запись больше не идет после выполнения функции
	defer m.mutex.Unlock()
	// записываем данные в map
	m.data[key] = value

}

// СПОСОБ 2:
//......................................................................................................................

// MyDataModel2 оборачивает map, чтобы обеспечить доступ на чтение\запись через методы структуры, которые гарантируют использование примитивов синхронизации
type MyDataModel2 struct {
	data map[string]int
	// RWMutex обеспечивает последовательную запись и позволяет осуществить параллельное чтение.
	// Более эффективно если приложение чаще читает данные чем записывает
	mutex sync.RWMutex
}

// NewMyDataModel2 Требуется конструктор тк map не инициализируется пустой, а имеет значение nil при простом объявлении объекта структуры
func NewMyDataModel2() *MyDataModel2 {
	return &MyDataModel2{data: make(map[string]int)}
}

// ReadUseRWMutex читает данные из map при этом несколько горутины могут делать это параллельно.
func (m *MyDataModel2) ReadUseRWMutex(key string) (int, bool) {
	// RLock не блокируется даже если она уже вызвана в другой горутине, операция чтения продолжится параллельно с иной горутиной
	m.mutex.RLock()
	// Сообщаем что операция чтения завершена
	defer m.mutex.RUnlock()
	// ok говорит, существует ли такой ключ в map, Если не использовать его, при отсутствии ключа будет возвращаться 0 (для int), а это может быть валидным значением в нашей модели данных.
	value, ok := m.data[key]
	return value, ok
}

// WriteUseRWMutex записывает данные в map (горутины делают это по очереди, даже если вызвали в одно время)
func (m *MyDataModel2) WriteUseRWMutex(key string, value int) {
	// Если Lock не используется, то продолжает работу. Иначе (значит что другая горутина сейчас записывает данные) блокируется на этом месте пока иная горутина не вызовет Unlock
	m.mutex.Lock()
	// Сообщаем что запись больше не идет после выполнения функции
	defer m.mutex.Unlock()
	// записываем данные в map
	m.data[key] = value

}

func main() {
	//go run -race .\tasks\7\

	// раскомментировать одно из трех и запустить:
	//withMutex()   // проверка способа 1
	withRWMutex() // проверка способа 2

	// можно еще свой mutex сделать используя int и sync/atomic, но тут уже не про мапу, а про mutex поэтому реализовывать не буду
}

// withMutex спавнит множество ридеров и врайтеров в мапу. они бесконечно читают мапу и пишут в нее
func withMutex() {
	gReadersNum := 10
	gWritersNum := 3

	fmt.Println("map read/write and Mutex:")
	dataModel := NewMyDataModel1() // Тут потокобезопасность мапе обеспечивает Mutex
	mapKeys := [...]string{"a", "b", "c", "d", "e"}
	// spawn writers
	for i := 0; i < gWritersNum; i++ {
		go func(dm *MyDataModel1) {
			for {
				// writes random numbers into 5 keys
				dm.WriteUseMutex(mapKeys[rand.Intn(len(mapKeys))], rand.Intn(1000)) // mapKeys доступна тут тк переменная захвачена
			}
		}(dataModel)
	}

	// spawn readers
	for i := 0; i < gReadersNum; i++ {
		go func(dm *MyDataModel1) {
			for {
				// reads random key from map
				key := mapKeys[rand.Intn(len(mapKeys))] // mapKeys доступна тут тк переменная захвачена.
				val, _ := dm.ReadUseMutex(key)
				fmt.Printf("%s: %d\n", key, val)

			}
		}(dataModel)
	}

	select {} // тут мы навсегда

}

// withRWMutex спавнит множество ридеров и врайтеров в мапу. они бесконечно читают мапу и пишут в нее
func withRWMutex() {
	gReadersNum := 10
	gWritersNum := 3

	fmt.Println("map read/write and RWMutex:")
	dataModel := NewMyDataModel2() // Тут потокобезопасность мапе обеспечивает RWMutex
	mapKeys := [...]string{"a", "b", "c", "d", "e"}
	// spawn writers
	for i := 0; i < gWritersNum; i++ {
		go func(dm *MyDataModel2) {
			for {
				// writes random numbers into 5 keys
				dm.WriteUseRWMutex(mapKeys[rand.Intn(len(mapKeys))], rand.Intn(1000)) // mapKeys доступна тут тк переменная захвачена
			}
		}(dataModel)
	}

	// spawn readers
	for i := 0; i < gReadersNum; i++ {
		go func(dm *MyDataModel2) {
			for {
				// reads random key from map
				key := mapKeys[rand.Intn(len(mapKeys))] // mapKeys доступна тут тк переменная захвачена.
				val, _ := dm.ReadUseRWMutex(key)
				fmt.Printf("%s: %d\n", key, val)

			}
		}(dataModel)
	}

	select {} // тут мы навсегда

}

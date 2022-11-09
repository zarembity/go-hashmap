/*
Создание хэш таблицы (спринт 4, финальная задача 2)

Посылка:
https://contest.yandex.ru/contest/24414/run-report/59436874/

Алгоритм работает за O(1) сложности, если исключить моменты коллизий.
1. Создаем таблицу с кол-вом корзин в рассчете на на кол-во элементов. В данном случе для 100000 элементов взято простое число
100049
2. Добавление эелемнета (метод put) - рассчитываем номер корзины исходя из теории 4 главы яндекс практикума - делением на остаток
3. Решение коллизий - если элемент существует в корзине (колизия) добавлем его в начало связного списка
4. Удаление элемента (метод delete) - пересоздаем связный список, исключая удаляемый элемент.
5. Вывод элемента (метод get) - заходим в нужную корзину и походим по связному списку и выводим элемент, если такой имеется.

 */
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	hashStack struct {
		m         int
		hashTable []*listNode
	}

	listNode struct {
		key      int
		val      int
		previews *listNode
	}
)

var (
	scanner *bufio.Scanner
	errNone = errors.New("None")
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024*5)
	scanner.Buffer(buf, 1024*1024*5)
}

func NewHT(m int) *hashStack {
	return &hashStack{
		m:         m,
		hashTable: make([]*listNode, m),
	}
}

func main() {
	var (
		n = scanInt()
	)

	ht := NewHT(100049)
	for i := 0; i < n; i++ {
		ht.execCommand(scanStr())
	}

}

func (ht *hashStack) execCommand(str string) {
	commandList := strings.Split(str, " ")
	switch commandList[0] {
	case "put":
		ht.put(strToInt(commandList[1]), strToInt(commandList[2]))
	case "get":
		ht.get(strToInt(commandList[1]))
	case "delete":
		ht.delete(strToInt(commandList[1]))
	}
}

func (ht *hashStack) put(key, val int) {
	b := ht.bucket(key)
	list := ht.hashTable[b]
	if list == nil {
		ht.hashTable[b] = &listNode{key: key, val: val}
		return
	}

	if key == list.key {
		ht.hashTable[b].val = val
		return
	}

	ht.hashTable[b] = &listNode{val: val, key: key, previews: list}
}

func (ht *hashStack) get(key int) {
	b := ht.bucket(key)
	list := ht.hashTable[b]
	if list == nil {
		fmt.Println(errNone.Error())
		return
	}

	for {
		if list.key == key {
			fmt.Println(list.val)
			return
		}
		if list.previews == nil {
			break
		}
		list = list.previews
	}

	fmt.Println(errNone.Error())
}

func (ht *hashStack) delete(key int) {
	b := ht.bucket(key)
	list := ht.hashTable[b]
	if list == nil {
		fmt.Println(errNone.Error())
		return
	}

	var (
		newList = &listNode{}
		isFound bool
		isNew   = true
	)

	for {
		if list.key == key {
			fmt.Println(list.val)
			isFound = true
		} else {
			if isNew {
				newList.key = list.key
				newList.val = list.val
				isNew = false
			} else {
				newList = &listNode{key: list.key, val: list.val, previews: newList}
			}
		}

		if list.previews == nil {
			break
		}

		list = list.previews
	}

	if !isFound {
		fmt.Println(errNone.Error())
	}

	ht.hashTable[b] = newList
}

func (ht *hashStack) bucket(k int) (result int) {
	if k < 0 {
		ost := k % ht.m
		tRes := ost * ht.m
		return k - tRes
	} else {
		return k % ht.m
	}
}

func strToInt(val string) int {
	intVar, _ := strconv.Atoi(val)
	return intVar
}

func scanInt() int {
	scanner.Scan()
	return strToInt(scanner.Text())
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

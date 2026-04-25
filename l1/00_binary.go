package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func binary() {
	// Загружаем DLL и функции для выделения и освобождения памяти
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	virtualAlloc := kernel32.NewProc("VirtualAlloc")
	virtualFree := kernel32.NewProc("VirtualFree")

	// Читаем бинарный файл с машинным кодом
	code, err := os.ReadFile("calc.bin")
	if err != nil {
		panic(err)
	}

	// Выделяем исполняемую память
	addr, _, err := virtualAlloc.Call(
		0,
		uintptr(len(code)),
		0x3000,
		0x40,
	)
	if addr == 0 {
		panic(err)
	}
	defer virtualFree.Call(addr, 0, 0x8000)

	// Создаем срез который смотрит прямо на выделенную нами память начиная с числового
	// адреса из VirtualAlloc адреса длиной len(code) байт
	dst := unsafe.Slice((*byte)(unsafe.Pointer(addr)), len(code))
	// Копируем байты в срез памяти
	copy(dst, code)

	// SyscallN вызывает адрес
	ret, _, _ := syscall.SyscallN(addr)
	fmt.Println("Результат:", int32(ret))
}

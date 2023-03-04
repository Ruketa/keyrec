package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {		
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		log.Info().Msg(fmt.Sprintf("You pressed: rune %q, key", char, key))
		//fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
   	if key == keyboard.KeyEsc {
			break
		}
	}	
}
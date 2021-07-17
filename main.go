package main
import  "vermont/crepes/*"

func Crepe() {
	rp := NewReactParser()
	rp.Scrape()

	dep := NewDockerParser()
	dep.Scrape()
}

func main() {


	Crepe()
	
}

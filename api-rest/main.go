package main

func main() {
    a := App{} 
    
    a.Initialize("root", "", "go-api")

    a.Run(":3000")
}

package main


func main () {
	
}

























// package main

// import (
// 	"fmt"
// )

// const BaseUrl string = "http://localhost:8000/"
// const test_access_token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhZGlpIiwiZXhwIjoxNzg4NDUyMjY5fQ.Oyvu2BGsACOt2JtQjDIsbvdYjpl-yro-j9vi_rb2yYw"
// func main(){
// 	fmt.Println("lets handle some web request with specific method")
// 	// fmt.Println("::Fetching Access Token::")
// 	// token := getToken()
// 	// fmt.Println(token)
// 	//



// }

// // this is getting frustrating, why do i need to create http.request for cookie instead of just post and cookie
// // func getMe(token string) string{
// // 	reqBody := strings.NewReader(`

// // 		`)

// // }

// // func getToken() string {
// // 	reqBody := strings.NewReader(`
// // {
// // "username" : "adii",
// // "password" : "1234"
// // }
// // 		`)
// // 	res, err := http.Post(BaseUrl + "/auth/login" , "application/json", reqBody)
// // 	if err != nil{
// // 		fmt.Println(err)
// // 		return ""
// // 	}
// // 	defer res.Body.Close()

// // 	body, err := io.ReadAll(res.Body)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return ""
// // 	}
// // 	fmt.Println(string(body))
// // 	return string(body)
// // }

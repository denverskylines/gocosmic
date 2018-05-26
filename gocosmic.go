package gocosmic

import "github.com/gocosmic/common"

//// OOP GoCosmic Type Should contain all Cosmic.js Api end points
type GoCosmic struct { }

/// POST JSON 
func GoCosmic(url) Post(url) string {
    Common(url).Post()
    	//
    // Process response
    //
    println("")
    if resp.Status() == 201 {
      fmt.Printf("CosmicJs Response:", res.body)
    } else {
      fmt.Println("Bad response status from CosmicJs server")
      fmt.Printf("\t Status:  %v\n", resp.Status())
      fmt.Printf("\t Message: %v\n", e.Message)
      fmt.Printf("\t Errors: %v\n", e.Message)
      pretty.Println(e.Errors)
    }
    println("")
    return res.body
}

//// GET JSON
func GoCosmic(url) Get(url) string {
  resp = Common(url).Get()
	//
	// Process response
	//
	println("")
	if resp.Status() == 201 {
		fmt.Printf("CosmicJs Response:", res.body)
	} else {
		fmt.Println("Bad response status from CosmicJs server")
		fmt.Printf("\t Status:  %v\n", resp.Status())
		fmt.Printf("\t Message: %v\n", e.Message)
		fmt.Printf("\t Errors: %v\n", e.Message)
		pretty.Println(e.Errors)
	}
	println("")
}


// func GoCosmic(url) GetImage(url) string {
//   resp = Common(url).Get()
// 	//
// 	// Process response
// 	//
// 	println("")
// 	if resp.Status() == 201 {
// 		fmt.Printf("CosmicJs Response:", res.body)
// 	} else {
// 		fmt.Println("Bad response status from CosmicJs server")
// 		fmt.Printf("\t Status:  %v\n", resp.Status())
// 		fmt.Printf("\t Message: %v\n", e.Message)
// 		fmt.Printf("\t Errors: %v\n", e.Message)
// 		pretty.Println(e.Errors)
// 	}
// 	println("")
// }

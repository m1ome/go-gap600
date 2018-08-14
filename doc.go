// Package g600 is used for communicating from Go with gap600 service:
// http://www.gap600.com/
//
// Example:
//	package main
//
//	import (
//		"fmt"
//
//		gap "github.com/m1ome/go-gap600"
//	)
//
//
//	func main() {
//		client, err := gap.New(gap.Options{
//			APIKey: "YOUR_TOKEN_HERE",
//		})
//
//		if err != nil {
//			panic(err)
//		}
//
//		fee, rec, err := client.RecommendedFee()
//		if err != nil {
//			panic(err)
//		}
//
//		fmt.Printf("Fee: %d, Recommendation: %s\n", fee, rec)
//
//		res, err := client.TransactionConfirm("92db07c2a31b2677dffdf82467693c33eeaba5ced81edd6d9126c697703ab26b", "1NgNmnGTwqjGvQKtqQF8dpBzQUDH45xbiH")
//		if err != nil {
//			panic(err)
//		}
//
//		fmt.Printf("Confirmation: %v\n", res)
//	}
package g600

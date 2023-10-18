package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dpc-sdp/go-section-io"
)

func main() {
	cfg := sectionio.NewConfiguration()
	client := sectionio.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), sectionio.ContextAPIKey, sectionio.APIKey{
		Key: os.Getenv("SECTION_IO_APIKEY"),
	})
	in := &sectionio.AccountApiAccountDomainListOpts{}
	domainlist, _, err := client.AccountApi.AccountDomainList(auth, os.Getenv("SECTION_IO_ACCOUNT_ID"), in)
	if err != nil {
		panic(err)
	}
	fmt.Println(domainlist)
}

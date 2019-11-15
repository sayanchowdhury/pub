package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/devigned/pub/pkg/service"

	"github.com/devigned/pub/pkg/partner"
	"github.com/devigned/pub/pkg/xcobra"
)

func newListCommand(sl service.CommandServicer) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all publishers",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := sl.GetCloudPartnerService()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			publishers, err := client.ListPublishers(ctx)

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			if err := sl.GetPrinter().Print(publishers); err != nil {
				log.Fatalf("unable to print publishers: %v", err)
			}
		}),
	}
	return cmd, nil
}

func printPublishers(publishers []partner.Publisher) {
	bits, err := json.Marshal(publishers)
	if err != nil {
		log.Fatalf("failed to print publishers: %v", err)
	}
	fmt.Print(string(bits))
}

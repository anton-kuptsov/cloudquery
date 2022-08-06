package fetch

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/internal/plugin"
	"github.com/cloudquery/cq-provider-sdk/clients"
	"github.com/cloudquery/cq-provider-sdk/plugins"
	"github.com/cloudquery/cq-provider-sdk/spec"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	fetchShort = "Fetch resources from configured source plugins to destination"
	fetchLong  = `Fetch resources from configured source plugins to destination
	
	This requires a cloudquery.yml file which can be generated by "cloudquery init"
	`
	fetchExample = `  # Fetch configured providers to PostgreSQL as configured in cloudquery.yml
	cloudquery fetch`
)

func NewCmdFetch() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fetch [directory]",
		Short:   fetchShort,
		Long:    fetchLong,
		Example: fetchExample,
		Args:    cobra.RangeArgs(0, 1),
		RunE:    fetch,
	}
	return cmd
}

func fetch(cmd *cobra.Command, args []string) error {
	directory := "."
	if len(args) == 1 {
		directory = args[0]
	}
	fmt.Println("Loading specs from directory: ", directory)
	specReader, err := spec.NewSpecReader(directory)
	if err != nil {
		return errors.Wrapf(err, "failed to load specs from directory %s", directory)
	}
	if len(specReader.Connections()) == 0 {
		fmt.Println("No connections specs found in directory: ", directory)
		return nil
	}
	pm := plugin.NewPluginManager()
	defer pm.CloseAll(cmd.Context())
	for _, connSpec := range specReader.Connections() {
		if err := fetchConnection(cmd.Context(), specReader, connSpec, pm); err != nil {
			return errors.Wrapf(err, "failed to fetch connection %s->%s", connSpec.Source, connSpec.Destination)
		}
		// destSpec := specs.GetDestinatinoByName(connSpec.Destination)
	}

	return nil
}

func fetchConnection(ctx context.Context, specReader *spec.SpecReader, connSpec spec.ConnectionSpec, pm *plugin.PluginManager) error {
	sourceSpec := specReader.GetSourceByName(connSpec.Source)

	sourceClient, err := pm.GetSourcePluginClient(ctx, sourceSpec)
	if err != nil {
		return errors.Wrap(err, "failed to get source plugin client")
	}

	destinationClient, err := pm.GetDestinationClient(
		ctx,
		specReader.GetDestinatinoByName(connSpec.Destination),
		plugins.DestinationPluginOptions{},
	)
	if err != nil {
		return errors.Wrap(err, "failed to get destination plugin client")
	}
	if err := destinationClient.Configure(ctx, specReader.GetDestinatinoByName(connSpec.Destination)); err != nil {
		return errors.Wrap(err, "failed to configure destination plugin client")
	}
	tables, err := sourceClient.GetTables(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tables")
	}
	if err := destinationClient.CreateTables(ctx, tables); err != nil {
		return errors.Wrap(err, "failed to create tables")
	}

	validationResult, err := sourceClient.Configure(ctx, sourceSpec)
	if err != nil {
		return errors.Wrap(err, "failed to configure source plugin client")
	}
	if validationResult != nil && !validationResult.Valid() {
		for _, err := range validationResult.Errors() {
			fmt.Printf("Spec Error at Field: %s, %s\n", err.Field(), err.Description())
		}
		return errors.New("spec validation failed")
	}
	resources := make(chan *clients.FetchResultMessage)
	g, ctx := errgroup.WithContext(ctx)
	fmt.Println("Starting fetch for: ", connSpec.Source, "->", connSpec.Destination)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Fetch(ctx, sourceSpec, resources); err != nil {
			return errors.Wrap(err, "failed to fetch resources")
		}
		return nil
	})

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Fetching"),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
	)

	g.Go(func() error {
		for _ = range resources {
			// fmt.Println("fetched")
			bar.Add(1)
			// if err := destinationClient.Save(ctx, resource); err != nil {
			// 	return errors.Wrapf(err, "failed to save resources")
			// }
			// fmt.Println("save")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		bar.Finish()
		return errors.Wrap(err, "failed to fetch resources")
	}
	bar.Finish()

	fmt.Println("Fetch completed successfully")
	return nil
}

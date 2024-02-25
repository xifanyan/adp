package axc

import (
	"fmt"

	"github.com/xifanyan/adp/pkg/adp"
)

type ApplicationBuilder struct {
	*Application
}

func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{
		&Application{},
	}
}

func (b *ApplicationBuilder) WithID(id string) *ApplicationBuilder {
	b.ID = id
	return b
}

func (b *ApplicationBuilder) WithClient(client *adp.Client) *ApplicationBuilder {
	b.client = client
	return b
}

func (b *ApplicationBuilder) Build() *Application {
	return b.Application
}

type Application struct {
	ID     string
	client *adp.Client
}

type ECA struct {
	Application
}

type RNA struct {
	Application
}

func (app *Application) SingleMindServers() (adp.ListEntitiesResult, error) {
	req := adp.NewListEntitiesRequest(
		adp.WithListEntitiesType("singleMindServer"),
		adp.WithListEntitiesRelatedEntity(app.ID),
	)

	task := &adp.ListEntitiesTask{
		Task: adp.NewTask().WithClient(app.client).WithRequest(req),
	}

	return task.GetResult()
}

func (app *Application) MergingMeta() (adp.ListEntitiesResult, error) {
	req := adp.NewListEntitiesRequest(
		adp.WithListEntitiesType("mergingMeta"),
		adp.WithListEntitiesRelatedEntity(app.ID),
	)

	task := &adp.ListEntitiesTask{
		Task: adp.NewTask().WithClient(app.client).WithRequest(req),
	}

	return task.GetResult()
}

func (app *Application) Datasources() (adp.ListEntitiesResult, error) {
	req := adp.NewListEntitiesRequest(
		adp.WithListEntitiesType("dataSources"),
		adp.WithListEntitiesRelatedEntity(app.ID),
	)

	task := &adp.ListEntitiesTask{
		Task: adp.NewTask().WithClient(app.client).WithRequest(req),
	}

	return task.GetResult()
}

func (app *Application) Taxonomy(engineTaxonomies string, outputTaxonomies string) (adp.TaxonomyStatisticResult, error) {
	req := adp.NewTaxonomyStatisticRequest(
		adp.WithTaxonomyStatisticApplicationIdentifier(app.ID),
		adp.WithTaxonomyStatisticEngineTaxonomies(engineTaxonomies),
		adp.WithTaxonomyStatisticOutputTaxonomies(outputTaxonomies),
	)

	task := &adp.TaxonomyStatisticTask{
		Task: adp.NewTask().WithClient(app.client).WithRequest(req),
	}

	return task.GetResult()
}

func (app *Application) TaxonomyWithProperties(engineTaxonomies string, outputTaxonomies string) (adp.TaxonomyStatisticResult, error) {
	req := adp.NewTaxonomyStatisticRequest(
		adp.WithTaxonomyStatisticApplicationIdentifier(app.ID),
		adp.WithTaxonomyStatisticEngineTaxonomies(engineTaxonomies),
		adp.WithTaxonomyStatisticOutputTaxonomies(outputTaxonomies),
		adp.WithTaxonomyStatisticListCategoryProperties("true"),
	)

	task := &adp.TaxonomyStatisticTask{
		Task: adp.NewTask().WithClient(app.client).WithRequest(req),
	}

	return task.GetResult()
}

func (app *Application) GetCategoriesByTaxonomy(taxonomy string) ([]adp.Category, error) {
	var err error

	taxonomies, err := app.Taxonomy("", taxonomy)
	if err != nil {
		return nil, err
	}

	return taxonomies[0].Category, err
}

func (app *Application) Matters() ([]adp.Category, error) {
	return app.GetCategoriesByTaxonomy("rm_document_hold")
}

func (app *Application) GetCustodiansByMatter(matterID string) ([]adp.Category, error) {
	engineTaxonomies := fmt.Sprintf("rm_document_hold=%s", matterID)

	taxonomies, err := app.Taxonomy(engineTaxonomies, "rm_custodian")
	if err != nil {
		return nil, err
	}

	return taxonomies[0].Category, err
}

func (app *Application) Query(engineTaxonomies string, engineQuery string) (adp.QueryEngineResult, error) {

	req := adp.NewQueryEngineRequest(
		adp.WithQueryEngineApplicationIdentifier(app.ID),
		adp.WithQueryEngineEngineTaxonomies(engineTaxonomies),
		adp.WithQueryEngineEngineQuery(engineQuery),
	)

	task := &adp.QueryEngineTask{
		Task: adp.NewTask().WithClient(app.client).WithRequest(req),
	}

	return task.GetResult()
}

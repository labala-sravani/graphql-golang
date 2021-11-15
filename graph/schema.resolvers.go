package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"projectSchema/generated"
	"projectSchema/models"
)

func (r *characteristicsResolver) Cuisines(ctx context.Context, obj *models.Characteristics) ([]*models.Cuisines, error) {
	fmt.Println("in cuisines")
	name := "sweet"
	urlKey := "some sweet url"
	return []*models.Cuisines{{Name: &name, URLKey: &urlKey}}, nil
}

func (r *characteristicsResolver) FoodCharacteristics(ctx context.Context, obj *models.Characteristics) ([]*models.FoodCharacteristics, error) {
	fmt.Println("in food characteristics")
	isVegetarian := true
	name := "potato"
	return []*models.FoodCharacteristics{{Name: &name, IsVegetarian: &isVegetarian}, nil}, nil
}

func (r *queryResolver) Vendor(ctx context.Context, id string) (*models.Vendor, error) {
	fmt.Println("in vendor")
	return &models.Vendor{Description: "abc", PlatformVendorId: "123"}, nil
}

func (r *vendorResolver) Tes(ctx context.Context, obj *models.Vendor) (*models.TES, error) {
	fmt.Println("in tes")
	return &models.TES{MinimumDeliverTime: "123"}, nil
}

func (r *vendorResolver) Dps(ctx context.Context, obj *models.Vendor) (*models.DPS, error) {
	fmt.Println("in dps")
	return &models.DPS{MinimumOrderAmount: "dps"}, nil
}

func (r *vendorResolver) Schedules(ctx context.Context, obj *models.Vendor) ([]*models.Schedule, error) {
	fmt.Println("in schedule")
	id := "123"
	weekday := "monday"
	return []*models.Schedule{{ID: &id, Weekday: &weekday}}, nil
}

func (r *vendorResolver) Characteristics(ctx context.Context, obj *models.Vendor) (*models.Characteristics, error) {
	return &models.Characteristics{}, nil
}

// Characteristics returns generated.CharacteristicsResolver implementation.
func (r *Resolver) Characteristics() generated.CharacteristicsResolver {
	return &characteristicsResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Vendor returns generated.VendorResolver implementation.
func (r *Resolver) Vendor() generated.VendorResolver { return &vendorResolver{r} }

type characteristicsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type vendorResolver struct{ *Resolver }

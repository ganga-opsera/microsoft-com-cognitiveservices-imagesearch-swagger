package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RelatedCollectionsModule represents the RelatedCollectionsModule schema from the OpenAPI specification
type RelatedCollectionsModule struct {
	Value []ImageGallery `json:"value,omitempty"` // A list of webpages that contain related images.
}

// WebPage represents the WebPage schema from the OpenAPI specification
type WebPage struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// PivotSuggestions represents the PivotSuggestions schema from the OpenAPI specification
type PivotSuggestions struct {
	Pivot string `json:"pivot"` // The segment from the original query to pivot on.
	Suggestions []Query `json:"suggestions"` // A list of suggested queries for the pivot.
}

// InsightsTag represents the InsightsTag schema from the OpenAPI specification
type InsightsTag struct {
	Name string `json:"name,omitempty"` // The name of the characteristic. For example, cat, kitty, calico cat.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// RecipesModule represents the RecipesModule schema from the OpenAPI specification
type RecipesModule struct {
	Value []Recipe `json:"value,omitempty"` // A list of recipes.
}

// RelatedSearchesModule represents the RelatedSearchesModule schema from the OpenAPI specification
type RelatedSearchesModule struct {
	Value []Query `json:"value,omitempty"` // A list of related searches.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// Person represents the Person schema from the OpenAPI specification
type Person struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// ImagesImageMetadata represents the ImagesImageMetadata schema from the OpenAPI specification
type ImagesImageMetadata struct {
	Aggregateoffer AggregateOffer `json:"aggregateOffer,omitempty"` // Defines a list of offers from merchants that are related to the image.
	Recipesourcescount int `json:"recipeSourcesCount,omitempty"` // The number of websites that offer recipes of the food seen in the image.
	Shoppingsourcescount int `json:"shoppingSourcesCount,omitempty"` // The number of websites that offer goods of the products seen in the image.
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// PropertiesItem represents the PropertiesItem schema from the OpenAPI specification
type PropertiesItem struct {
	TypeField string `json:"_type"`
	Text string `json:"text,omitempty"` // Text representation of an item.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
}

// NormalizedRectangle represents the NormalizedRectangle schema from the OpenAPI specification
type NormalizedRectangle struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// TrendingImagesCategory represents the TrendingImagesCategory schema from the OpenAPI specification
type TrendingImagesCategory struct {
	Title string `json:"title"` // The name of the image category. For example, Popular People Searches.
	Tiles []TrendingImagesTile `json:"tiles"` // A list of images that are trending in the category. Each tile contains an image and a URL that returns more images of the subject. For example, if the category is Popular People Searches, the image is of a popular person and the URL would return more images of that person.
}

// TrendingImagesTile represents the TrendingImagesTile schema from the OpenAPI specification
type TrendingImagesTile struct {
	Query Query `json:"query"` // Defines a search query.
	Image ImageObject `json:"image"` // Defines an image
}

// Rating represents the Rating schema from the OpenAPI specification
type Rating struct {
	TypeField string `json:"_type"`
	Text string `json:"text,omitempty"` // Text representation of an item.
}

// RecognizedEntity represents the RecognizedEntity schema from the OpenAPI specification
type RecognizedEntity struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// TrendingImages represents the TrendingImages schema from the OpenAPI specification
type TrendingImages struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// MediaObject represents the MediaObject schema from the OpenAPI specification
type MediaObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// Images represents the Images schema from the OpenAPI specification
type Images struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
}

// ImagesModule represents the ImagesModule schema from the OpenAPI specification
type ImagesModule struct {
	Value []ImageObject `json:"value,omitempty"` // A list of images.
}

// RecognizedEntityRegion represents the RecognizedEntityRegion schema from the OpenAPI specification
type RecognizedEntityRegion struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// ImageInsights represents the ImageInsights schema from the OpenAPI specification
type ImageInsights struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// ImageObject represents the ImageObject schema from the OpenAPI specification
type ImageObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Hostpageurl string `json:"hostPageUrl,omitempty"` // URL of the page that hosts the media object.
	Width int `json:"width,omitempty"` // The width of the source media object, in pixels.
	Contentsize string `json:"contentSize,omitempty"` // Size of the media object content (use format "value unit" e.g "1024 B").
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g the source URL for the image).
	Encodingformat string `json:"encodingFormat,omitempty"` // Encoding format (e.g mp3, mp4, jpeg, etc).
	Height int `json:"height,omitempty"` // The height of the source media object, in pixels.
	Hostpagedisplayurl string `json:"hostPageDisplayUrl,omitempty"` // Display URL of the page that hosts the media object.
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
}

// Intangible represents the Intangible schema from the OpenAPI specification
type Intangible struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
}

// Offer represents the Offer schema from the OpenAPI specification
type Offer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
}

// Organization represents the Organization schema from the OpenAPI specification
type Organization struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// ImageTagsModule represents the ImageTagsModule schema from the OpenAPI specification
type ImageTagsModule struct {
	Value []InsightsTag `json:"value"` // A list of tags that describe the characteristics of the content found in the image. For example, if the image is of a musical artist, the list might include Female, Dress, and Music to indicate the person is female music artist that's wearing a dress.
}

// StructuredValue represents the StructuredValue schema from the OpenAPI specification
type StructuredValue struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
}

// CollectionPage represents the CollectionPage schema from the OpenAPI specification
type CollectionPage struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// ImageInsightsImageCaption represents the ImageInsightsImageCaption schema from the OpenAPI specification
type ImageInsightsImageCaption struct {
	Caption string `json:"caption"` // A caption about the image.
	Datasourceurl string `json:"dataSourceUrl"` // The URL to the website where the caption was found. You must attribute the caption to the source. For example, by displaying the domain name from the URL next to the caption and using the URL to link to the source website.
	Relatedsearches []Query `json:"relatedSearches"` // A list of entities found in the caption. Use the contents of the Query object to find the entity in the caption and create a link. The link takes the user to images of the entity.
}

// RecognizedEntityGroup represents the RecognizedEntityGroup schema from the OpenAPI specification
type RecognizedEntityGroup struct {
	Name string `json:"name"` // The name of the group where images of the entity were also found. The following are possible groups. CelebRecognitionAnnotations: Similar to CelebrityAnnotations but provides a higher probability of an accurate match. CelebrityAnnotations: Contains celebrities such as actors, politicians, athletes, and historical figures.
	Recognizedentityregions []RecognizedEntityRegion `json:"recognizedEntityRegions"` // The regions of the image that contain entities.
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Thumbnail ImageObject `json:"thumbnail,omitempty"` // Defines an image
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL that takes the user to the Bing search results page for the query.Only related search results include this field.
	Displaytext string `json:"displayText,omitempty"` // The display version of the query term. This version of the query term may contain special characters that highlight the search term found in the query string. The string contains the highlighting characters only if the query enabled hit highlighting
	Searchlink string `json:"searchLink,omitempty"` // The URL that you use to get the results of the related search. Before using the URL, you must append query parameters as appropriate and include the Ocp-Apim-Subscription-Key header. Use this URL if you're displaying the results in your own user interface. Otherwise, use the webSearchUrl URL.
	Text string `json:"text"` // The query string. Use this string as the query term in a new search request.
}

// AggregateRating represents the AggregateRating schema from the OpenAPI specification
type AggregateRating struct {
	TypeField string `json:"_type"`
	Text string `json:"text,omitempty"` // Text representation of an item.
	Bestrating float32 `json:"bestRating,omitempty"` // The highest rated review. The possible values are 1.0 through 5.0.
	Ratingvalue float32 `json:"ratingValue"` // The mean (average) rating. The possible values are 1.0 through 5.0.
}

// ImageGallery represents the ImageGallery schema from the OpenAPI specification
type ImageGallery struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// Recipe represents the Recipe schema from the OpenAPI specification
type Recipe struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Datepublished string `json:"datePublished,omitempty"` // The date on which the CreativeWork was published.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// AggregateOffer represents the AggregateOffer schema from the OpenAPI specification
type AggregateOffer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Alternatename string `json:"alternateName,omitempty"` // An alias for the item
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Price float32 `json:"price,omitempty"` // The item's price.
	Pricecurrency string `json:"priceCurrency,omitempty"` // The monetary currency. For example, USD.
	Seller Organization `json:"seller,omitempty"` // Defines an organization.
	Aggregaterating AggregateRating `json:"aggregateRating,omitempty"` // Defines the metrics that indicate how well an item was rated by others.
	Availability string `json:"availability,omitempty"` // The item's availability. The following are the possible values: Discontinued, InStock, InStoreOnly, LimitedAvailability, OnlineOnly, OutOfStock, PreOrder, SoldOut
	Lastupdated string `json:"lastUpdated,omitempty"` // The last date that the offer was updated. The date is in the form YYYY-MM-DD.
}

// RecognizedEntitiesModule represents the RecognizedEntitiesModule schema from the OpenAPI specification
type RecognizedEntitiesModule struct {
	Value []RecognizedEntityGroup `json:"value,omitempty"` // A list of recognized entities.
}

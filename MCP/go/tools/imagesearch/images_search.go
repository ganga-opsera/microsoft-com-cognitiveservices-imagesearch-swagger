package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/image-search-client/mcp-server/config"
	"github.com/image-search-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Images_searchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["aspect"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("aspect=%v", val))
		}
		if val, ok := args["color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("color=%v", val))
		}
		if val, ok := args["cc"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cc=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["freshness"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("freshness=%v", val))
		}
		if val, ok := args["height"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("height=%v", val))
		}
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["imageContent"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("imageContent=%v", val))
		}
		if val, ok := args["imageType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("imageType=%v", val))
		}
		if val, ok := args["license"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("license=%v", val))
		}
		if val, ok := args["mkt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mkt=%v", val))
		}
		if val, ok := args["maxFileSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxFileSize=%v", val))
		}
		if val, ok := args["maxHeight"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxHeight=%v", val))
		}
		if val, ok := args["maxWidth"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxWidth=%v", val))
		}
		if val, ok := args["minFileSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minFileSize=%v", val))
		}
		if val, ok := args["minHeight"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minHeight=%v", val))
		}
		if val, ok := args["minWidth"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("minWidth=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["safeSearch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("safeSearch=%v", val))
		}
		if val, ok := args["size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("size=%v", val))
		}
		if val, ok := args["setLang"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("setLang=%v", val))
		}
		if val, ok := args["width"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("width=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/images/search%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Ocp-Apim-Subscription-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-BingApis-SDK"]; ok {
			req.Header.Set("X-BingApis-SDK", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Accept"]; ok {
			req.Header.Set("Accept", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Accept-Language"]; ok {
			req.Header.Set("Accept-Language", fmt.Sprintf("%v", val))
		}
		if val, ok := args["User-Agent"]; ok {
			req.Header.Set("User-Agent", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-MSEdge-ClientID"]; ok {
			req.Header.Set("X-MSEdge-ClientID", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-MSEdge-ClientIP"]; ok {
			req.Header.Set("X-MSEdge-ClientIP", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Search-Location"]; ok {
			req.Header.Set("X-Search-Location", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Images
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateImages_searchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_images_search",
		mcp.WithDescription("The Image Search API lets you send a search query to Bing and get back a list of relevant images. This section provides technical details about the query parameters and headers that you use to request images and the JSON response objects that contain them. For examples that show how to make requests, see [Searching the Web for Images](https://docs.microsoft.com/azure/cognitive-services/bing-image-search/search-the-web)."),
		mcp.WithString("X-BingApis-SDK", mcp.Required(), mcp.Description("Activate swagger compliance")),
		mcp.WithString("Accept", mcp.Description("The default media type is application/json. To specify that the response use [JSON-LD](http://json-ld.org/), set the Accept header to application/ld+json.")),
		mcp.WithString("Accept-Language", mcp.Description("A comma-delimited list of one or more languages to use for user interface strings. The list is in decreasing order of preference. For additional information, including expected format, see [RFC2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). This header and the [setLang](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#setlang) query parameter are mutually exclusive; do not specify both. If you set this header, you must also specify the [cc](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#cc) query parameter. To determine the market to return results for, Bing uses the first supported language it finds from the list and combines it with the cc parameter value. If the list does not include a supported language, Bing finds the closest language and market that supports the request or it uses an aggregated or default market for the results. To determine the market that Bing used, see the BingAPIs-Market header. Use this header and the cc query parameter only if you specify multiple languages. Otherwise, use the [mkt](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#mkt) and [setLang](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#setlang) query parameters. A user interface string is a string that's used as a label in a user interface. There are few user interface strings in the JSON response objects. Any links to Bing.com properties in the response objects apply the specified language.")),
		mcp.WithString("User-Agent", mcp.Description("The user agent originating the request. Bing uses the user agent to provide mobile users with an optimized experience. Although optional, you are encouraged to always specify this header. The user-agent should be the same string that any commonly used browser sends. For information about user agents, see [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). The following are examples of user-agent strings. Windows Phone: Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 822). Android: Mozilla / 5.0 (Linux; U; Android 2.3.5; en - us; SCH - I500 Build / GINGERBREAD) AppleWebKit / 533.1 (KHTML; like Gecko) Version / 4.0 Mobile Safari / 533.1. iPhone: Mozilla / 5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit / 536.26 (KHTML; like Gecko) Mobile / 10B142 iPhone4; 1 BingWeb / 3.03.1428.20120423. PC: Mozilla / 5.0 (Windows NT 6.3; WOW64; Trident / 7.0; Touch; rv:11.0) like Gecko. iPad: Mozilla / 5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit / 537.51.1 (KHTML, like Gecko) Version / 7.0 Mobile / 11A465 Safari / 9537.53")),
		mcp.WithString("X-MSEdge-ClientID", mcp.Description("Bing uses this header to provide users with consistent behavior across Bing API calls. Bing often flights new features and improvements, and it uses the client ID as a key for assigning traffic on different flights. If you do not use the same client ID for a user across multiple requests, then Bing may assign the user to multiple conflicting flights. Being assigned to multiple conflicting flights can lead to an inconsistent user experience. For example, if the second request has a different flight assignment than the first, the experience may be unexpected. Also, Bing can use the client ID to tailor web results to that client ID’s search history, providing a richer experience for the user. Bing also uses this header to help improve result rankings by analyzing the activity generated by a client ID. The relevance improvements help with better quality of results delivered by Bing APIs and in turn enables higher click-through rates for the API consumer. IMPORTANT: Although optional, you should consider this header required. Persisting the client ID across multiple requests for the same end user and device combination enables 1) the API consumer to receive a consistent user experience, and 2) higher click-through rates via better quality of results from the Bing APIs. Each user that uses your application on the device must have a unique, Bing generated client ID. If you do not include this header in the request, Bing generates an ID and returns it in the X-MSEdge-ClientID response header. The only time that you should NOT include this header in a request is the first time the user uses your app on that device. Use the client ID for each Bing API request that your app makes for this user on the device. Persist the client ID. To persist the ID in a browser app, use a persistent HTTP cookie to ensure the ID is used across all sessions. Do not use a session cookie. For other apps such as mobile apps, use the device's persistent storage to persist the ID. The next time the user uses your app on that device, get the client ID that you persisted. Bing responses may or may not include this header. If the response includes this header, capture the client ID and use it for all subsequent Bing requests for the user on that device. If you include the X-MSEdge-ClientID, you must not include cookies in the request.")),
		mcp.WithString("X-MSEdge-ClientIP", mcp.Description("The IPv4 or IPv6 address of the client device. The IP address is used to discover the user's location. Bing uses the location information to determine safe search behavior. Although optional, you are encouraged to always specify this header and the X-Search-Location header. Do not obfuscate the address (for example, by changing the last octet to 0). Obfuscating the address results in the location not being anywhere near the device's actual location, which may result in Bing serving erroneous results.")),
		mcp.WithString("X-Search-Location", mcp.Description("A semicolon-delimited list of key/value pairs that describe the client's geographical location. Bing uses the location information to determine safe search behavior and to return relevant local content. Specify the key/value pair as <key>:<value>. The following are the keys that you use to specify the user's location. lat (required): The latitude of the client's location, in degrees. The latitude must be greater than or equal to -90.0 and less than or equal to +90.0. Negative values indicate southern latitudes and positive values indicate northern latitudes. long (required): The longitude of the client's location, in degrees. The longitude must be greater than or equal to -180.0 and less than or equal to +180.0. Negative values indicate western longitudes and positive values indicate eastern longitudes. re (required): The radius, in meters, which specifies the horizontal accuracy of the coordinates. Pass the value returned by the device's location service. Typical values might be 22m for GPS/Wi-Fi, 380m for cell tower triangulation, and 18,000m for reverse IP lookup. ts (optional): The UTC UNIX timestamp of when the client was at the location. (The UNIX timestamp is the number of seconds since January 1, 1970.) head (optional): The client's relative heading or direction of travel. Specify the direction of travel as degrees from 0 through 360, counting clockwise relative to true north. Specify this key only if the sp key is nonzero. sp (optional): The horizontal velocity (speed), in meters per second, that the client device is traveling. alt (optional): The altitude of the client device, in meters. are (optional): The radius, in meters, that specifies the vertical accuracy of the coordinates. Specify this key only if you specify the alt key. Although many of the keys are optional, the more information that you provide, the more accurate the location results are. Although optional, you are encouraged to always specify the user's geographical location. Providing the location is especially important if the client's IP address does not accurately reflect the user's physical location (for example, if the client uses VPN). For optimal results, you should include this header and the X-MSEdge-ClientIP header, but at a minimum, you should include this header.")),
		mcp.WithString("aspect", mcp.Description("Filter images by the following aspect ratios. All: Do not filter by aspect.Specifying this value is the same as not specifying the aspect parameter. Square: Return images with standard aspect ratio. Wide: Return images with wide screen aspect ratio. Tall: Return images with tall aspect ratio.")),
		mcp.WithString("color", mcp.Description("Filter images by the following color options. ColorOnly: Return color images. Monochrome: Return black and white images. Return images with one of the following dominant colors: Black, Blue, Brown, Gray, Green, Orange, Pink, Purple, Red, Teal, White, Yellow")),
		mcp.WithString("cc", mcp.Description("A 2-character country code of the country where the results come from. For a list of possible values, see [Market Codes](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#market-codes). If you set this parameter, you must also specify the [Accept-Language](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#acceptlanguage) header. Bing uses the first supported language it finds from the languages list, and combine that language with the country code that you specify to determine the market to return results for. If the languages list does not include a supported language, Bing finds the closest language and market that supports the request, or it may use an aggregated or default market for the results instead of a specified one. You should use this query parameter and the Accept-Language query parameter only if you specify multiple languages; otherwise, you should use the mkt and setLang query parameters. This parameter and the [mkt](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#mkt) query parameter are mutually exclusive—do not specify both.")),
		mcp.WithNumber("count", mcp.Description("The number of images to return in the response. The actual number delivered may be less than requested. The default is 35. The maximum value is 150. You use this parameter along with the offset parameter to page results.For example, if your user interface displays 20 images per page, set count to 20 and offset to 0 to get the first page of results.For each subsequent page, increment offset by 20 (for example, 0, 20, 40). Use this parameter only with the Image Search API.Do not specify this parameter when calling the Insights, Trending Images, or Web Search APIs.")),
		mcp.WithString("freshness", mcp.Description("Filter images by the following discovery options. Day: Return images discovered by Bing within the last 24 hours. Week: Return images discovered by Bing within the last 7 days. Month: Return images discovered by Bing within the last 30 days.")),
		mcp.WithNumber("height", mcp.Description("Filter images that have the specified height, in pixels. You may use this filter with the size filter to return small images that have a height of 150 pixels.")),
		mcp.WithString("id", mcp.Description("An ID that uniquely identifies an image. Use this parameter to ensure that the specified image is the first image in the list of images that Bing returns. The [Image](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#image) object's imageId field contains the ID that you set this parameter to.")),
		mcp.WithString("imageContent", mcp.Description("Filter images by the following content types. Face: Return images that show only a person's face. Portrait: Return images that show only a person's head and shoulders.")),
		mcp.WithString("imageType", mcp.Description("Filter images by the following image types. AnimatedGif: Return only animated GIFs. Clipart: Return only clip art images. Line: Return only line drawings. Photo: Return only photographs(excluding line drawings, animated Gifs, and clip art). Shopping: Return only images that contain items where Bing knows of a merchant that is selling the items. This option is valid in the en - US market only.Transparent: Return only images with a transparent background.")),
		mcp.WithString("license", mcp.Description("Filter images by the following license types. All: Do not filter by license type.Specifying this value is the same as not specifying the license parameter. Any: Return images that are under any license type. The response doesn't include images that do not specify a license or the license is unknown. Public: Return images where the creator has waived their exclusive rights, to the fullest extent allowed by law. Share: Return images that may be shared with others. Changing or editing the image might not be allowed. Also, modifying, sharing, and using the image for commercial purposes might not be allowed. Typically, this option returns the most images. ShareCommercially: Return images that may be shared with others for personal or commercial purposes. Changing or editing the image might not be allowed. Modify: Return images that may be modified, shared, and used. Changing or editing the image might not be allowed. Modifying, sharing, and using the image for commercial purposes might not be allowed. ModifyCommercially: Return images that may be modified, shared, and used for personal or commercial purposes. Typically, this option returns the fewest images. For more information about these license types, see [Filter Images By License Type](http://go.microsoft.com/fwlink/?LinkId=309768).")),
		mcp.WithString("mkt", mcp.Description("The market where the results come from. Typically, mkt is the country where the user is making the request from. However, it could be a different country if the user is not located in a country where Bing delivers results. The market must be in the form <language code>-<country code>. For example, en-US. The string is case insensitive. For a list of possible market values, see [Market Codes](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#market-codes). NOTE: If known, you are encouraged to always specify the market. Specifying the market helps Bing route the request and return an appropriate and optimal response. If you specify a market that is not listed in [Market Codes](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#market-codes), Bing uses a best fit market code based on an internal mapping that is subject to change. This parameter and the [cc](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#cc) query parameter are mutually exclusive—do not specify both.")),
		mcp.WithNumber("maxFileSize", mcp.Description("Filter images that are less than or equal to the specified file size. The maximum file size that you may specify is 520,192 bytes. If you specify a larger value, the API uses 520,192. It is possible that the response may include images that are slightly larger than the specified maximum. You may specify this filter and minFileSize to filter images within a range of file sizes.")),
		mcp.WithNumber("maxHeight", mcp.Description("Filter images that have a height that is less than or equal to the specified height. Specify the height in pixels. You may specify this filter and minHeight to filter images within a range of heights. This filter and the height filter are mutually exclusive.")),
		mcp.WithNumber("maxWidth", mcp.Description("Filter images that have a width that is less than or equal to the specified width. Specify the width in pixels. You may specify this filter and maxWidth to filter images within a range of widths. This filter and the width filter are mutually exclusive.")),
		mcp.WithNumber("minFileSize", mcp.Description("Filter images that are greater than or equal to the specified file size. The maximum file size that you may specify is 520,192 bytes. If you specify a larger value, the API uses 520,192. It is possible that the response may include images that are slightly smaller than the specified minimum. You may specify this filter and maxFileSize to filter images within a range of file sizes.")),
		mcp.WithNumber("minHeight", mcp.Description("Filter images that have a height that is greater than or equal to the specified height. Specify the height in pixels. You may specify this filter and maxHeight to filter images within a range of heights. This filter and the height filter are mutually exclusive.")),
		mcp.WithNumber("minWidth", mcp.Description("Filter images that have a width that is greater than or equal to the specified width. Specify the width in pixels. You may specify this filter and maxWidth to filter images within a range of widths. This filter and the width filter are mutually exclusive.")),
		mcp.WithNumber("offset", mcp.Description("The zero-based offset that indicates the number of images to skip before returning images. The default is 0. The offset should be less than ([totalEstimatedMatches](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#totalestimatedmatches) - count). Use this parameter along with the count parameter to page results. For example, if your user interface displays 20 images per page, set count to 20 and offset to 0 to get the first page of results. For each subsequent page, increment offset by 20 (for example, 0, 20, 40). It is possible for multiple pages to include some overlap in results. To prevent duplicates, see [nextOffset](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#nextoffset). Use this parameter only with the Image API. Do not specify this parameter when calling the Trending Images API or the Web Search API.")),
		mcp.WithString("q", mcp.Required(), mcp.Description("The user's search query term. The term cannot be empty. The term may contain [Bing Advanced Operators](http://msdn.microsoft.com/library/ff795620.aspx). For example, to limit images to a specific domain, use the [site:](http://msdn.microsoft.com/library/ff795613.aspx) operator. To help improve relevance of an insights query (see [insightsToken](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#insightstoken)), you should always include the user's query term. Use this parameter only with the Image Search API.Do not specify this parameter when calling the Trending Images API.")),
		mcp.WithString("safeSearch", mcp.Description("Filter images for adult content. The following are the possible filter values. Off: May return images with adult content. If the request is through the Image Search API, the response includes thumbnail images that are clear (non-fuzzy). However, if the request is through the Web Search API, the response includes thumbnail images that are pixelated (fuzzy). Moderate: If the request is through the Image Search API, the response doesn't include images with adult content. If the request is through the Web Search API, the response may include images with adult content (the thumbnail images are pixelated (fuzzy)). Strict: Do not return images with adult content. The default is Moderate. If the request comes from a market that Bing's adult policy requires that safeSearch is set to Strict, Bing ignores the safeSearch value and uses Strict. If you use the site: query operator, there is the chance that the response may contain adult content regardless of what the safeSearch query parameter is set to. Use site: only if you are aware of the content on the site and your scenario supports the possibility of adult content.")),
		mcp.WithString("size", mcp.Description("Filter images by the following sizes. All: Do not filter by size. Specifying this value is the same as not specifying the size parameter. Small: Return images that are less than 200x200 pixels. Medium: Return images that are greater than or equal to 200x200 pixels but less than 500x500 pixels. Large: Return images that are 500x500 pixels or larger. Wallpaper: Return wallpaper images. You may use this parameter along with the height or width parameters. For example, you may use height and size to request small images that are 150 pixels tall.")),
		mcp.WithString("setLang", mcp.Description("The language to use for user interface strings. Specify the language using the ISO 639-1 2-letter language code. For example, the language code for English is EN. The default is EN (English). Although optional, you should always specify the language. Typically, you set setLang to the same language specified by mkt unless the user wants the user interface strings displayed in a different language. This parameter and the [Accept-Language](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#acceptlanguage) header are mutually exclusive; do not specify both. A user interface string is a string that's used as a label in a user interface. There are few user interface strings in the JSON response objects. Also, any links to Bing.com properties in the response objects apply the specified language.")),
		mcp.WithNumber("width", mcp.Description("Filter images that have the specified width, in pixels. You may use this filter with the size filter to return small images that have a width of 150 pixels.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Images_searchHandler(cfg),
	}
}

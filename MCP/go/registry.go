package main

import (
	"github.com/image-search-client/mcp-server/config"
	"github.com/image-search-client/mcp-server/models"
	tools_imagedetailsearch "github.com/image-search-client/mcp-server/tools/imagedetailsearch"
	tools_imagesearch "github.com/image-search-client/mcp-server/tools/imagesearch"
	tools_imagetrendingsearch "github.com/image-search-client/mcp-server/tools/imagetrendingsearch"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_imagedetailsearch.CreateImages_detailsTool(cfg),
		tools_imagesearch.CreateImages_searchTool(cfg),
		tools_imagetrendingsearch.CreateImages_trendingTool(cfg),
	}
}
